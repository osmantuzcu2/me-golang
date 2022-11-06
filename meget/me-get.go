package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api-mainnet.magiceden.io/v2/instructions/buy_now?buyer=ERTAk2yxUkQgqayhZ35oKEGgnX1L8e9n6hT3Euf4X2eB&seller=G4825KocqbypVoReCrdDrHfGzYB21StggoTH7mTGzgCW&auctionHouseAddress=E8cU1WiRWjanGxmn96ewBgk9vPTcL6AEZ1t6F6fkgUWe&tokenMint=2KEytXKMmhmGv586pDNJTsrt47gw3EVvGi9mgKc53w3H&price=0.1&buyerReferral=&expiry"
	//url2 := "https://api-mainnet.magiceden.io/v2/tokens/6nwmzg6Xqber6Zg1ZcGj4uNMom7zTNFrUVQ9sC6kvbVv"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("3")
		fmt.Println(err)
		return
	}
	fmt.Println("4")
	fmt.Println(string(body))
}
