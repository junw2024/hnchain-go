// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"hnchain/product/rpc/internal/logic"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

// 查询产品信息
func (s *ProductServer) Product(ctx context.Context, in *product.ProductItemReq) (*product.ProductItem, error) {
	l := logic.NewProductLogic(ctx, s.svcCtx)
	return l.Product(in)
}

// 查询集合的商品信息
func (s *ProductServer) Products(ctx context.Context, in *product.ProductReq) (*product.ProductRsp, error) {
	l := logic.NewProductsLogic(ctx, s.svcCtx)
	return l.Products(in)
}

// 查询指定分类的商品分页
func (s *ProductServer) ProductList(ctx context.Context, in *product.ProductListReq) (*product.ProductListRsp, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}

// 刷新运营产品缓存
func (s *ProductServer) OperationProducts(ctx context.Context, in *product.OperationProductsReq) (*product.OperationProductsRsp, error) {
	l := logic.NewOperationProductsLogic(ctx, s.svcCtx)
	return l.OperationProducts(in)
}

// 根据目标量，增加库存
func (s *ProductServer) UpdateProductStock(ctx context.Context, in *product.UpdateProductStockReq) (*product.UpdateProductStockRsp, error) {
	l := logic.NewUpdateProductStockLogic(ctx, s.svcCtx)
	return l.UpdateProductStock(in)
}

// 扣减库存(缓存)
func (s *ProductServer) CheckAndUpdateStock(ctx context.Context, in *product.CheckAndUpdateStockReq) (*product.CheckAndUpdateStockRsp, error) {
	l := logic.NewCheckAndUpdateStockLogic(ctx, s.svcCtx)
	return l.CheckAndUpdateStock(in)
}

// 检查库存量
func (s *ProductServer) CheckProductStock(ctx context.Context, in *product.UpdateProductStockReq) (*product.UpdateProductStockRsp, error) {
	l := logic.NewCheckProductStockLogic(ctx, s.svcCtx)
	return l.CheckProductStock(in)
}

func (s *ProductServer) RollbackProductStock(ctx context.Context, in *product.UpdateProductStockReq) (*product.UpdateProductStockRsp, error) {
	l := logic.NewRollbackProductStockLogic(ctx, s.svcCtx)
	return l.RollbackProductStock(in)
}

// 事务:扣减库存
func (s *ProductServer) DecrStock(ctx context.Context, in *product.DecrStockReq) (*product.DecrStockRsp, error) {
	l := logic.NewDecrStockLogic(ctx, s.svcCtx)
	return l.DecrStock(in)
}

// 事务:扣减库存,再加回去
func (s *ProductServer) DecrStockRevert(ctx context.Context, in *product.DecrStockReq) (*product.DecrStockRsp, error) {
	l := logic.NewDecrStockRevertLogic(ctx, s.svcCtx)
	return l.DecrStockRevert(in)
}

// 根据类父ID,查询分类列表
func (s *ProductServer) CategoryList(ctx context.Context, in *product.CategoryItemReq) (*product.CategoryItemRsp, error) {
	l := logic.NewCategoryListLogic(ctx, s.svcCtx)
	return l.CategoryList(in)
}

// 目标数量的推荐商品
func (s *ProductServer) ProductRecommends(ctx context.Context, in *product.ProductRecommendReq) (*product.ProductRecommendRsp, error) {
	l := logic.NewProductRecommendsLogic(ctx, s.svcCtx)
	return l.ProductRecommends(in)
}
