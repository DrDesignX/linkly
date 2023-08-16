/*
 Navicat Premium Data Transfer

 Source Server         : linkly
 Source Server Type    : MySQL
 Source Server Version : 100424
 Source Host           : localhost:3306
 Source Schema         : linkly

 Target Server Type    : MySQL
 Target Server Version : 100424
 File Encoding         : 65001

 Date: 20/07/2023 00:56:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for links
-- ----------------------------
DROP TABLE IF EXISTS `links`;
CREATE TABLE `links`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `redirect` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `linkly` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `clicked` int UNSIGNED NOT NULL DEFAULT 0,
  `is_available` smallint NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `user_id` int NOT NULL,
  `expiration_date` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`, `linkly`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  CONSTRAINT `links_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `is_active` tinyint(1) NULL DEFAULT NULL,
  `is_admin` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `email`(`email` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
