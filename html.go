package main

import (
	"fmt"
	"os"
	"time"
)

func catalogHTML(rows string, path string) string {
	body := fmt.Sprintf(`
<html>
<head>
<title>Catalog %s</title>
<link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css">
</head>
<body>
<div class="w3-container">
	<h1>%s</h1>
</div>
<table class="w3-table">
<tr>
<th/>
<th>Name</th>
<th>Size</th>
<th>Modified</th>
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

func notFound() []byte {
	var body []byte = []byte("\nSorry")
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
	return fmt.Sprintf(`<tr>
	<td>
		%s
	</td>
	<td>
		<a href="%s%s">%s</a>
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

var mockHTML string = `HTTP/1.1 200 OK
Server: deforlis/prealpha
Content-Type: text/html; charset=UTF-8


<html>
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
<p/>
<hr>
</body>
</html>

`
