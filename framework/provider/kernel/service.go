package kernel

import (
	"github.com/BugKillerPro/less/framework/gin"
	"net/http"
)

// 引擎服务
type lessKernelService struct {
	engine *gin.Engine
}

// 初始化web引擎服务实例
func NewlessKernelService(params ...interface{}) (interface{}, error) {
	httpEngine := params[0].(*gin.Engine)
	return &lessKernelService{engine: httpEngine}, nil
}

// 返回web引擎
func (s *lessKernelService) HttpEngine() http.Handler {
	return s.engine
}
