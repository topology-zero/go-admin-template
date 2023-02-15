/*
 Navicat Premium Data Transfer

 Source Server         : localhost-3309
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : localhost:3309
 Source Schema         : admin-template

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 07/02/2023 15:42:43
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `realname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `role_id` int(11) NULL DEFAULT NULL COMMENT '角色ID',
  `status` int(11) NULL DEFAULT 0 COMMENT '状态 0:未启用 1:正常',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台用户' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', 'admin', '$2a$10$jlxETsTS1zLajnxqhojtIuMdGHGvEX5vKlLHzKx4LLx4Qj2vujMKq', '10086', 1, 1, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for admin_auth
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth`;
CREATE TABLE `admin_auth`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NULL DEFAULT NULL COMMENT '上级ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点名',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `is_menu` int(11) NULL DEFAULT NULL COMMENT '是否是菜单栏 0：否 1：是',
  `api` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '接口',
  `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作方法',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `key`(`key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_auth
-- ----------------------------
INSERT INTO `admin_auth` VALUES (1, 0, '后台管理', 'admin', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (2, 1, '用户管理', 'admin:user', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (3, 2, '用户列表', 'admin:user:list', 0, '/user', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (4, 2, '添加用户', 'admin:user:add', 0, '/user', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (7, 1, '角色管理', 'admin:role', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (8, 7, '角色列表', 'admin:role:list', 0, '/role', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (9, 2, '编辑用户', 'admin:user:edit', 0, '/user/:id', 'put', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (10, 2, '获取用户详情', 'admin:user:detail', 0, '/user/:id', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (11, 2, '删除用户', 'admin:user:del', 0, '/user/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (12, 7, '角色详情', 'admin:role:detail', 0, '/role/:id', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (13, 7, '添加角色', 'admin:role:add', 0, '/role', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (14, 7, '编辑角色', 'admin:role:edit', 0, '/role/:id', 'put', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (15, 7, '删除角色', 'admin:role:del', 0, '/role/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (16, 1, '权限管理', 'admin:auth', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (17, 16, '权限列表', 'admin:auth:list', 0, '/auth', 'get', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (18, 16, '添加权限', 'admin:auth:add', 0, '/auth', 'post', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (19, 16, '编辑权限', 'admin:auth:edit', 0, '/auth/:id', 'put', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (20, 16, '删除权限', 'admin:auth:del', 0, '/auth/:id', 'delete', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (21, 0, '用户管理', 'member', 1, '', '', NULL, NULL, NULL);
INSERT INTO `admin_auth` VALUES (22, 0, '订单管理', 'order', 1, '', '', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for admin_casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `admin_casbin_rule`;
CREATE TABLE `admin_casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v7` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_admin_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'casbin 权限管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_casbin_rule
-- ----------------------------

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `auth` json NULL COMMENT '权限ID',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, '超级管理员', NULL, NULL, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
