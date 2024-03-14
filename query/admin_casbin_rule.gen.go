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

	"go-admin-template/model"
)

func newAdminCasbinRuleModel(db *gorm.DB, opts ...gen.DOOption) adminCasbinRuleModel {
	_adminCasbinRuleModel := adminCasbinRuleModel{}

	_adminCasbinRuleModel.adminCasbinRuleModelDo.UseDB(db, opts...)
	_adminCasbinRuleModel.adminCasbinRuleModelDo.UseModel(&model.AdminCasbinRuleModel{})

	tableName := _adminCasbinRuleModel.adminCasbinRuleModelDo.TableName()
	_adminCasbinRuleModel.ALL = field.NewAsterisk(tableName)
	_adminCasbinRuleModel.ID = field.NewInt64(tableName, "id")
	_adminCasbinRuleModel.Ptype = field.NewString(tableName, "ptype")
	_adminCasbinRuleModel.V0 = field.NewString(tableName, "v0")
	_adminCasbinRuleModel.V1 = field.NewString(tableName, "v1")
	_adminCasbinRuleModel.V2 = field.NewString(tableName, "v2")
	_adminCasbinRuleModel.V3 = field.NewString(tableName, "v3")
	_adminCasbinRuleModel.V4 = field.NewString(tableName, "v4")
	_adminCasbinRuleModel.V5 = field.NewString(tableName, "v5")
	_adminCasbinRuleModel.V6 = field.NewString(tableName, "v6")
	_adminCasbinRuleModel.V7 = field.NewString(tableName, "v7")

	_adminCasbinRuleModel.fillFieldMap()

	return _adminCasbinRuleModel
}

// adminCasbinRuleModel casbin 权限管理
type adminCasbinRuleModel struct {
	adminCasbinRuleModelDo adminCasbinRuleModelDo

	ALL   field.Asterisk
	ID    field.Int64
	Ptype field.String
	V0    field.String
	V1    field.String
	V2    field.String
	V3    field.String
	V4    field.String
	V5    field.String
	V6    field.String
	V7    field.String

	fieldMap map[string]field.Expr
}

func (a adminCasbinRuleModel) Table(newTableName string) *adminCasbinRuleModel {
	a.adminCasbinRuleModelDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminCasbinRuleModel) As(alias string) *adminCasbinRuleModel {
	a.adminCasbinRuleModelDo.DO = *(a.adminCasbinRuleModelDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminCasbinRuleModel) updateTableName(table string) *adminCasbinRuleModel {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Ptype = field.NewString(table, "ptype")
	a.V0 = field.NewString(table, "v0")
	a.V1 = field.NewString(table, "v1")
	a.V2 = field.NewString(table, "v2")
	a.V3 = field.NewString(table, "v3")
	a.V4 = field.NewString(table, "v4")
	a.V5 = field.NewString(table, "v5")
	a.V6 = field.NewString(table, "v6")
	a.V7 = field.NewString(table, "v7")

	a.fillFieldMap()

	return a
}

func (a *adminCasbinRuleModel) WithContext(ctx context.Context) IAdminCasbinRuleModelDo {
	return a.adminCasbinRuleModelDo.WithContext(ctx)
}

func (a adminCasbinRuleModel) TableName() string { return a.adminCasbinRuleModelDo.TableName() }

func (a adminCasbinRuleModel) Alias() string { return a.adminCasbinRuleModelDo.Alias() }

func (a adminCasbinRuleModel) Columns(cols ...field.Expr) gen.Columns {
	return a.adminCasbinRuleModelDo.Columns(cols...)
}

func (a *adminCasbinRuleModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminCasbinRuleModel) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["ptype"] = a.Ptype
	a.fieldMap["v0"] = a.V0
	a.fieldMap["v1"] = a.V1
	a.fieldMap["v2"] = a.V2
	a.fieldMap["v3"] = a.V3
	a.fieldMap["v4"] = a.V4
	a.fieldMap["v5"] = a.V5
	a.fieldMap["v6"] = a.V6
	a.fieldMap["v7"] = a.V7
}

func (a adminCasbinRuleModel) clone(db *gorm.DB) adminCasbinRuleModel {
	a.adminCasbinRuleModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminCasbinRuleModel) replaceDB(db *gorm.DB) adminCasbinRuleModel {
	a.adminCasbinRuleModelDo.ReplaceDB(db)
	return a
}

type adminCasbinRuleModelDo struct{ gen.DO }

