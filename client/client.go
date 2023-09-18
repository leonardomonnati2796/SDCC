package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Errore nella connessione al server:", err)
	}
	defer client.Close()

	args := &Args{A: 10, B: 5}
	var risposta int

	// Chiamata RPC per l'addizione
	err = client.Call("CalcolatriceRPC.Addiziona", args, &risposta)
	if err != nil {
		log.Fatal("Errore nella chiamata RPC:", err)
	}
	fmt.Printf("Risultato Addiziona: %d\n", risposta)

	// Chiamata RPC per la sottrazione
	err = client.Call("CalcolatriceRPC.Sottrai", args, &risposta)
	if err != nil {
		log.Fatal("Errore nella chiamata RPC:", err)
	}
	fmt.Printf("Risultato Sottrai: %d\n", risposta)
}
