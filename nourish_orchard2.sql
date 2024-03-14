/*
 Navicat Premium Data Transfer

 Source Server         : 本地MySQL测试
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : nourish_orchard2

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 14/03/2024 15:44:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号名',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `promise` int NOT NULL COMMENT '权限，用户2458，员工3705，管理员8092',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`username`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account
-- ----------------------------
INSERT INTO `account` VALUES ('CZKJ111706348185', '$2a$10$Mdmy6WViAAvcbPYM7n8U2.xgbIOe1d2zLiN1Rw4F3xNqBRMVXreW.', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ11710304892', '$2a$10$/ZmmSH0y76WSNgsVAKZ.6uqdrs8hrrjiSPmqtgh7DGVOSKMn0Qh5W', 2458, '2024-03-13 12:41:32', '2024-03-13 12:41:32', NULL);
INSERT INTO `account` VALUES ('CZKJ141706348185', '$2a$10$YxcBKcNasmbu/sJr.7rTVOoWOg2A9eOeQ05NUmFZtlz6hTb80O1QW', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ151710304943', '$2a$10$PRznGzAPo9REHus3pFjc1eABn24vxpuKzsPnnKu7c17no6W0FQsmi', 2458, '2024-03-13 12:42:23', '2024-03-13 12:42:23', NULL);
INSERT INTO `account` VALUES ('CZKJ201709005881', '$2a$10$LeJBxeJvMcXxjQum88RrvOnoEPfe/5xs.twtvO0sE.6vyRBnojU7S', 2458, '2024-02-27 11:51:22', '2024-03-12 11:39:59', '2024-03-12 11:40:00');
INSERT INTO `account` VALUES ('CZKJ261706344878', '$2a$10$froOrMspKgquGAuwniGy5uGF3EJRp.RIVI9jdMBA2h.B.kQOi7go2', 8092, '2024-01-27 16:41:19', '2024-01-27 16:41:19', NULL);
INSERT INTO `account` VALUES ('CZKJ291706344878', '$2a$10$W.Sa7H7aKaZof6MRpfx7WOIr9juyakZQRvW4NE86dhpH1rMdD3.Fq', 8092, '2024-01-27 16:41:19', '2024-01-27 16:41:19', NULL);
INSERT INTO `account` VALUES ('CZKJ31706348185', '$2a$10$06j2m5cAzBnDCmV8Cs4LaO2Og/tTjRpM5OR9vWXV8faMMgcIZm.na', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ31706348186', '$2a$10$RuXlZEcHPqWSRpv/YDUc..SgPwEJQHYdXy5EEbeopmE5g5DGoKJzq', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ331706348186', '$2a$10$gM9lL11s48XOfgWdDfzkI.jmyw8g8dNeWQ3fYX5UTsHWCAqhSibnW', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ341708940332', '$2a$10$D8gQ1Aqvhg0TbQt2.i8RueITfX2wn4dgW1FlCFbwu6tG84/3lBmn2', 2458, '2024-02-26 17:38:53', '2024-03-12 11:33:21', NULL);
INSERT INTO `account` VALUES ('CZKJ351708869915', '$2a$10$.QzPeTzzYm.Sc0WpaS7p9OCYFryX9ugXNAdgFjFrwbQyWhO3SZQDG', 2458, '2024-02-25 22:05:16', '2024-03-12 11:26:48', '2024-03-12 11:26:49');
INSERT INTO `account` VALUES ('CZKJ381706348186', '$2a$10$aegw0iJnyFrHRegvVwuQGebFmbONHmS1kPCseNfHsvpOpDkxcyxTe', 3705, '2024-01-27 17:36:27', '2024-01-27 17:36:27', NULL);
INSERT INTO `account` VALUES ('CZKJ441710302534', '$2a$10$jIuu5FQBj7RN1gGgdhu5l.2ZW0y6y5cbpcaw.xFa45RkKjQeyAhW2', 2458, '2024-03-13 12:02:14', '2024-03-13 12:02:14', NULL);
INSERT INTO `account` VALUES ('CZKJ511710304763', '$2a$10$7/1zHKQ72h3BIfzFODp1H.GAsyyKuW2aU1onOEpb2OXTG0BYZW/X.', 2458, '2024-03-13 12:39:23', '2024-03-13 12:39:23', NULL);
INSERT INTO `account` VALUES ('CZKJ521706348185', '$2a$10$YQck8aC9a42QKTxz0JAC2usmRd9Fet19Lpx4Eu9Hi7t0hgbB9D6l.', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ531709006038', '$2a$10$T5c5qMfcj8JHICW6mwz/OuU6HbppCmZAnkp1zSoKtZ5Yx07uICcqm', 2458, '2024-02-27 11:53:58', '2024-03-11 17:45:18', '2024-03-11 17:45:19');
INSERT INTO `account` VALUES ('CZKJ531710305097', '$2a$10$cr07HcBy/8ceWLf34Hcvcu0ZHRjStNH4UukgzlqPvO1EradhbhZha', 2458, '2024-03-13 12:44:57', '2024-03-13 12:44:57', NULL);
INSERT INTO `account` VALUES ('CZKJ541710305023', '$2a$10$t8v6lh1kHP/KWRGoSZtgseXbx9xWhTgXY62ia7YvPrn.YewpY3cGq', 2458, '2024-03-13 12:43:44', '2024-03-13 12:43:44', NULL);
INSERT INTO `account` VALUES ('CZKJ551706348186', '$2a$10$eUZ7l0uJBrx5YMChbhQ8Pucz.gBv7hCV8WS7F4xBAdBBnj/xcgkDu', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ561710304527', '$2a$10$jxaDJkGoKgE0WdU25NA.MOQG6knJg0VQO9mLJQC3v7MgloQSS2UX6', 2458, '2024-03-13 12:35:27', '2024-03-13 12:35:27', NULL);
INSERT INTO `account` VALUES ('CZKJ571706348390', '$2a$10$TIaYOpBJkpOxP0V7FGuOKuBgk9fpQMMq5VCrcwNkUX7C8gQhipzIa', 3705, '2024-01-27 17:39:51', '2024-01-27 17:39:51', NULL);
INSERT INTO `account` VALUES ('CZKJ641706344878', '$2a$10$zjiz24Nz7xNjIc8USeAXPuK0yG8a0J/.ghA76A7n1FO5OECwu.9wy', 8092, '2024-01-27 16:41:19', '2024-01-27 16:41:19', NULL);
INSERT INTO `account` VALUES ('CZKJ681706348185', '$2a$10$.WMU7yd/XX.RXgEEY5Lvd.7Le8VTlsXNoN7gmZvtf5kTdNPq/6AVG', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ691706348186', '$2a$10$YZfwlrOfQyERItZEUR3qQOpOGKA4rOtOycGww54oukLS4878upTja', 3705, '2024-01-27 17:36:27', '2024-01-27 17:36:27', NULL);
INSERT INTO `account` VALUES ('CZKJ691710304250', '$2a$10$b76l6msBufOOBmCbQLt1.O7gSn9GtgsKpDgSdVEASxEYWg3Ps3Yha', 2458, '2024-03-13 12:30:50', '2024-03-13 12:30:50', NULL);
INSERT INTO `account` VALUES ('CZKJ711706348186', '$2a$10$klUMfIrRA2UwL6hN0X4MceKKKjgXfsqwiSfGBAHN9n2vQf7tMWmjG', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ731706348186', '$2a$10$uVbmX8pVXENsD8Ku7xgqquaN1taWszTHlv6gWGO.f2gzlkxwx4Rmi', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ741706348185', '$2a$10$vb5endCGUzhggCWpzhT.ouMqvV7f2MRblR2z4y3qjIX9.9lDnqILe', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ751706348186', '$2a$10$NESbsihbRQ30toP59q9C5e9sDF/bfROPks2ueoyapERj2FSuVTyc2', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ801706348186', '$2a$10$p/.Unv3SRTfv8HgwMlzVx.YDon9k/IFyfi8DRott84CFodetdnIyO', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ831706348186', '$2a$10$xzQ5B1xac32kNG.WVyuubOmlETu7sorfDamMJNVSkcTY0/oIKnM9m', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ841710304988', '$2a$10$9NkHqkd.VFWwsiJGknaWr.Qa0yTSW1n8XjkYOL/en3EXuStZ8JmjK', 2458, '2024-03-13 12:43:09', '2024-03-13 12:43:09', NULL);
INSERT INTO `account` VALUES ('CZKJ861706344879', '$2a$10$8il.lSCtPVlv5Y/DcCnJ/ONztJ3Ar5YoT4ptCozFVsQ9Cdx/K4msG', 8092, '2024-01-27 16:41:19', '2024-01-27 16:41:19', NULL);
INSERT INTO `account` VALUES ('CZKJ901704561805', '$2a$10$hHHfHVhY1GoL9UNQYok.Euec4o1LWTiWWhHMxHRUhM31wwL44M/oG', 2458, '2024-01-07 01:23:26', '2024-03-12 11:15:34', NULL);
INSERT INTO `account` VALUES ('CZKJ911706348185', '$2a$10$QANO2DTkkcDZUit1FXK8tOo2PP8GTWVhizlLF7NZQgdOhpJ9VogK6', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ941706337500', '$2a$10$ygJi7mDhVLjzUzCb..CoouYXNgZGFCkKBHASJzbS3Q9q6z0udrOM6', 2458, '2024-01-27 14:38:21', '2024-03-12 11:37:04', '2024-03-12 11:37:05');
INSERT INTO `account` VALUES ('CZKJ991706348185', '$2a$10$O7xY65n7Uw4QWAcGUJtKOeJ5/sAvtX1KQiRyk6zJhj5Wc./NadRG.', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员主键ID',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号名',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '管理员实名',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '管理员邮箱',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE COMMENT '账号'
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 'CZKJ641706344878', 'Cukor', 'cukorzhong@gmail.com');
INSERT INTO `admin` VALUES (2, 'CZKJ261706344878', 'Jack', 'Jack763@gmail.com');
INSERT INTO `admin` VALUES (3, 'CZKJ291706344878', 'Betty', 'Betty844@gmail.com');
INSERT INTO `admin` VALUES (4, 'CZKJ861706344879', 'Lukas', 'lukas164@out-look.com');

-- ----------------------------
-- Table structure for employee
-- ----------------------------
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '员工主键ID',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号名',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '员工实名',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '员工电话',
  `position` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '员工职位',
  `salary` int NULL DEFAULT NULL COMMENT '薪资',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE COMMENT '账号'
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of employee
-- ----------------------------
INSERT INTO `employee` VALUES (1, 'CZKJ521706348185', 'Mary', '13798452063', '采购员', 3000);
INSERT INTO `employee` VALUES (2, 'CZKJ741706348185', '李晓', '18777620849', '仓库管理员', 5000);
INSERT INTO `employee` VALUES (3, 'CZKJ681706348185', '苏敏', '15277820403', '仓库管理员', 3500);
INSERT INTO `employee` VALUES (4, 'CZKJ31706348185', '王鑫悦', '18277620305', '仓库管理员', 4000);
INSERT INTO `employee` VALUES (5, 'CZKJ911706348185', '李沁月', '18577630145', '采购员', 3200);
INSERT INTO `employee` VALUES (6, 'CZKJ111706348185', '吴浩明', '13789720312', '采购员', 3200);
INSERT INTO `employee` VALUES (7, 'CZKJ991706348185', 'Coco', '13489427501', '仓库管理员', 3500);
INSERT INTO `employee` VALUES (8, 'CZKJ141706348185', '赵林', '13675760842', '采购员', 3200);
INSERT INTO `employee` VALUES (9, 'CZKJ831706348186', '黄杰', '18289230654', '仓库管理员', 3500);
INSERT INTO `employee` VALUES (10, 'CZKJ801706348186', '陈邹郭', '19449474438', '采购员', 7400);
INSERT INTO `employee` VALUES (11, 'CZKJ551706348186', '杜季胡', '18369582593', '仓库管理员', 3000);
INSERT INTO `employee` VALUES (12, 'CZKJ751706348186', '马吕梁', '13482333812', '采购员', 6700);
INSERT INTO `employee` VALUES (13, 'CZKJ711706348186', '郭袁席', '13690301079', '仓库管理员', 7400);
INSERT INTO `employee` VALUES (14, 'CZKJ731706348186', '罗孔邱', '15436880023', '采购员', 3300);
INSERT INTO `employee` VALUES (15, 'CZKJ331706348186', '梅盛颜', '18686039042', '采购员', 6600);
INSERT INTO `employee` VALUES (16, 'CZKJ571706348390', '夏姚黄', '15026352199', '仓库管理员', 6700);
INSERT INTO `employee` VALUES (17, 'CZKJ31706348186', '贾危丁', '17130160885', '仓库管理员', 7900);
INSERT INTO `employee` VALUES (18, 'CZKJ691706348186', '姚麻赵', '13157578418', '采购员', 5400);
INSERT INTO `employee` VALUES (19, 'CZKJ381706348186', '燕傅沈', '13912663427', '采购员', 6600);

-- ----------------------------
-- Table structure for employee_status
-- ----------------------------
DROP TABLE IF EXISTS `employee_status`;
CREATE TABLE `employee_status`  (
  `id` bigint NOT NULL COMMENT '主键ID',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号',
  `status` tinyint NOT NULL COMMENT '状态，0离职，1在职，2调离',
  `mark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注，离职原因，正常在岗，调离原因调离去向',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of employee_status
-- ----------------------------
INSERT INTO `employee_status` VALUES (1, 'CZKJ521706348185', 1, NULL);
INSERT INTO `employee_status` VALUES (2, 'CZKJ741706348185', 1, NULL);
INSERT INTO `employee_status` VALUES (3, 'CZKJ681706348185', 1, NULL);
INSERT INTO `employee_status` VALUES (4, 'CZKJ31706348185', 1, NULL);
INSERT INTO `employee_status` VALUES (5, 'CZKJ911706348185', 1, NULL);
INSERT INTO `employee_status` VALUES (6, 'CZKJ111706348185', 1, NULL);
INSERT INTO `employee_status` VALUES (7, 'CZKJ991706348185', 1, NULL);
INSERT INTO `employee_status` VALUES (8, 'CZKJ141706348185', 1, NULL);
INSERT INTO `employee_status` VALUES (9, 'CZKJ831706348186', 0, '能力不够');
INSERT INTO `employee_status` VALUES (10, 'CZKJ801706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (11, 'CZKJ551706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (12, 'CZKJ751706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (13, 'CZKJ711706348186', 2, '能力突出，调离到总部');
INSERT INTO `employee_status` VALUES (14, 'CZKJ731706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (15, 'CZKJ331706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (16, 'CZKJ571706348390', 1, NULL);
INSERT INTO `employee_status` VALUES (17, 'CZKJ31706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (18, 'CZKJ691706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (19, 'CZKJ381706348186', 1, NULL);

-- ----------------------------
-- Table structure for fruits
-- ----------------------------
DROP TABLE IF EXISTS `fruits`;
CREATE TABLE `fruits`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '水果主键ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '水果名称',
  `water` float NULL DEFAULT NULL COMMENT '含水量',
  `sugar` float NULL DEFAULT NULL COMMENT '含糖量',
  `shelf_life` int NULL DEFAULT NULL COMMENT '保质期，天数',
  `origin` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '来源，产地',
  `supplier_id` bigint NULL DEFAULT NULL COMMENT '供应商ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fruits
-- ----------------------------
INSERT INTO `fruits` VALUES (1, '砂糖桔', 90, 80, 15, '广西百色', 1);
INSERT INTO `fruits` VALUES (2, '沃柑', 80, 90, 15, '广西南宁武鸣', 2);

-- ----------------------------
-- Table structure for inventory
-- ----------------------------
DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '库存ID',
  `commodity_id` bigint NULL DEFAULT NULL COMMENT '商品ID，水果ID',
  `quantity` int NULL DEFAULT NULL COMMENT '数量',
  `employee_id` bigint NULL DEFAULT NULL COMMENT '员工ID',
  `warehouse_id` bigint NULL DEFAULT NULL COMMENT '仓库ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of inventory
-- ----------------------------

-- ----------------------------
-- Table structure for logout_users
-- ----------------------------
DROP TABLE IF EXISTS `logout_users`;
CREATE TABLE `logout_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户主键ID',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号名',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户实名',
  `gender` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '性别',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '家庭地址',
  `birthday` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生日',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of logout_users
-- ----------------------------
INSERT INTO `logout_users` VALUES (4, 'CZKJ941706337500', '李丽梅', '女', '13746510502', '广东省深圳市', '1996-08-26 23:03:10');
INSERT INTO `logout_users` VALUES (5, 'CZKJ351708869915', '卢漫琪', '女', '18792930842', '广东潮汕', '2001-10-08 00:00:00');
INSERT INTO `logout_users` VALUES (7, 'CZKJ201709005881', '吴宣仪', '女', '18777263907', '海南省海口市', '1995-03-14 00:00:00');
INSERT INTO `logout_users` VALUES (8, 'CZKJ531709006038', '杨超越', '女', '13408740942', '湖南省长沙市', '1998-12-15 00:00:00');

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '订单描述',
  `status` int NULL DEFAULT 0 COMMENT '订单状态',
  `commodity_id` bigint NULL DEFAULT NULL COMMENT '商品ID，水果ID',
  `quantity` int NULL DEFAULT NULL COMMENT '数量',
  `buyers_id` bigint NULL DEFAULT NULL COMMENT '买家',
  `admin_id` bigint NULL DEFAULT NULL COMMENT '管理员ID',
  `warehouse_id` int NULL DEFAULT NULL COMMENT '仓库ID',
  `receiver_id` bigint NULL DEFAULT NULL COMMENT '收货人ID',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '收货地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orders
-- ----------------------------

-- ----------------------------
-- Table structure for purchase
-- ----------------------------
DROP TABLE IF EXISTS `purchase`;
CREATE TABLE `purchase`  (
  `id` bigint NOT NULL COMMENT '采购ID',
  `staus` int NULL DEFAULT 0 COMMENT '采购状态',
  `employee_id` bigint NULL DEFAULT NULL COMMENT '员工ID',
  `fruit_id` bigint NULL DEFAULT NULL COMMENT '水果ID',
  `quantity` int NULL DEFAULT NULL COMMENT '采购数量',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of purchase
-- ----------------------------

-- ----------------------------
-- Table structure for suppliers
-- ----------------------------
DROP TABLE IF EXISTS `suppliers`;
CREATE TABLE `suppliers`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '供应商主键ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '供应商名',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '供应商地址',
  `contact_person` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系人名',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  `reputation` int NULL DEFAULT NULL COMMENT '信誉',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of suppliers
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户主键ID',
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账号名',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户实名',
  `gender` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '性别',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '家庭地址',
  `birthday` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生日',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (3, 'CZKJ901704561805', 'Cukor', '男', '13746510506', '广东省深圳市福田区下梅林A栋516', '2000-08-07 00:00:00');
INSERT INTO `users` VALUES (6, 'CZKJ341708940332', 'Coco', '女', '13789720842', '广东省深圳市福田区下梅林', '2005-11-17 00:00:00');
INSERT INTO `users` VALUES (9, 'CZKJ441710302534', '方大同', '男', '18789720403', '中国香港', '1982-07-27 00:00:00');
INSERT INTO `users` VALUES (10, 'CZKJ691710304250', '陶喆', '男', '18808426872', '北京市', '1969-07-11 00:00:00');
INSERT INTO `users` VALUES (11, 'CZKJ561710304527', '王力宏', '男', '13589750842', '美国洛杉矶', '1972-06-28 00:00:00');
INSERT INTO `users` VALUES (12, 'CZKJ511710304763', '周杰伦', '男', '15208724949', '中国台湾', '1979-12-18 00:00:00');
INSERT INTO `users` VALUES (13, 'CZKJ11710304892', 'KhalilFong', '男', '18787720872', '云南省怒江市', '1982-05-26 00:00:00');
INSERT INTO `users` VALUES (14, 'CZKJ151710304943', 'DavidTao', '男', '17208760506', '中国台湾', '1969-07-11 00:00:00');
INSERT INTO `users` VALUES (15, 'CZKJ841710304988', 'LeehomWang', '男', '17208760403', '中国台湾', '1973-07-21 00:00:00');
INSERT INTO `users` VALUES (16, 'CZKJ541710305023', 'JayChou', '男', '17285730403', '中国台湾', '1979-08-15 00:00:00');
INSERT INTO `users` VALUES (17, 'CZKJ531710305097', 'viva', '女', '13788660412', '日本东京', '1975-07-20 00:00:00');

-- ----------------------------
-- Table structure for warehouse
-- ----------------------------
DROP TABLE IF EXISTS `warehouse`;
CREATE TABLE `warehouse`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '仓库ID',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '仓库地址',
  `capacity` float NULL DEFAULT NULL COMMENT '仓库容量',
  `status` int NULL DEFAULT 0 COMMENT '仓库状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of warehouse
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
