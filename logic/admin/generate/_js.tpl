import request from '@/utils/request'

// 获取{{.name}}列表
export function getList(params) {
    return request({
        url: '/{{.url}}',
        params
    })
}

// 获取{{.name}}详情
export function getDetail(id) {
    return request({
        url: '/{{.url}}/' + id,
        method: 'get'
    })
}

// 添加{{.name}}
export function add(data) {
    return request({
        url: '/{{.url}}',
        method: 'post',
        data
    })
}

// 编辑{{.name}}
export function edit(id, data) {
    return request({
        url: '/{{.url}}/' + id,
        method: 'put',
        data
    })
}

// 删除{{.name}}
export function del(id) {
    return request({
        url: '/{{.url}}/' + id,
        method: 'delete'
    })
}
