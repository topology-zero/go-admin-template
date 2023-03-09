package generate

import (
	"bytes"
	_ "embed"
	"html/template"
	"strings"

	"go-admin-template/dto"
	"go-admin-template/model"
	"go-admin-template/pkg/util"
	"go-admin-template/svc"
	"go-admin-template/types/admin/generate"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/mysql"
	_ "github.com/pingcap/tidb/types/parser_driver"
	driver "github.com/pingcap/tidb/types/parser_driver"
	"github.com/pkg/errors"
)

//go:embed _js.tpl
var jsTpl string

//go:embed _vue.tpl
var vueTpl string

//go:embed _api.tpl
var apiTpl string

type genFile struct {
	ctx       *svc.ServiceContext
	tableName string
	comment   string
	hasId     bool
	cols      []dto.Cols
}

// Generate 生成前端文件
func Generate(req *generate.GenerateRequest, ctx *svc.ServiceContext) (resp generate.GenerateResponse, err error) {
	g := genFile{
		ctx:       ctx,
		tableName: req.Table,
	}
	g.tableParse()

	resp.Js, err = g.jsParse()
	if err != nil {
		return
	}

	resp.Vue, err = g.vueParse()
	if err != nil {
		return
	}

	resp.Api, err = g.apiParse()
	return
}

// 解析 sql
func (g *genFile) tableParse() {
	var res dto.CreateTable
	model.DB().Raw("show create table " + g.tableName).Scan(&res)
	p := parser.New()
	stmt, _, _ := p.Parse(res.CreateTable, "", "")
	tableStmt := stmt[0].(*ast.CreateTableStmt)
	for _, option := range tableStmt.Options {
		if option.Tp == ast.TableOptionComment {
			g.comment = option.StrValue
		}
	}

	if len(g.comment) == 0 {
		g.comment = g.tableName
	}

	for _, col := range tableStmt.Cols {
		l := dto.Cols{Name: col.Name.Name.L}
		if l.Name == "id" {
			g.hasId = true
			continue
		}

		if filterCols(l.Name) {
			continue
		}

		for _, option := range col.Options {
			if option.Tp == ast.ColumnOptionComment {
				expr := option.Expr.(*driver.ValueExpr)
				l.Comment = expr.GetString()
			}
		}
		if len(l.Comment) == 0 {
			l.Comment = strings.ToUpper(l.Name)
		}

		_, l.GoType = getColumnType(col.Tp.Tp)

		g.cols = append(g.cols, l)
	}
}

// js 模板解析
func (g *genFile) jsParse() (string, error) {
	parse, err := template.New("js").Parse(jsTpl)
	if err != nil {
		g.ctx.Log.Errorf("%+v", errors.WithStack(err))
		return "", errors.New("模板解析错误")
	}

	data := map[string]string{
		"name": g.comment,
		"url":  g.tableName,
	}

	buffer := new(bytes.Buffer)
	err = parse.Execute(buffer, &data)
	if err != nil {
		g.ctx.Log.Errorf("%+v", errors.WithStack(err))
		return "", errors.New("模板解析错误")
	}

	return buffer.String(), nil
}

// vue 模板解析
func (g *genFile) vueParse() (string, error) {
	parse, err := template.New("vue").Parse(vueTpl)
	if err != nil {
		g.ctx.Log.Errorf("%+v", errors.WithStack(err))
		return "", errors.New("模板解析错误")
	}

	data := map[string]any{
		"name":        g.comment,
		"table":       g.tableName,
		"cols":        g.cols,
		"colsLen":     len(g.cols) - 1,
		"hasId":       g.hasId,
		"startScript": template.HTML("<script>"),
		"endScript":   template.HTML("</script>"),
	}

	buffer := new(bytes.Buffer)
	err = parse.Execute(buffer, data)
	if err != nil {
		g.ctx.Log.Errorf("%+v", errors.WithStack(err))
		return "", errors.New("模板解析错误")
	}

	return buffer.String(), nil
}

// api 模板解析
func (g *genFile) apiParse() (string, error) {
	parse, err := template.New("api").Funcs(template.FuncMap{
		"title": util.Title,
		"camel": util.Camel,
	}).Parse(apiTpl)
	if err != nil {
		g.ctx.Log.Errorf("%+v", errors.WithStack(err))
		return "", errors.New("模板解析错误")
	}

	data := map[string]any{
		"name":  g.comment,
		"table": g.tableName,
		"cols":  g.cols,
	}

	buffer := new(bytes.Buffer)
	err = parse.Execute(buffer, data)
	if err != nil {
		g.ctx.Log.Errorf("%+v", errors.WithStack(err))
		return "", errors.New("模板解析错误")
	}

	return buffer.String(), nil
}

func getColumnType(t byte) (mysqlType, goType string) {
	switch t {
	case mysql.TypeUnspecified, mysql.TypeNull, mysql.TypeNewDate:
		// 母鸡这些字段是干嘛的
		return "", "string"
	case mysql.TypeFloat, mysql.TypeDouble:
		return "float", "float64"
	case mysql.TypeTimestamp:
		return "timestamp", "string"
	case mysql.TypeDate:
		return "date", "string"
	case mysql.TypeDuration:
		return "time", "string"
	case mysql.TypeDatetime:
		return "datetime", "string"
	case mysql.TypeYear, mysql.TypeTiny, mysql.TypeLong, mysql.TypeShort, mysql.TypeLonglong, mysql.TypeInt24:
		return "int", "int"
	case mysql.TypeBit:
		return "bit", "string"
	case mysql.TypeJSON:
		return "json", "string"
	case mysql.TypeNewDecimal:
		return "decimal", "float64"
	case mysql.TypeEnum:
		return "enum", "string"
	case mysql.TypeSet:
		return "set", "string"
	case mysql.TypeTinyBlob, mysql.TypeMediumBlob, mysql.TypeLongBlob, mysql.TypeBlob:
		return "blob", "string"
	case mysql.TypeVarString, mysql.TypeVarchar, mysql.TypeString:
		return "string", "string"
	case mysql.TypeGeometry:
		return "point", "string"
	default:
		return "", "string"
	}
}

func filterCols(colName string) bool {
	del := []string{"create", "update", "delete"}
	for _, s := range del {
		if strings.HasPrefix(colName, s) {
			return true
		}
	}
	return false
}
