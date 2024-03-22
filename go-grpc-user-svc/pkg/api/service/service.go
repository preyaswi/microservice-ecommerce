package service

import (
	"context"

	"github.com/preyaswi/go-grpc-user-svc/pkg/pb"
	interfaces "github.com/preyaswi/go-grpc-user-svc/pkg/usecase/interface"
	"github.com/preyaswi/go-grpc-user-svc/pkg/utils/models"
)
type UserServer struct{
	userUseCase interfaces.UserUseCase
	pb.UnimplementedUserServer
}
func NewUserServer(useCase interfaces.UserUseCase)pb.UserServer  {
	return &UserServer{
		userUseCase: useCase,
	}
	
}
func (u *UserServer)UserSignUp(ctx context.Context,usersignupdetails *pb.UserSignUpRequest) (*pb.UserSignUpResponse,error) {
	userCreateDetails:=models.UserSignUp{
		Firstname: usersignupdetails.Firstname,
		Lastname:  usersignupdetails.Lastname,
		Email:     usersignupdetails.Email,
		Phone:     usersignupdetails.Phone,
		Password:  usersignupdetails.Password,
	}
	data, err := u.userUseCase.UsersSignUp(userCreateDetails)
	if err != nil {
		return &pb.UserSignUpResponse{}, err
	}
	userDetails := &pb.UserDetails{Id: uint64(data.User.ID), Firstname: data.User.Firstname, Lastname: data.User.Lastname, Email: data.User.Email, Phone: data.User.Phone}
	return &pb.UserSignUpResponse{
		Status:       201,
		UserDetails:  userDetails,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}, nil
}
func (a *UserServer) UserLogin(ctx context.Context, loginDeatails *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	login := models.UserLogin{
		Email:    loginDeatails.Email,
		Password: loginDeatails.Password,
	}
	data, err := a.userUseCase.UsersLogin(login)
	if err != nil {
		return &pb.UserLoginResponse{}, err
	}
	userDetails := &pb.UserDetails{
		Id:        uint64(data.User.ID),
		Firstname: data.User.Firstname,
		Lastname:  data.User.Lastname,
		Email:     data.User.Email,
		Phone:     data.User.Phone,
	}
	return &pb.UserLoginResponse{
		Status:       201,
		UserDetails:  userDetails,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
	}, nil
}
