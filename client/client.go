package main

import (
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Errore nella connessione al server:", err)
	}
	defer client.Close()

	args := &Args{A: 10, B: 5}
	var risposta int

	for {
		fmt.Print("\nInserisci valore:\n\n1 per Addizione\n2 per Sottrazione\n\n ")
		var input string
		fmt.Scanln(&input)

		if strings.EqualFold(input, "1") {
			// Chiamata RPC per l'addizione
			err = client.Call("CalcolatriceRPC.Addiziona", args, &risposta)
			if err != nil {
				log.Fatal("Errore nella chiamata RPC:", err)
			}
			fmt.Printf("Risultato Addiziona: %d\n", risposta)
		}

		if strings.EqualFold(input, "2") {
			// Chiamata RPC per la sottrazione
			err = client.Call("CalcolatriceRPC.Sottrai", args, &risposta)
			if err != nil {
				log.Fatal("Errore nella chiamata RPC:", err)
			}
			fmt.Printf("Risultato Sottrai: %d\n", risposta)
		}
	}
}