type IAdminCasbinRuleModelDo interface {
	gen.SubQuery
	Debug() IAdminCasbinRuleModelDo
	WithContext(ctx context.Context) IAdminCasbinRuleModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminCasbinRuleModelDo
	WriteDB() IAdminCasbinRuleModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminCasbinRuleModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminCasbinRuleModelDo
	Not(conds ...gen.Condition) IAdminCasbinRuleModelDo
	Or(conds ...gen.Condition) IAdminCasbinRuleModelDo
	Select(conds ...field.Expr) IAdminCasbinRuleModelDo
	Where(conds ...gen.Condition) IAdminCasbinRuleModelDo
	Order(conds ...field.Expr) IAdminCasbinRuleModelDo
	Distinct(cols ...field.Expr) IAdminCasbinRuleModelDo
	Omit(cols ...field.Expr) IAdminCasbinRuleModelDo
	Join(table schema.Tabler, on ...field.Expr) IAdminCasbinRuleModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminCasbinRuleModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminCasbinRuleModelDo
	Group(cols ...field.Expr) IAdminCasbinRuleModelDo
	Having(conds ...gen.Condition) IAdminCasbinRuleModelDo
	Limit(limit int) IAdminCasbinRuleModelDo
	Offset(offset int) IAdminCasbinRuleModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminCasbinRuleModelDo
	Unscoped() IAdminCasbinRuleModelDo
	Create(values ...*model.AdminCasbinRuleModel) error
	CreateInBatches(values []*model.AdminCasbinRuleModel, batchSize int) error
	Save(values ...*model.AdminCasbinRuleModel) error
	First() (*model.AdminCasbinRuleModel, error)
	Take() (*model.AdminCasbinRuleModel, error)
	Last() (*model.AdminCasbinRuleModel, error)
	Find() ([]*model.AdminCasbinRuleModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminCasbinRuleModel, err error)
	FindInBatches(result *[]*model.AdminCasbinRuleModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminCasbinRuleModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminCasbinRuleModelDo
	Assign(attrs ...field.AssignExpr) IAdminCasbinRuleModelDo
	Joins(fields ...field.RelationField) IAdminCasbinRuleModelDo
	Preload(fields ...field.RelationField) IAdminCasbinRuleModelDo
	FirstOrInit() (*model.AdminCasbinRuleModel, error)
	FirstOrCreate() (*model.AdminCasbinRuleModel, error)
	FindByPage(offset int, limit int) (result []*model.AdminCasbinRuleModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminCasbinRuleModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminCasbinRuleModelDo) Debug() IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Debug())
}

func (a adminCasbinRuleModelDo) WithContext(ctx context.Context) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminCasbinRuleModelDo) ReadDB() IAdminCasbinRuleModelDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminCasbinRuleModelDo) WriteDB() IAdminCasbinRuleModelDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminCasbinRuleModelDo) Session(config *gorm.Session) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminCasbinRuleModelDo) Clauses(conds ...clause.Expression) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminCasbinRuleModelDo) Returning(value interface{}, columns ...string) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminCasbinRuleModelDo) Not(conds ...gen.Condition) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminCasbinRuleModelDo) Or(conds ...gen.Condition) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminCasbinRuleModelDo) Select(conds ...field.Expr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminCasbinRuleModelDo) Where(conds ...gen.Condition) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminCasbinRuleModelDo) Order(conds ...field.Expr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminCasbinRuleModelDo) Distinct(cols ...field.Expr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminCasbinRuleModelDo) Omit(cols ...field.Expr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminCasbinRuleModelDo) Join(table schema.Tabler, on ...field.Expr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminCasbinRuleModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminCasbinRuleModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminCasbinRuleModelDo) Group(cols ...field.Expr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminCasbinRuleModelDo) Having(conds ...gen.Condition) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminCasbinRuleModelDo) Limit(limit int) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminCasbinRuleModelDo) Offset(offset int) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminCasbinRuleModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminCasbinRuleModelDo) Unscoped() IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminCasbinRuleModelDo) Create(values ...*model.AdminCasbinRuleModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminCasbinRuleModelDo) CreateInBatches(values []*model.AdminCasbinRuleModel, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminCasbinRuleModelDo) Save(values ...*model.AdminCasbinRuleModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminCasbinRuleModelDo) First() (*model.AdminCasbinRuleModel, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminCasbinRuleModel), nil
	}
}

func (a adminCasbinRuleModelDo) Take() (*model.AdminCasbinRuleModel, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminCasbinRuleModel), nil
	}
}

func (a adminCasbinRuleModelDo) Last() (*model.AdminCasbinRuleModel, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminCasbinRuleModel), nil
	}
}

func (a adminCasbinRuleModelDo) Find() ([]*model.AdminCasbinRuleModel, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminCasbinRuleModel), err
}

func (a adminCasbinRuleModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminCasbinRuleModel, err error) {
	buf := make([]*model.AdminCasbinRuleModel, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminCasbinRuleModelDo) FindInBatches(result *[]*model.AdminCasbinRuleModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminCasbinRuleModelDo) Attrs(attrs ...field.AssignExpr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminCasbinRuleModelDo) Assign(attrs ...field.AssignExpr) IAdminCasbinRuleModelDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminCasbinRuleModelDo) Joins(fields ...field.RelationField) IAdminCasbinRuleModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminCasbinRuleModelDo) Preload(fields ...field.RelationField) IAdminCasbinRuleModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminCasbinRuleModelDo) FirstOrInit() (*model.AdminCasbinRuleModel, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminCasbinRuleModel), nil
	}
}

func (a adminCasbinRuleModelDo) FirstOrCreate() (*model.AdminCasbinRuleModel, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminCasbinRuleModel), nil
	}
}

func (a adminCasbinRuleModelDo) FindByPage(offset int, limit int) (result []*model.AdminCasbinRuleModel, count int64, err error) {
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

func (a adminCasbinRuleModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminCasbinRuleModelDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminCasbinRuleModelDo) Delete(models ...*model.AdminCasbinRuleModel) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminCasbinRuleModelDo) withDO(do gen.Dao) *adminCasbinRuleModelDo {
	a.DO = *do.(*gen.DO)
	return a
}
