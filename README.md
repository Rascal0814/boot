# boot

A Framework Based on Kratos Packaging

### Installation

```bash
go get -u github.com/Rascal0814/boot
```

### Quickstart

```go
go install github.com/Rascal0814/boot/tools/boot@latest
```

and enter the following in your terminal for getting help
```shell
boot --help

# generate app template 
boot new 

# generate database table crud
# --dsn is database conn like: mysql://user:password@tcp(xxx:3306)/xxx?charset=utf8&parseTime=True&loc=Local
# --pkg is you model pkg
# -o --output output file path. if not specified while use current dir to generate file

boot crud --dsn xxx --pkg xxx -o xxx

```
