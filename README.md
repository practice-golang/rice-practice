# Practice for go.rice

Bundling practice using [go.rice](https://github.com/GeertJohan/go.rice)

## Prepare and build

```sh
go get github.com/GeertJohan/go.rice/rice

./bin/rice.exe embed-go

go mod init
go mod tidy
go build
```

* `http://localhost:8080/` loads index.html which is not a bundle.
* `http://localhost:8080/www` loads index.html which is a bundle.

Source: https://www.thepolyglotdeveloper.com/2017/03/bundle-html-css-javascript-served-golang-application/
