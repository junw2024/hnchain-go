// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"hnchain/user/rpc/internal/logic"
	"hnchain/user/rpc/internal/svc"
	"hnchain/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 登录
func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 获得用户信息
func (s *UserServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

// 注册用户
func (s *UserServer) RegisterUser(ctx context.Context, in *user.RegisterUserReq) (*user.RegisterUserRes, error) {
	l := logic.NewRegisterUserLogic(ctx, s.svcCtx)
	return l.RegisterUser(in)
}

// 添加收获地址
func (s *UserServer) AddUserRevAddr(ctx context.Context, in *user.UserRevAddrAddReq) (*user.UserRevAddrAddRes, error) {
	l := logic.NewAddUserRevAddrLogic(ctx, s.svcCtx)
	return l.AddUserRevAddr(in)
}

// 编辑收获地址
func (s *UserServer) EditUserRevAddr(ctx context.Context, in *user.UserRevAddrEditReq) (*user.UserRevAddrEditRes, error) {
	l := logic.NewEditUserRevAddrLogic(ctx, s.svcCtx)
	return l.EditUserRevAddr(in)
}

// 删除收获地址
func (s *UserServer) DelUserRevAddr(ctx context.Context, in *user.UserRevAddrDelReq) (*user.UserCollectionDelRes, error) {
	l := logic.NewDelUserRevAddrLogic(ctx, s.svcCtx)
	return l.DelUserRevAddr(in)
}

// 获取收获地址列表
func (s *UserServer) GetUserRevAddrList(ctx context.Context, in *user.UserRevAddrListReq) (*user.UserRevAddrListRes, error) {
	l := logic.NewGetUserRevAddrListLogic(ctx, s.svcCtx)
	return l.GetUserRevAddrList(in)
}

// 添加收藏
func (s *UserServer) AddUserCollection(ctx context.Context, in *user.UserCollectionAddReq) (*user.UserCollectionAddRes, error) {
	l := logic.NewAddUserCollectionLogic(ctx, s.svcCtx)
	return l.AddUserCollection(in)
}

// 删除收藏
func (s *UserServer) DelUserCollection(ctx context.Context, in *user.UserCollectionDelReq) (*user.UserCollectionDelRes, error) {
	l := logic.NewDelUserCollectionLogic(ctx, s.svcCtx)
	return l.DelUserCollection(in)
}

// 收藏列表
func (s *UserServer) GetUserCollectionList(ctx context.Context, in *user.UserCollectionListReq) (*user.UserCollectionListRes, error) {
	l := logic.NewGetUserCollectionListLogic(ctx, s.svcCtx)
	return l.GetUserCollectionList(in)
}

// 根据主键id,查询收获地址
func (s *UserServer) GetUserRevAddrInfo(ctx context.Context, in *user.UserRevAddrInfoReq) (*user.UserRevAddr, error) {
	l := logic.NewGetUserRevAddrInfoLogic(ctx, s.svcCtx)
	return l.GetUserRevAddrInfo(in)
}
