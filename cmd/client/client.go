package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/regismartiny/go-expert-desafio-client-server-api/internal/usdbrlquotesserver"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	client := usdbrlquotesserver.NewClient()

	quote, err := client.GetQuote(&ctx)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(quote)

	saveQuoteToFile(quote)
}

func saveQuoteToFile(quote float64) {
	log.Println("Salvando cotação no arquivo:", quote)

	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	file.WriteString(fmt.Sprintf("Dólar: %f\n", quote))
}
