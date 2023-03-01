package main

import (
	"fmt"
	"github.com/jakezhu9/chatgpt-lark-bot/internal/config"
	"github.com/jakezhu9/chatgpt-lark-bot/internal/gpt"
	"github.com/jakezhu9/chatgpt-lark-bot/internal/larkbot"
	"github.com/jakezhu9/chatgpt-lark-bot/internal/util"
	"log"
)

func main() {
	configFile := "./config.yaml"
	exist, err := util.FileExists(configFile)
	if err != nil {
		log.Fatalf("failed to read configuration file: %s", err)
	}
	if !exist {
		err = config.WriteDefaultConfig(configFile)
		if err != nil {
			log.Fatalf("failed to create configuration file: %s", err)
		}
		log.Println("create default config to " + configFile)
		return
	}

	conf, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("failed to read configuration file: %s", err)
	}
	log.Println(conf.EventEncryptKey)

	g := gpt.New(conf.OpenAIKey)
	bot := larkbot.New(larkbot.Config{
		AppID:             conf.AppID,
		AppSecret:         conf.AppSecret,
		VerificationToken: conf.VerificationToken,
		EventEncryptKey:   conf.EventEncryptKey,
	})
	handler := func(msg larkbot.Message) {
		if msg.Content == "" {
			return
		}
		log.Printf("receive message: %s %s\n", msg.ID, msg.Content)
		res, err := g.Handle(fmt.Sprintf("Q:%s\nA:", msg.Content))
		if err != nil {
			log.Printf("gpt error: %s\n", err)
			return
		}
		log.Printf("gpt: %s %s\n", msg.Content, res)
		err = bot.Reply(msg.ID, res)
		if err != nil {
			log.Printf("reply error: %s\n", err)
			return
		}
	}

	log.Fatal(bot.Run(handler))
}
