package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
	"time"
)

type encryption struct {
	Enabled        bool
	Scheme         string
	PrivateKeyURI  string
	CertificateURI string
	ExpirityCheck  expirityCheck
}

type expirityCheck struct {
	Enabled         bool
	CheckReqSeconds int
	ToleranceHours  int
}

func loopExpirityCheck(needReRead chan bool, encBlock encryption) {
	for {
		certificate := getKey(encBlock.CertificateURI)
		block, _ := pem.Decode([]byte(certificate))
		if block == nil {
			panic("Не вдалося декодувати сертіфікат в PEM під час перевіркі терміну дії сертіфікату")
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			panic("Не вдалося розпарсити PEM під час перевіркі терміну дії сертіфікату" + err.Error())
		}
		if int(time.Until(cert.NotAfter).Hours()) < encBlock.ExpirityCheck.ToleranceHours {
			var certId string
			for _, n := range cert.DNSNames {
				certId += n + " "
			}
			fmt.Println("Час оновлювати сертіфткат для [ " + certId + "]")
		}
		time.Sleep(time.Duration(encBlock.ExpirityCheck.CheckReqSeconds) * time.Second)
	}
}

func craftTLSconfig(encBlock encryption) *tls.Config {
	// fmt.Println(tls.CipherSuites())
	// fmt.Println(tls.InsecureCipherSuites())
	// fmt.Println(encBlock.Enabled)
	privateKey := getKey(encBlock.PrivateKeyURI)
	if privateKey == nil {
		return nil
	}
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
		return nil
		// panic("Вкажи протокол доступа до ключу. Зараз бачу " + uri)
	}
}
