# 插入父级权限
INSERT INTO `go-admin-template`.`admin_auth` (`id`, `pid`, `name`, `key`, `is_menu`, `api`, `action`)
VALUES ({{.maxId}}, 0, '{{.name}}', '{{.url}}', 1, '', '');

# 插入子级列表权限
INSERT INTO `go-admin-template`.`admin_auth` (`pid`, `name`, `key`, `is_menu`, `api`, `action`)
VALUES ({{.maxId}}, '{{.name}}列表', '{{.url}}:list', 0, '/{{.url}}', 'get');

# 插入子级详情权限
INSERT INTO `go-admin-template`.`admin_auth` (`pid`, `name`, `key`, `is_menu`, `api`, `action`)
VALUES ({{.maxId}}, '{{.name}}详情', '{{.url}}:detail', 0, '/{{.url}}/:id', 'get');

# 插入子级添加权限
INSERT INTO `go-admin-template`.`admin_auth` (`pid`, `name`, `key`, `is_menu`, `api`, `action`)
VALUES ({{.maxId}}, '新增{{.name}}', '{{.url}}:add', 0, '/{{.url}}', 'post');

# 插入子级编辑权限
INSERT INTO `go-admin-template`.`admin_auth` (`pid`, `name`, `key`, `is_menu`, `api`, `action`)
VALUES ({{.maxId}}, '编辑{{.name}}', '{{.url}}:edit', 0, '/{{.url}}/:id', 'put');

# 插入子级删除权限
INSERT INTO `go-admin-template`.`admin_auth` (`pid`, `name`, `key`, `is_menu`, `api`, `action`)
VALUES ({{.maxId}}, '删除{{.name}}', '{{.url}}:del', 0, '/{{.url}}/:id', 'delete');
