# tin-mp

## install sqlite3 driver
```
go get -u github.com/mattn/go-sqlite3@v1.12.0
```

## running
```
go run .
```

## building
```
go get -u github.com/gobuffalo/packr/v2/packr2
packr2 # collect boxes
go build
packr2 clean # do not check out packr files
```

## cross building for windows
```
brew install mingw-w64
GOOS=windows CC=/usr/local/bin/x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build .
```