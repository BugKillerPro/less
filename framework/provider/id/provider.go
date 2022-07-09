package id

import (
	"github.com/BugKillerPro/less/framework"
	"github.com/BugKillerPro/less/framework/contract"
)

type lessIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *lessIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewlessIDService
}

// Boot will called when the service instantiate
func (provider *lessIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *lessIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *lessIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *lessIDProvider) Name() string {
	return contract.IDKey
}
