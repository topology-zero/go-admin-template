syntax = "v1"

type (
    {{title .table}}ListRequest {
        Page int `form:"page"`         // 分页
        PageSize int `form:"pageSize"` // 每页条数
    }

    {{title .table}}ListResponse {
        Page int `json:"page"`         // 分页
        PageSize int `json:"pageSize"` // 每页条数
        Total int `json:"total"`       // 总条数
        Data []{{title .table}}List `json:"data"`  // 具体数据
    }

    {{title .table}}List {
        {{- if .hasId }}
        Id int `json:"id"`
        {{- end}}
        {{- range .cols}}
        {{title .Name}} {{.GoType}} `json:"{{camel .Name}}"` // {{.Comment}}
        {{- end}}
    }

    {{title .table}}DetailRequest {
        Id int `path:"id"`
    }

    {{title .table}}DetailResponse {
        {{- if .hasId }}
        Id int `json:"id"`
        {{- end}}
        {{- range .cols}}
        {{title .Name}} {{.GoType}} `json:"{{camel .Name}}"` // {{.Comment}}
        {{- end}}
    }

    {{title .table}}AddRequest {
        {{- range .cols}}
        {{title .Name}} {{.GoType}} `json:"{{camel .Name}}" binding:"required"` // {{.Comment}}
        {{- end}}
    }

    {{title .table}}EditRequest {
        Id int `path:"id"`
        {{- range .cols}}
        {{title .Name}} {{.GoType}} `json:"{{camel .Name}}" binding:"required"` // {{.Comment}}
        {{- end}}
    }

    {{title .table}}DeleteRequest {
        Id int `path:"id"`
    }
)

@server(
    jwt: Jwt
    middleware: Auth
    group: {{title .table}}
    swtags: {{.name}}相关
)

service go-admin-template {

    @doc "{{.name}}列表"
    @handler list
    get /{{.table}} ({{title .table}}ListRequest) returns ({{title .table}}ListResponse)

    @doc "{{.name}}详情"
    @handler detail
    get /{{.table}}/:id ({{title .table}}DetailRequest) returns ({{title .table}}DetailResponse)

    @doc "添加{{.name}}"
    @handler add
    post /{{.table}} ({{title .table}}AddRequest)

    @doc "编辑{{.name}}"
    @handler edit
    put /{{.table}}/:id ({{title .table}}EditRequest)

    @doc "删除{{.name}}"
    @handler del
    delete /{{.table}}/:id ({{title .table}}DeleteRequest)
}
