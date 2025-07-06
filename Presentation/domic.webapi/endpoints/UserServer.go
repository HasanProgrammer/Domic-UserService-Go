package endpoints

import (
	"context"
	userRpc "github.com/HasanProgrammer/Domic-GrpcService-Go/UserService/compile/service"
)

type UserServer struct {
	userRpc.UnimplementedUserServiceServer
}

func (server *UserServer) CheckExist(ctx context.Context, req *userRpc.CheckExistRequest) (*userRpc.CheckExistResponse, error) {
	return nil, nil
}

func (server *UserServer) ReadOne(ctx context.Context, req *userRpc.ReadOneRequest) (*userRpc.ReadOneResponse, error) {
	return nil, nil
}

func (server *UserServer) ReadAllPaginated(ctx context.Context, req *userRpc.ReadAllPaginatedRequest) (*userRpc.ReadAllPaginatedResponse, error) {
	return nil, nil
}

func (server *UserServer) Create(ctx context.Context, req *userRpc.CreateRequest) (*userRpc.CreateResponse, error) {
	return nil, nil
}

func (server *UserServer) Update(ctx context.Context, req *userRpc.UpdateRequest) (*userRpc.UpdateResponse, error) {
	return nil, nil
}

func (server *UserServer) Active(ctx context.Context, req *userRpc.ActiveRequest) (*userRpc.ActiveResponse, error) {
	return nil, nil
}

func (server *UserServer) InActive(ctx context.Context, req *userRpc.InActiveRequest) (*userRpc.InActiveResponse, error) {
	return nil, nil
}
