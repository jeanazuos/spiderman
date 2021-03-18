package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.webmotors.com.br/carros/estoque/hyundai?tipoveiculo=carros&marca1=HYUNDAI`),
		chromedp.Text(`#root > main > div.container > div.Search-result.Search-result--container-right > div:nth-child(4) > div > div:nth-child(1) > div`, &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(res))
}
