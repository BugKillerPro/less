package trace

import (
	"github.com/BugKillerPro/less/framework"
	"github.com/BugKillerPro/less/framework/contract"
)

type lessTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *lessTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewlessTraceService
}

// Boot will called when the service instantiate
func (provider *lessTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *lessTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *lessTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *lessTraceProvider) Name() string {
	return contract.TraceKey
}
