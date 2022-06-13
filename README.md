# less

less is a web framework written in Go (Golang).
less is more.

## Table of Contents

- [Installation](#background)
- [Quick start](#install)

## Installation

To install Gin package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed (**version 1.14+ is required**), then you can use the below Go command to install Gin.

```sh
$ go get -u https://github.com/BugKillerPro/less
```

2. Import it in your code:

```go
import "https://github.com/BugKillerPro/less"
```

3. (Optional) Import `net/http`. This is required for example if using constants such as `http.StatusOK`.

```go
import "net/http"
```

## Quick start

```sh
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import (
	"net/http"

	"https://github.com/BugKillerPro/less"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
