<template>
    <div class="app-container">
        <p class="page-title">{{.name}}列表</p>
        <div class="filter-container">
            <el-form ref="search"
                     :model="params"
                     inline
                     size="small">
                {{- range .cols}}
                <el-form-item label="{{.Comment}}" prop="{{.Name}}">
                    <el-input v-model="params.{{.Name}}"
                              placeholder="请输入{{.Comment}}"
                              clearable />
                </el-form-item>
                {{- end}}
                <el-form-item>
                    <el-button type="primary"
                               @click="search()">搜索</el-button>
                    <el-button @click="$refs.search.resetFields()">重置</el-button>
                </el-form-item>
            </el-form>
            <el-button v-permission="`{{.table}}:add`"
                       type="primary"
                       plain
                       class="filter-item"
                       @click="handleAdd()">
                添加{{.name}}
            </el-button>
        </div>
        <el-table v-loading="table_loading"
                  element-loading-text="加载中..."
                  border
                  :data="list">
            {{- if .hasId }}
            <el-table-column label="ID"
                             prop="id"
                             align="center"
                             width="60px" />
            {{- end}}
            {{- range .cols}}
            <el-table-column align="center"
                             label="{{.Comment}}"
                             prop="{{.Name}}" />
            {{- end}}
            <el-table-column align="center"
                             width="170px"
                             label="操作">
                <template slot-scope="{row}">
                    <el-button v-permission="`{{.table}}:edit`"
                               size="mini"
                               type="primary"
                               @click="handleEdit(row)">编辑
                    </el-button>
                    <el-button v-if="row.id != 1"
                               v-permission="`{{.table}}:del`"
                               size="mini"
                               type="danger"
                               @click="handleDel(row)">
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="page-container">
            <el-pagination background
                           :current-page.sync="params.page"
                           :page-sizes="page_sizes"
                           :page-size="params.pageSize"
                           layout="total, sizes, prev, pager, next, jumper"
                           :total="total"
                           @size-change="handleSizeChange"
                           @current-change="handleCurrentChange" />
        </div>
        <el-dialog :title="formDataId == 0 ? '添加{{.name}}' : '编辑{{.name}}'"
                   :visible.sync="dialogVisible"
                   width="30%">
            <el-form ref="elForm"
                     :model="formData"
                     label-width="100px">
                {{- range .cols}}
                <el-form-item label="{{.Comment}}"
                              :rules="{ required: true, trigger: 'blur', message: '请输入{{.Comment}}' }"
                              prop="{{.Name}}">
                    <el-input v-model="formData.{{.Name}}" />
                </el-form-item>
                {{- end}}
            </el-form>
            <span slot="footer"
                  class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button type="primary"
                           @click="submitForm()">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
    const defaultFormData = () => {
        return {
            {{- range $i, $v := .cols}}
            {{.Name}}: ''{{if ne $i $.colsLen}},{{end}}
            {{- end}}
        }
    }
    import { getList, add, edit, del, getDetail } from './api'
    import { pageMixin } from '@/utils/mixin'
    export default {
        name: '{{.table}}',
        mixins: [pageMixin],
        data() {
            return {
                params: {
                    name: ''
                },
                dialogVisible: false,
                formDataId: 0,
                formData: defaultFormData(),
            }
        },
        methods: {
            {{.dd}} 添加{{.name}}
            handleAdd() {
                this.dialogVisible = true
                this.formDataId = 0
                this.formData = defaultFormData()
            },
            {{.dd}} 编辑{{.name}}
            async handleEdit(info) {
                this.dialogVisible = true
                this.formDataId = info.id
                const { data } = await getDetail(info.id)
                this.formData = data
            },
            {{.dd}} 添加 or 编辑{{.name}}提交
            async submitForm() {
                await this.$refs.elForm.validate()
                try {
                    if (this.formDataId == 0) {
                        await add(this.formData)
                    } else {
                        await edit(this.formDataId, this.formData)
                    }
                    this._getData()
                    this.$message.success('成功')
                } finally {
                    this.dialogVisible = false
                }
            },
            {{.dd}} 删除{{.name}}
            async handleDel(info) {
                await this.$confirm('删除用户不可恢复', '警告')
                const { message } = await del(info.id)
                this.$message.success(message)
                this._getData(this.params)
            },
            {{.dd}} 获取数据
            async _getData() {
                const res = await getList(this.params)
                this.list = res.data.data
                this.total = res.data.total
                this.pageSize = res.data.per_page
                this.page = res.data.current_page
                this.table_loading = false
            }
        }
    }
</script>
