package echo

import (
	"context"
	"fmt"
	"github.com/dk-sirius/example/cmd/e-grpc/api"
)

type EchoServie struct {
	Port string
	api.UnimplementedEchoServer
}

func (e *EchoServie) Say(ctx context.Context, in *api.RequestEcho) (*api.ReplyEcho, error) {
	return &api.ReplyEcho{
		Name: fmt.Sprintf("Hello %v in Port %v", in.GetName(), e.Port),
	}, nil
}
