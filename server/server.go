package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// CalcolatriceRPC è l'interfaccia del servizio RPC.
type CalcolatriceRPC interface {
	Addiziona(args *Args, risposta *int) error
	Sottrai(args *Args, risposta *int) error
}

// Args contiene gli argomenti per le operazioni Addiziona e Sottrai.
type Args struct {
	A, B int
}

type Calcolatrice struct{}

func (c *Calcolatrice) Addiziona(args *Args, risposta *int) error {
	*risposta = args.A + args.B
	return nil
}

func (c *Calcolatrice) Sottrai(args *Args, risposta *int) error {
	*risposta = args.A - args.B
	return nil
}

func main() {
	calcolatrice := new(Calcolatrice)

	// Registra il server RPC
	server := rpc.NewServer()
	err := server.Register(calcolatrice)
	if err != nil {
		log.Fatal("Errore nella registrazione del servizio:", err)
	}

	// Crea un listener sulla porta 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Errore nella creazione del listener:", err)
	}
	defer listener.Close()

	fmt.Println("In ascolto sulla porta 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Errore nella connessione:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
