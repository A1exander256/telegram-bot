package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/a1exander256/telegram-bot/config"
	"github.com/a1exander256/telegram-bot/logger"
	"github.com/a1exander256/telegram-bot/telegram"
	"github.com/a1exander256/telegram-bot/telegram/pkg/handler"
	"github.com/a1exander256/telegram-bot/telegram/pkg/service"
	"github.com/a1exander256/telegram-bot/telegram/pkg/storage"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	conf = kingpin.Flag("conf", "path to the config file").Short('c').Default("./config.json").String()
)

func main() {
	kingpin.Parse()

	config, err := config.LoadConfig(*conf)
	if err != nil {
		log.Fatalln(err)
	}

	log := logger.InitLogger(config.Logger.Level)

	tgbot, err := telegram.Init(config.Telegram.Token, config.Telegram.ModeDebug)
	if err != nil {
		log.Fatalln(err)
	}

	log.Info("authorized on account ", tgbot.Self.UserName)

	//поднимаем коннект postgres и инициализируем слои обработки данных
	db, err := storage.NewPostgresDB(&config.Postgres)
	if err != nil {
		log.Fatalln("error occured while running http server : ", err)
	}
	//слой для работы с бд
	storageTG := storage.NewStorage(db, log)
	//слой бизнес логики
	serviceTG := service.NewService(storageTG, log)
	//слой получения и структурирования данных из бота и отправки данных из сервиса клиенту
	handlerTG := handler.NewHandler(serviceTG, tgbot, log)

	go func() {
		handlerTG.ListenUpdates()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("app shutting down")
}
