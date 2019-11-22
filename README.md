# tin-mp

# install sqlite3 driver
```
go get -u github.com/mattn/go-sqlite3@v1.12.0
```

# building
```
go get -u github.com/gobuffalo/packr/v2/packr2
packr2 # collect boxes
go build
packr2 clean # do not check out packr files
```
