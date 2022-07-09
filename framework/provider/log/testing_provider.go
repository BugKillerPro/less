package log

import (
	"github.com/BugKillerPro/less/framework"
	"github.com/BugKillerPro/less/framework/contract"
)

type lessTestingLogProvider struct {
}

// Register registe a new function for make a service instance
func (provider *lessTestingLogProvider) Register(c framework.Container) framework.NewInstance {
	return NewlessTestingLog
}

// Boot will called when the service instantiate
func (provider *lessTestingLogProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *lessTestingLogProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *lessTestingLogProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *lessTestingLogProvider) Name() string {
	return contract.LogKey
}
