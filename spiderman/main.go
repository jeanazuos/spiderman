package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/jeanazuos/buscakr_v2/spiderman/model"
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

	lines := strings.TrimSpace(res)
	lines2 := strings.ReplaceAll(lines, "\n", ",")
	lines3 := strings.ReplaceAll(lines2, ",,", ";")
	carAdvertising := strings.Split(lines3, ";")

	var payloadCar []model.Car

	var carAttributes []string
	for _, car := range carAdvertising {
		carSlice := strings.Split(car, ",")
		//Limpa atributos inválidos
		carAttributes = cleanAttributes(carSlice)

		// Verifica se o len é 6, pois webmotors retorna esta qtd de valores uteis
		if len(carAttributes) == 6 {
			x := model.Car{carAttributes[0], carAttributes[1], carAttributes[2], carAttributes[3], carAttributes[4], carAttributes[5]}
			payloadCar = append(payloadCar, x)
		}

		result, err := json.Marshal(payloadCar)
		if err != nil {
			log.Println(err)
		}

		//preciso criar uma coluna indicando que os dados provem do webmotors para fazermos filtros
		//posteriormente
		fmt.Println("result=> ", string(result))

	}

}
func cleanAttributes(carAttributes []string) []string {

	//ATRIBUTOS PARA SEREM REMOVIDOS
	// attributeToRemove := []string{"Troca + Troco", "Car Delivery"}

	for index, attribute := range carAttributes {
		if bytes.Equal([]byte(attribute), []byte("Car Delivery")) {
			carAttributes = removeIndex(carAttributes, index)
			break
		}
	}

	for index, attribute := range carAttributes {
		if bytes.Equal([]byte(attribute), []byte("Troca + Troco")) {
			carAttributes = removeIndex(carAttributes, index)
			break
		}
	}

	for index, attribute := range carAttributes {
		if bytes.Equal([]byte(attribute), []byte("Alerta para grandes ofertas:")) {
			carAttributes = removeIndex(carAttributes, index)
			break
		}
	}

	for index, attribute := range carAttributes {
		if bytes.Equal([]byte(attribute), []byte("este pode ser um ótimo negócio!")) {
			carAttributes = removeIndex(carAttributes, index)
			break
		}
	}

	return carAttributes
}

func removeIndex(carAttributes []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, carAttributes[:index]...)
	return append(ret, carAttributes[index+1:]...)
}
