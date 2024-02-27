// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"hnchain/common/distributedid"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productCategoryFieldNames          = builder.RawFieldNames(&ProductCategory{}, true)
	productCategoryRows                = strings.Join(productCategoryFieldNames, ",")
	productCategoryRowsExpectAutoSet   = strings.Join(stringx.Remove(productCategoryFieldNames), ",")
	productCategoryRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(productCategoryFieldNames, "id"))

	cachePublicProductCategoryIdPrefix = "cache:public:productCategory:id:"
)

type (
	productCategoryModel interface {
		Insert(ctx context.Context, data *ProductCategory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductCategory, error)
		Update(ctx context.Context, data *ProductCategory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	ProductCategory struct {
		Id         int64     `db:"id"`
		Parentid   int64     `db:"parentid"` // 父类别id当id=0时说明是根节点,一级类别
		Name       string    `db:"name"`     // 类别名称
		Status     int64     `db:"status"`   // 类别状态1-正常,2-已废弃
		Createtime time.Time `db:"createtime"`
		Updatetime time.Time `db:"updatetime"`
	}
)

func newProductCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProductCategoryModel {
	return &defaultProductCategoryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."product_category"`,
	}
}

func (m *defaultProductCategoryModel) Delete(ctx context.Context, id int64) error {
	publicProductCategoryIdKey := fmt.Sprintf("%s%v", cachePublicProductCategoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicProductCategoryIdKey)
	return err
}

func (m *defaultProductCategoryModel) FindOne(ctx context.Context, id int64) (*ProductCategory, error) {
	publicProductCategoryIdKey := fmt.Sprintf("%s%v", cachePublicProductCategoryIdPrefix, id)
	var resp ProductCategory
	err := m.QueryRowCtx(ctx, &resp, publicProductCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", productCategoryRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductCategoryModel) Insert(ctx context.Context, data *ProductCategory) (sql.Result, error) {
	
	idgenerator := distributedid.NewSnowflake(int64(1))
	data.Id = idgenerator.GenerateId()
	
	publicProductCategoryIdKey := fmt.Sprintf("%s%v", cachePublicProductCategoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6)", m.table, productCategoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Parentid, data.Name, data.Status, data.Createtime, data.Updatetime)
	}, publicProductCategoryIdKey)
	return ret, err
}

func (m *defaultProductCategoryModel) Update(ctx context.Context, data *ProductCategory) error {
	publicProductCategoryIdKey := fmt.Sprintf("%s%v", cachePublicProductCategoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, productCategoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Parentid, data.Name, data.Status, data.Createtime, data.Updatetime)
	}, publicProductCategoryIdKey)
	return err
}

func (m *defaultProductCategoryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicProductCategoryIdPrefix, primary)
}

func (m *defaultProductCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", productCategoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductCategoryModel) tableName() string {
	
	return m.table
}
