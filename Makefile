all:
	go get -u github.com/gobuffalo/packr/v2/packr2
	~/go/bin/packr2
	go build
	GOOS=windows CC=/usr/local/bin/x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build .
	~/go/bin/packr2 clean
