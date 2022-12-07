# less

less is a web framework written in Go (Golang).
less is more.

## Installation

To install less package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed (**version 1.16+ is required**), then you can use the below Go command to install less.

```sh
$ go get -u https://github.com/BugKillerPro/less
```

2. Import it in your code:

```go
import "https://github.com/BugKillerPro/less"
```

## Quick start


```go
package main

import (
    "github.com/BugKillerPro/less/app/console"
    "github.com/BugKillerPro/less/app/http"
    "github.com/BugKillerPro/less/framework"
    "github.com/BugKillerPro/less/framework/provider/app"
    "github.com/BugKillerPro/less/framework/provider/kernel"
)

func main() {
    // 初始化服务容器
    container := framework.NewlessContainer()
    // 绑定App服务提供者
    container.Bind(&app.LessAppProvider{})

    // 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
    if engine, err := http.NewHttpEngine(container); err == nil {
        container.Bind(&kernel.LessKernelProvider{HttpEngine: engine})
    }

    // 运行root命令
    console.RunCommand(container)
}
```
