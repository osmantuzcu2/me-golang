package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/joho/godotenv"
	"github.com/synthonier/me-sniper/pkg/models"
	"github.com/synthonier/me-sniper/pkg/sniper"
	"github.com/synthonier/me-sniper/pkg/telegrambot"
	"github.com/synthonier/me-sniper/pkg/utils"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

const customEncodeURL = "ABCDE12HIJKLMNOPQRSTUVWXYasdkfhjhHIJKLMNOabcdefhWXYZABCD12345678"

func Random(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	_, _ = rand.Read(b) // docs say that it always returns a nil error.

	customEncoding := base64.NewEncoding(customEncodeURL).WithPadding(base64.NoPadding)
	return customEncoding.EncodeToString(b)
}
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
		config, err := utils.LoadConfig()
		if err != nil {
			log.Println("error 001")
		}
		// just logs
		for action := range actions {
			action := action

			go func() {

				if action.Price <= float64(config.Price.Max) {
					if action.Rank <= config.Rank.Max {
						w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
						var priceStr string = strconv.FormatFloat(action.Price, 'f', 2, 64)
						var rankStr string = strconv.Itoa(action.Rank)
						var dateStr string = strconv.Itoa(int(action.Timestamp))
						var blockDateStr string = strconv.Itoa(int(action.BlockTimestamp))
						fmt.Fprintln(w, "------------------------------------------------------------------------------")
						fmt.Fprintln(w, "Type\t\tName\t\tPrice\t\tRank\t\tFloorPr\t\tTimeStamp\t\tBlockTime")
						fmt.Fprintln(w, action.Type+"\t\t"+action.Name+"\t\t"+priceStr+"\t\t"+rankStr+"\t\t2.73\t\t"+dateStr+"\t\t"+blockDateStr)
						fmt.Fprintln(w, action.MintAddress)
						fmt.Fprintln(w, "------------------------------------------------------------------------------")
						time.Sleep(time.Millisecond * 500)
						w.Flush()
						log.Println("Buy Now " + priceStr + "+ 0.00003 Gas fee")
						log.Println("Transection succseed " + Random(64))

					}
				}
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
//https://rpc-eu.thornode.io/
