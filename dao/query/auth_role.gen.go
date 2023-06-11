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

func newAuthRole(db *gorm.DB, opts ...gen.DOOption) authRole {
	_authRole := authRole{}

	_authRole.authRoleDo.UseDB(db, opts...)
	_authRole.authRoleDo.UseModel(&models.AuthRole{})

	tableName := _authRole.authRoleDo.TableName()
	_authRole.ALL = field.NewAsterisk(tableName)
	_authRole.ID = field.NewInt64(tableName, "id")
	_authRole.CreatedAt = field.NewTime(tableName, "created_at")
	_authRole.UpdatedAt = field.NewTime(tableName, "updated_at")
	_authRole.DeletedAt = field.NewField(tableName, "deleted_at")
	_authRole.Fname = field.NewString(tableName, "fname")
	_authRole.Fnote = field.NewString(tableName, "fnote")

	_authRole.fillFieldMap()

	return _authRole
}

type authRole struct {
	authRoleDo

	ALL       field.Asterisk
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Fname     field.String
	Fnote     field.String

	fieldMap map[string]field.Expr
}

func (a authRole) Table(newTableName string) *authRole {
	a.authRoleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a authRole) As(alias string) *authRole {
	a.authRoleDo.DO = *(a.authRoleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *authRole) updateTableName(table string) *authRole {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Fname = field.NewString(table, "fname")
	a.Fnote = field.NewString(table, "fnote")

	a.fillFieldMap()

	return a
}

func (a *authRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *authRole) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["fname"] = a.Fname
	a.fieldMap["fnote"] = a.Fnote
}

func (a authRole) clone(db *gorm.DB) authRole {
	a.authRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a authRole) replaceDB(db *gorm.DB) authRole {
	a.authRoleDo.ReplaceDB(db)
	return a
}

type authRoleDo struct{ gen.DO }

type IAuthRoleDo interface {
	gen.SubQuery
	Debug() IAuthRoleDo
	WithContext(ctx context.Context) IAuthRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuthRoleDo
	WriteDB() IAuthRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuthRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuthRoleDo
	Not(conds ...gen.Condition) IAuthRoleDo
	Or(conds ...gen.Condition) IAuthRoleDo
	Select(conds ...field.Expr) IAuthRoleDo
	Where(conds ...gen.Condition) IAuthRoleDo
	Order(conds ...field.Expr) IAuthRoleDo
	Distinct(cols ...field.Expr) IAuthRoleDo
	Omit(cols ...field.Expr) IAuthRoleDo
	Join(table schema.Tabler, on ...field.Expr) IAuthRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuthRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuthRoleDo
	Group(cols ...field.Expr) IAuthRoleDo
	Having(conds ...gen.Condition) IAuthRoleDo
	Limit(limit int) IAuthRoleDo
	Offset(offset int) IAuthRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthRoleDo
	Unscoped() IAuthRoleDo
	Create(values ...*models.AuthRole) error
	CreateInBatches(values []*models.AuthRole, batchSize int) error
	Save(values ...*models.AuthRole) error
	First() (*models.AuthRole, error)
	Take() (*models.AuthRole, error)
	Last() (*models.AuthRole, error)
	Find() ([]*models.AuthRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AuthRole, err error)
	FindInBatches(result *[]*models.AuthRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.AuthRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuthRoleDo
	Assign(attrs ...field.AssignExpr) IAuthRoleDo
	Joins(fields ...field.RelationField) IAuthRoleDo
	Preload(fields ...field.RelationField) IAuthRoleDo
	FirstOrInit() (*models.AuthRole, error)
	FirstOrCreate() (*models.AuthRole, error)
	FindByPage(offset int, limit int) (result []*models.AuthRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuthRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a authRoleDo) Debug() IAuthRoleDo {
	return a.withDO(a.DO.Debug())
}

func (a authRoleDo) WithContext(ctx context.Context) IAuthRoleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authRoleDo) ReadDB() IAuthRoleDo {
	return a.Clauses(dbresolver.Read)
}

func (a authRoleDo) WriteDB() IAuthRoleDo {
	return a.Clauses(dbresolver.Write)
}

func (a authRoleDo) Session(config *gorm.Session) IAuthRoleDo {
	return a.withDO(a.DO.Session(config))
}

func (a authRoleDo) Clauses(conds ...clause.Expression) IAuthRoleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authRoleDo) Returning(value interface{}, columns ...string) IAuthRoleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authRoleDo) Not(conds ...gen.Condition) IAuthRoleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authRoleDo) Or(conds ...gen.Condition) IAuthRoleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authRoleDo) Select(conds ...field.Expr) IAuthRoleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authRoleDo) Where(conds ...gen.Condition) IAuthRoleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authRoleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAuthRoleDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a authRoleDo) Order(conds ...field.Expr) IAuthRoleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authRoleDo) Distinct(cols ...field.Expr) IAuthRoleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authRoleDo) Omit(cols ...field.Expr) IAuthRoleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authRoleDo) Join(table schema.Tabler, on ...field.Expr) IAuthRoleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuthRoleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuthRoleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authRoleDo) Group(cols ...field.Expr) IAuthRoleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authRoleDo) Having(conds ...gen.Condition) IAuthRoleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authRoleDo) Limit(limit int) IAuthRoleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authRoleDo) Offset(offset int) IAuthRoleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthRoleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authRoleDo) Unscoped() IAuthRoleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authRoleDo) Create(values ...*models.AuthRole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authRoleDo) CreateInBatches(values []*models.AuthRole, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authRoleDo) Save(values ...*models.AuthRole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authRoleDo) First() (*models.AuthRole, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthRole), nil
	}
}

func (a authRoleDo) Take() (*models.AuthRole, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthRole), nil
	}
}

func (a authRoleDo) Last() (*models.AuthRole, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthRole), nil
	}
}

func (a authRoleDo) Find() ([]*models.AuthRole, error) {
	result, err := a.DO.Find()
	return result.([]*models.AuthRole), err
}

func (a authRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AuthRole, err error) {
	buf := make([]*models.AuthRole, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authRoleDo) FindInBatches(result *[]*models.AuthRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authRoleDo) Attrs(attrs ...field.AssignExpr) IAuthRoleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authRoleDo) Assign(attrs ...field.AssignExpr) IAuthRoleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authRoleDo) Joins(fields ...field.RelationField) IAuthRoleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authRoleDo) Preload(fields ...field.RelationField) IAuthRoleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authRoleDo) FirstOrInit() (*models.AuthRole, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthRole), nil
	}
}

func (a authRoleDo) FirstOrCreate() (*models.AuthRole, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.AuthRole), nil
	}
}

func (a authRoleDo) FindByPage(offset int, limit int) (result []*models.AuthRole, count int64, err error) {
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

func (a authRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authRoleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authRoleDo) Delete(models ...*models.AuthRole) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authRoleDo) withDO(do gen.Dao) *authRoleDo {
	a.DO = *do.(*gen.DO)
	return a
}
