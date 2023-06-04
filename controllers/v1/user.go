package v1

import (
	"context"
	"log"

	"github.com/golang/protobuf/proto"

	"UserServer/constant"
	"UserServer/postgresql"
	"UserServer/postgresql/model"
	"UserServer/proto/UserServer"
)

func (s *UserServe) CreateUser(ctx context.Context, req *UserServer.CreateUserReq) (*UserServer.CreateUserResp, error) {
	log.Println("get createUser request")
	var resp UserServer.CreateUserResp

	u := model.User{
		Account:  &req.Account,
		Password: &req.Password,
	}

	if err := u.Upsert(ctx, postgresql.PoolWr.Write()); err != nil {
		return nil, err
	}

	shopResp, err := CreateShop(ctx, *u.Account)
	if err != nil {
		return nil, err
	}

	byteD, _ := proto.Marshal(shopResp)

	proto.Unmarshal(byteD, &resp)

	return &resp, nil
}

func (s *UserServe) GetUser(ctx context.Context, req *UserServer.GetUserReq) (*UserServer.GetUserResp, error) {
	log.Println("get user request")
	var resp = UserServer.GetUserResp{
		Code: constant.SUCCESS,
		Msg:  "Ok",
	}

	u := model.User{
		Account:  &req.Account,
		Password: &req.Password,
	}
	log.Println("start get")
	_, err := u.Get(ctx, postgresql.PoolWr.Read())
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
