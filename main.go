package main

import (
	"fmt"
	"log"

	"github.com/ireneufia/cwAPI"
)

func main() {
	//Get global market data
	getGlobal, err := cwAPI.GetGlobalData()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getGlobal.ActiveMarkets)
	}
	//Get info about one asset
	getBTC, err := cwAPI.GetCoinInfo("bitcoin")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getBTC)

	}
	//GetAllCoinInfo(0) for all coins, GetAllCoinInfo(10) for top 10 coins & etc.
	getCoins, err := cwAPI.GetAllCoinInfo(0)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(getCoins["ethereum"])
	}
}
