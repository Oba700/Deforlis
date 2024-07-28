package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type handler struct {
	Description string
	Type        string
	Path        string
}

func handlingDispatcher(conn net.Conn, Handler handler, BufferSize int) {
	switch Handler.Type {
	case "mock":
		go mock(conn)
	case "catalog":
		go catalog(conn, Handler, BufferSize)
	}
}

func handlingTerminator(resp []byte, needConnClose bool, conn net.Conn, Handler handler, BufferSize int) {
	conn.Write(resp)
	if needConnClose {
		conn.Close()
	} else {
		handlingDispatcher(conn, Handler, BufferSize)
	}

}

func catalog(conn net.Conn, handler handler, BufferSize int) {
	buf := make([]byte, BufferSize)
	_, err := conn.Read((buf))
	if err != nil {
		if err == io.EOF {
			fmt.Println(conn.RemoteAddr(), "З'ЄДНАННЯ РОЗІРВАНО")
		} else {
			fmt.Println(err)
		}
		return
	}
	request := strings.Split(string(buf), "\n")
	//fmt.Println(request)
	firstReqStr := strings.Split(request[0], " ")
	reqPath := firstReqStr[1]
	method := firstReqStr[0]
	if method != "GET" {
		resp := []byte("HTTP/1.1 405 Method Not Allowed\nServer: deforlis/prealpha\nContent-Type: text/html; charset=UTF-8\n\nСкористайтеся методом GET\n")
		handlingTerminator(resp, true, conn, handler, BufferSize)
	}
	quPath, quErr := url.QueryUnescape(reqPath)
	osPath := handler.Path + quPath
	if quErr != nil {
		fmt.Println("URL запита попердолено")
	}
	//fmt.Println(quPath)
	var needConnClose bool = true
	for _, header := range request {
		if strings.HasPrefix(header, "Connection:") && strings.Contains(header, "keep-alive") {
			needConnClose = false
		}
	}
	reqStuffStat, err := os.Lstat(osPath)
	if err != nil {
		handlingTerminator(notFound(quPath), true, conn, handler, BufferSize)
		return
	}
	switch mode := reqStuffStat.Mode(); {
	case mode.IsRegular():
		dat, err := os.ReadFile(osPath)
		if err != nil {
			panic(err)
		}
		mimeType := http.DetectContentType(dat)
		headers := []byte(fmt.Sprintf(`HTTP/1.1 200 OK
Server: deforlis/prealpha
Content-Type: %s
Content-Length: %d

`, mimeType, len(dat)+1))
		resp := append(headers, dat...)
		handlingTerminator(append(resp, []byte("\n")...), needConnClose, conn, handler, BufferSize)
	case mode.IsDir():
		var htmlTableRows string
		entries, err := os.ReadDir(osPath)
		if err != nil {
			panic(err)
		}
		for _, e := range entries {
			htmlTableRows += catalogEntrieHTML(quPath, e)
		}

		handlingTerminator([]byte(catalogHTML(htmlTableRows, quPath)), needConnClose, conn, handler, BufferSize)
	default:
		fmt.Println("Похуй")
	}

}

func mock(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read((buf))
	if err != nil {
		fmt.Println(err)
		return
	}
	request := strings.Split(string(buf), "\n")
	var Method string = strings.Split(request[0], " ")[0]
	var Path string = strings.Split(request[0], " ")[1]
	var clientPort = conn.RemoteAddr()
	var Host string
	for _, s := range request {
		if strings.HasPrefix(s, "Host: ") {
			Host = strings.Split(s, " ")[1]
		}
	}

	conn.Write([]byte(mockHTML(Method, Host, Path, clientPort)))
}
