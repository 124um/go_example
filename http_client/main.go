package main

import (
	"coincap"
	"fmt"
	"log"
	"time"
)

func main() {
	coincapClient, err := coincap.NewClient(time.Second * 10)
	fatal(err)

	assets, err := coincapClient.GetAssets()
	fatal(err)

	for _, asset := range assets {
		fmt.Println(asset.Info())
	}

	bitcoin, err := coincapClient.GetAsset("bitcoin")
	fatal(err)

	fmt.Println(bitcoin.Info())
}

func fatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}