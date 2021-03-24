package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
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

	lines := strings.TrimSpace(res)

	lines2 := strings.ReplaceAll(lines, "\n", ",")

	lines3 := strings.ReplaceAll(lines2, ",,", ";")

	carAdvertising := strings.Split(lines3, ";")
	if carAdvertising != nil {
		fmt.Println("----")
	}

	//slice[0] representa um anuncio completo
	// carAdvertising é o antigo slice

	// fmt.Println("anuncio=> ", carAdvertising[12])
	//MOCK
	var carAdvertisingMocked = []string{"HYUNDAI TUCSON,2.0 MPFI GLS 16V 143CV 2WD FLEX 4P AUTOMÁTICO,Car Delivery,Troca + Troco,R$ 58.790,2016/2016,49722 km,São Paulo - SP"}
	// fmt.Println("anunciomocked=>", carAdvertisingMocked[0])

	//Precisamos iterar depois carAdvertising para pegar todos os anuncios
	carAttributes := strings.Split(carAdvertisingMocked[0], ",")

	fmt.Println("ANTES=> ", carAttributes)

	// carAttributes = cleanAttributes(value, "Car Delivery", carAttributes, index)
	carAttributes = cleanAttributes(carAttributes)

	fmt.Println("DEPOIS=> ", carAttributes)

	os.Exit(1)

	// fmt.Print(slice)

	// for index, value := range slice {
	// 	// fmt.Print(index, value)

	// 	//check para remocao
	// 	res := bytes.Compare([]byte(value), []byte("este pode ser um ótimo negócio!"))
	// 	if res == 0 {
	// 		fmt.Println(value)
	// 		slice = append(slice[:index], slice[index+1:]...)
	// 	}
	// fmt.Println(slice)

	// res = bytes.Compare([]byte(value), []byte("Troca + Troco"))
	// if res == 0 {
	// 	slice = append(slice[:index], slice[index+1:]...)
	// }

	// res = bytes.Compare([]byte(value), []byte("Car Delivery"))
	// if res == 0 {
	// 	slice = append(slice[:index], slice[index+1:]...)
	// }

	// res = bytes.Compare([]byte(value), []byte("Alerta para grandes ofertas:"))
	// if res == 0 {
	// 	slice = append(slice[:index], slice[index+1:]...)
	// }
	// fmt.Println(slice)

	// os.Exit(1)

	// }
	// for _, x := range slice {
	// 	fmt.Println(x)

	// }
}
func cleanAttributes(carAttributes []string) []string {

	//ATRIBUTOS PARA SEREM REMOVIDOS
	// attributeToRemove := []string{"Troca + Troco", "Car Delivery"}

	for index, attribute := range carAttributes {
		if bytes.Equal([]byte(attribute), []byte("Car Delivery")) {
			carAttributes = removeIndex(carAttributes, index)
			fmt.Println("x1 ", carAttributes, "index", index)
			break
		}
	}

	for index, attribute := range carAttributes {
		if bytes.Equal([]byte(attribute), []byte("Troca + Troco")) {
			carAttributes = removeIndex(carAttributes, index)
			fmt.Println("x2 ", carAttributes, "index", index)
			break
		}
	}

	return carAttributes
}

func removeIndex(carAttributes []string, index int) []string {
	fmt.Println("FUNCremoveIndex", "carAttributes=> ", carAttributes, "index=>", index)
	ret := make([]string, 0)
	ret = append(ret, carAttributes[:index]...)
	return append(ret, carAttributes[index+1:]...)
}
