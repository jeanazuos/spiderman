package utils

import (
	"log"
	"strconv"
	"strings"
)

func webmmotorsPriceConverter(price string) int64 {
	price = strings.TrimSpace(price)
	price = strings.ReplaceAll(price, ".", "")
	priceInt, err := strconv.ParseInt(price, 10, 64)
	if err != nil {
		log.Println("Nao foi possivel converter price para inteiro")
	}
	return priceInt
}
