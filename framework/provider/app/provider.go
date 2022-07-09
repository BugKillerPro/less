package app

import (
	"github.com/BugKillerPro/less/framework"
	"github.com/BugKillerPro/less/framework/contract"
)

// LessAppProvider 提供App的具体实现方法
type LessAppProvider struct {
	BaseFolder string
}

// Register 注册lessApp方法
func (h *LessAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewlessApp
}

// Boot 启动调用
func (h *LessAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (h *LessAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (h *LessAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, h.BaseFolder}
}

// Name 获取字符串凭证
func (h *LessAppProvider) Name() string {
	return contract.AppKey
}
