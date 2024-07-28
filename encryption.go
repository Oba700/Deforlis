package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"
)

type encryption struct {
	Enabled        bool
	Scheme         string
	PrivateKeyURI  string
	CertificateURI string
}

func craftTLSconfig(encBlock encryption) *tls.Config {
	fmt.Println(tls.CipherSuites())
	fmt.Println(tls.InsecureCipherSuites())
	privateKey := getKey(encBlock.PrivateKeyURI)
	certificate := getKey(encBlock.CertificateURI)
	pair, err := tls.X509KeyPair(certificate, privateKey)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Пара ключів попердолена")
	}
	var conf tls.Config
	conf.Certificates = []tls.Certificate{pair}
	conf.CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384}
	return &conf
}

func getKey(uri string) []byte {
	uriSplit := strings.Split(uri, "://")
	if len(uriSplit) == 2 {
		proto := uriSplit[0]
		if len(proto) < 1 {
			panic("Вкажи протокол доступа до ключу. Зараз бачу" + uri)
		}
		switch proto {
		case "file":
			key, err := os.ReadFile(uriSplit[1])
			if err != nil {
				panic("Не можу прочитати файл з ключем" + uri)
			}
			return key
		case "env":
			env := os.Getenv(uriSplit[1])
			fmt.Println("ENV: ", os.Environ())
			return []byte(env)
		default:
			panic("Не знаю такого протокола" + uri)
		}
	} else {
		panic("Вкажи протокол доступа до ключу. Зараз бачу " + uri)
	}
}
