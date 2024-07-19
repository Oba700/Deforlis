package main

import (
	"fmt"
	"net"
	"net/http"
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

func catalog(conn net.Conn, handler handler, BufferSize int) {
	//defer conn.Close()
	buf := make([]byte, BufferSize)
	_, err := conn.Read((buf))
	if err != nil {
		fmt.Println(err)
		return
	}
	request := strings.Split(string(buf), "\n")
	Path := strings.Split(request[0], " ")[1]
	var reqPath string
	if strings.HasSuffix(handler.Path, "/") {
		reqPath = fmt.Sprintf("%s%s", handler.Path, Path[1:])
	} else {
		reqPath = fmt.Sprintf("%s%s", handler.Path, Path)
	}
	reqStuffStat, err := os.Lstat(reqPath)
	if err != nil {
		fmt.Println(Path)
		conn.Write(notFound())
		handlingDispatcher(conn, handler, BufferSize)
	}
	switch mode := reqStuffStat.Mode(); {
	case mode.IsRegular():
		dat, err := os.ReadFile(reqPath)
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
		conn.Write(append(resp, []byte("\n")...))
		handlingDispatcher(conn, handler, BufferSize)
	case mode.IsDir():
		// fmt.Println("directory")
		var htmlTableRows string
		entries, err := os.ReadDir(reqPath)
		if err != nil {
			panic(err)
		}
		for _, e := range entries {
			htmlTableRows += catalogEntrieHTML(e, Path)
		}
		conn.Write([]byte(catalogHTML(htmlTableRows, Path)))
		handlingDispatcher(conn, handler, BufferSize)
		// fmt.Println(os.ReadDir(reqPath))
	default:
		fmt.Println("Похуй")
		// case mode&fs.ModeSymlink != 0:
		// 	fmt.Println("symbolic link")
		// case mode&fs.ModeNamedPipe != 0:
		// 	fmt.Println("named pipe")
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
