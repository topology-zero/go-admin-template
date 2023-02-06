package internal

import "admin_template/types/admin/base"

type AdminUser interface {

	// GetUserInfo 获取用户基本信息
	//
	// SELECT
	//		u.id,
	//		u.username,
	//		u.realname,
	//		u.phone,
	//		r.NAME AS rolename,
	//		( SELECT GROUP_CONCAT(`key`) FROM auth WHERE JSON_CONTAINS( r.auth, JSON_ARRAY(id))) as authkeys
	//	FROM
	//		admin_user u
	//		LEFT JOIN role r ON u.role_id = r.id
	//	WHERE
	//		u.id = @id
	GetUserInfo(id int) base.UserInfoResponse
}
