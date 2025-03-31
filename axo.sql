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
  `cover_id` bigint unsigned DEFAULT NULL COMMENT '封面ID',
  `path` varchar(255) NOT NULL COMMENT '路径',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `profile` json NOT NULL COMMENT '简介',
  `admin` tinyint(1) NOT NULL COMMENT '仅允许管理员',
  `pinned` tinyint(1) NOT NULL COMMENT '是否置顶',
  `creator_id` bigint unsigned DEFAULT NULL COMMENT '创建者',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_albums_title` (`title`),
  UNIQUE KEY `uni_albums_path` (`path`),
  KEY `idx_albums_cover_id` (`cover_id`),
  KEY `idx_albums_creator_id` (`creator_id`),
  CONSTRAINT `fk_albums_cover_id_images` FOREIGN KEY (`cover_id`) REFERENCES `images` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_albums_creator_id_users` FOREIGN KEY (`creator_id`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `albums`
--

LOCK TABLES `albums` WRITE;
/*!40000 ALTER TABLE `albums` DISABLE KEYS */;
/*!40000 ALTER TABLE `albums` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `forum_groups`
--

DROP TABLE IF EXISTS `forum_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `forum_groups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `label` varchar(255) NOT NULL COMMENT '标题',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_forum_groups_label` (`label`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forum_groups`
--

LOCK TABLES `forum_groups` WRITE;
/*!40000 ALTER TABLE `forum_groups` DISABLE KEYS */;
INSERT INTO `forum_groups` VALUES (3,'反馈分区'),(2,'游戏交流分区'),(1,'闲聊分区');
/*!40000 ALTER TABLE `forum_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `forums`
--

DROP TABLE IF EXISTS `forums`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `forums` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `forum_group_id` bigint unsigned NOT NULL COMMENT '论坛组ID',
  `path` varchar(255) NOT NULL COMMENT '路径',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `profile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '简介',
  `cover_id` bigint unsigned DEFAULT NULL COMMENT '封面ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_forums_path` (`path`),
  UNIQUE KEY `uni_forums_title` (`title`),
  KEY `idx_forums_forum_group_id` (`forum_group_id`),
  KEY `idx_forums_cover_id` (`cover_id`),
  CONSTRAINT `fk_forums_cover_id_images` FOREIGN KEY (`cover_id`) REFERENCES `images` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_forums_forum_group_id_forum_groups` FOREIGN KEY (`forum_group_id`) REFERENCES `forum_groups` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `forums`
--

LOCK TABLES `forums` WRITE;
/*!40000 ALTER TABLE `forums` DISABLE KEYS */;
INSERT INTO `forums` VALUES (1,1,'chat','聊天室','在这里闲聊',1),(2,2,'mcbe','我的世界基岩版','讨论我的世界基岩版',2),(3,2,'mcje','我的世界Java版','讨论我的世界Java版',3),(4,2,'dst','饥荒联机版','讨论饥荒联机版',4),(5,2,'terraria','泰拉瑞亚','讨论泰拉瑞亚',5),(6,2,'stardew','星露谷物语','讨论星露谷物语',6),(7,2,'sky','光遇','讨论光遇',7),(8,3,'feedback','意见反馈','在这里反馈意见',8),(9,3,'bug','BUG报告','在这里报告BUG',9);
/*!40000 ALTER TABLE `forums` ENABLE KEYS */;
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
  `name` varchar(255) NOT NULL COMMENT '公会名',
  `count` bigint unsigned NOT NULL COMMENT '公会人数',
  `avatar_id` bigint unsigned DEFAULT NULL COMMENT 'LOGO图片ID',
  `cover_id` bigint unsigned DEFAULT NULL COMMENT '背景图片ID',
  `profile` json NOT NULL COMMENT '公会介绍',
  `notice` json NOT NULL COMMENT '公会公告',
  `money` bigint unsigned NOT NULL COMMENT '公会资金',
  `level` bigint unsigned NOT NULL COMMENT '公会等级',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_guilds_name` (`name`),
  KEY `idx_guilds_avatar_id` (`avatar_id`),
  KEY `idx_guilds_cover_id` (`cover_id`),
  CONSTRAINT `fk_guilds_avatar_id_images` FOREIGN KEY (`avatar_id`) REFERENCES `images` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_guilds_cover_id_images` FOREIGN KEY (`cover_id`) REFERENCES `images` (`id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `guilds`
--

LOCK TABLES `guilds` WRITE;
/*!40000 ALTER TABLE `guilds` DISABLE KEYS */;
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
  `filename` varchar(255) NOT NULL COMMENT '文件名',
  `label` varchar(255) DEFAULT NULL COMMENT '标题',
  `profile` varchar(255) DEFAULT NULL COMMENT '简介',
  `likes` bigint unsigned DEFAULT NULL COMMENT '点赞',
  `uploader_id` bigint unsigned DEFAULT NULL COMMENT '上传者用户ID',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户照片ID',
  `album_id` bigint unsigned DEFAULT NULL COMMENT '相册ID',
  `guild_id` bigint unsigned DEFAULT NULL COMMENT '公会ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_images_filename` (`filename`),
  KEY `idx_images_user_id` (`user_id`),
  KEY `idx_images_album_id` (`album_id`),
  KEY `idx_images_guild_id` (`guild_id`),
  KEY `idx_images_uploader_id` (`uploader_id`),
  CONSTRAINT `fk_images_album_id_albums` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_images_guild_id_guilds` FOREIGN KEY (`guild_id`) REFERENCES `guilds` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_images_uploader_id_users` FOREIGN KEY (`uploader_id`) REFERENCES `users` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_images_user_id_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `images`
--

LOCK TABLES `images` WRITE;
/*!40000 ALTER TABLE `images` DISABLE KEYS */;
INSERT INTO `images` VALUES (1,'2025-03-29 03:57:26.000','forum-chat.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(2,'2025-03-29 03:57:26.000','forum-mcbe.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(3,'2025-03-29 03:57:26.000','forum-mcje.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(4,'2025-03-29 03:57:26.000','forum-dst.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(5,'2025-03-29 03:57:26.000','forum-terraria.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(6,'2025-03-29 03:57:26.000','forum-stardew.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(7,'2025-03-29 03:57:26.000','forum-sky.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(8,'2025-03-29 03:57:26.000','forum-feedback.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL),(9,'2025-03-29 03:57:26.000','forum-bug.jpg',NULL,NULL,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `images` ENABLE KEYS */;
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
  PRIMARY KEY (`id`),
  KEY `idx_onlines_server` (`server`)
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
  `pinned` tinyint(1) NOT NULL COMMENT '是否置顶',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `forum_id` bigint unsigned DEFAULT NULL COMMENT '论坛ID',
  `guild_id` bigint unsigned DEFAULT NULL COMMENT '公会ID',
  `content` json NOT NULL COMMENT '原内容',
  `author_id` bigint unsigned DEFAULT NULL COMMENT '作者ID',
  PRIMARY KEY (`id`),
  KEY `idx_posts_forum_id` (`forum_id`),
  KEY `idx_posts_guild_id` (`guild_id`),
  KEY `idx_posts_author_id` (`author_id`),
  CONSTRAINT `fk_posts_author_id_users` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_posts_forum_id_forums` FOREIGN KEY (`forum_id`) REFERENCES `forums` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_posts_guild_id_guilds` FOREIGN KEY (`guild_id`) REFERENCES `guilds` (`id`) ON DELETE CASCADE
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
-- Table structure for table `prop_types`
--

DROP TABLE IF EXISTS `prop_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `prop_types` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `label` varchar(255) NOT NULL COMMENT '道具名称',
  `profile` text NOT NULL COMMENT '道具简介',
  `icon_id` bigint unsigned DEFAULT NULL COMMENT '图标ID',
  `function` varchar(255) NOT NULL COMMENT '功能',
  `params` json NOT NULL COMMENT '功能参数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_prop_types_label` (`label`),
  KEY `idx_prop_types_icon_id` (`icon_id`),
  CONSTRAINT `fk_prop_types_icon_id_images` FOREIGN KEY (`icon_id`) REFERENCES `images` (`id`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prop_types`
--

LOCK TABLES `prop_types` WRITE;
/*!40000 ALTER TABLE `prop_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `prop_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `props`
--

DROP TABLE IF EXISTS `props`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `props` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `owner_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `guild_id` bigint unsigned DEFAULT NULL COMMENT '公会ID',
  `prop_type_id` bigint unsigned NOT NULL COMMENT '道具类型ID',
  `count` bigint unsigned NOT NULL COMMENT '数量',
  PRIMARY KEY (`id`),
  KEY `idx_props_prop_type_id` (`prop_type_id`),
  KEY `idx_props_owner_id` (`owner_id`),
  KEY `idx_props_guild_id` (`guild_id`),
  CONSTRAINT `fk_props_guild_id_guilds` FOREIGN KEY (`guild_id`) REFERENCES `guilds` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_props_owner_id_users` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_props_prop_type_id_prop_types` FOREIGN KEY (`prop_type_id`) REFERENCES `prop_types` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `props`
--

LOCK TABLES `props` WRITE;
/*!40000 ALTER TABLE `props` DISABLE KEYS */;
/*!40000 ALTER TABLE `props` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reviews`
--

DROP TABLE IF EXISTS `reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reviews` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `content` json NOT NULL COMMENT '源内容',
  `attitude` tinyint(1) DEFAULT NULL COMMENT '态度',
  `author_id` bigint unsigned DEFAULT NULL COMMENT '作者ID',
  `album_id` bigint unsigned DEFAULT NULL COMMENT '相册ID',
  `post_id` bigint unsigned DEFAULT NULL COMMENT '帖子ID',
  `review_id` bigint unsigned DEFAULT NULL COMMENT '评论ID',
  PRIMARY KEY (`id`),
  KEY `idx_reviews_album_id` (`album_id`),
  KEY `idx_reviews_post_id` (`post_id`),
  KEY `idx_reviews_review_id` (`review_id`),
  KEY `idx_reviews_author_id` (`author_id`),
  CONSTRAINT `fk_reviews_author_id_users` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_reviews_post_id_posts` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_reviews_review_id_reviews` FOREIGN KEY (`review_id`) REFERENCES `reviews` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviews`
--

LOCK TABLES `reviews` WRITE;
/*!40000 ALTER TABLE `reviews` DISABLE KEYS */;
/*!40000 ALTER TABLE `reviews` ENABLE KEYS */;
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
  `name` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `avatar_id` bigint unsigned DEFAULT NULL COMMENT '头像图片',
  `cover_id` bigint unsigned DEFAULT NULL COMMENT '封面图片',
  `admin` tinyint(1) NOT NULL COMMENT '是否管理员',
  `voter` tinyint(1) NOT NULL COMMENT '是否议员',
  `is_male` tinyint(1) DEFAULT NULL COMMENT '性别',
  `profile` json NOT NULL COMMENT '个人介绍',
  `birthday` datetime(3) DEFAULT NULL COMMENT '生日',
  `location` varchar(255) DEFAULT NULL COMMENT '地址',
  `daily_coin` bigint unsigned NOT NULL COMMENT '签到币',
  `honor_coin` bigint unsigned NOT NULL COMMENT '贡献币',
  `checkin` bigint NOT NULL COMMENT '签到记录',
  `email` varchar(255) NOT NULL COMMENT '邮箱',
  `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
  `qq` varchar(255) DEFAULT NULL COMMENT 'QQ号',
  `mcbe_name` varchar(255) DEFAULT NULL COMMENT 'MCBE用户名',
  `mcje_name` varchar(255) DEFAULT NULL COMMENT 'MCJE用户名',
  `guild_id` bigint unsigned DEFAULT NULL COMMENT '所属公会',
  `guild_role` bigint unsigned DEFAULT NULL COMMENT '公会身份角色',
  `donation` bigint unsigned NOT NULL COMMENT '捐赠数额',
  `exp` bigint unsigned NOT NULL COMMENT '经验值',
  `setting` json NOT NULL COMMENT '用户设置',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_name` (`name`),
  UNIQUE KEY `uni_users_email` (`email`),
  UNIQUE KEY `uni_users_phone` (`phone`),
  UNIQUE KEY `uni_users_qq` (`qq`),
  UNIQUE KEY `uni_users_mcbe_name` (`mcbe_name`),
  UNIQUE KEY `uni_users_mcje_name` (`mcje_name`),
  KEY `idx_users_avatar_id` (`avatar_id`),
  KEY `idx_users_cover_id` (`cover_id`),
  KEY `idx_users_guild_id` (`guild_id`),
  CONSTRAINT `fk_users_avatar_id_images` FOREIGN KEY (`avatar_id`) REFERENCES `images` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_users_cover_id_images` FOREIGN KEY (`cover_id`) REFERENCES `images` (`id`) ON DELETE RESTRICT,
  CONSTRAINT `fk_users_guild_id_guilds` FOREIGN KEY (`guild_id`) REFERENCES `guilds` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wiki_groups`
--

DROP TABLE IF EXISTS `wiki_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wiki_groups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `label` varchar(255) NOT NULL COMMENT '名称',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_wiki_groups_label` (`label`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wiki_groups`
--

LOCK TABLES `wiki_groups` WRITE;
/*!40000 ALTER TABLE `wiki_groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `wiki_groups` ENABLE KEYS */;
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
  `path` varchar(255) NOT NULL COMMENT '路径',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `content` json NOT NULL COMMENT '内容',
  `wiki_group_id` bigint unsigned DEFAULT NULL COMMENT '文档组',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_wikis_path` (`path`),
  UNIQUE KEY `uni_wikis_title` (`title`),
  KEY `idx_wikis_wiki_group_id` (`wiki_group_id`),
  CONSTRAINT `fk_wikis_wiki_group_id_wiki_groups` FOREIGN KEY (`wiki_group_id`) REFERENCES `wiki_groups` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wikis`
--

LOCK TABLES `wikis` WRITE;
/*!40000 ALTER TABLE `wikis` DISABLE KEYS */;
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

-- Dump completed on 2025-03-29 18:29:28
