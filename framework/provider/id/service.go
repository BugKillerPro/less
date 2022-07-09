package id

import (
	"github.com/rs/xid"
)

type lessIDService struct {
}

func NewlessIDService(params ...interface{}) (interface{}, error) {
	return &lessIDService{}, nil
}

func (s *lessIDService) NewID() string {
	return xid.New().String()
}
