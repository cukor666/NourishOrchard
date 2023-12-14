/*
 Navicat Premium Data Transfer

 Source Server         : 本地测试
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : nourish_orchard

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 03/12/2023 19:20:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for fruits
-- ----------------------------
DROP TABLE IF EXISTS `fruits`;
CREATE TABLE `fruits`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `fruit_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `water` tinyint UNSIGNED NULL DEFAULT NULL,
  `sugar` tinyint UNSIGNED NULL DEFAULT NULL,
  `inventory` bigint NOT NULL,
  `price` double NOT NULL,
  `suppher_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_fruits_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fruits
-- ----------------------------
INSERT INTO `fruits` VALUES (1, '2023-11-16 16:44:27.905', '2023-11-16 16:44:27.905', NULL, '苹果', 20, 20, 300, 4.5, 4);
INSERT INTO `fruits` VALUES (2, '2023-11-16 16:45:34.898', '2023-11-16 16:45:34.898', NULL, '香蕉', 8, 60, 200, 4, 3);
INSERT INTO `fruits` VALUES (3, '2023-11-16 16:46:27.496', '2023-11-16 16:46:27.496', NULL, '橙子', 80, 90, 800, 3.5, 3);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '123456',
  `gender` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '男',
  `birthday` datetime(3) NULL DEFAULT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `promise` tinyint NOT NULL DEFAULT 1,
  `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name` ASC) USING BTREE,
  UNIQUE INDEX `name_2`(`name` ASC) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2023-10-24 13:52:36.214', '2023-11-11 17:45:50.725', NULL, 'Cukor', '123456', '男', '2010-11-15 00:00:00.000', '852146', '广东省广州市天河区', 3, 'CukorZhong');
INSERT INTO `users` VALUES (2, '2023-10-24 13:57:02.821', '2023-11-04 16:02:40.139', NULL, 'David', '54aa12', '男', '2010-03-16 00:00:00.000', '18877981101', '广东省深圳市', 2, '陶喆');
INSERT INTO `users` VALUES (3, '2023-10-24 14:19:16.403', '2023-11-08 16:12:26.559', NULL, 'Leehom', '1sa132', '男', '1974-01-16 00:00:00.000', '15278621140', '美国洛杉矶', 1, '五旬老汉王力宏');
INSERT INTO `users` VALUES (4, '2023-10-24 19:09:35.977', '2023-11-08 16:10:50.314', NULL, 'Khalil', '841aag2', '男', '2023-10-26 00:00:00.000', '18877451241', '云南省', 2, '云南老农方大同');
INSERT INTO `users` VALUES (5, '2023-10-24 19:11:29.202', '2023-11-11 19:11:34.748', NULL, 'Ziqi', 'dzq5512', '女', '2018-10-09 00:00:00.000', '18877107892', '中国香港', 1, '金鱼嘴');
INSERT INTO `users` VALUES (6, '2023-10-24 19:12:59.020', '2023-11-18 23:18:59.820', NULL, 'Wenwen', '23414', '女', '1993-10-21 00:00:00.000', '18277419632', '连市', 1, '大连美女于文文');
INSERT INTO `users` VALUES (7, '2023-10-24 19:20:33.402', '2023-10-27 16:07:16.879', NULL, 'Liangying', '5113410w', '女', '1985-06-27 00:00:00.000', '13788457521', '湖南长沙省', 2, '张靓颖');
INSERT INTO `users` VALUES (8, '2023-10-24 19:26:59.329', '2023-11-11 17:53:34.840', NULL, 'Candy', '887744', '男', '2022-10-16 00:00:00.000', '15913145621', '中国台湾省台北市', 1, '王心凌');
INSERT INTO `users` VALUES (9, '2023-11-05 20:17:41.973', '2023-11-11 19:10:52.566', NULL, 'Betty', '5431542', '女', '2023-11-07 00:00:00.000', '15461243510', '海南省海口市', 1, '吴宣仪');
INSERT INTO `users` VALUES (10, '2023-11-05 20:24:53.986', '2023-11-08 16:14:37.559', NULL, 'Ronghao', '852147', '男', '2023-11-10 00:00:00.000', '5664441234', '安徽省蚌埠市', 2, '小眼睛李荣浩');
INSERT INTO `users` VALUES (11, '2023-11-05 20:34:46.091', '2023-11-06 20:26:07.037', NULL, 'Zining', '987654', '女', '2023-11-07 00:00:00.000', '12345678910', '湖南省长沙市', 1, '紫宁');
INSERT INTO `users` VALUES (12, '2023-11-07 16:21:26.000', '2023-11-07 16:21:26.000', NULL, 'zhangsan', '123456', '男', '2023-11-07 00:00:00.000', '12345678910', '天津', 1, '老张');
INSERT INTO `users` VALUES (13, '2023-11-07 16:22:56.000', '2023-11-07 16:22:56.000', NULL, 'lisi', '123456', '男', '2023-11-07 16:22:56.000', '12345678910', '天津', 1, '老李');
INSERT INTO `users` VALUES (14, '2023-11-07 16:23:55.000', '2023-11-10 13:17:03.171', NULL, 'wangwu', '123456', '女', '2023-11-07 16:23:55.000', '10987654321', '天津', 1, '老王');
INSERT INTO `users` VALUES (15, '2023-11-07 16:24:23.000', '2023-11-07 16:24:23.000', NULL, 'zhaoliu', '123456', '男', '2023-11-07 16:24:23.000', '12345678910', '天津', 1, '老六');
INSERT INTO `users` VALUES (17, '2023-11-07 16:25:28.000', '2023-11-07 16:25:28.000', NULL, 'wenzhang', '123456', '男', '2023-11-07 16:25:28.000', '12345678910', '禅达', 2, '龙文章');
INSERT INTO `users` VALUES (18, '2023-11-07 16:26:00.000', '2023-11-07 16:26:00.000', NULL, 'fanliao', '123456', '男', '2023-11-07 16:26:00.000', '12345678910', '北京', 1, '孟烦了');
INSERT INTO `users` VALUES (19, '2023-11-07 16:26:31.000', '2023-11-07 16:26:31.000', NULL, 'milong', '123456', '男', '2023-11-07 16:26:31.000', '12345678910', '禅达', 1, '张迷龙');
INSERT INTO `users` VALUES (20, '2023-11-07 16:27:16.000', '2023-11-07 16:27:16.000', NULL, 'chaoyue', '123456', '女', '2023-11-07 16:27:16.000', '12345678910', '成都', 1, '啥也不会');
INSERT INTO `users` VALUES (21, '2023-11-07 17:44:49.000', '2023-11-07 17:44:49.000', NULL, 'Jay', '123456', '男', '2023-11-07 17:44:49.000', '12345678910', '中国台湾', 1, '第一名');
INSERT INTO `users` VALUES (22, '2023-11-10 12:02:38.713', '2023-11-10 12:02:38.713', NULL, 'Coco', '987654', '女', '1972-06-22 00:00:00.000', '15987329197', '湖北省武汉市', 2, '李玟');
INSERT INTO `users` VALUES (23, '2023-11-10 12:18:45.699', '2023-11-10 12:18:45.699', NULL, 'CaiXiao', '9874631', '女', '2023-11-20 00:00:00.000', '12345678910', '广西百色', 1, '彩肖');
INSERT INTO `users` VALUES (24, '2023-11-10 12:29:06.435', '2023-11-10 12:49:55.278', NULL, 'Yier', '123456', '女', '1992-10-02 18:20:21.000', '18864329812', '广东省汕头市', 2, '伊尔');
INSERT INTO `users` VALUES (28, '2023-11-10 12:38:57.990', '2023-11-10 12:52:51.080', NULL, 'Carl', '123456', '男', '1992-10-02 18:20:21.000', '13777269848', '广东省深圳市', 1, '代码随想录');

SET FOREIGN_KEY_CHECKS = 1;
