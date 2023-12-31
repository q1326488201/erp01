// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"app/dao/models"
)

func newAuthUser(db *gorm.DB, opts ...gen.DOOption) authUser {
	_authUser := authUser{}

	_authUser.authUserDo.UseDB(db, opts...)
	_authUser.authUserDo.UseModel(&models.AuthUser{})

	tableName := _authUser.authUserDo.TableName()
	_authUser.ALL = field.NewAsterisk(tableName)
	_authUser.ID = field.NewInt64(tableName, "id")
	_authUser.CreatedAt = field.NewTime(tableName, "created_at")
	_authUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_authUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_authUser.Username = field.NewString(tableName, "username")
	_authUser.Password = field.NewString(tableName, "password")
	_authUser.Nickname = field.NewString(tableName, "nickname")
	_authUser.Ftype = field.NewInt64(tableName, "ftype")
	_authUser.Fnote = field.NewString(tableName, "fnote")
	_authUser.AvatarURL = field.NewString(tableName, "avatar_url")

	_authUser.fillFieldMap()

	return _authUser
}

type authUser struct {
	authUserDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Username  field.String
	Password  field.String
	Nickname  field.String
	Ftype     field.Int64
	Fnote     field.String
	AvatarURL field.String

	fieldMap map[string]field.Expr
}

func (a authUser) Table(newTableName string) *authUser {
	a.authUserDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a authUser) As(alias string) *authUser {
	a.authUserDo.DO = *(a.authUserDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *authUser) updateTableName(table string) *authUser {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Username = field.NewString(table, "username")
	a.Password = field.NewString(table, "password")
	a.Nickname = field.NewString(table, "nickname")
	a.Ftype = field.NewInt64(table, "ftype")
	a.Fnote = field.NewString(table, "fnote")
	a.AvatarURL = field.NewString(table, "avatar_url")

	a.fillFieldMap()

	return a
}

func (a *authUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *authUser) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["username"] = a.Username
	a.fieldMap["password"] = a.Password
	a.fieldMap["nickname"] = a.Nickname
	a.fieldMap["ftype"] = a.Ftype
	a.fieldMap["fnote"] = a.Fnote
	a.fieldMap["avatar_url"] = a.AvatarURL
}

func (a authUser) clone(db *gorm.DB) authUser {
	a.authUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a authUser) replaceDB(db *gorm.DB) authUser {
	a.authUserDo.ReplaceDB(db)
	return a
}

type authUserDo struct{ gen.DO }

type IAuthUserDo interface {
	gen.SubQuery
	Debug() IAuthUserDo
	WithContext(ctx context.Context) IAuthUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuthUserDo
	WriteDB() IAuthUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuthUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuthUserDo
	Not(conds ...gen.Condition) IAuthUserDo
	Or(conds ...gen.Condition) IAuthUserDo
	Select(conds ...field.Expr) IAuthUserDo
	Where(conds ...gen.Condition) IAuthUserDo
	Order(conds ...field.Expr) IAuthUserDo
	Distinct(cols ...field.Expr) IAuthUserDo
	Omit(cols ...field.Expr) IAuthUserDo
	Join(table schema.Tabler, on ...field.Expr) IAuthUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuthUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuthUserDo
	Group(cols ...field.Expr) IAuthUserDo
	Having(conds ...gen.Condition) IAuthUserDo
	Limit(limit int) IAuthUserDo
	Offset(offset int) IAuthUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthUserDo
	Unscoped() IAuthUserDo
	Create(values ...*models.AuthUser) error
	CreateInBatches(values []*models.AuthUser, batchSize int) error
	Save(values ...*models.AuthUser) error
	First() (*models.AuthUser, error)
	Take() (*models.AuthUser, error)
	Last() (*models.AuthUser, error)
	Find() ([]*models.AuthUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AuthUser, err error)
	FindInBatches(result *[]*models.AuthUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.AuthUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuthUserDo
	Assign(attrs ...field.AssignExpr) IAuthUserDo
	Joins(fields ...field.RelationField) IAuthUserDo
	Preload(fields ...field.RelationField) IAuthUserDo
	FirstOrInit() (*models.AuthUser, error)
	FirstOrCreate() (*models.AuthUser, error)
	FindByPage(offset int, limit int) (result []*models.AuthUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuthUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a authUserDo) Debug() IAuthUserDo {
	return a.withDO(a.DO.Debug())
}

func (a authUserDo) WithContext(ctx context.Context) IAuthUserDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authUserDo) ReadDB() IAuthUserDo {
	return a.Clauses(dbresolver.Read)
}

func (a authUserDo) WriteDB() IAuthUserDo {
	return a.Clauses(dbresolver.Write)
}

func (a authUserDo) Session(config *gorm.Session) IAuthUserDo {
	return a.withDO(a.DO.Session(config))
}

func (a authUserDo) Clauses(conds ...clause.Expression) IAuthUserDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authUserDo) Returning(value interface{}, columns ...string) IAuthUserDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authUserDo) Not(conds ...gen.Condition) IAuthUserDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authUserDo) Or(conds ...gen.Condition) IAuthUserDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authUserDo) Select(conds ...field.Expr) IAuthUserDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authUserDo) Where(conds ...gen.Condition) IAuthUserDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAuthUserDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a authUserDo) Order(conds ...field.Expr) IAuthUserDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authUserDo) Distinct(cols ...field.Expr) IAuthUserDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authUserDo) Omit(cols ...field.Expr) IAuthUserDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authUserDo) Join(table schema.Tabler, on ...field.Expr) IAuthUserDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuthUserDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuthUserDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authUserDo) Group(cols ...field.Expr) IAuthUserDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authUserDo) Having(conds ...gen.Condition) IAuthUserDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authUserDo) Limit(limit int) IAuthUserDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authUserDo) Offset(offset int) IAuthUserDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthUserDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authUserDo) Unscoped() IAuthUserDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authUserDo) Create(values ...*models.AuthUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authUserDo) CreateInBatches(values []*models.AuthUser, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authUserDo) Save(values ...*models.AuthUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authUserDo) First() (*models.AuthUser, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUser), nil
	}
}

func (a authUserDo) Take() (*models.AuthUser, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUser), nil
	}
}

func (a authUserDo) Last() (*models.AuthUser, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUser), nil
	}
}

func (a authUserDo) Find() ([]*models.AuthUser, error) {
	result, err := a.DO.Find()
	return result.([]*models.AuthUser), err
}

func (a authUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AuthUser, err error) {
	buf := make([]*models.AuthUser, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authUserDo) FindInBatches(result *[]*models.AuthUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authUserDo) Attrs(attrs ...field.AssignExpr) IAuthUserDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authUserDo) Assign(attrs ...field.AssignExpr) IAuthUserDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authUserDo) Joins(fields ...field.RelationField) IAuthUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authUserDo) Preload(fields ...field.RelationField) IAuthUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authUserDo) FirstOrInit() (*models.AuthUser, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUser), nil
	}
}

func (a authUserDo) FirstOrCreate() (*models.AuthUser, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUser), nil
	}
}

func (a authUserDo) FindByPage(offset int, limit int) (result []*models.AuthUser, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a authUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authUserDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authUserDo) Delete(models ...*models.AuthUser) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authUserDo) withDO(do gen.Dao) *authUserDo {
	a.DO = *do.(*gen.DO)
	return a
}
