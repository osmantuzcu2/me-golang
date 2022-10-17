package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/synthonier/me-sniper/pkg/models"
	"github.com/synthonier/me-sniper/pkg/sniper"
	"github.com/synthonier/me-sniper/pkg/telegrambot"
)

func main() {
	log.Println("Bot started")
	err := godotenv.Load()
	checkError(err)

	var actions = make(chan *models.Token, 5)

	// create sniper instance
	s, err := sniper.New(os.Getenv("NODE_ENDPOINT"), actions)
	checkError(err)

	go func() {
		err = s.Start()
		checkError(err)
	}()

	TELEGRAM_APIKEY := os.Getenv("TELEGRAM_APIKEY")
	if TELEGRAM_APIKEY != "" {
		// create and start telegram bot
		tgbot, err := telegrambot.New(TELEGRAM_APIKEY, actions)
		checkError(err)

		err = tgbot.Start()
		checkError(err)
	} else {
		// just logs
		for action := range actions {
			action := action

			go func() {
				log.Println(action)
			}()
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//{ ws: 'wss://ssc-dao.genesysgo.net', rpc: 'https://ssc-dao.genesysgo.net/' },
// { ws: 'wss://solana-api.projectserum.com/', rpc: 'https://solana-api.projectserum.com/' },
//{ ws: 'wss://solana-mainnet.phantom.tech/', rpc: 'https://solana-mainnet.phantom.tech/' },
