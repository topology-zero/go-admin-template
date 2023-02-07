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

	"admin_template/model"
)

func newAuthModel(db *gorm.DB, opts ...gen.DOOption) authModel {
	_authModel := authModel{}

	_authModel.authModelDo.UseDB(db, opts...)
	_authModel.authModelDo.UseModel(&model.AuthModel{})

	tableName := _authModel.authModelDo.TableName()
	_authModel.ALL = field.NewAsterisk(tableName)
	_authModel.ID = field.NewInt(tableName, "id")
	_authModel.Pid = field.NewInt(tableName, "pid")
	_authModel.Name = field.NewString(tableName, "name")
	_authModel.Key = field.NewString(tableName, "key")
	_authModel.IsMenu = field.NewInt(tableName, "is_menu")
	_authModel.API = field.NewString(tableName, "api")
	_authModel.Action = field.NewString(tableName, "action")
	_authModel.CreateTime = field.NewTime(tableName, "create_time")
	_authModel.UpdateTime = field.NewTime(tableName, "update_time")
	_authModel.DeleteTime = field.NewField(tableName, "delete_time")

	_authModel.fillFieldMap()

	return _authModel
}

type authModel struct {
	authModelDo

	ALL        field.Asterisk
	ID         field.Int
	Pid        field.Int    // 上级ID
	Name       field.String // 节点名
	Key        field.String // 权限标识
	IsMenu     field.Int    // 是否是菜单栏 0：否 1：是
	API        field.String // 接口
	Action     field.String // 操作方法
	CreateTime field.Time
	UpdateTime field.Time
	DeleteTime field.Field

	fieldMap map[string]field.Expr
}

func (a authModel) Table(newTableName string) *authModel {
	a.authModelDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a authModel) As(alias string) *authModel {
	a.authModelDo.DO = *(a.authModelDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *authModel) updateTableName(table string) *authModel {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt(table, "id")
	a.Pid = field.NewInt(table, "pid")
	a.Name = field.NewString(table, "name")
	a.Key = field.NewString(table, "key")
	a.IsMenu = field.NewInt(table, "is_menu")
	a.API = field.NewString(table, "api")
	a.Action = field.NewString(table, "action")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.DeleteTime = field.NewField(table, "delete_time")

	a.fillFieldMap()

	return a
}

func (a *authModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *authModel) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 10)
	a.fieldMap["id"] = a.ID
	a.fieldMap["pid"] = a.Pid
	a.fieldMap["name"] = a.Name
	a.fieldMap["key"] = a.Key
	a.fieldMap["is_menu"] = a.IsMenu
	a.fieldMap["api"] = a.API
	a.fieldMap["action"] = a.Action
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["delete_time"] = a.DeleteTime
}

func (a authModel) clone(db *gorm.DB) authModel {
	a.authModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a authModel) replaceDB(db *gorm.DB) authModel {
	a.authModelDo.ReplaceDB(db)
	return a
}

type authModelDo struct{ gen.DO }

