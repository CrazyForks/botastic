// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"github.com/pandodao/botastic/core"
)

func newApp(db *gorm.DB, opts ...gen.DOOption) app {
	_app := app{}

	_app.appDo.UseDB(db, opts...)
	_app.appDo.UseModel(&core.App{})

	tableName := _app.appDo.TableName()
	_app.ALL = field.NewAsterisk(tableName)
	_app.ID = field.NewUint64(tableName, "id")
	_app.AppID = field.NewString(tableName, "app_id")
	_app.AppSecret = field.NewString(tableName, "app_secret")
	_app.AppSecretEncrypted = field.NewString(tableName, "app_secret_encrypted")
	_app.CreatedAt = field.NewTime(tableName, "created_at")
	_app.UpdatedAt = field.NewTime(tableName, "updated_at")
	_app.DeletedAt = field.NewTime(tableName, "deleted_at")

	_app.fillFieldMap()

	return _app
}

type app struct {
	appDo

	ALL                field.Asterisk
	ID                 field.Uint64
	AppID              field.String
	AppSecret          field.String
	AppSecretEncrypted field.String
	CreatedAt          field.Time
	UpdatedAt          field.Time
	DeletedAt          field.Time

	fieldMap map[string]field.Expr
}

func (a app) Table(newTableName string) *app {
	a.appDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a app) As(alias string) *app {
	a.appDo.DO = *(a.appDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *app) updateTableName(table string) *app {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint64(table, "id")
	a.AppID = field.NewString(table, "app_id")
	a.AppSecret = field.NewString(table, "app_secret")
	a.AppSecretEncrypted = field.NewString(table, "app_secret_encrypted")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewTime(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *app) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *app) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["app_secret"] = a.AppSecret
	a.fieldMap["app_secret_encrypted"] = a.AppSecretEncrypted
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a app) clone(db *gorm.DB) app {
	a.appDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a app) replaceDB(db *gorm.DB) app {
	a.appDo.ReplaceDB(db)
	return a
}

type appDo struct{ gen.DO }

type IAppDo interface {
	WithContext(ctx context.Context) IAppDo

	GetApp(ctx context.Context, id uint64) (result *core.App, err error)
	GetApps(ctx context.Context) (result []*core.App, err error)
	GetAppByAppID(ctx context.Context, appID string) (result *core.App, err error)
	CreateApp(ctx context.Context, appID string, appSecretEncrypted string) (result uint64, err error)
	UpdateAppSecret(ctx context.Context, id uint64, appSecretEncrypted string) (err error)
}

// SELECT
//
//	"id", "app_id", "app_secret_encrypted", "created_at", "updated_at"
//
// FROM @@table WHERE
//
//	"id"=@id AND "deleted_at" IS NULL
//
// LIMIT 1
func (a appDo) GetApp(ctx context.Context, id uint64) (result *core.App, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT \"id\", \"app_id\", \"app_secret_encrypted\", \"created_at\", \"updated_at\" FROM apps WHERE \"id\"=? AND \"deleted_at\" IS NULL LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT
//
//	"id", "app_id", "app_secret_encrypted", "created_at", "updated_at"
//
// FROM @@table WHERE
//
//	"deleted_at" IS NULL
func (a appDo) GetApps(ctx context.Context) (result []*core.App, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT \"id\", \"app_id\", \"app_secret_encrypted\", \"created_at\", \"updated_at\" FROM apps WHERE \"deleted_at\" IS NULL ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT
//
//	"id", "app_id", "app_secret_encrypted", "created_at", "updated_at"
//
// FROM @@table WHERE
//
//	"app_id"=@appID AND "deleted_at" IS NULL
//
// LIMIT 1
func (a appDo) GetAppByAppID(ctx context.Context, appID string) (result *core.App, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, appID)
	generateSQL.WriteString("SELECT \"id\", \"app_id\", \"app_secret_encrypted\", \"created_at\", \"updated_at\" FROM apps WHERE \"app_id\"=? AND \"deleted_at\" IS NULL LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// INSERT INTO @@table
//
//	("app_id", "app_secret_encrypted", "created_at", "updated_at")
//
// VALUES
//
//	(@appID, @appSecretEncrypted, NOW(), NOW())
//
// ON CONFLICT ("app_id") DO NOTHING
// RETURNING "id"
func (a appDo) CreateApp(ctx context.Context, appID string, appSecretEncrypted string) (result uint64, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, appID)
	params = append(params, appSecretEncrypted)
	generateSQL.WriteString("INSERT INTO apps (\"app_id\", \"app_secret_encrypted\", \"created_at\", \"updated_at\") VALUES (?, ?, NOW(), NOW()) ON CONFLICT (\"app_id\") DO NOTHING RETURNING \"id\" ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// UPDATE @@table
//
//	{{set}}
//		"app_secret_encrypted"=@appSecretEncrypted,
//		"updated_at"=NOW()
//	{{end}}
//
// WHERE
//
//	"id"=@id
func (a appDo) UpdateAppSecret(ctx context.Context, id uint64, appSecretEncrypted string) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("UPDATE apps ")
	var setSQL0 strings.Builder
	params = append(params, appSecretEncrypted)
	setSQL0.WriteString("\"app_secret_encrypted\"=?, \"updated_at\"=NOW() ")
	helper.JoinSetBuilder(&generateSQL, setSQL0)
	params = append(params, id)
	generateSQL.WriteString("WHERE \"id\"=? ")

	var executeSQL *gorm.DB
	executeSQL = a.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (a appDo) WithContext(ctx context.Context) IAppDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a *appDo) withDO(do gen.Dao) *appDo {
	a.DO = *do.(*gen.DO)
	return a
}
