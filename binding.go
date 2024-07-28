package main

import (
	"crypto/tls"
	"fmt"
	"net"
)

type binding struct {
	Description string
	Address     string
	Handler     string
	Encryption  encryption
	BufferSize  int
}

func bind(Binding binding, Handler handler, BufferSize int) {
	var ln net.Listener
	var err error
	if Binding.Encryption.Enabled {
		ln, err = tls.Listen("tcp", Binding.Address, craftTLSconfig(Binding.Encryption))
		if err != nil {
			panic(err)
		}
		fmt.Println("🔒", Binding.Address)
	} else {
		ln, err = net.Listen("tcp", Binding.Address)
		if err != nil {
			panic(err)
		}
		fmt.Println("🔓", Binding.Address)
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
