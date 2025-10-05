package endpoints

import (
	"context"
	"domic.domain/commons/contracts/interfaces"
	"domic.usecase/user/commands"
	"domic.usecase/user/queries"
	"encoding/json"
	userRpc "github.com/HasanProgrammer/Domic-GrpcService-Go/UserService/compile/service"
)

type UserServer struct {
	userRpc.UnimplementedUserServiceServer

	UnitOfWork  interfaces.IUnitOfWork
	IdGenerator interfaces.IIdentityGenerator
}

func (server *UserServer) CheckExist(ctx context.Context, req *userRpc.CheckExistRequest) (*userRpc.CheckExistResponse, error) {

	query := queries.CheckExistQuery{Id: req.TargetId.GetValue()}

	queryHandler := queries.CheckExistQueryHandler{UnitOfWork: server.UnitOfWork}

	result := queryHandler.Handle(&query, ctx)

	return &userRpc.CheckExistResponse{Result: result.Result}, nil

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

	var roles []string

	errRoles := json.Unmarshal([]byte(req.Roles.GetValue()), &roles)

	if errRoles != nil {
		return nil, errRoles
	}

	var permissions []string

	err := json.Unmarshal([]byte(req.Permissions.GetValue()), &permissions)

	if err != nil {
		return nil, err
	}

	command := commands.CreateUserCommand{
		FirstName:   req.FirstName.GetValue(),
		LastName:    req.LastName.GetValue(),
		Username:    req.Username.GetValue(),
		Password:    req.Password.GetValue(),
		PhoneNumber: req.PhoneNumber.GetValue(),
		EMail:       req.Email.GetValue(),
		Description: req.Description.GetValue(),
		Roles:       roles,
		Permissions: permissions,
	}

	handler := commands.NewCreateUserCommandHandler(server.UnitOfWork, server.IdGenerator)

	handler.Handle(&command, ctx)

	return nil, nil

}

func (server *UserServer) Active(ctx context.Context, req *userRpc.ActiveRequest) (*userRpc.ActiveResponse, error) {
	return nil, nil
}

func (server *UserServer) InActive(ctx context.Context, req *userRpc.InActiveRequest) (*userRpc.InActiveResponse, error) {
	return nil, nil
}
