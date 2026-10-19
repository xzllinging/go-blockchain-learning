/*
 Navicat Premium Data Transfer

 Source Server         : 1111
 Source Server Type    : MySQL
 Source Server Version : 80036
 Source Host           : localhost:3306
 Source Schema         : xzl_blog

 Target Server Type    : MySQL
 Target Server Version : 80036
 File Encoding         : 65001

 Date: 03/01/2025 10:50:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_articles_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `fk_articles_user`(`user_id` ASC) USING BTREE,
  CONSTRAINT `articles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of articles
-- ----------------------------
INSERT INTO `articles` VALUES (1, '2024-12-24 23:09:48.411', '2024-12-24 23:09:48.411', NULL, '《音符间的花开花落：音乐里的四季诗篇》', '维瓦尔第的《四季》宛如流动的诗篇。《春》中，小提琴跳跃的音符似春日苏醒的万物，轻快活泼；《冬》里，急促旋律如凛冽寒风，让人感受冬日的冷峻，音乐将四季之美娓娓道来。', 1);
INSERT INTO `articles` VALUES (2, '2024-12-24 23:09:48.414', '2024-12-24 23:09:48.414', NULL, '《从混音技巧看音乐质感提升》', '混音时，合理调节各乐器音量平衡是关键。比如，摇滚乐队混音，鼓和贝斯需奠定坚实节奏基础，吉他和键盘层次要分明，恰当使用效果器，能让音乐空间感与立体感增强。', 1);
INSERT INTO `articles` VALUES (3, '2024-12-24 23:09:48.416', '2024-12-24 23:09:48.416', NULL, '《音乐疗愈：困境中重生的旋律力量》', 'Imagine Dragons的《Demons》，在失落时听，主唱沧桑嗓音与激昂旋律交织，似在诉说困境，又传递力量。它能让听众明白痛苦可战胜，借音乐走出阴霾。', 1);
INSERT INTO `articles` VALUES (4, '2024-12-24 23:09:48.418', '2024-12-24 23:09:48.418', NULL, '《K - Pop新风尚：席卷全球的韩流音乐》', 'Blackpink的音乐融合多元风格，《DDU - DDU - DDU》节奏强烈，舞蹈与舞台魅力十足，MV视觉效果震撼，以独特造型与潮流元素，引领K - Pop新潮流。', 1);
INSERT INTO `articles` VALUES (5, '2024-12-24 23:09:48.421', '2024-12-24 23:09:48.421', NULL, '《港乐旧时光：粤语金曲的光辉岁月》', '张国荣的《Monica》，动感旋律与粤语歌词完美结合，是港乐经典。它不仅是一代人的青春记忆，更见证香港乐坛辉煌，让我们怀念那个充满活力的港乐时代。', 1);
INSERT INTO `articles` VALUES (6, '2024-12-24 23:09:48.423', '2024-12-24 23:09:48.423', NULL, '《时光留声机：80年代迪斯科音乐的狂欢》', '80年代的迪斯科音乐是一个时代的标志。例如比吉斯乐队（Bee Gees）的《Stayin\' Alive》，这首歌的旋律具有强烈的节奏感，贝斯线条流畅且极具动感，主唱的假音演唱风格独特。歌曲一经推出，就迅速在全球的迪斯科舞厅中流行开来，成为当时年轻人在舞池中尽情狂欢的必备曲目，每当听到这首歌，都仿佛能看到舞池中人们穿着喇叭裤、烫着卷发，随着音乐尽情摇摆的热闹场景。', 1);
INSERT INTO `articles` VALUES (7, '2024-12-24 23:09:48.426', '2024-12-24 23:09:48.426', NULL, '《音乐中的亲情纽带：那些温暖心灵的乐章》', '久石让为宫崎骏动画创作的音乐，很多都充满了亲情的味道。例如《龙猫》中的配乐，那悠扬的旋律，在小梅和小月与龙猫相遇时响起，音符间洋溢着姐妹间纯真的情谊和对神秘龙猫的好奇，如同一条温暖的亲情纽带，将观众与动画中的美好场景紧密相连，每次听到都能让人回忆起动画中的温馨画面，心中充满暖意。', 1);
INSERT INTO `articles` VALUES (8, '2024-12-24 23:09:48.428', '2024-12-24 23:09:48.428', NULL, '《旋律如墨，绘出音乐的山水画卷》', '中国传统民乐有着独特的韵味。比如琵琶曲《十面埋伏》，乐曲如同展开一幅古代战争的画卷，琵琶的轮指好似万马奔腾，扫弦仿佛战场上的金戈交鸣，从旋律中能感受到紧张的战斗氛围和磅礴的气势，仿佛穿越时空，置身于楚汉相争的垓下战场。', 1);
INSERT INTO `articles` VALUES (9, '2024-12-24 23:09:48.431', '2024-12-24 23:09:48.431', NULL, '《音乐与心灵的共鸣：那些触动泪点的旋律》', '阿黛尔的《Someone Like You》，其嗓音极具穿透力，唱出失恋痛苦，旋律由轻到强，勾起听众伤心回忆。《泰坦尼克号》的《My Heart Will Go On》深情诠释生死爱情，令人感动。', 1);
INSERT INTO `articles` VALUES (10, '2024-12-24 23:09:48.433', '2024-12-24 23:09:48.433', NULL, '《弦上月光：音乐中的诗意角落》', '古典音乐如肖邦的《夜曲》，音符似深夜飘落的花瓣，旋律舒缓，能让我们感受肖邦对祖国的眷恋。德彪西的《月光》营造出朦胧梦幻氛围，和声如湖波，带我们进入诗意世界。', 1);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `email` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2024-12-24 23:09:48.387', '2024-12-24 23:09:48.387', NULL, 'Random name', 'random@email.com', '$2a$12$Znuh8wS17A5HyTvN2AN4AOmtrUWykT1u21MdWrzGviOo/XrXuxy9q');
INSERT INTO `users` VALUES (2, '2024-12-31 23:46:11.263', '2024-12-31 23:46:11.263', NULL, '木木ayya', '2128283202@qq.com', '$2a$12$9mrS8tlmBZPQR4sMNGFUFux/jgv63XsbQfLn76bb8XmipYaEkiE3G');

SET FOREIGN_KEY_CHECKS = 1;
