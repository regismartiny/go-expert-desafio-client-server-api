package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/regismartiny/go-expert-desafio-client-server-api/internal/awesomeapi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /cotacao", handleGet)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Server is listening on http://localhost:8080")
}

func handleGet(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	client := awesomeapi.NewClient()

	quote, err := client.GetUSDBRLQuote(&ctx)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(quote)

	quoteDB := quoteToQuoteDB(quote)
	persist(&quoteDB)

	w.Write([]byte(quote.Bid))
}

func quoteToQuoteDB(quote awesomeapi.Quote) QuoteDB {

	return QuoteDB{
		Code:       quote.Code,
		Codein:     quote.Codein,
		Name:       quote.Name,
		High:       quote.High,
		Low:        quote.Low,
		VarBid:     quote.VarBid,
		PctChange:  quote.PctChange,
		Bid:        quote.Bid,
		Ask:        quote.Ask,
		Timestamp:  quote.Timestamp,
		CreateDate: quote.CreateDate,
	}
}

func persist(quote *QuoteDB) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	db, err := connectToSQLite(&ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&QuoteDB{})
	if err != nil {
		log.Fatal(err)
	}

	err = createQuote(db, quote)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Quote:", quote)
}

func connectToSQLite(ctx *context.Context) (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("quotes.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db.WithContext(*ctx), nil
}

func createQuote(db *gorm.DB, quote *QuoteDB) error {

	result := db.Create(quote)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type QuoteDB struct {
	ID         uint `gorm:"primaryKey"`
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
}
