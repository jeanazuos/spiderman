package utils

import (
	"errors"
	"strconv"
	"strings"
)

func WebmmotorsMileageConverter(mileage string) (int64, error) {
	mileage = strings.ReplaceAll(mileage, "km", "")
	mileage = strings.TrimSpace(mileage)

	mileageInt, err := strconv.ParseInt(mileage, 10, 64)

	if err != nil {
		return 0, errors.New("Nao foi possivel converter mileage para inteiro")
	}
	return mileageInt, nil
}
