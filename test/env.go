package test

import (
	"github.com/BugKillerPro/less/framework"
	"github.com/BugKillerPro/less/framework/provider/app"
	"github.com/BugKillerPro/less/framework/provider/env"
)

const (
	BasePath = "/Users/bugkillerpro/Documents/UGit/blogdemo/"
)

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewlessContainer()
	// 绑定App服务提供者
	container.Bind(&app.LessAppProvider{BaseFolder: BasePath})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.LessTestingEnvProvider{})
	return container
}
