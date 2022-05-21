package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/emilgalimov/homework-2/pkg/api/v1"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/app"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/config"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/repository"
	"gitlab.ozon.dev/emilgalimov/homework-2_2/internal/router"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		panic(fmt.Sprint("error load config ", err))
	}

	ctx := context.Background()

	connectString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DbName,
	)

	pgConn, _ := pgxpool.Connect(ctx, connectString)

	if err := pgConn.Ping(ctx); err != nil {
		log.Fatal("error pinging db: ", err)
	}

	repo := repository.NewRepository(pgConn)

	smConn, err := grpc.Dial(cfg.SmartCalendar.Host+":"+cfg.SmartCalendar.Port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer func(smConn *grpc.ClientConn) {
		_ = smConn.Close()
	}(smConn)

	smClient := api.NewSmartCalendarClient(smConn)

	service := app.NewService(repo, smClient)

	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	rt := router.NewRouter(service)

	for update := range updates {
		answer := rt.ProcessMessage(update, ctx)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)

		_, err2 := bot.Send(msg)

		if err2 != nil {
			log.Printf("error send message %v", err2)
		}
	}
}
