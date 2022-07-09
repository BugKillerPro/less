// Copyright 2021 bugkillerpro.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/BugKillerPro/less/app/console"
	"github.com/BugKillerPro/less/app/http"
	"github.com/BugKillerPro/less/framework"
	"github.com/BugKillerPro/less/framework/provider/app"
	"github.com/BugKillerPro/less/framework/provider/cache"
	"github.com/BugKillerPro/less/framework/provider/config"
	"github.com/BugKillerPro/less/framework/provider/distributed"
	"github.com/BugKillerPro/less/framework/provider/env"
	"github.com/BugKillerPro/less/framework/provider/id"
	"github.com/BugKillerPro/less/framework/provider/kernel"
	"github.com/BugKillerPro/less/framework/provider/log"
	"github.com/BugKillerPro/less/framework/provider/orm"
	"github.com/BugKillerPro/less/framework/provider/redis"
	"github.com/BugKillerPro/less/framework/provider/ssh"
	"github.com/BugKillerPro/less/framework/provider/trace"
)

func main() {
	// 初始化服务容器
	container := framework.NewlessContainer()
	// 绑定App服务提供者
	container.Bind(&app.lessAppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.lessEnvProvider{})
	container.Bind(&distributed.LocalDistributedProvider{})
	container.Bind(&config.lessConfigProvider{})
	container.Bind(&id.lessIDProvider{})
	container.Bind(&trace.lessTraceProvider{})
	container.Bind(&log.lessLogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&redis.RedisProvider{})
	container.Bind(&cache.lessCacheProvider{})
	container.Bind(&ssh.SSHProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.lessKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
