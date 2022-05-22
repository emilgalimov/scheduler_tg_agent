#!/bin/bash

include .env.default
include .env

up:
	docker-compose up -d
	go mod tidy
	go run cmd/bot/main.go

down:
	docker-compose down

migrate:
	cd migrations; ./goose postgres "host=${DB_HOST} port=5433 user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable" up