type IAuthModelDo interface {
	gen.SubQuery
	Debug() IAuthModelDo
	WithContext(ctx context.Context) IAuthModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuthModelDo
	WriteDB() IAuthModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuthModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuthModelDo
	Not(conds ...gen.Condition) IAuthModelDo
	Or(conds ...gen.Condition) IAuthModelDo
	Select(conds ...field.Expr) IAuthModelDo
	Where(conds ...gen.Condition) IAuthModelDo
	Order(conds ...field.Expr) IAuthModelDo
	Distinct(cols ...field.Expr) IAuthModelDo
	Omit(cols ...field.Expr) IAuthModelDo
	Join(table schema.Tabler, on ...field.Expr) IAuthModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuthModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuthModelDo
	Group(cols ...field.Expr) IAuthModelDo
	Having(conds ...gen.Condition) IAuthModelDo
	Limit(limit int) IAuthModelDo
	Offset(offset int) IAuthModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthModelDo
	Unscoped() IAuthModelDo
	Create(values ...*model.AuthModel) error
	CreateInBatches(values []*model.AuthModel, batchSize int) error
	Save(values ...*model.AuthModel) error
	First() (*model.AuthModel, error)
	Take() (*model.AuthModel, error)
	Last() (*model.AuthModel, error)
	Find() ([]*model.AuthModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuthModel, err error)
	FindInBatches(result *[]*model.AuthModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AuthModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuthModelDo
	Assign(attrs ...field.AssignExpr) IAuthModelDo
	Joins(fields ...field.RelationField) IAuthModelDo
	Preload(fields ...field.RelationField) IAuthModelDo
	FirstOrInit() (*model.AuthModel, error)
	FirstOrCreate() (*model.AuthModel, error)
	FindByPage(offset int, limit int) (result []*model.AuthModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuthModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a authModelDo) Debug() IAuthModelDo {
	return a.withDO(a.DO.Debug())
}

func (a authModelDo) WithContext(ctx context.Context) IAuthModelDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authModelDo) ReadDB() IAuthModelDo {
	return a.Clauses(dbresolver.Read)
}

func (a authModelDo) WriteDB() IAuthModelDo {
	return a.Clauses(dbresolver.Write)
}

func (a authModelDo) Session(config *gorm.Session) IAuthModelDo {
	return a.withDO(a.DO.Session(config))
}

func (a authModelDo) Clauses(conds ...clause.Expression) IAuthModelDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authModelDo) Returning(value interface{}, columns ...string) IAuthModelDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authModelDo) Not(conds ...gen.Condition) IAuthModelDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authModelDo) Or(conds ...gen.Condition) IAuthModelDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authModelDo) Select(conds ...field.Expr) IAuthModelDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authModelDo) Where(conds ...gen.Condition) IAuthModelDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authModelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAuthModelDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a authModelDo) Order(conds ...field.Expr) IAuthModelDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authModelDo) Distinct(cols ...field.Expr) IAuthModelDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authModelDo) Omit(cols ...field.Expr) IAuthModelDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authModelDo) Join(table schema.Tabler, on ...field.Expr) IAuthModelDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuthModelDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuthModelDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authModelDo) Group(cols ...field.Expr) IAuthModelDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authModelDo) Having(conds ...gen.Condition) IAuthModelDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authModelDo) Limit(limit int) IAuthModelDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authModelDo) Offset(offset int) IAuthModelDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthModelDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authModelDo) Unscoped() IAuthModelDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authModelDo) Create(values ...*model.AuthModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authModelDo) CreateInBatches(values []*model.AuthModel, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authModelDo) Save(values ...*model.AuthModel) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authModelDo) First() (*model.AuthModel, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthModel), nil
	}
}

func (a authModelDo) Take() (*model.AuthModel, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthModel), nil
	}
}

func (a authModelDo) Last() (*model.AuthModel, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthModel), nil
	}
}

func (a authModelDo) Find() ([]*model.AuthModel, error) {
	result, err := a.DO.Find()
	return result.([]*model.AuthModel), err
}

func (a authModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuthModel, err error) {
	buf := make([]*model.AuthModel, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authModelDo) FindInBatches(result *[]*model.AuthModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authModelDo) Attrs(attrs ...field.AssignExpr) IAuthModelDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authModelDo) Assign(attrs ...field.AssignExpr) IAuthModelDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authModelDo) Joins(fields ...field.RelationField) IAuthModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authModelDo) Preload(fields ...field.RelationField) IAuthModelDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authModelDo) FirstOrInit() (*model.AuthModel, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthModel), nil
	}
}

func (a authModelDo) FirstOrCreate() (*model.AuthModel, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthModel), nil
	}
}

func (a authModelDo) FindByPage(offset int, limit int) (result []*model.AuthModel, count int64, err error) {
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

func (a authModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authModelDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authModelDo) Delete(models ...*model.AuthModel) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authModelDo) withDO(do gen.Dao) *authModelDo {
	a.DO = *do.(*gen.DO)
	return a
}
