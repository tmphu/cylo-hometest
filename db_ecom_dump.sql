-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: db_ecom
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `category` longtext,
  `price` bigint unsigned DEFAULT NULL,
  `featured_image` longtext,
  `rating` float DEFAULT NULL,
  `gender` longtext,
  `brand` longtext,
  `color` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (3,'2023-03-16 22:21:05.468','2023-03-16 22:21:05.468',NULL,'Male jacket premium black','Jacket',1500,'https://salt.tikicdn.com/cache/750x750/ts/product/95/d2/30/5b26cfc7c646326a13d30d3e9e1148db.jpg.webp',4.5,'Male','Nike','Black'),(4,'2023-03-16 22:23:22.054','2023-03-16 22:23:22.054',NULL,'Male jacket premium gray','Jacket',1700,'https://salt.tikicdn.com/cache/750x750/ts/product/4a/4b/11/d2717e65c7829314381f4240b4450a1f.jpg.webp',4.2,'Male','Nike','Gray'),(5,'2023-03-16 22:24:34.858','2023-03-16 22:24:34.858',NULL,'Male jacket premium white','Jacket',1900,'https://salt.tikicdn.com/cache/750x750/ts/product/c2/10/48/76aef762b7be402b914ffa9c39886b3e.jpg.webp',4.3,'Male','Nike','White'),(6,'2023-03-16 22:25:08.181','2023-03-16 22:25:08.181',NULL,'Male jacket premium navy','Jacket',1850,'https://salt.tikicdn.com/cache/750x750/ts/product/06/ed/31/98e1abbd025f049031768eb75f51aabf.jpg.webp',4.6,'Male','Nike','Navy'),(7,'2023-03-16 22:26:28.374','2023-03-16 22:26:28.374',NULL,'Male short gray','Short',750,'https://salt.tikicdn.com/cache/750x750/ts/product/56/90/d9/1162f4e3139acfc2f0b96efee82e8a9c.jpg.webp',4.9,'Male','Coolmate','Gray'),(8,'2023-03-16 22:26:58.276','2023-03-16 22:26:58.276',NULL,'Male short white','Short',800,'https://salt.tikicdn.com/cache/750x750/ts/product/24/4e/3a/c7357c8919dc6f0d7edaf949b0aa9e36.jpg.webp',4.9,'Male','Coolmate','White'),(9,'2023-03-16 22:28:33.926','2023-03-16 22:28:33.926',NULL,'Female T-shirt navy','Shirt',1200,'https://salt.tikicdn.com/cache/750x750/ts/product/f2/b2/07/bf792bbe4044fdd37f88e68b38ec4bb9.jpg.webp',4.8,'Female','Uniqlo','Navy'),(10,'2023-03-16 22:29:05.373','2023-03-16 22:29:05.373',NULL,'Female T-shirt orange','Shirt',1300,'https://salt.tikicdn.com/cache/750x750/ts/product/ca/bd/60/724c61488f8c4c4bc1873e68741f2c4e.jpg.webp',4.7,'Female','Uniqlo','Orange'),(11,'2023-03-16 22:29:33.874','2023-03-16 22:29:33.874',NULL,'Female T-shirt black','Shirt',1400,'https://salt.tikicdn.com/cache/750x750/ts/product/b6/41/3b/9837f3f0bc324d5727e4a738b170b66a.jpg.webp',4.7,'Female','Uniqlo','Black'),(12,'2023-03-16 22:30:01.135','2023-03-16 22:30:01.135',NULL,'Female T-shirt brown','Shirt',1450,'https://salt.tikicdn.com/cache/750x750/ts/product/7b/be/c0/0d62c77b66ed18d5a54ad56767e6b980.jpg.webp',4.7,'Female','Uniqlo','Brown'),(13,'2023-03-16 22:34:37.875','2023-03-16 22:34:37.875',NULL,'Female pants pink ','Pants',2400,'https://salt.tikicdn.com/cache/750x750/ts/product/ef/4d/dc/4a1e32e115bb5c86d35d7daabc149e43.jpg.webp',4.7,'Female','Zara','Pink'),(14,'2023-03-16 22:35:06.427','2023-03-16 22:35:06.427',NULL,'Female pants brown','Pants',2500,'https://salt.tikicdn.com/cache/750x750/ts/product/28/8f/2f/54640d833358a13372bac73fd718a046.jpg.webp',4.3,'Female','Zara','Brown'),(15,'2023-03-16 22:35:47.394','2023-03-16 22:35:47.394',NULL,'Female pants orange','Pants',2600,'https://salt.tikicdn.com/cache/750x750/ts/product/d1/c3/f3/d20779af7d350f49cab82b2e1b7dae52.jpg.webp',4.4,'Female','Zara','Orange'),(16,'2023-03-16 23:13:27.726','2023-03-16 23:13:27.726',NULL,'Female pants black','Pants',2600,'https://salt.tikicdn.com/cache/750x750/ts/product/3a/7a/a0/f6f85289cab70a186ebf919009032459.jpg.webp',4.6,'Female','Zara','Black'),(17,'2023-03-17 10:27:59.012','2023-03-17 10:27:59.012',NULL,'Female skinny jeans denim','Pants',2850,'https://salt.tikicdn.com/cache/750x750/ts/product/c5/de/c4/98908747ec727582f964d725e1b04ac9.jpg.webp',4.9,'Female','Gumac','Denim'),(18,'2023-03-17 10:28:51.624','2023-03-17 10:28:51.624',NULL,'Female skinny jeans navy','Pants',2850,'https://salt.tikicdn.com/cache/750x750/ts/product/ca/84/e2/457f82c4230a367c2213bee003a9da7a.jpg.webp',4.8,'Female','Gumac','Navy'),(19,'2023-03-17 10:31:40.757','2023-03-17 10:31:40.757',NULL,'Female skinny jeans black','Pants',2900,'https://salt.tikicdn.com/cache/750x750/ts/product/b0/6e/a6/1f9f0b01d5152cb51b3c046d646e4ac3.jpg.webp',4.8,'Female','Gumac','Black'),(29,'2023-03-19 23:09:51.386','2023-03-19 23:09:51.386',NULL,'Male skinny jeans denim','Pants',2100,'https://salt.tikicdn.com/cache/750x750/ts/product/13/73/92/90b540d58e6497d780fa897acb649209.jpg.webp',4.7,'Male','Lados','Denim');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-20  0:04:05
