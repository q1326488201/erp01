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

func newAuthUserRole(db *gorm.DB, opts ...gen.DOOption) authUserRole {
	_authUserRole := authUserRole{}

	_authUserRole.authUserRoleDo.UseDB(db, opts...)
	_authUserRole.authUserRoleDo.UseModel(&models.AuthUserRole{})

	tableName := _authUserRole.authUserRoleDo.TableName()
	_authUserRole.ALL = field.NewAsterisk(tableName)
	_authUserRole.ID = field.NewInt64(tableName, "id")
	_authUserRole.CreatedAt = field.NewTime(tableName, "created_at")
	_authUserRole.UpdatedAt = field.NewTime(tableName, "updated_at")
	_authUserRole.DeletedAt = field.NewField(tableName, "deleted_at")
	_authUserRole.UserID = field.NewInt64(tableName, "userID")
	_authUserRole.RoleID = field.NewInt64(tableName, "roleID")

	_authUserRole.fillFieldMap()

	return _authUserRole
}

type authUserRole struct {
	authUserRoleDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserID    field.Int64
	RoleID    field.Int64

	fieldMap map[string]field.Expr
}

func (a authUserRole) Table(newTableName string) *authUserRole {
	a.authUserRoleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a authUserRole) As(alias string) *authUserRole {
	a.authUserRoleDo.DO = *(a.authUserRoleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *authUserRole) updateTableName(table string) *authUserRole {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.UserID = field.NewInt64(table, "userID")
	a.RoleID = field.NewInt64(table, "roleID")

	a.fillFieldMap()

	return a
}

func (a *authUserRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *authUserRole) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["userID"] = a.UserID
	a.fieldMap["roleID"] = a.RoleID
}

func (a authUserRole) clone(db *gorm.DB) authUserRole {
	a.authUserRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a authUserRole) replaceDB(db *gorm.DB) authUserRole {
	a.authUserRoleDo.ReplaceDB(db)
	return a
}

type authUserRoleDo struct{ gen.DO }

