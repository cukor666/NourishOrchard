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

 Date: 07/05/2024 21:38:45
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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of account
-- ----------------------------
INSERT INTO `account` VALUES ('CZKJ111706348185', '$2a$10$Mdmy6WViAAvcbPYM7n8U2.xgbIOe1d2zLiN1Rw4F3xNqBRMVXreW.', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ11710304892', '$2a$10$/ZmmSH0y76WSNgsVAKZ.6uqdrs8hrrjiSPmqtgh7DGVOSKMn0Qh5W', 2458, '2024-03-13 12:41:32', '2024-03-13 12:41:32', NULL);
INSERT INTO `account` VALUES ('CZKJ141706348185', '$2a$10$YxcBKcNasmbu/sJr.7rTVOoWOg2A9eOeQ05NUmFZtlz6hTb80O1QW', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ151710304943', '$2a$10$PRznGzAPo9REHus3pFjc1eABn24vxpuKzsPnnKu7c17no6W0FQsmi', 2458, '2024-03-13 12:42:23', '2024-03-13 12:42:23', NULL);
INSERT INTO `account` VALUES ('CZKJ171713011483', '$2a$10$dVgC3xBjzC6lPe0KhgoQT.RkJlOES3nzwG5mYFs2GtPKPDzNcB8/C', 2458, '2024-04-13 20:31:23', '2024-04-13 20:31:23', NULL);
INSERT INTO `account` VALUES ('CZKJ201709005881', '$2a$10$LeJBxeJvMcXxjQum88RrvOnoEPfe/5xs.twtvO0sE.6vyRBnojU7S', 2458, '2024-02-27 11:51:22', '2024-04-02 23:53:30', NULL);
INSERT INTO `account` VALUES ('CZKJ261706344878', '$2a$10$froOrMspKgquGAuwniGy5uGF3EJRp.RIVI9jdMBA2h.B.kQOi7go2', 8092, '2024-01-27 16:41:19', '2024-01-27 16:41:19', NULL);
INSERT INTO `account` VALUES ('CZKJ291706344878', '$2a$10$W.Sa7H7aKaZof6MRpfx7WOIr9juyakZQRvW4NE86dhpH1rMdD3.Fq', 8092, '2024-01-27 16:41:19', '2024-01-27 16:41:19', NULL);
INSERT INTO `account` VALUES ('CZKJ31706348185', '$2a$10$06j2m5cAzBnDCmV8Cs4LaO2Og/tTjRpM5OR9vWXV8faMMgcIZm.na', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ31706348186', '$2a$10$RuXlZEcHPqWSRpv/YDUc..SgPwEJQHYdXy5EEbeopmE5g5DGoKJzq', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ321713016563', '$2a$10$yuFcvUYIgnj/7fMm7y7/7.A27pQHnPXn/Iy28NzthfkfFFWEdAPiq', 2458, '2024-04-13 21:56:03', '2024-04-13 21:56:03', NULL);
INSERT INTO `account` VALUES ('CZKJ331706348186', '$2a$10$gM9lL11s48XOfgWdDfzkI.jmyw8g8dNeWQ3fYX5UTsHWCAqhSibnW', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ341708940332', '$2a$10$D8gQ1Aqvhg0TbQt2.i8RueITfX2wn4dgW1FlCFbwu6tG84/3lBmn2', 2458, '2024-02-26 17:38:53', '2024-03-12 11:33:21', NULL);
INSERT INTO `account` VALUES ('CZKJ351708869915', '$2a$10$.QzPeTzzYm.Sc0WpaS7p9OCYFryX9ugXNAdgFjFrwbQyWhO3SZQDG', 2458, '2024-02-25 22:05:16', '2024-03-12 11:26:48', '2024-03-12 11:26:49');
INSERT INTO `account` VALUES ('CZKJ381706348186', '$2a$10$aegw0iJnyFrHRegvVwuQGebFmbONHmS1kPCseNfHsvpOpDkxcyxTe', 3705, '2024-01-27 17:36:27', '2024-01-27 17:36:27', NULL);
INSERT INTO `account` VALUES ('CZKJ441710302534', '$2a$10$jIuu5FQBj7RN1gGgdhu5l.2ZW0y6y5cbpcaw.xFa45RkKjQeyAhW2', 2458, '2024-03-13 12:02:14', '2024-03-13 12:02:14', NULL);
INSERT INTO `account` VALUES ('CZKJ491712129368', '$2a$10$98x6Xf4TU1l6d75y.8WaPunKVzrd1.g9ayuexYYuSLmppH1wb9VJy', 2458, '2024-04-03 15:29:28', '2024-04-03 15:32:45', NULL);
INSERT INTO `account` VALUES ('CZKJ511710304763', '$2a$10$7/1zHKQ72h3BIfzFODp1H.GAsyyKuW2aU1onOEpb2OXTG0BYZW/X.', 2458, '2024-03-13 12:39:23', '2024-04-03 00:36:44', '2024-04-03 00:36:44');
INSERT INTO `account` VALUES ('CZKJ521706348185', '$2a$10$YQck8aC9a42QKTxz0JAC2usmRd9Fet19Lpx4Eu9Hi7t0hgbB9D6l.', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ531709006038', '$2a$10$T5c5qMfcj8JHICW6mwz/OuU6HbppCmZAnkp1zSoKtZ5Yx07uICcqm', 2458, '2024-02-27 11:53:58', '2024-03-11 17:45:18', '2024-03-11 17:45:19');
INSERT INTO `account` VALUES ('CZKJ531710305097', '$2a$10$cr07HcBy/8ceWLf34Hcvcu0ZHRjStNH4UukgzlqPvO1EradhbhZha', 2458, '2024-03-13 12:44:57', '2024-03-13 12:44:57', NULL);
INSERT INTO `account` VALUES ('CZKJ541710305023', '$2a$10$t8v6lh1kHP/KWRGoSZtgseXbx9xWhTgXY62ia7YvPrn.YewpY3cGq', 2458, '2024-03-13 12:43:44', '2024-03-13 12:43:44', NULL);
INSERT INTO `account` VALUES ('CZKJ551706348186', '$2a$10$eUZ7l0uJBrx5YMChbhQ8Pucz.gBv7hCV8WS7F4xBAdBBnj/xcgkDu', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ561710304527', '$2a$10$jxaDJkGoKgE0WdU25NA.MOQG6knJg0VQO9mLJQC3v7MgloQSS2UX6', 2458, '2024-03-13 12:35:27', '2024-03-13 12:35:27', NULL);
INSERT INTO `account` VALUES ('CZKJ571706348390', '$2a$10$TIaYOpBJkpOxP0V7FGuOKuBgk9fpQMMq5VCrcwNkUX7C8gQhipzIa', 8092, '2024-01-27 17:39:51', '2024-04-04 19:00:40', NULL);
INSERT INTO `account` VALUES ('CZKJ641706344878', '$2a$10$1gQz.sLlThsUyiypbgHPEe1YiE.0VMPxCJMWf2Fe3tj.CKqOX.w6.', 8092, '2024-01-27 16:41:19', '2024-04-13 22:41:23', NULL);
INSERT INTO `account` VALUES ('CZKJ641712129852', '$2a$10$R4xp5KPhK2NqRTRw6k6p8OKapPEVtWOhNal416TQiw2DV9ccSOXa6', 2458, '2024-04-03 15:37:33', '2024-04-03 15:37:33', NULL);
INSERT INTO `account` VALUES ('CZKJ681706348185', '$2a$10$.WMU7yd/XX.RXgEEY5Lvd.7Le8VTlsXNoN7gmZvtf5kTdNPq/6AVG', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ691706348186', '$2a$10$YZfwlrOfQyERItZEUR3qQOpOGKA4rOtOycGww54oukLS4878upTja', 3705, '2024-01-27 17:36:27', '2024-01-27 17:36:27', NULL);
INSERT INTO `account` VALUES ('CZKJ691710304250', '$2a$10$b76l6msBufOOBmCbQLt1.O7gSn9GtgsKpDgSdVEASxEYWg3Ps3Yha', 2458, '2024-03-13 12:30:50', '2024-03-13 12:30:50', NULL);
INSERT INTO `account` VALUES ('CZKJ711706348186', '$2a$10$klUMfIrRA2UwL6hN0X4MceKKKjgXfsqwiSfGBAHN9n2vQf7tMWmjG', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ731706348186', '$2a$10$uVbmX8pVXENsD8Ku7xgqquaN1taWszTHlv6gWGO.f2gzlkxwx4Rmi', 8092, '2024-01-27 17:36:26', '2024-04-13 22:59:03', NULL);
INSERT INTO `account` VALUES ('CZKJ741706348185', '$2a$10$vb5endCGUzhggCWpzhT.ouMqvV7f2MRblR2z4y3qjIX9.9lDnqILe', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ751706348186', '$2a$10$NESbsihbRQ30toP59q9C5e9sDF/bfROPks2ueoyapERj2FSuVTyc2', 3705, '2024-01-27 17:36:26', '2024-01-27 17:36:26', NULL);
INSERT INTO `account` VALUES ('CZKJ801706348186', '$2a$10$p/.Unv3SRTfv8HgwMlzVx.YDon9k/IFyfi8DRott84CFodetdnIyO', 8092, '2024-01-27 17:36:26', '2024-04-04 18:59:20', NULL);
INSERT INTO `account` VALUES ('CZKJ831706348186', '$2a$10$xzQ5B1xac32kNG.WVyuubOmlETu7sorfDamMJNVSkcTY0/oIKnM9m', 8092, '2024-01-27 17:36:26', '2024-04-04 18:18:50', NULL);
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
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 'CZKJ641706344878', 'Cukor', 'cukorzhong@gmail.com');
INSERT INTO `admin` VALUES (2, 'CZKJ261706344878', 'Jack', 'Jack763@gmail.com');
INSERT INTO `admin` VALUES (3, 'CZKJ291706344878', 'Betty', 'Betty844@gmail.com');
INSERT INTO `admin` VALUES (4, 'CZKJ861706344879', 'Lukas', 'lukas164@out-look.com');
INSERT INTO `admin` VALUES (16, 'CZKJ831706348186', '黄杰', 'huangjie@gmail.com');
INSERT INTO `admin` VALUES (17, 'CZKJ801706348186', '陈邹郭', 'chentest@gmail.com');
INSERT INTO `admin` VALUES (18, 'CZKJ571706348390', '夏姚黄', 'xiayaohuang@qq.com');
INSERT INTO `admin` VALUES (19, 'CZKJ731706348186', '罗孔邱', 'luokongqiu@gmail.com');

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
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `employee_status` VALUES (9, 'CZKJ831706348186', 2, '表现突出晋升管理员');
INSERT INTO `employee_status` VALUES (10, 'CZKJ801706348186', 2, '测试晋升管理员');
INSERT INTO `employee_status` VALUES (11, 'CZKJ551706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (12, 'CZKJ751706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (13, 'CZKJ711706348186', 2, '能力突出，调离到总部');
INSERT INTO `employee_status` VALUES (14, 'CZKJ731706348186', 2, '表现突出晋升管理员');
INSERT INTO `employee_status` VALUES (15, 'CZKJ331706348186', 1, NULL);
INSERT INTO `employee_status` VALUES (16, 'CZKJ571706348390', 2, '');
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
  `water` int NULL DEFAULT NULL COMMENT '含水量',
  `sugar` int NULL DEFAULT NULL COMMENT '含糖量',
  `shelf_life` int NULL DEFAULT NULL COMMENT '保质期，天数',
  `origin` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '来源，产地',
  `supplier_id` bigint NULL DEFAULT NULL COMMENT '供应商ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of fruits
-- ----------------------------
INSERT INTO `fruits` VALUES (1, '砂糖桔', 90, 80, 15, '广西百色', 1);
INSERT INTO `fruits` VALUES (2, '沃柑', 80, 90, 15, '广西南宁武鸣', 2);
INSERT INTO `fruits` VALUES (3, '香蕉', 74, 14, 3, '海南省海口市龙华区', 3);
INSERT INTO `fruits` VALUES (4, '橙子', 87, 9, 30, '江西省南昌市东湖区', 4);
INSERT INTO `fruits` VALUES (5, '柠檬', 89, 3, 40, '云南省昆明市五华区', 5);
INSERT INTO `fruits` VALUES (6, '草莓', 91, 5, 5, '四川省成都市武侯区', 6);
INSERT INTO `fruits` VALUES (7, '葡萄', 81, 16, 7, '新疆维吾尔自治区乌鲁木齐市天山区', 7);
INSERT INTO `fruits` VALUES (8, '菠萝', 86, 10, 3, '广东省广州市越秀区', 8);
INSERT INTO `fruits` VALUES (9, '樱桃', 82, 13, 20, '河北省石家庄市长安区', 9);
INSERT INTO `fruits` VALUES (10, '猕猴桃', 83, 9, 30, '湖北省武汉市江岸区', 6);
INSERT INTO `fruits` VALUES (11, '芒果', 83, 14, 4, '广西壮族自治区南宁市青秀区', 10);
INSERT INTO `fruits` VALUES (12, '杏子', 86, 4, 20, '甘肃省兰州市城关区', 8);
INSERT INTO `fruits` VALUES (15, '西瓜', 92, 6, 7, '甘肃省张掖市甘州区', 8);
INSERT INTO `fruits` VALUES (16, '哈密瓜', 90, 8, 20, '新疆维吾尔自治区哈密市伊州区', 1);
INSERT INTO `fruits` VALUES (17, '柿子', 81, 18, 30, '山西省太原市小店区', 17);
INSERT INTO `fruits` VALUES (18, '椰子', 47, 6, 50, '海南省三亚市天涯区', 2);
INSERT INTO `fruits` VALUES (19, '橄榄', 80, 1, 180, '广东省深圳市福田区', 3);
INSERT INTO `fruits` VALUES (20, '杨梅', 85, 8, 3, '浙江省杭州市上城区', 11);
INSERT INTO `fruits` VALUES (22, '苹果', 85, 10, 30, '山东省济南市历下区', 2);
INSERT INTO `fruits` VALUES (23, '荔枝', 82, 16, 5, '广东省广州市海珠区', 1);
INSERT INTO `fruits` VALUES (24, '龙眼', 82, 15, 7, '福建省福州市鼓楼区', 2);
INSERT INTO `fruits` VALUES (25, '柚子', 91, 2, 30, '广东省梅州市梅江区', 3);
INSERT INTO `fruits` VALUES (26, '蓝莓', 85, 10, 7, '江苏省南京市玄武区', 4);
INSERT INTO `fruits` VALUES (27, '黑莓', 88, 5, 3, '浙江省宁波市海曙区', 5);
INSERT INTO `fruits` VALUES (28, '石榴', 83, 13, 30, '山西省大同市城区', 6);
INSERT INTO `fruits` VALUES (29, '木瓜', 88, 8, 5, '海南省三亚市吉阳区', 7);
INSERT INTO `fruits` VALUES (30, '榴莲', 65, 17, 7, '广东省深圳市罗湖区', 8);
INSERT INTO `fruits` VALUES (31, '火龙果', 87, 8, 15, '广东省珠海市香洲区', 9);
INSERT INTO `fruits` VALUES (32, '百香果', 83, 11, 20, '广西壮族自治区桂林市秀峰区', 10);
INSERT INTO `fruits` VALUES (33, '山楂', 86, 9, 7, '河北省邯郸市邯山区', 11);
INSERT INTO `fruits` VALUES (34, '杨桃', 91, 4, 20, '广东省佛山市禅城区', 12);
INSERT INTO `fruits` VALUES (35, '柿子', 81, 18, 30, '山西省太原市小店区', 13);
INSERT INTO `fruits` VALUES (36, '椰子', 47, 6, 50, '海南省三亚市天涯区', 14);
INSERT INTO `fruits` VALUES (38, '杨梅', 85, 8, 3, '浙江省杭州市上城区', 16);
INSERT INTO `fruits` VALUES (40, '桂圆', 79, 15, 365, '广东省东莞市莞城区', 18);
INSERT INTO `fruits` VALUES (42, '莲雾', 83, 7, 7, '福建省厦门市思明区', 20);
INSERT INTO `fruits` VALUES (43, '水蜜桃', 75, 20, 7, '四川省攀枝花市', 19);

-- ----------------------------
-- Table structure for fruits_commodity
-- ----------------------------
DROP TABLE IF EXISTS `fruits_commodity`;
CREATE TABLE `fruits_commodity`  (
  `id` bigint NOT NULL COMMENT '水果ID，商品ID',
  `price` decimal(10, 2) NOT NULL COMMENT '水果价格',
  `imgs` json NOT NULL COMMENT '水果图集',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '水果描述',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fruits_commodity
-- ----------------------------
INSERT INTO `fruits_commodity` VALUES (1, 3.50, '[\"https://momo-wp.s3.ap-northeast-1.amazonaws.com/wp-content/uploads/2018/11/%E5%9C%967.jpg\", \"https://p2.itc.cn/images01/20210114/5646cfc03530484f9360fcc736e07f68.jpeg\", \"https://img95.699pic.com/photo/50072/1365.jpg_wh300.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (2, 3.50, '[\"https://lh4.googleusercontent.com/proxy/G4gESZ6u1LTsQIt_FhUwiGzH6J5xVLe3VuIwGG7nqUv4sxz1c-Wd3KrAu2kfb5L1VLXiUHrTpQ3SEUq4zq-DLuw9nZkxmaKu6XttH7aR9l9sxbXVEtXp5FFX_jXZjA6xHRjLi2fyhYIr\", \"https://img95.699pic.com/photo/50080/9753.jpg_wh300.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (3, 3.50, '[\"https://www.bowtie.com.hk/blog/wp-content/uploads/2023/11/07015921/%E9%A6%99%E8%95%89.jpg\", \"https://c.files.bbci.co.uk/477F/production/_116230381_gettyimages-492757968.jpg\", \"https://images.agriharvest.tw/wp-content/uploads/2023/05/7305-8-0-1024x766.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (4, 3.50, '[\"https://exp-picture.cdn.bcebos.com/dd58d02c5b1b1ede65cd568e981fceecd2d90fa9.jpg?x-bce-process=image%2Fcrop%2Cx_0%2Cy_0%2Cw_820%2Ch_516%2Fformat%2Cf_auto%2Fquality%2Cq_80\", \"https://k.sinaimg.cn/n/front/277/w640h437/20181212/U3dE-hqackaa5877093.jpg/w700d1q75cms.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (5, 3.50, '[\"https://image.healthjd.com/da/jfs/t1/175013/33/14234/1201959/60c3664eEdda4aa17/703b12235f04461b.jpg\", \"https://celeplate.co.uk/cdn/shop/files/lemon1.png?v=1705321312&width=1946\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (6, 3.50, '[\"https://cdn.pixabay.com/photo/2022/05/27/10/35/strawberry-7224875_640.jpg\", \"https://global-blog.cpcdn.com/tw/2021/12/AdobeStock_343373754--1-.jpeg\", \"https://theoita.com/wp-content/uploads/2020/07/%EF%BE%8D%EF%BE%9E%EF%BE%98%EF%BD%B0%EF%BE%82-1.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (7, 3.50, '[\"https://images.agriharvest.tw/wp-content/uploads/2022/06/0-18-1024x768.jpg\", \"https://vegemap.merit-times.com/_i/assets/upload/nutrition/d0bf9672ffe71de8c6b371f806674023.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (8, 3.50, '[\"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTd3IngD293w4FPcX-Jd4d_WnH8CLLKJww8yZLngrAPuA&s\", \"https://p8.itc.cn/q_70/images01/20220402/97467c7372ee4022ad75e9a0862d23fe.jpeg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (9, 3.50, '[\"https://img95.699pic.com/photo/60019/9969.jpg_wh300.jpg\", \"https://images.sbs.com.au/dims4/default/92fda7a/2147483647/strip/true/crop/1920x1080+0+107/resize/1280x720!/quality/90/?url=http%3A%2F%2Fsbs-au-brightspot.s3.amazonaws.com%2Fdrupal%2Fyourlanguage%2Fpublic%2Fcherry-2554364_1920.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (10, 3.50, '[\"https://www.ecoportal.net/wp-content/uploads/2012/10/kiwi-rdn-801x534.webp\", \"https://k.sinaimg.cn/n/sinacn20190815ac/277/w640h437/20190815/9399-ichcymv3193252.jpg/w700d1q75cms.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (11, 3.50, '[\"https://k.sinaimg.cn/n/front/277/w640h437/20190417/kYcN-hvvuiym7277966.jpg/w700d1q75cms.jpg\", \"https://global-blog.cpcdn.com/tw/2020/06/76d3f7d1c30e2e48889aa9cffdb9c28c.jpeg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (12, 3.50, '[\"https://img95.699pic.com/photo/60029/2558.jpg_wh300.jpg\", \"https://img95.699pic.com/photo/60115/7307.jpg_wh300.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (15, 3.50, '[\"https://goods.watchinese.com/img/cms/Blog/202306-watermelon/0.jpeg\", \"https://pica.zhimg.com/v2-efb229deccd9b6cfc6187bbfd8514dfe_720w.jpg?source=172ae18b\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (16, 3.50, '[\"https://www.ttvc.com.tw/data/images/202304/dreamstime_s_14781029.jpg\", \"https://p4.itc.cn/images01/20230102/da9fb1102537468fb1d26ca37318ae0d.jpeg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (17, 3.50, '[\"https://img.ltn.com.tw/Upload/health/page/800/2022/10/08/php0Src0k.jpg\", \"https://cdn.esence.travel/%E5%AE%98%E7%B6%B2%E6%96%87%E7%AB%A0%E5%B0%81%E9%9D%A2-6-1628x1000.png\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (18, 3.50, '[\"https://as.chdev.tw/web/article/d/9/4/d3b7ff83-99d1-4903-8463-79806b53f7191677059699.jpg\", \"https://www.leezen.com.tw/uploads/article/1058_l.jpg?r=615663115\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (19, 3.50, '[\"https://img95.699pic.com/photo/60014/4791.jpg_wh300.jpg\", \"https://p9.itc.cn/images01/20220622/2582bd168df947158f7a36467a713913.jpeg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (20, 3.50, '[\"https://mediabluk.cnr.cn/record/img/cnr/CNRCDP/2023/0602/803fbea5de9d4b168df04d9f7cb7d28610.jpg\", \"https://img95.699pic.com/photo/50046/2335.jpg_wh300.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (22, 3.50, '[\"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRqA9lgtsgQ7nyE-PtDw1vNBWBQ9yeFKgcFybDXj4atvw&s\", \"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTsPyjzoPK3d2tkqn_Yn6vRvtQy4fr4buT2kIG2jNQ7Nw&s\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (23, 3.50, '[\"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSEZ0amGbq-MDOp_FiM6c1SrBPLXTeE2QvRNDgZYeh21Q&s\", \"https://cdn.chinafruitportal.com/2020/04/%E5%A6%83%E5%AD%90%E7%AC%91%E8%8D%94%E6%9E%9D/%E5%A6%83%E5%AD%90%E7%AC%91%E8%8D%94%E6%9E%9D-1.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (24, 3.50, '[\"https://lh5.googleusercontent.com/proxy/m_iwmbZ3n8C3uDco10eHYwhMPY0S3IXB-aaupjFWieJAEpVQic8Pk6XyJ4fWBIL0I4DDLF9SEvMX_34RC2Px_5uViQ\", \"https://www.gz.gov.cn/img/0/210/210935/8178185.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (25, 3.50, '[\"https://www.uho.com.tw/userfiles/sm/sm1200675_images_A1/2022082368008497.jpg\", \"https://s3.mababy.com/mababy-production/knowledge/8937a5bb-6191-489a-861e-6126185694a2\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (26, 3.50, '[\"https://img95.699pic.com/photo/50127/1939.jpg_wh300.jpg\", \"https://hips.hearstapps.com/hmg-prod/images/gettyimages-1910897591-1-65eae89db2d9c.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (27, 3.50, '[\"https://img95.699pic.com/photo/60052/2421.jpg_wh860.jpg\", \"https://pic4.zhimg.com/v2-7ebb2a5b63c994fba1f85ae8b13beaff_b.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (28, 3.50, '[\"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTD0oLaJLexyW2ZJvoGXVkhOGoGLzsAiwI-PR0KlwIing&s\", \"https://cdn-news.org/WebView/GetMedia.ashx?PK=0000000021befc7732a7acce17de434297cde4764099c16f&VideoSize=1\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (29, 3.50, '[\"https://www.gz.gov.cn/img/0/210/210959/8178181.png\", \"https://professorlindotcom.files.wordpress.com/2017/05/farm-seeds-100-hybrid-papaya-2-packet-seeds-original-imaefec4n9rwk7mh.jpeg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (30, 3.50, '[\"https://cdn-assets-eu.frontify.com/s3/frontify-enterprise-files-eu/eyJvYXV0aCI6eyJjbGllbnRfaWQiOiJmcm9udGlmeS1maW5kZXIifSwicGF0aCI6ImloaC1oZWFsdGhjYXJlLWJlcmhhZFwvYWNjb3VudHNcL2MzXC80MDAwNjI0XC9wcm9qZWN0c1wvMjA5XC9hc3NldHNcLzgyXC8zNzk2OFwvNGFiZmRmNGZlMDA1N2JjNTU2MThjYTU0ZTA2OTIwOWQtMTY1ODMwMDAwOS5qcGcifQ:ihh-healthcare-berhad:q8WUI77npjt2OB4jweFlIdzZGLziKLUEiZ265-19-18?format=webp\", \"https://static01.nyt.com/images/2013/12/08/travel/08-pursuits-span/08-pursuits-span-master1050.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (31, 3.50, '[\"https://pgw.udn.com.tw/gw/photo.php?u=https://uc.udn.com.tw/photo/2020/06/03/99/7977261.jpg&x=0&y=0&sw=1209&sh=806&s=Y&exp=3600\", \"https://product.hstatic.net/200000325223/product/thanh_long-01_c635795659f64d07a67a5900d47e2f7b_large_d464adf0306e40a6958bfea66ba9e46a_large.png\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (32, 3.50, '[\"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT6jNS6_ewajSNOia_wfKgRXHQt_MxJsYExf07dAJbIQQ&s\", \"https://product.hstatic.net/200000325223/product/44f920810a271a3bdd789bf275c2cad_cc8529666234418993f4d483c9f9939f_large_27566d30a81e41a39de3b1718e3db29e_master.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (33, 3.50, '[\"https://p6.itc.cn/images01/20201027/cda311de69f8471bbcce50015f533c03.jpeg\", \"https://p2.itc.cn/q_70/images01/20220617/852f7cb558e74c1485e2971b134811fa.png\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (34, 3.50, '[\"https://pic3.zhimg.com/80/v2-9b90ec39f6abc95206e7b13e3ff7fe42_1440w.webp\", \"https://p4.itc.cn/images01/20220316/e5cd4ada228c4d80a764b1b72131e621.jpeg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (35, 3.50, '[\"https://img.ltn.com.tw/Upload/health/page/800/2022/10/08/php0Src0k.jpg\", \"https://cdn.esence.travel/%E5%AE%98%E7%B6%B2%E6%96%87%E7%AB%A0%E5%B0%81%E9%9D%A2-6-1628x1000.png\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (36, 3.50, '[\"https://as.chdev.tw/web/article/d/9/4/d3b7ff83-99d1-4903-8463-79806b53f7191677059699.jpg\", \"https://www.leezen.com.tw/uploads/article/1058_l.jpg?r=615663115\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (38, 3.50, '[\"https://mediabluk.cnr.cn/record/img/cnr/CNRCDP/2023/0602/803fbea5de9d4b168df04d9f7cb7d28610.jpg\", \"https://img95.699pic.com/photo/50046/2335.jpg_wh300.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (40, 3.50, '[\"https://k.sinaimg.cn/n/front/277/w640h437/20190222/F0l5-htknpmh5372771.jpg/w700d1q75cms.jpg\", \"https://pic1.zhimg.com/v2-0e8307dea7d5ea521531c6236dbb369c_720w.jpg?source=172ae18b\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (42, 3.50, '[\"https://p7.itc.cn/q_70/images01/20230605/f436134d342f4bf5a272ae9c89c41e1b.jpeg\", \"https://k.sinaimg.cn/n/sinacn20200121ac/205/w612h393/20200121/c5dd-innckce4819732.jpg/w700d1q75cms.jpg\"]', '香甜可口，美味十足，值得拥有');
INSERT INTO `fruits_commodity` VALUES (43, 3.50, '[\"https://img95.699pic.com/photo/50135/6620.jpg_wh300.jpg\", \"https://superbuy-hifamily.cdn.hinet.net/superbuy_admin/images/product/525/PR1606160003_main.jpg\"]', '香甜可口，美味十足，值得拥有');

-- ----------------------------
-- Table structure for inventory
-- ----------------------------
DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '库存ID',
  `commodity_id` bigint NULL DEFAULT NULL COMMENT '商品ID，水果ID',
  `quantity` double NULL DEFAULT NULL COMMENT '数量，单位：吨',
  `employee_id` bigint NULL DEFAULT NULL COMMENT '员工ID',
  `warehouse_id` bigint NULL DEFAULT NULL COMMENT '仓库ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of inventory
-- ----------------------------
INSERT INTO `inventory` VALUES (1, 1, 30, 3, 1, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (2, 2, 20, 1, 1, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (3, 3, 25, 2, 2, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (4, 4, 30, 3, 3, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (5, 5, 15, 4, 4, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (6, 6, 35, 5, 5, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (7, 7, 40, 6, 6, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (8, 8, 45, 7, 7, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (9, 9, 50, 8, 8, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (10, 10, 55, 9, 9, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (11, 11, 60, 10, 10, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (12, 12, 65, 11, 1, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (13, 13, 70, 12, 2, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (14, 14, 75, 13, 3, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (15, 15, 80, 14, 4, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (16, 16, 85, 1, 5, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (17, 17, 90, 2, 6, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (18, 18, 95, 3, 7, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (19, 19, 100, 4, 8, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (20, 20, 105, 5, 9, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (21, 21, 110, 6, 10, '2024-04-02 22:56:46', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (22, 7, 8, 11, 2, '2024-04-05 22:59:24', '2024-04-17 10:30:23', NULL);
INSERT INTO `inventory` VALUES (23, 1, 8, 11, 4, '2024-04-05 23:38:12', '2024-04-17 10:30:23', '2024-04-06 18:21:51');
INSERT INTO `inventory` VALUES (24, 15, 14, 7, 3, '2024-04-15 12:53:32', '2024-04-17 10:30:23', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of logout_users
-- ----------------------------
INSERT INTO `logout_users` VALUES (4, 'CZKJ941706337500', '李丽梅', '女', '13746510502', '广东省深圳市', '1996-08-26 23:03:10');
INSERT INTO `logout_users` VALUES (5, 'CZKJ351708869915', '卢漫琪', '女', '18792930842', '广东潮汕', '2001-10-08 00:00:00');
INSERT INTO `logout_users` VALUES (8, 'CZKJ531709006038', '杨超越', '女', '13408740942', '湖南省长沙市', '1998-12-15 00:00:00');
INSERT INTO `logout_users` VALUES (12, 'CZKJ511710304763', '周杰伦', '男', '15208724949', '中国台湾', '1979-12-18 00:00:00');

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '订单描述',
  `status` int NULL DEFAULT 0 COMMENT '订单状态，0全部，1待付款，2待发货，3待收货，4已完成，5已取消',
  `commodity_id` bigint NULL DEFAULT NULL COMMENT '商品ID，水果ID',
  `quantity` double NULL DEFAULT NULL COMMENT '数量，单位吨',
  `buyer_id` bigint NULL DEFAULT NULL COMMENT '买家',
  `admin_id` bigint NULL DEFAULT NULL COMMENT '管理员ID，如果等于0则表示还没有管理员接管',
  `warehouse_id` int NULL DEFAULT NULL COMMENT '仓库ID',
  `receiver_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '收货人名',
  `receiver_phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '收货人联系电话',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '收货地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES (1, '甜美西瓜', 0, 4, 3, 13, 2, 3, '黄旭', '17289490342', '广东省深圳市福田区景田地铁A口景蜜水果店', '急需', '2024-04-15 18:07:00', '2024-04-22 15:20:31', '2024-04-22 11:30:54');
INSERT INTO `orders` VALUES (3, '香甜百香果美味可口', 2, 32, 1.4, 7, 1, 1, '张贤', '13711042751', '广东省广州市天河区高塘石5巷16号', '等不及了', '2024-04-22 11:10:28', '2024-04-22 15:58:19', NULL);
INSERT INTO `orders` VALUES (4, '新鲜草莓美味可口', 1, 6, 2.3, 7, 0, 0, '李明', '13876543210', '上海市浦东新区世纪大道123号', '请尽快发货', '2024-04-22 16:22:38', '2024-04-22 16:34:39', NULL);
INSERT INTO `orders` VALUES (5, '香脆苹果多汁可口', 5, 12, 0.8, 7, 0, 0, '王小红', '15098765432', '北京市朝阳区建国路789号', '谢谢！', '2024-04-22 16:23:11', '2024-04-22 16:35:20', NULL);
INSERT INTO `orders` VALUES (6, '甜美木瓜清爽可口', 1, 29, 1, 7, 0, 0, '刘大山', '13712345678', '广西省南宁市青秀区花明路456号', '不要错过了', '2024-04-22 16:24:29', '2024-04-22 16:34:47', NULL);
INSERT INTO `orders` VALUES (7, '酸甜橙子汁多可口', 2, 4, 1.5, 7, 0, 0, '陈小芳', '13987654321', '湖北省武汉市江汉区中山大道789号', '急需配送', '2024-04-22 16:25:06', '2024-04-22 16:35:09', NULL);
INSERT INTO `orders` VALUES (8, '浓郁葡萄美味可口', 2, 7, 0.5, 7, 0, 0, '赵明阳', '13654321098', '四川省成都市锦江区春熙路123号', '谢谢您！', '2024-04-22 16:25:41', '2024-04-22 16:35:12', NULL);
INSERT INTO `orders` VALUES (9, '清新柠檬酸甜可口', 1, 15, 1.2, 7, 0, 0, '李华', '13912345678', '江苏省南京市玄武区紫金山路5号', '请注意保鲜', '2024-04-22 16:32:19', '2024-04-22 16:32:19', NULL);
INSERT INTO `orders` VALUES (10, '甘甜橙子多汁可口', 3, 22, 0.8, 7, 0, 0, '王丽', '15898765432', '浙江省杭州市西湖区断桥残雪1号', '送货前请电话联系', '2024-04-22 16:32:29', '2024-04-22 16:35:14', NULL);
INSERT INTO `orders` VALUES (11, '热带芒果香甜可口', 1, 34, 1.5, 7, 0, 0, '张伟', '17612344321', '福建省厦门市思明区鼓浪屿路2号', '急用，请加急处理', '2024-04-22 16:32:36', '2024-04-22 16:32:36', NULL);
INSERT INTO `orders` VALUES (12, '鲜嫩草莓甜美可口', 4, 6, 1, 7, 0, 0, '赵敏', '13109876543', '四川省重庆市渝中区解放碑23号', '周末到货', '2024-04-22 16:33:08', '2024-04-22 16:35:16', NULL);
INSERT INTO `orders` VALUES (13, '多汁西瓜清甜可口', 2, 15, 2, 7, 0, 0, '孙悟空', '14785236974', '湖北省武汉市武昌区黄鹤楼5号', '注意不要压坏', '2024-04-22 16:33:33', '2024-04-22 16:35:23', NULL);
INSERT INTO `orders` VALUES (14, '多汁西瓜清甜可口', 1, 15, 2, 7, 0, 0, '孙悟空', '14785236974', '湖北省武汉市武昌区黄鹤楼5号', '注意不要压坏', '2024-04-23 16:10:56', '2024-04-23 16:10:56', NULL);
INSERT INTO `orders` VALUES (15, '香甜水蜜桃清甜可口', 1, 15, 2, 7, 0, 0, '孙悟空', '14785236974', '湖北省武汉市武昌区黄鹤楼5号', '注意不要压坏', '2024-04-23 16:11:30', '2024-04-23 16:11:30', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of suppliers
-- ----------------------------
INSERT INTO `suppliers` VALUES (1, '广东深圳星灿有限公司', '广东省深圳市龙华区简上村', '钟航飞', '13789035512', 99);
INSERT INTO `suppliers` VALUES (2, '鲜果缘有限公司', '广东省深圳市福田区', '何蓝', '13100102030', 89);
INSERT INTO `suppliers` VALUES (3, '甜蜜果园有限公司', '广西壮族自治区桂林市七星区', '姚冬梅', '13212345678', 87);
INSERT INTO `suppliers` VALUES (4, '果香天地有限公司', '海南省三亚市天涯区', '孟浩然', '13398765432', 91);
INSERT INTO `suppliers` VALUES (5, '绿意盎然有限公司', '云南省昆明市五华区', '赵晓风', '13401020304', 87);
INSERT INTO `suppliers` VALUES (6, '丰收农庄有限公司', '福建省厦门市思明区', '李白云', '13599887766', 90);
INSERT INTO `suppliers` VALUES (7, '翠绿园艺有限公司', '江西省南昌市西湖区', '王翠花', '13611223344', 86);
INSERT INTO `suppliers` VALUES (8, '新鲜直供有限公司', '贵州省贵阳市南明区', '张小凡', '13777665544', 85);
INSERT INTO `suppliers` VALUES (9, '丰年农产品有限公司', '湖南省长沙市芙蓉区', '刘星月', '13866554433', 92);
INSERT INTO `suppliers` VALUES (10, '翠花农业有限公司', '广东省珠海市香洲区', '周莹', '18500112233', 86);
INSERT INTO `suppliers` VALUES (11, '盛世果蔬有限公司', '江苏省苏州市姑苏区', '吴梦', '18200223344', 88);
INSERT INTO `suppliers` VALUES (12, '天然农产品有限公司', '浙江省温州市鹿城区', '陈思', '15200334455', 87);
INSERT INTO `suppliers` VALUES (13, '绿色家园有限公司', '福建省泉州市丰泽区', '林峰', '18500445566', 89);
INSERT INTO `suppliers` VALUES (14, '丰硕农庄有限公司', '湖南省衡阳市石鼓区', '黄蓉', '18200556677', 90);
INSERT INTO `suppliers` VALUES (16, '绿茵果蔬有限公司', '江苏省扬州市广陵区', '黄继光', '18200778899', 87);
INSERT INTO `suppliers` VALUES (17, '天然绿色有限公司', '浙江省绍兴市越城区', '孙悦', '15200889911', 86);
INSERT INTO `suppliers` VALUES (18, '丰硕农业有限公司', '湖北省宜昌市西陵区', '李波', '18500991122', 88);
INSERT INTO `suppliers` VALUES (19, '绿色食品有限公司', '四川省成都市武侯区', '张薇', '18200112233', 90);
INSERT INTO `suppliers` VALUES (20, '新鲜农庄有限公司', '云南省大理白族自治州', '王小明', '15200223344', 89);
INSERT INTO `suppliers` VALUES (21, '花蜜水果集团', '广西壮族自治区百色市靖西市', '钟灵韵', '15287320419', 90);

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
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (3, 'CZKJ901704561805', 'Cukor', '男', '13746510506', '广东省深圳市福田区下梅林A栋516', '2000-08-07 00:00:00');
INSERT INTO `users` VALUES (6, 'CZKJ341708940332', 'Coco', '女', '13789720842', '广东省深圳市福田区下梅林', '2005-11-17 00:00:00');
INSERT INTO `users` VALUES (7, 'CZKJ201709005881', '吴宣仪', '女', '18777263907', '海南省海口市', '1995-03-14 00:00:00');
INSERT INTO `users` VALUES (9, 'CZKJ441710302534', '方大同', '男', '18789720403', '中国香港', '1982-07-27 00:00:00');
INSERT INTO `users` VALUES (10, 'CZKJ691710304250', '陶喆', '男', '18808426872', '北京市', '1969-07-11 00:00:00');
INSERT INTO `users` VALUES (11, 'CZKJ561710304527', '王力宏', '男', '13589750842', '美国洛杉矶', '1972-06-28 00:00:00');
INSERT INTO `users` VALUES (13, 'CZKJ11710304892', 'KhalilFong', '男', '18787720872', '云南省怒江市', '1982-05-26 00:00:00');
INSERT INTO `users` VALUES (14, 'CZKJ151710304943', 'DavidTao', '男', '17208760506', '中国台湾', '1969-07-11 00:00:00');
INSERT INTO `users` VALUES (15, 'CZKJ841710304988', 'LeehomWang', '男', '17208760403', '中国台湾', '1973-07-21 00:00:00');
INSERT INTO `users` VALUES (16, 'CZKJ541710305023', 'JayChou', '男', '17285730403', '中国台湾', '1979-08-15 00:00:00');
INSERT INTO `users` VALUES (17, 'CZKJ531710305097', 'viva', '女', '13788660412', '日本东京', '1975-07-20 00:00:00');
INSERT INTO `users` VALUES (18, 'CZKJ491712129368', '李心艾', '女', '13403421319', '广东省深圳市光明区', '2001-10-28 00:00:00');
INSERT INTO `users` VALUES (19, 'CZKJ641712129852', '龙玉颖', '女', '18777620379', '广西壮族自治区柳州市城中区', '2002-10-06 00:00:00');
INSERT INTO `users` VALUES (20, 'CZKJ171713011483', '陆漫漫', '女', '13788145032', '广东省潮州市', '2001-08-26 23:03:10');
INSERT INTO `users` VALUES (21, 'CZKJ321713016563', '陆琪琪', '女', '15789741430', '广东省潮州市', '2001-08-26 23:03:10');

-- ----------------------------
-- Table structure for warehouse
-- ----------------------------
DROP TABLE IF EXISTS `warehouse`;
CREATE TABLE `warehouse`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '仓库ID',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '仓库地址',
  `capacity` double NULL DEFAULT NULL COMMENT '仓库容量，单位立方米',
  `remaining` double NULL DEFAULT NULL COMMENT '仓库剩余容量',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of warehouse
-- ----------------------------
INSERT INTO `warehouse` VALUES (1, '广西壮族自治区百色市', 100.5, 60.2);
INSERT INTO `warehouse` VALUES (2, '广东省汕头市', 121.7, 80.3);
INSERT INTO `warehouse` VALUES (3, '海南省海口市', 150, 99.7);
INSERT INTO `warehouse` VALUES (4, '福建省泉州市', 80, 63.4);
INSERT INTO `warehouse` VALUES (5, '江西省南昌市', 97.2, 32.4);
INSERT INTO `warehouse` VALUES (6, '湖南省长沙市', 110.3, 38.2);
INSERT INTO `warehouse` VALUES (7, '贵州省贵阳市', 130, 50.7);
INSERT INTO `warehouse` VALUES (8, '云南省昆明市', 140, 75.2);
INSERT INTO `warehouse` VALUES (9, '四川省成都市', 110, 20);
INSERT INTO `warehouse` VALUES (11, '广东省深圳市', 103.2, 62.7);
INSERT INTO `warehouse` VALUES (12, '广西壮族自治区靖西市', 150, 150);

SET FOREIGN_KEY_CHECKS = 1;
