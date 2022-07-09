package env

import (
    "github.com/BugKillerPro/less/framework"
    "github.com/BugKillerPro/less/framework/contract"
)

type LessTestingEnvProvider struct {
    Folder string
}

// Register registe a new function for make a service instance
func (provider *LessTestingEnvProvider) Register(c framework.Container) framework.NewInstance {
    return NewlessTestingEnv
}

// Boot will called when the service instantiate
func (provider *LessTestingEnvProvider) Boot(c framework.Container) error {
    return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *LessTestingEnvProvider) IsDefer() bool {
    return false
}

// Params define the necessary params for NewInstance
func (provider *LessTestingEnvProvider) Params(c framework.Container) []interface{} {
    return []interface{}{}
}

/// Name define the name for this service
func (provider *LessTestingEnvProvider) Name() string {
    return contract.EnvKey
}
