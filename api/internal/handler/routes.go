// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	mall "hnchain/api/internal/handler/mall"
	order "hnchain/api/internal/handler/order"
	user "hnchain/api/internal/handler/user"
	"hnchain/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/cartList",
				Handler: mall.CartListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/categoryList",
				Handler: mall.CategoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/categorys",
				Handler: mall.CategorysHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/flashsale",
				Handler: mall.FlashSaleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/home/banner",
				Handler: mall.HomeBannerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/product/comment",
				Handler: mall.ProductCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/product/detail",
				Handler: mall.ProductDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommend",
				Handler: mall.RecommendHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/mall"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: order.AddOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: order.OrderListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/order"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/addCollection",
				Handler: user.UserCollectionAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addReveAddr",
				Handler: user.AddReveAddrHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delCollection",
				Handler: user.UserCollectionDelHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delRevAddress",
				Handler: user.DelRevAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/editRevAddr",
				Handler: user.EditRevAddrHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getCollectionList",
				Handler: user.UserCollectionListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/info",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/userRevAddrList",
				Handler: user.UserRevAddrListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/user"),
	)
}
