// Code generated by Kitex v0.7.1. DO NOT EDIT.

package register

import (
	api "SimpleDouyin/user/kitex_gen/api"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Echo(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error)
	Register(ctx context.Context, rreq *api.RegisterRequest, callOptions ...callopt.Option) (r *api.RegisterResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRegisterClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRegisterClient struct {
	*kClient
}

func (p *kRegisterClient) Echo(ctx context.Context, req *api.Request, callOptions ...callopt.Option) (r *api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Echo(ctx, req)
}

func (p *kRegisterClient) Register(ctx context.Context, rreq *api.RegisterRequest, callOptions ...callopt.Option) (r *api.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, rreq)
}