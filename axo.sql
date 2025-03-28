-- MySQL dump 10.13  Distrib 8.0.41, for Linux (x86_64)
--
-- Host: localhost    Database: axo
-- ------------------------------------------------------
-- Server version	8.0.41-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `albums`
--

DROP TABLE IF EXISTS `albums`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `albums` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `cover` varchar(255) DEFAULT NULL COMMENT '封面文件名',
  `path` varchar(255) NOT NULL COMMENT '路径',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `description` text NOT NULL COMMENT '简介',
  `only_admin` tinyint(1) NOT NULL COMMENT '仅允许管理员',
  `pinned` tinyint(1) NOT NULL COMMENT '是否置顶',
  `guild_id` bigint unsigned DEFAULT NULL COMMENT '公会ID',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_albums_path` (`path`),
  UNIQUE KEY `uni_albums_title` (`title`),
  KEY `idx_albums_guild_id` (`guild_id`),
  KEY `idx_albums_user_id` (`user_id`),
  KEY `idx_albums_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_guilds_albums` FOREIGN KEY (`guild_id`) REFERENCES `guilds` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_users_albums` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `albums`
--

LOCK TABLES `albums` WRITE;
/*!40000 ALTER TABLE `albums` DISABLE KEYS */;
INSERT INTO `albums` VALUES (1,'2025-01-13 16:58:47.000','2025-01-13 16:58:47.000',NULL,'2023-new-year-001.jpg','2023-new-year','2023新年','2023新年',0,0,NULL,1),(2,'2025-01-21 01:17:40.000','2025-01-21 01:17:40.000',NULL,'carousel-1.jpg','carousel','走马灯','走马灯',1,1,NULL,NULL);
/*!40000 ALTER TABLE `albums` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `black_lists`
--

DROP TABLE IF EXISTS `black_lists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `black_lists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type` varchar(255) NOT NULL COMMENT '账号类型',
  `value` varchar(255) NOT NULL COMMENT '账号',
  `expiry` datetime(3) NOT NULL COMMENT '解禁时间',
  PRIMARY KEY (`id`),
  KEY `idx_black_lists_value` (`value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `black_lists`
--

LOCK TABLES `black_lists` WRITE;
/*!40000 ALTER TABLE `black_lists` DISABLE KEYS */;
/*!40000 ALTER TABLE `black_lists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `source` text NOT NULL COMMENT '源内容',
  `content` text NOT NULL COMMENT '内容',
  `attitude` tinyint NOT NULL COMMENT '态度',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '作者ID',
  `image_id` bigint unsigned DEFAULT NULL COMMENT '图片ID',
  `post_id` bigint unsigned DEFAULT NULL COMMENT '帖子ID',
  `comment_id` bigint unsigned DEFAULT NULL COMMENT '评论ID',
  PRIMARY KEY (`id`),
  KEY `idx_comments_comment_id` (`comment_id`),
  KEY `idx_comments_deleted_at` (`deleted_at`),
  KEY `idx_comments_user_id` (`user_id`),
  KEY `idx_comments_image_id` (`image_id`),
  KEY `idx_comments_post_id` (`post_id`),
  CONSTRAINT `fk_comments_comments` FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_images_comments` FOREIGN KEY (`image_id`) REFERENCES `images` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_posts_comments` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `guilds`
--

DROP TABLE IF EXISTS `guilds`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `guilds` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) NOT NULL COMMENT '公会名',
  `number` bigint unsigned NOT NULL COMMENT '公会人数',
  `logo` varchar(255) DEFAULT NULL COMMENT 'LOGO路径',
  `profile` text COMMENT '公会介绍',
  `money` bigint unsigned NOT NULL COMMENT '公会资金',
  `level` bigint unsigned NOT NULL COMMENT '公会等级',
  `notice` text COMMENT '公会公告',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_guilds_name` (`name`),
  KEY `idx_guilds_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `guilds`
--

LOCK TABLES `guilds` WRITE;
/*!40000 ALTER TABLE `guilds` DISABLE KEYS */;
INSERT INTO `guilds` VALUES (1,'2025-01-14 15:59:38.000','2025-01-14 15:59:38.000',NULL,'理塘公会',1,'guild-test.jpeg','默认公会介绍',0,0,NULL);
/*!40000 ALTER TABLE `guilds` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `images`
--

DROP TABLE IF EXISTS `images`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `images` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `filename` varchar(255) NOT NULL COMMENT '文件名',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `description` text NOT NULL COMMENT '简介',
  `likes` bigint unsigned NOT NULL COMMENT '点赞',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '上传者',
  `album_id` bigint unsigned NOT NULL COMMENT '相册ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_images_title` (`title`),
  KEY `idx_images_deleted_at` (`deleted_at`),
  KEY `idx_images_user_id` (`user_id`),
  KEY `idx_images_album_id` (`album_id`),
  CONSTRAINT `fk_albums_images` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_images_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `images`
--

LOCK TABLES `images` WRITE;
/*!40000 ALTER TABLE `images` DISABLE KEYS */;
INSERT INTO `images` VALUES (1,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-001.jpg','2023-new-year-001','',0,NULL,1),(2,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-002.jpg','2023-new-year-002','',0,NULL,1),(3,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-003.jpg','2023-new-year-003','',0,NULL,1),(4,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-004.jpg','2023-new-year-004','',0,NULL,1),(5,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-005.jpg','2023-new-year-005','',0,NULL,1),(6,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-006.jpg','2023-new-year-006','',0,NULL,1),(7,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-007.jpg','2023-new-year-007','',0,NULL,1),(8,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-008.jpg','2023-new-year-008','',0,NULL,1),(9,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-009.jpg','2023-new-year-009','',0,NULL,1),(10,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-010.jpg','2023-new-year-010','',0,NULL,1),(11,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-011.jpg','2023-new-year-011','',0,NULL,1),(12,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-012.jpg','2023-new-year-012','',0,NULL,1),(13,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-013.jpg','2023-new-year-013','',0,NULL,1),(14,'2025-01-13 17:03:26.000','2025-01-13 17:03:26.000',NULL,'2023-new-year-014.jpg','2023-new-year-014','',0,NULL,1),(15,'2025-01-21 01:23:37.000','2025-01-21 01:23:37.000',NULL,'carousel-1.jpg','走马灯一','走马灯一',0,NULL,2),(16,'2025-01-21 01:24:47.000','2025-01-21 01:24:47.000',NULL,'carousel-2.jpg','走马灯二','走马灯二',0,NULL,2),(17,'2025-01-21 01:25:19.000','2025-01-21 01:25:19.000',NULL,'carousel-3.jpg','走马灯三','走马灯三',0,NULL,2),(18,'2025-01-21 01:25:48.000','2025-01-21 01:25:48.000',NULL,'carousel-4.jpg','走马灯四','走马灯四',0,NULL,2),(19,'2025-01-21 01:25:48.000','2025-01-21 01:25:48.000',NULL,'carousel-5.jpg','走马灯五','走马灯五',0,NULL,2);
/*!40000 ALTER TABLE `images` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `limiter_records`
--

DROP TABLE IF EXISTS `limiter_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `limiter_records` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user` varchar(255) NOT NULL COMMENT '用户',
  `action` varchar(255) NOT NULL COMMENT '行为',
  `time` datetime(3) NOT NULL COMMENT '触发时间',
  PRIMARY KEY (`id`),
  KEY `idx_limiter_records_user` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `limiter_records`
--

LOCK TABLES `limiter_records` WRITE;
/*!40000 ALTER TABLE `limiter_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `limiter_records` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `onlines`
--

DROP TABLE IF EXISTS `onlines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `onlines` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `server` varchar(255) NOT NULL COMMENT '服务器',
  `count` bigint DEFAULT NULL COMMENT '在线人数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `onlines`
--

LOCK TABLES `onlines` WRITE;
/*!40000 ALTER TABLE `onlines` DISABLE KEYS */;
/*!40000 ALTER TABLE `onlines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `source` text NOT NULL COMMENT '原内容',
  `content` text NOT NULL COMMENT '内容',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '作者ID',
  `guild_id` bigint unsigned DEFAULT NULL COMMENT '公会ID',
  `forum` varchar(255) DEFAULT NULL COMMENT '论坛',
  `category` varchar(255) DEFAULT NULL COMMENT '分类',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_posts_title` (`title`),
  KEY `idx_posts_deleted_at` (`deleted_at`),
  KEY `idx_posts_user_id` (`user_id`),
  KEY `idx_posts_guild_id` (`guild_id`),
  CONSTRAINT `fk_guilds_posts` FOREIGN KEY (`guild_id`) REFERENCES `guilds` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_posts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `properties`
--

DROP TABLE IF EXISTS `properties`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `properties` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `property` longtext NOT NULL COMMENT '道具ID',
  `count` bigint unsigned NOT NULL COMMENT '数量',
  PRIMARY KEY (`id`),
  KEY `idx_properties_deleted_at` (`deleted_at`),
  KEY `idx_properties_user_id` (`user_id`),
  CONSTRAINT `fk_users_props` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `properties`
--

LOCK TABLES `properties` WRITE;
/*!40000 ALTER TABLE `properties` DISABLE KEYS */;
/*!40000 ALTER TABLE `properties` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `questions`
--

DROP TABLE IF EXISTS `questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `questions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `content` text NOT NULL COMMENT '题目内容',
  `class` varchar(255) NOT NULL COMMENT '题目分类',
  `answer` varchar(255) NOT NULL COMMENT '正确答案',
  PRIMARY KEY (`id`),
  KEY `idx_questions_deleted_at` (`deleted_at`),
  KEY `idx_questions_class` (`class`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `questions`
--

LOCK TABLES `questions` WRITE;
/*!40000 ALTER TABLE `questions` DISABLE KEYS */;
/*!40000 ALTER TABLE `questions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `profile` text COMMENT '个人简介',
  `admin` tinyint(1) NOT NULL COMMENT '管理员',
  `temp_coin` bigint unsigned NOT NULL COMMENT '签到币',
  `perm_coin` bigint unsigned NOT NULL COMMENT '蝾螈币',
  `checkin` bigint NOT NULL COMMENT '签到记录',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
  `qq` varchar(255) DEFAULT NULL COMMENT 'QQ号',
  `bedrock_name` varchar(255) DEFAULT NULL COMMENT '基岩版用户名',
  `java_name` varchar(255) DEFAULT NULL COMMENT 'Java版用户名',
  `guild_id` bigint unsigned DEFAULT NULL COMMENT '公会ID',
  `guild_role` bigint unsigned NOT NULL COMMENT '公会角色',
  `donation` bigint unsigned NOT NULL COMMENT '捐赠数额',
  `exp` bigint unsigned NOT NULL COMMENT '经验值',
  `str_meta` json NOT NULL COMMENT '字符串元数据',
  `bool_meta` json NOT NULL COMMENT '布尔元数据',
  `setting` bigint NOT NULL COMMENT '设置',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_name` (`name`),
  UNIQUE KEY `uni_users_email` (`email`),
  UNIQUE KEY `uni_users_phone` (`phone`),
  UNIQUE KEY `uni_users_qq` (`qq`),
  UNIQUE KEY `uni_users_bedrock_name` (`bedrock_name`),
  UNIQUE KEY `uni_users_java_name` (`java_name`),
  KEY `idx_users_deleted_at` (`deleted_at`),
  KEY `idx_users_guild_id` (`guild_id`),
  CONSTRAINT `fk_guilds_users` FOREIGN KEY (`guild_id`) REFERENCES `guilds` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2025-01-13 08:34:02.644','2025-01-18 12:49:40.259',NULL,'Nerakolo','1a32654cb32760ce26fc6fd08d94e48fc1feb9f7ae357aaea3cdf6672de65e6c','Nerakolo_avatar.jpg','妈妈生的',1,102,100,290873,'nerakolo@qq.com','15828517721','1285607932','Nerakolox','Nerakolo',1,4,21348971,0,'{}','{}',0),(2,'2025-01-13 17:50:39.000','2025-03-28 07:04:41.442',NULL,'Bestcb233','4e6ff5a3b3fdc65cf98c56f1734d9f9d734b3fc413a974f762e6cfcd9764ab76','Bestcb233_avatar.png','鸡你太美',1,101,100,268435556,'2892709432@qq.com','15623013552','2892709432','bestcb5843','Bestcb233',1,3,100,0,'{}','{\"PubQQ\": false, \"PubEmail\": false, \"PubGuild\": false, \"PubPhone\": false}',0);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wikis`
--

DROP TABLE IF EXISTS `wikis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wikis` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `path` varchar(255) NOT NULL COMMENT '路径',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `markdown` text NOT NULL COMMENT '内容',
  `html` text NOT NULL COMMENT 'HTML内容',
  `category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类ID',
  PRIMARY KEY (`id`),
  KEY `idx_wikis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wikis`
--

LOCK TABLES `wikis` WRITE;
/*!40000 ALTER TABLE `wikis` DISABLE KEYS */;
INSERT INTO `wikis` VALUES (1,'2025-01-21 05:17:48.000','2025-01-21 05:17:48.000',NULL,'eula','最终用户许可协议','最终用户许可协议','最终用户许可协议',''),(2,'2025-01-21 05:18:17.000','2025-01-21 05:18:17.000',NULL,'join','怎么进服务器','怎么进服务器','怎么进服务器',''),(3,'2025-01-22 18:55:16.000','2025-01-22 18:55:16.000',NULL,'be_join','怎么进基岩服','怎么进基岩服','怎么进基岩服','基岩版'),(4,'2025-01-22 18:55:52.000','2025-01-22 18:55:52.000',NULL,'be_claim','怎么圈地','怎么圈地','怎么圈地','基岩版'),(5,'2025-01-29 09:22:05.000','2025-01-29 09:22:05.000',NULL,'index','主页','## 主页','<h2>主页</h2>','');
/*!40000 ALTER TABLE `wikis` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-03-28 23:39:25
