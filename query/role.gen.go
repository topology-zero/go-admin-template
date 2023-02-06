// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"admin_template/model"
)

func newRoleModel(db *gorm.DB, opts ...gen.DOOption) roleModel {
	_roleModel := roleModel{}

	_roleModel.roleModelDo.UseDB(db, opts...)
	_roleModel.roleModelDo.UseModel(&model.RoleModel{})

	tableName := _roleModel.roleModelDo.TableName()
	_roleModel.ALL = field.NewAsterisk(tableName)
	_roleModel.ID = field.NewInt(tableName, "id")
	_roleModel.Name = field.NewString(tableName, "name")
	_roleModel.Auth = field.NewString(tableName, "auth")
	_roleModel.CreatedAt = field.NewTime(tableName, "created_at")
	_roleModel.UpdatedAt = field.NewTime(tableName, "updated_at")
	_roleModel.DeletedAt = field.NewField(tableName, "deleted_at")

	_roleModel.fillFieldMap()

	return _roleModel
}

type roleModel struct {
	roleModelDo

	ALL       field.Asterisk
	ID        field.Int
	Name      field.String // 角色名
	Auth      field.String // 权限ID
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (r roleModel) Table(newTableName string) *roleModel {
	r.roleModelDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r roleModel) As(alias string) *roleModel {
	r.roleModelDo.DO = *(r.roleModelDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *roleModel) updateTableName(table string) *roleModel {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt(table, "id")
	r.Name = field.NewString(table, "name")
	r.Auth = field.NewString(table, "auth")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *roleModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *roleModel) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 6)
	r.fieldMap["id"] = r.ID
	r.fieldMap["name"] = r.Name
	r.fieldMap["auth"] = r.Auth
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r roleModel) clone(db *gorm.DB) roleModel {
	r.roleModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r roleModel) replaceDB(db *gorm.DB) roleModel {
	r.roleModelDo.ReplaceDB(db)
	return r
}

type roleModelDo struct{ gen.DO }

type IRoleModelDo interface {
	gen.SubQuery
	Debug() IRoleModelDo
	WithContext(ctx context.Context) IRoleModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRoleModelDo
	WriteDB() IRoleModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRoleModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRoleModelDo
	Not(conds ...gen.Condition) IRoleModelDo
	Or(conds ...gen.Condition) IRoleModelDo
	Select(conds ...field.Expr) IRoleModelDo
	Where(conds ...gen.Condition) IRoleModelDo
	Order(conds ...field.Expr) IRoleModelDo
	Distinct(cols ...field.Expr) IRoleModelDo
	Omit(cols ...field.Expr) IRoleModelDo
	Join(table schema.Tabler, on ...field.Expr) IRoleModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRoleModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRoleModelDo
	Group(cols ...field.Expr) IRoleModelDo
	Having(conds ...gen.Condition) IRoleModelDo
	Limit(limit int) IRoleModelDo
	Offset(offset int) IRoleModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleModelDo
	Unscoped() IRoleModelDo
	Create(values ...*model.RoleModel) error
	CreateInBatches(values []*model.RoleModel, batchSize int) error
	Save(values ...*model.RoleModel) error
	First() (*model.RoleModel, error)
	Take() (*model.RoleModel, error)
	Last() (*model.RoleModel, error)
	Find() ([]*model.RoleModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RoleModel, err error)
	FindInBatches(result *[]*model.RoleModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RoleModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRoleModelDo
	Assign(attrs ...field.AssignExpr) IRoleModelDo
	Joins(fields ...field.RelationField) IRoleModelDo
	Preload(fields ...field.RelationField) IRoleModelDo
	FirstOrInit() (*model.RoleModel, error)
	FirstOrCreate() (*model.RoleModel, error)
	FindByPage(offset int, limit int) (result []*model.RoleModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRoleModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetRoleByAuthId(authId int) (result *model.RoleModel)
}

// GetRoleByAuthId 根据 auth_id 获取角色
//
// where(json_contains( auth, json_array( @authId )) )
func (r roleModelDo) GetRoleByAuthId(authId int) (result *model.RoleModel) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, authId)
	generateSQL.WriteString("json_contains( auth, json_array( ? )) ")

	var executeSQL *gorm.DB
	executeSQL = r.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	_ = executeSQL

	return
}

func (r roleModelDo) Debug() IRoleModelDo {
	return r.withDO(r.DO.Debug())
}

func (r roleModelDo) WithContext(ctx context.Context) IRoleModelDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roleModelDo) ReadDB() IRoleModelDo {
	return r.Clauses(dbresolver.Read)
}

func (r roleModelDo) WriteDB() IRoleModelDo {
	return r.Clauses(dbresolver.Write)
}

func (r roleModelDo) Session(config *gorm.Session) IRoleModelDo {
	return r.withDO(r.DO.Session(config))
}

func (r roleModelDo) Clauses(conds ...clause.Expression) IRoleModelDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roleModelDo) Returning(value interface{}, columns ...string) IRoleModelDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roleModelDo) Not(conds ...gen.Condition) IRoleModelDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roleModelDo) Or(conds ...gen.Condition) IRoleModelDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roleModelDo) Select(conds ...field.Expr) IRoleModelDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roleModelDo) Where(conds ...gen.Condition) IRoleModelDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roleModelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IRoleModelDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r roleModelDo) Order(conds ...field.Expr) IRoleModelDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roleModelDo) Distinct(cols ...field.Expr) IRoleModelDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roleModelDo) Omit(cols ...field.Expr) IRoleModelDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roleModelDo) Join(table schema.Tabler, on ...field.Expr) IRoleModelDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roleModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRoleModelDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roleModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IRoleModelDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roleModelDo) Group(cols ...field.Expr) IRoleModelDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roleModelDo) Having(conds ...gen.Condition) IRoleModelDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roleModelDo) Limit(limit int) IRoleModelDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roleModelDo) Offset(offset int) IRoleModelDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roleModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleModelDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roleModelDo) Unscoped() IRoleModelDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roleModelDo) Create(values ...*model.RoleModel) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roleModelDo) CreateInBatches(values []*model.RoleModel, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roleModelDo) Save(values ...*model.RoleModel) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roleModelDo) First() (*model.RoleModel, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleModel), nil
	}
}

func (r roleModelDo) Take() (*model.RoleModel, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleModel), nil
	}
}

func (r roleModelDo) Last() (*model.RoleModel, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleModel), nil
	}
}

func (r roleModelDo) Find() ([]*model.RoleModel, error) {
	result, err := r.DO.Find()
	return result.([]*model.RoleModel), err
}

func (r roleModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RoleModel, err error) {
	buf := make([]*model.RoleModel, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roleModelDo) FindInBatches(result *[]*model.RoleModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roleModelDo) Attrs(attrs ...field.AssignExpr) IRoleModelDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roleModelDo) Assign(attrs ...field.AssignExpr) IRoleModelDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roleModelDo) Joins(fields ...field.RelationField) IRoleModelDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roleModelDo) Preload(fields ...field.RelationField) IRoleModelDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roleModelDo) FirstOrInit() (*model.RoleModel, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleModel), nil
	}
}

func (r roleModelDo) FirstOrCreate() (*model.RoleModel, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RoleModel), nil
	}
}

func (r roleModelDo) FindByPage(offset int, limit int) (result []*model.RoleModel, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r roleModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roleModelDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roleModelDo) Delete(models ...*model.RoleModel) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roleModelDo) withDO(do gen.Dao) *roleModelDo {
	r.DO = *do.(*gen.DO)
	return r
}
