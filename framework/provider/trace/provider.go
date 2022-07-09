package trace

import (
	"github.com/BugKillerPro/less/framework"
	"github.com/BugKillerPro/less/framework/contract"
)

type LessTraceProvider struct {
	c framework.Container
}

// Register registe a new function for make a service instance
func (provider *LessTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewlessTraceService
}

// Boot will called when the service instantiate
func (provider *LessTraceProvider) Boot(c framework.Container) error {
	provider.c = c
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *LessTraceProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *LessTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.c}
}

/// Name define the name for this service
func (provider *LessTraceProvider) Name() string {
	return contract.TraceKey
}
