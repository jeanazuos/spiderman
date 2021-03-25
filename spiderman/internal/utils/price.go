package utils

import (
	"errors"
	"strconv"
	"strings"
)

func WebmmotorsPriceConverter(price string) (int64, error) {
	price = strings.TrimSpace(price)
	price = strings.ReplaceAll(price, ".", "")
	priceInt, err := strconv.ParseInt(price, 10, 64)
	if err != nil {
		return 0, errors.New("Nao foi possivel converter price para inteiro")
	}
	return priceInt, nil
}
