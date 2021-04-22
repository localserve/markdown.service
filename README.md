# markdown.service

A markdown to html service in Go

```shell

go run service.go

curl -d "body=# abc" -X "POST" http://localhost:40001/markdown
<h1>abc</h1>

```