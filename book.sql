/*
 Navicat Premium Data Transfer

 Source Server         : 密码(biglight)
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : book

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 06/04/2020 20:32:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bk_bookshelf
-- ----------------------------
DROP TABLE IF EXISTS `bk_bookshelf`;
CREATE TABLE `bk_bookshelf` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `book_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书本名称',
  `author` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '作者',
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `hub_id` tinyint(1) DEFAULT NULL COMMENT '书源ID',
  `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '章节链接,已读为已读章节链接,未读为第一章链接',
  `domain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书本链接',
  `img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书本封面',
  `created` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `link` (`link`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户书架表';

-- ----------------------------
-- Records of bk_bookshelf
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for bk_content
-- ----------------------------
DROP TABLE IF EXISTS `bk_content`;
CREATE TABLE `bk_content` (
  `id` int NOT NULL AUTO_INCREMENT,
  `root` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '根',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '名字',
  `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '正文',
  `s_page` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '上一页',
  `x_page` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '下一页',
  `list` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '章节链接',
  `domain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '域名',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='文章详情节点表';

-- ----------------------------
-- Records of bk_content
-- ----------------------------
BEGIN;
INSERT INTO `bk_content` VALUES (1, '//div[@class=\"box_con\"]', '//div[@class=\"bookname\"]/h1', '//div[@id=\"content\"]', '//div[@class=\"bottem1\"]/a[2]/@href', '//div[@class=\"bottem1\"]/a[4]/@href', '//div[@class=\"bottem1\"]/a[3]/@href', NULL, '2020-03-30 17:25:19', '2020-03-30 17:26:34');
INSERT INTO `bk_content` VALUES (2, '//div[@class=\"box_con\"]', '//div[@class=\"bookname\"]/h1', '//div[@id=\"content\"]', '//div[@class=\"bottem1\"]/a[2]/@href', '//div[@class=\"bottem1\"]/a[4]/@href', '//div[@class=\"bottem1\"]/a[3]/@href', NULL, '2020-03-31 09:32:44', '2020-03-31 09:39:45');
INSERT INTO `bk_content` VALUES (3, '//div[@class=\"box_con\"]', '//div[@class=\"bookname\"]/h1', '//div[@id=\"content\"]', '//div[@class=\"bottem1\"]/a[2]/@href', '//div[@class=\"bottem1\"]/a[4]/@href', '//div[@class=\"bottem1\"]/a[3]/@href', 'http://www.vipzw.com', '2020-03-31 09:32:44', '2020-03-31 10:00:42');
COMMIT;

-- ----------------------------
-- Table structure for bk_hub
-- ----------------------------
DROP TABLE IF EXISTS `bk_hub`;
CREATE TABLE `bk_hub` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `book_hub` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书源链接',
  `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '链接',
  `root` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '根节点',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书名节点',
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '作者节点',
  `new_list` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '最新章节',
  `new_list_link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '最新章节链接',
  `renew_time` varchar(225) DEFAULT NULL COMMENT '最后更新时间',
  `status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书本状态',
  `image` varchar(255) DEFAULT NULL COMMENT '图片节点',
  `suffix` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书源完整链接',
  `mark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书源名称',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='书源仓库';

-- ----------------------------
-- Records of bk_hub
-- ----------------------------
BEGIN;
INSERT INTO `bk_hub` VALUES (1, 'https://www.kuxiaoshuo.com', './td[1]/a/@href', '//tbody/tr[1]/following-sibling::tr', './td[1]/a', './td[3]', './td[2]/a', './td[2]/a/@href', '/td[5]', '/td[6]', NULL, 'https://www.kuxiaoshuo.com/modules/article/search.php?searchkey=', '酷小说', '2020-03-30 17:49:58', '2020-04-05 12:26:34');
INSERT INTO `bk_hub` VALUES (2, 'https://www.biquge.info', './td[1]/a/@href', '//tbody/tr[1]/following-sibling::tr', './td[1]/a', './td[3]', './td[2]/a', './td[2]/a/@href', '/td[5]', '/td[6]', NULL, 'https://www.biquge.info/modules/article/search.php?searchkey=', '笔趣阁', '2020-03-31 09:06:00', '2020-04-05 12:26:39');
INSERT INTO `bk_hub` VALUES (3, 'http://www.vipzw.com', './dl/dt/a/@href', '//div[@class=\"item\"]', './dl/dt/a', './dl/dt/span', NULL, NULL, NULL, NULL, './div/a/img/@src', 'http://www.vipzw.com/search.php?searchkey=', 'VIP笔趣阁', '2020-03-31 09:07:27', '2020-04-05 15:39:54');
COMMIT;

-- ----------------------------
-- Table structure for bk_list
-- ----------------------------
DROP TABLE IF EXISTS `bk_list`;
CREATE TABLE `bk_list` (
  `id` int NOT NULL AUTO_INCREMENT,
  `root` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '根',
  `link` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '链接',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '章节名字',
  `domain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '域名',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='书库目录节点表';

-- ----------------------------
-- Records of bk_list
-- ----------------------------
BEGIN;
INSERT INTO `bk_list` VALUES (1, '//dt[2]/following-sibling::dd', './a/@href', './a', 'https://www.kuxiaoshuo.com/', '2020-03-30 17:03:35', '2020-03-31 09:30:53');
INSERT INTO `bk_list` VALUES (2, '//div[@id=\"list\"]/dl/dd', './a/@href', './a', 'https://www.biquge.info/', '2020-03-31 09:15:14', '2020-03-31 09:30:50');
INSERT INTO `bk_list` VALUES (3, '//dt[2]/following-sibling::dd', './a/@href', './a', 'http://www.vipzw.com', '2020-03-31 09:42:49', '2020-03-31 09:49:13');
COMMIT;

-- ----------------------------
-- Table structure for bk_synopsis
-- ----------------------------
DROP TABLE IF EXISTS `bk_synopsis`;
CREATE TABLE `bk_synopsis` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `root` varchar(255) DEFAULT NULL COMMENT '根节点',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '书名',
  `writer` varchar(255) DEFAULT NULL COMMENT '作者节点',
  `img` varchar(255) DEFAULT NULL COMMENT '图片节点',
  `synopsis` varchar(255) DEFAULT NULL COMMENT '简介节点',
  `renew_time` varchar(255) DEFAULT NULL COMMENT '更新时间节点',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of bk_synopsis
-- ----------------------------
BEGIN;
INSERT INTO `bk_synopsis` VALUES (1, '//div[@id=\"wrapper\"]', '//div[@id=\"info\"]/h1', '//div[@id=\"info\"]/p[1]', '//div[@id=\"fmimg\"]/img/@src', '//div[@id=\"intro\"]/p[1]', '//div[@id=\"info\"]/p[3]', '2020-04-05 15:53:55', '2020-04-05 16:20:54');
INSERT INTO `bk_synopsis` VALUES (2, '//div[@id=\"wrapper\"]', '//div[@id=\"info\"]/h1', '//div[@id=\"info\"]/p[1]', '//div[@id=\"fmimg\"]/img/@src', '//div[@id=\"intro\"]/p[1]', '//div[@id=\"info\"]/p[3]', '2020-04-05 15:53:55', '2020-04-05 16:20:50');
INSERT INTO `bk_synopsis` VALUES (3, '//div[@id=\"wrapper\"]', '//div[@id=\"info\"]/h1', '//div[@id=\"info\"]/p[1]', '//div[@id=\"fmimg\"]/img/@src', '//div[@id=\"intro\"]/p[1]', '//div[@id=\"info\"]/p[3]', '2020-04-05 15:53:55', '2020-04-05 16:20:46');
COMMIT;

-- ----------------------------
-- Table structure for bk_users
-- ----------------------------
DROP TABLE IF EXISTS `bk_users`;
CREATE TABLE `bk_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户名',
  `pwd` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '密码',
  `mobile` int DEFAULT NULL COMMENT '手机号',
  `email` varchar(225) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '邮箱',
  `created` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`id`) USING BTREE,
  KEY `name` (`name`) USING BTREE,
  KEY `mobile` (`mobile`) USING BTREE,
  KEY `email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户表';

-- ----------------------------
-- Records of bk_users
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
