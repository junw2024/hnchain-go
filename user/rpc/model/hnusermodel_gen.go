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
	hnuserFieldNames          = builder.RawFieldNames(&Hnuser{}, true)
	hnuserRows                = strings.Join(hnuserFieldNames, ",")
	hnuserRowsExpectAutoSet   = strings.Join(stringx.Remove(hnuserFieldNames), ",")
	hnuserRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(hnuserFieldNames, "id"))

	cachePublicHnuserIdPrefix       = "cache:public:hnuser:id:"
	cachePublicHnuserPhonePrefix    = "cache:public:hnuser:phone:"
	cachePublicHnuserUsernamePrefix = "cache:public:hnuser:username:"
)

type (
	hnuserModel interface {
		Insert(ctx context.Context, data *Hnuser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Hnuser, error)
		FindOneByPhone(ctx context.Context, phone string) (*Hnuser, error)
		FindOneByUsername(ctx context.Context, username string) (*Hnuser, error)
		Update(ctx context.Context, data *Hnuser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHnuserModel struct {
		sqlc.CachedConn
		table string
	}

	Hnuser struct {
		Id           int64          `db:"id"`
		Identitytype string         `db:"identitytype"` // 证件类型 0:身份证 1:护照 2:驾照
		Identity     string         `db:"identity"`     // 证件号码
		Name         string         `db:"name"`         // 名称
		Nick         string         `db:"nick"`         // 昵称
		Sex          string         `db:"sex"`          // 性别 0:未知 1:男 2:女
		Phone        string         `db:"phone"`
		Username     string         `db:"username"` // 用户名称
		Question     string         `db:"question"` // 找回密码问题
		Answer       string         `db:"answer"`   // 问题答案
		Password     string         `db:"password"`
		Address      string         `db:"address"`
		Email        string         `db:"email"`
		Wxopenid     string         `db:"wxopenid"`   // 微信openid
		Wxunionid    string         `db:"wxunionid"`  // 微信unionid
		Loginaddr    string         `db:"loginaddr"`  // 最后登录ip
		Createtime   time.Time      `db:"createtime"` // 创建时间
		Updatetime   time.Time      `db:"updatetime"` // 更新时间
	}
)

func newHnuserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultHnuserModel {
	return &defaultHnuserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."hnuser"`,
	}
}

func (m *defaultHnuserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicHnuserIdKey := fmt.Sprintf("%s%v", cachePublicHnuserIdPrefix, id)
	publicHnuserPhoneKey := fmt.Sprintf("%s%v", cachePublicHnuserPhonePrefix, data.Phone)
	publicHnuserUsernameKey := fmt.Sprintf("%s%v", cachePublicHnuserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicHnuserIdKey, publicHnuserPhoneKey, publicHnuserUsernameKey)
	return err
}

func (m *defaultHnuserModel) FindOne(ctx context.Context, id int64) (*Hnuser, error) {
	publicHnuserIdKey := fmt.Sprintf("%s%v", cachePublicHnuserIdPrefix, id)
	var resp Hnuser
	err := m.QueryRowCtx(ctx, &resp, publicHnuserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", hnuserRows, m.table)
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

func (m *defaultHnuserModel) FindOneByPhone(ctx context.Context, phone string) (*Hnuser, error) {
	publicHnuserPhoneKey := fmt.Sprintf("%s%v", cachePublicHnuserPhonePrefix, phone)
	var resp Hnuser
	err := m.QueryRowIndexCtx(ctx, &resp, publicHnuserPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where phone = $1 limit 1", hnuserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, phone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHnuserModel) FindOneByUsername(ctx context.Context, username string) (*Hnuser, error) {
	publicHnuserUsernameKey := fmt.Sprintf("%s%v", cachePublicHnuserUsernamePrefix, username)
	var resp Hnuser
	err := m.QueryRowIndexCtx(ctx, &resp, publicHnuserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where username = $1 limit 1", hnuserRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHnuserModel) Insert(ctx context.Context, data *Hnuser) (sql.Result, error) {

	idgenerator := distributedid.NewSnowflake(int64(1))
	data.Id = idgenerator.GenerateId()

	publicHnuserIdKey := fmt.Sprintf("%s%v", cachePublicHnuserIdPrefix, data.Id)
	publicHnuserPhoneKey := fmt.Sprintf("%s%v", cachePublicHnuserPhonePrefix, data.Phone)
	publicHnuserUsernameKey := fmt.Sprintf("%s%v", cachePublicHnuserUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)", m.table, hnuserRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Identitytype, data.Identity, data.Name, data.Nick, data.Sex, data.Phone, data.Username, data.Question, data.Answer, data.Password, data.Address, data.Email, data.Wxopenid, data.Wxunionid, data.Loginaddr, data.Createtime, data.Updatetime)
	}, publicHnuserIdKey, publicHnuserPhoneKey, publicHnuserUsernameKey)
	return ret, err
}

func (m *defaultHnuserModel) Update(ctx context.Context, newData *Hnuser) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicHnuserIdKey := fmt.Sprintf("%s%v", cachePublicHnuserIdPrefix, data.Id)
	publicHnuserPhoneKey := fmt.Sprintf("%s%v", cachePublicHnuserPhonePrefix, data.Phone)
	publicHnuserUsernameKey := fmt.Sprintf("%s%v", cachePublicHnuserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, hnuserRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Identitytype, newData.Identity, newData.Name, newData.Nick, newData.Sex, newData.Phone, newData.Username, newData.Question, newData.Answer, newData.Password, newData.Address, newData.Email, newData.Wxopenid, newData.Wxunionid, newData.Loginaddr, newData.Createtime, newData.Updatetime)
	}, publicHnuserIdKey, publicHnuserPhoneKey, publicHnuserUsernameKey)
	return err
}

func (m *defaultHnuserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicHnuserIdPrefix, primary)
}

func (m *defaultHnuserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", hnuserRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultHnuserModel) tableName() string {
	return m.table
}
