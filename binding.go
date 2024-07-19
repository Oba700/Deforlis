package main

import (
	"fmt"
	"net"
)

type binding struct {
	Description string
	Address     string
	Handler     string
	Encryption  encryption
	Buffer      int
}

func bind(Binding binding, Handler handler, BufferSize int) {
	ln, err := net.Listen("tcp", Binding.Address)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Помилка прийому")
			fmt.Println(err)
			continue
		}
		handlingDispatcher(conn, Handler, BufferSize)
	}
}