type IAuthUserRoleDo interface {
	gen.SubQuery
	Debug() IAuthUserRoleDo
	WithContext(ctx context.Context) IAuthUserRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuthUserRoleDo
	WriteDB() IAuthUserRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuthUserRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuthUserRoleDo
	Not(conds ...gen.Condition) IAuthUserRoleDo
	Or(conds ...gen.Condition) IAuthUserRoleDo
	Select(conds ...field.Expr) IAuthUserRoleDo
	Where(conds ...gen.Condition) IAuthUserRoleDo
	Order(conds ...field.Expr) IAuthUserRoleDo
	Distinct(cols ...field.Expr) IAuthUserRoleDo
	Omit(cols ...field.Expr) IAuthUserRoleDo
	Join(table schema.Tabler, on ...field.Expr) IAuthUserRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuthUserRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuthUserRoleDo
	Group(cols ...field.Expr) IAuthUserRoleDo
	Having(conds ...gen.Condition) IAuthUserRoleDo
	Limit(limit int) IAuthUserRoleDo
	Offset(offset int) IAuthUserRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthUserRoleDo
	Unscoped() IAuthUserRoleDo
	Create(values ...*models.AuthUserRole) error
	CreateInBatches(values []*models.AuthUserRole, batchSize int) error
	Save(values ...*models.AuthUserRole) error
	First() (*models.AuthUserRole, error)
	Take() (*models.AuthUserRole, error)
	Last() (*models.AuthUserRole, error)
	Find() ([]*models.AuthUserRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AuthUserRole, err error)
	FindInBatches(result *[]*models.AuthUserRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.AuthUserRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuthUserRoleDo
	Assign(attrs ...field.AssignExpr) IAuthUserRoleDo
	Joins(fields ...field.RelationField) IAuthUserRoleDo
	Preload(fields ...field.RelationField) IAuthUserRoleDo
	FirstOrInit() (*models.AuthUserRole, error)
	FirstOrCreate() (*models.AuthUserRole, error)
	FindByPage(offset int, limit int) (result []*models.AuthUserRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuthUserRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a authUserRoleDo) Debug() IAuthUserRoleDo {
	return a.withDO(a.DO.Debug())
}

func (a authUserRoleDo) WithContext(ctx context.Context) IAuthUserRoleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authUserRoleDo) ReadDB() IAuthUserRoleDo {
	return a.Clauses(dbresolver.Read)
}

func (a authUserRoleDo) WriteDB() IAuthUserRoleDo {
	return a.Clauses(dbresolver.Write)
}

func (a authUserRoleDo) Session(config *gorm.Session) IAuthUserRoleDo {
	return a.withDO(a.DO.Session(config))
}

func (a authUserRoleDo) Clauses(conds ...clause.Expression) IAuthUserRoleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authUserRoleDo) Returning(value interface{}, columns ...string) IAuthUserRoleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authUserRoleDo) Not(conds ...gen.Condition) IAuthUserRoleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authUserRoleDo) Or(conds ...gen.Condition) IAuthUserRoleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authUserRoleDo) Select(conds ...field.Expr) IAuthUserRoleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authUserRoleDo) Where(conds ...gen.Condition) IAuthUserRoleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authUserRoleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAuthUserRoleDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a authUserRoleDo) Order(conds ...field.Expr) IAuthUserRoleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authUserRoleDo) Distinct(cols ...field.Expr) IAuthUserRoleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authUserRoleDo) Omit(cols ...field.Expr) IAuthUserRoleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authUserRoleDo) Join(table schema.Tabler, on ...field.Expr) IAuthUserRoleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authUserRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuthUserRoleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authUserRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuthUserRoleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authUserRoleDo) Group(cols ...field.Expr) IAuthUserRoleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authUserRoleDo) Having(conds ...gen.Condition) IAuthUserRoleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authUserRoleDo) Limit(limit int) IAuthUserRoleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authUserRoleDo) Offset(offset int) IAuthUserRoleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authUserRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthUserRoleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authUserRoleDo) Unscoped() IAuthUserRoleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authUserRoleDo) Create(values ...*models.AuthUserRole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authUserRoleDo) CreateInBatches(values []*models.AuthUserRole, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authUserRoleDo) Save(values ...*models.AuthUserRole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authUserRoleDo) First() (*models.AuthUserRole, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUserRole), nil
	}
}

func (a authUserRoleDo) Take() (*models.AuthUserRole, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUserRole), nil
	}
}

func (a authUserRoleDo) Last() (*models.AuthUserRole, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUserRole), nil
	}
}

func (a authUserRoleDo) Find() ([]*models.AuthUserRole, error) {
	result, err := a.DO.Find()
	return result.([]*models.AuthUserRole), err
}

func (a authUserRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AuthUserRole, err error) {
	buf := make([]*models.AuthUserRole, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authUserRoleDo) FindInBatches(result *[]*models.AuthUserRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authUserRoleDo) Attrs(attrs ...field.AssignExpr) IAuthUserRoleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authUserRoleDo) Assign(attrs ...field.AssignExpr) IAuthUserRoleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authUserRoleDo) Joins(fields ...field.RelationField) IAuthUserRoleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authUserRoleDo) Preload(fields ...field.RelationField) IAuthUserRoleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authUserRoleDo) FirstOrInit() (*models.AuthUserRole, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUserRole), nil
	}
}

func (a authUserRoleDo) FirstOrCreate() (*models.AuthUserRole, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthUserRole), nil
	}
}

func (a authUserRoleDo) FindByPage(offset int, limit int) (result []*models.AuthUserRole, count int64, err error) {
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

func (a authUserRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authUserRoleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authUserRoleDo) Delete(models ...*models.AuthUserRole) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authUserRoleDo) withDO(do gen.Dao) *authUserRoleDo {
	a.DO = *do.(*gen.DO)
	return a
}
