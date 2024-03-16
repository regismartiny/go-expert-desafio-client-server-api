package main

import (
	"context"
	"time"
)

func main() {

	//buscar cotacao em localhost:8080/cotacao
	//timeout 300ms

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	//salvar resultado arquivo cotacao.txt formato: Dólar: {valor}
}
