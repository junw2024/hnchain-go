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
	productOperationFieldNames          = builder.RawFieldNames(&ProductOperation{}, true)
	productOperationRows                = strings.Join(productOperationFieldNames, ",")
	productOperationRowsExpectAutoSet   = strings.Join(stringx.Remove(productOperationFieldNames), ",")
	productOperationRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(productOperationFieldNames, "id"))

	cachePublicProductOperationIdPrefix = "cache:public:productOperation:id:"
)

type (
	productOperationModel interface {
		Insert(ctx context.Context, data *ProductOperation) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProductOperation, error)
		Update(ctx context.Context, data *ProductOperation) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductOperationModel struct {
		sqlc.CachedConn
		table string
	}

	ProductOperation struct {
		Id         int64     `db:"id"`
		Productid  int64     `db:"productid"` // 运营商品状态 0-下线 1-上线
		Status     int64     `db:"status"`    // 商品ID
		Createtime time.Time `db:"createtime"`
		Updatetime time.Time `db:"updatetime"`
	}
)

func newProductOperationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProductOperationModel {
	return &defaultProductOperationModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."product_operation"`,
	}
}

func (m *defaultProductOperationModel) Delete(ctx context.Context, id int64) error {
	publicProductOperationIdKey := fmt.Sprintf("%s%v", cachePublicProductOperationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicProductOperationIdKey)
	return err
}

func (m *defaultProductOperationModel) FindOne(ctx context.Context, id int64) (*ProductOperation, error) {
	publicProductOperationIdKey := fmt.Sprintf("%s%v", cachePublicProductOperationIdPrefix, id)
	var resp ProductOperation
	err := m.QueryRowCtx(ctx, &resp, publicProductOperationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", productOperationRows, m.table)
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

func (m *defaultProductOperationModel) Insert(ctx context.Context, data *ProductOperation) (sql.Result, error) {
	idgenerator := distributedid.NewSnowflake(int64(1))
	data.Id = idgenerator.GenerateId()

	publicProductOperationIdKey := fmt.Sprintf("%s%v", cachePublicProductOperationIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5)", m.table, productOperationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Productid, data.Status, data.Createtime, data.Updatetime)
	}, publicProductOperationIdKey)
	return ret, err
}

func (m *defaultProductOperationModel) Update(ctx context.Context, data *ProductOperation) error {
	publicProductOperationIdKey := fmt.Sprintf("%s%v", cachePublicProductOperationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, productOperationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.Productid, data.Status, data.Createtime, data.Updatetime)
	}, publicProductOperationIdKey)
	return err
}

func (m *defaultProductOperationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicProductOperationIdPrefix, primary)
}

func (m *defaultProductOperationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", productOperationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductOperationModel) tableName() string {
	return m.table
}
