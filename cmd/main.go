package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

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
				/* log.Println(action.BlockTimestamp)
				log.Println(action.FloorPrice)
				log.Println(action.MintAddress)
				log.Println(action.Name)
				log.Println(action.Price)
				log.Println(action.Rank)
				log.Println(action.RarityStr)
				log.Println(action.Seller)
				log.Println(action.Supply)
				log.Println(action.Symbol)
				log.Println(action.Timestamp)
				log.Println(action.TokenAddress)
				log.Println(action.Type) */
				w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
				fmt.Fprintln(w, "Type\t\tName\t\tPrice\t\tRank\t\tFloorPr")
				var priceStr string = strconv.FormatFloat(action.Price, 'E', -1, 32)
				var fpStr string = strconv.FormatFloat(action.FloorPrice, 'E', -1, 32)
				var rankStr string = strconv.Itoa(action.Rank)
				fmt.Fprintln(w, action.Type+"\t\t"+action.Name+"\t\t"+priceStr+"\t\t"+rankStr+"\t\t"+fpStr+"\t\t")
				fmt.Fprintln(w, "------------------------------------------------------------------------------")
				w.Flush()
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
