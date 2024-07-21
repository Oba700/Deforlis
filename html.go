package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
)

func catalogHTML(rows string, path string) string {
	pathElements := strings.Split(path, "/")
	cwd := pathElements[len(pathElements)-1]
	if cwd == "" {
		cwd = "/"
	}
	body := fmt.Sprintf(`
<html>
<head>
<title>Catalog %s</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="icon" type="image/x-icon" href="/favicon.ico">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
</head>
<body>
<div class="w3-container">
	<h1>%s</h1>
</div>
<hr>
<table class="w3-table">
<tr>
<th/>
<th>Ім'я</th>
<th>Розмір</th>
<th>Змінено</th>
</tr>
%s
</table>
<hr>
</body>
</html>`, cwd, linkifyPath(path), rows)

	headers := fmt.Sprintf(`HTTP/1.1 200 OK
Server: deforlis/prealpha
Content-Type: text/html; charset=UTF-8
Content-Length: %d
`, len([]byte(body)))
	return headers + body + "\n"
}

func linkifyPath(path string) string {
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	pathElements := strings.Split(path, "/")
	lpath := "<a href=\"/\">🏠</a> / "
	if pathElements[0] != "" {
		for pei, pe := range pathElements {
			if pei == 0 {
				lpath += "<a href=\"/" + pe + "\">" + pe + "</a> / "
			} else {
				var cumpe string
				for _, subpe := range pathElements[0:pei] {
					cumpe += subpe + "/"
				}
				cumpe += pe + "/"
				lpath += "<a href=\"/" + cumpe + "\">" + pe + "</a> / "
			}
		}
	}

	return lpath
}

func notFound(path string) []byte {
	var body []byte = []byte(fmt.Sprintf(`
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Не знайдено</title>
<link rel="icon" type="image/x-icon" href="/favicon.ico">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
</head>
<body>
<div class="w3-container">
	<h1>404</h1>
	<p>Нічого не знайдено за шляхом %s</p>
</div>
</body>
</html>`, path))
	headers := fmt.Sprintf(`HTTP/1.1 404 Not Found
Server: deforlis/prealpha
Content-Type: text/html; charset=UTF-8
Content-Length: %d
`, len([]byte(body)))
	resp := append([]byte(headers), body...)
	return append(resp, []byte("\n")...)
}

func catalogEntrieHTML(path string, e os.DirEntry) string {
	info, _ := e.Info()
	var emoji string
	if e.Type().IsDir() {
		emoji = "📂"
	} else {
		emoji = "📃"
	}
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	link := "/"
	pathElements := strings.Split(path, "/")
	fmt.Println("pathElements[0] :" + pathElements[0])
	if pathElements[0] != "" {
		for _, pe := range pathElements {
			link += url.PathEscape(pe) + "/"
		}
	}
	link += url.PathEscape(e.Name())
	fmt.Println("link: " + link)
	return fmt.Sprintf(`<tr>
	<td>
		%s
	</td>
	<td>
		<a href="%s">%s</a>
	</td>
	<td>
		%d
	</td>
	<td>
		%s
	</td>
	</tr>
`, emoji, link, e.Name(), info.Size(), info.ModTime().UTC().Format(time.UnixDate))
}

func mockHTML(Method, Host, Path string, clientPort net.Addr) string {
	body := fmt.Sprintf(`<html>
<head>
<title>Hello there</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
</head>
<body>
<h1>Hello there</h1>
<p>
Prealpaha deforlis mock function here. Things seem wired Huh?
</p>
<br>Method: %s
<br>Host: %s
<br>Path: %s
<br>Remote: %s
<p/>
<hr>
</body>
</html>`, Method, Host, Path, clientPort)
	headers := fmt.Sprintf(`HTTP/1 200 OK
Server: deforlis/prealpha
Content-Type: text/html; charset=UTF-8
Content-Length: %d`, len([]byte(body))+2)
	return headers + "\n\n" + body + "\n\n"
}
