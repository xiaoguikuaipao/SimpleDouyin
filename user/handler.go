package main

import (
	api "SimpleDouyin/user/kitex_gen/api"
	"context"
)

// LoginImpl implements the last service interface defined in the IDL.
type LoginImpl struct{}

// Echo implements the LoginImpl interface.
func (s *LoginImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}

// Login implements the LoginImpl interface.
func (s *LoginImpl) Login(ctx context.Context, lreq *api.LoginRequest) (resp *api.LoginResponse, err error) {
	// TODO: Your code here...
	return
}
