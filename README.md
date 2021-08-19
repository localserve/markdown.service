# markdown.service

A markdown to html service in Go

```shell

go run service.go

curl -d "body=# abc" -X "POST" http://localhost:40001/markdown
<h1>abc</h1>

curl -d "body=# abc" -X "POST" http://localhost:40001/markdown
<h2>abc</h2>

```
