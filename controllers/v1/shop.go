package v1

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"UserServer/config"
	"UserServer/proto/ShopServer"
)

func CreateShop(ctx context.Context, account string) (*ShopServer.CreateShopResp, error) {

	req := ShopServer.CreateShopReq{
		Account: account,
	}
	log.Println("shop url", config.ShopServerUrl)
	conn, err := grpc.Dial(config.ShopServerUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	shopResp, err := ShopServer.NewShopServerClient(conn).CreateShop(ctx, &req)
	if err != nil {
		return nil, err
	}
	return shopResp, nil
}
