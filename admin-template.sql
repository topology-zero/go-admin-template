/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : 127.0.0.1:3306
 Source Schema         : admin-template

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 10/08/2022 19:43:21
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
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', 'admin', '$2a$10$QMKlSt.XRnsKXdNg4J2RnuyDyQT0TvPag/rzCjbc.uPjdX6yIorwm', '10086', 1, 1, NULL, '2022-08-10 17:15:57', NULL);

-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NULL DEFAULT NULL COMMENT '上级ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '节点名',
  `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `is_menu` int(11) NULL DEFAULT NULL COMMENT '是否是菜单栏 0：否 1：是',
  `api` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '接口',
  `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '操作方法',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `key`(`key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth
-- ----------------------------
INSERT INTO `auth` VALUES (1, 0, '后台管理', 'admin', '1', '', '', NULL, '2022-08-09 17:09:35', NULL);
INSERT INTO `auth` VALUES (2, 1, '用户管理', 'admin:user', '1', '', '', NULL, '2022-08-09 17:09:52', NULL);
INSERT INTO `auth` VALUES (3, 2, '用户列表', 'admin:user:list', '0', '/user', 'get', NULL, '2022-08-09 17:10:42', NULL);
INSERT INTO `auth` VALUES (4, 2, '添加用户', 'admin:user:add', '0', '/user', 'post', NULL, '2022-08-09 17:11:15', NULL);
INSERT INTO `auth` VALUES (7, 1, '角色管理', 'admin:role', '1', '', '', '2022-08-09 14:08:20', '2022-08-09 17:13:50', NULL);
INSERT INTO `auth` VALUES (8, 7, '角色列表', 'admin:role:list', '0', '/role', 'get', '2022-08-09 14:38:50', '2022-08-09 17:14:23', NULL);
INSERT INTO `auth` VALUES (9, 2, '编辑用户', 'admin:user:edit', '0', '/user/:id', 'put', '2022-08-09 17:12:16', '2022-08-09 17:12:16', NULL);
INSERT INTO `auth` VALUES (10, 2, '获取用户详情', 'admin:user:detail', '0', '/user/:id', 'get', '2022-08-09 17:12:45', '2022-08-09 17:12:45', NULL);
INSERT INTO `auth` VALUES (11, 2, '删除用户', 'admin:user:del', '0', '/user/:id', 'delete', '2022-08-09 17:13:01', '2022-08-09 17:13:01', NULL);
INSERT INTO `auth` VALUES (12, 7, '角色详情', 'admin:role:detail', '0', '/role/:id', 'get', '2022-08-09 17:14:56', '2022-08-09 17:14:56', NULL);
INSERT INTO `auth` VALUES (13, 7, '添加角色', 'admin:role:add', '0', '/role', 'post', '2022-08-09 17:15:17', '2022-08-09 17:15:17', NULL);
INSERT INTO `auth` VALUES (14, 7, '编辑角色', 'admin:role:edit', '0', '/role/:id', 'put', '2022-08-09 17:15:44', '2022-08-09 17:15:44', NULL);
INSERT INTO `auth` VALUES (15, 7, '删除角色', 'admin:role:del', '0', '/role/:id', 'delete', '2022-08-09 17:16:12', '2022-08-09 17:16:12', NULL);
INSERT INTO `auth` VALUES (16, 1, '权限管理', 'admin:auth', '1', '', '', '2022-08-09 17:16:41', '2022-08-09 17:16:41', NULL);
INSERT INTO `auth` VALUES (17, 16, '权限列表', 'admin:auth:list', '0', '/auth', 'get', '2022-08-09 17:17:10', '2022-08-09 17:17:10', NULL);
INSERT INTO `auth` VALUES (18, 16, '添加权限', 'admin:auth:add', '0', '/auth', 'post', '2022-08-09 17:17:35', '2022-08-09 17:17:35', NULL);
INSERT INTO `auth` VALUES (19, 16, '编辑权限', 'admin:auth:edit', '0', '/auth/:id', 'put', '2022-08-09 17:17:53', '2022-08-09 17:17:53', NULL);
INSERT INTO `auth` VALUES (20, 16, '删除权限', 'admin:auth:del', '0', '/auth/:id', 'delete', '2022-08-09 17:18:17', '2022-08-09 17:18:17', NULL);
INSERT INTO `auth` VALUES (21, 0, '用户管理', 'member', '1', '', '', '2022-08-10 14:38:44', '2022-08-10 14:38:44', NULL);
INSERT INTO `auth` VALUES (22, 0, '订单管理', 'order', '1', '', '', '2022-08-10 14:38:57', '2022-08-10 14:38:57', NULL);

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
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
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`, `v6`, `v7`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `auth` json NULL COMMENT '权限ID',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, '超级管理员', NULL, NULL, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
