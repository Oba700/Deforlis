package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func catalogHTML(rows string, path string) string {
	body := fmt.Sprintf(`
<html>
<head>
<title>Catalog %s</title>
<link rel="icon" type="image/x-icon" href="/favicon.ico">
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
</head>
<body>
<div class="w3-container">
	<h1>%s</h1>
</div>
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
</html>`, path, path, rows)

	headers := fmt.Sprintf(`HTTP/1.1 200 OK
Server: deforlis/prealpha
Content-Type: text/html; charset=UTF-8
Content-Length: %d
`, len([]byte(body)))
	return headers + body + "\n"
}

func notFound(path string) []byte {
	var body []byte = []byte(fmt.Sprintf(`
<html>
<head>
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

func catalogEntrieHTML(e os.DirEntry, path string) string {
	info, _ := e.Info()
	var emoji string
	if e.Type().IsDir() {
		emoji = "📂"
	} else {
		emoji = "📃"
	}
	if path == "/" {
		path = ""
	}
	return fmt.Sprintf(`<tr>
	<td>
		%s
	</td>
	<td>
		<a href="%s/%s">%s</a>
	</td>
	<td>
		%d
	</td>
	<td>
		%s
	</td>
	</tr>
`, emoji, path, e.Name(), e.Name(), info.Size(), info.ModTime().UTC().Format(time.UnixDate))

}

// var mockHTML string = `HTTP/1.1 200 OK
// Server: deforlis/prealpha
// Content-Type: text/html; charset=UTF-8

// <html>
// <head>
// <title>Hello there</title>
// <link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
// </head>
// <body>
// <h1>Hello there</h1>
// <p>
// Prealpaha deforlis mock function here. Things seem wired Huh?
// </p>
// <br>Method: %s
// <br>Host: %s
// <br>Path: %s
// <br>Remote: %s
// <p/>
// <hr>
// </body>
// </html>

// `
func mockHTML(Method, Host, Path string, clientPort net.Addr) string {
	body := fmt.Sprintf(`<html>
<head>
<title>Hello there</title>
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
