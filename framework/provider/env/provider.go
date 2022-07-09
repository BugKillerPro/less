package env

import (
	"github.com/BugKillerPro/less/framework"
	"github.com/BugKillerPro/less/framework/contract"
)

type LessEnvProvider struct {
	Folder string
}

// Register registe a new function for make a service instance
func (provider *LessEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewlessEnv
}

// Boot will called when the service instantiate
func (provider *LessEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *LessEnvProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *LessEnvProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

/// Name define the name for this service
func (provider *LessEnvProvider) Name() string {
	return contract.EnvKey
}
