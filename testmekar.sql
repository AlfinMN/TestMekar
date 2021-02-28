-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: testmekar
-- ------------------------------------------------------
-- Server version	8.0.20

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
-- Table structure for table `tb_account`
--

DROP TABLE IF EXISTS `tb_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_account` (
  `Id_account` varchar(36) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`Id_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_account`
--

LOCK TABLES `tb_account` WRITE;
/*!40000 ALTER TABLE `tb_account` DISABLE KEYS */;
INSERT INTO `tb_account` VALUES ('04c7dd7b-7962-11eb-b583-2cfda12ae845','apnchocs@gmail.com','f9fa7d940601a480b306e371d014a998');
/*!40000 ALTER TABLE `tb_account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_pekerjaan`
--

DROP TABLE IF EXISTS `tb_pekerjaan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_pekerjaan` (
  `Id_pekerjaan` varchar(36) NOT NULL,
  `pekerjaan` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Id_pekerjaan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_pekerjaan`
--

LOCK TABLES `tb_pekerjaan` WRITE;
/*!40000 ALTER TABLE `tb_pekerjaan` DISABLE KEYS */;
INSERT INTO `tb_pekerjaan` VALUES ('10e84107-798d-11eb-b583-2cfda12ae845','WIRSWASTA'),('75abe938-7960-11eb-b583-2cfda12ae845','PNS'),('f704c058-797d-11eb-b583-2cfda12ae845','WIRAUSAHA');
/*!40000 ALTER TABLE `tb_pekerjaan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_pendidikan`
--

DROP TABLE IF EXISTS `tb_pendidikan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_pendidikan` (
  `Id_pendidikan` varchar(36) NOT NULL,
  `pendidikan` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Id_pendidikan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_pendidikan`
--

LOCK TABLES `tb_pendidikan` WRITE;
/*!40000 ALTER TABLE `tb_pendidikan` DISABLE KEYS */;
INSERT INTO `tb_pendidikan` VALUES ('e206a586-7960-11eb-b583-2cfda12ae845','SD'),('e97d7e13-798c-11eb-b583-2cfda12ae845','SMA'),('eb24c61f-797d-11eb-b583-2cfda12ae845','SMP'),('ed2344d7-798c-11eb-b583-2cfda12ae845','DIPLOMA'),('f1853b73-798c-11eb-b583-2cfda12ae845','SARJANA'),('f85f2c61-798c-11eb-b583-2cfda12ae845','MAGISTER'),('fee51a51-798c-11eb-b583-2cfda12ae845','DOKTOR');
/*!40000 ALTER TABLE `tb_pendidikan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_users`
--

DROP TABLE IF EXISTS `tb_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_users` (
  `Id_user` varchar(36) NOT NULL,
  `username` varchar(55) DEFAULT NULL,
  `no_ktp` varchar(45) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `id_pekerjaan` varchar(36) DEFAULT NULL,
  `id_pendidikan` varchar(36) DEFAULT NULL,
  `user_status` int DEFAULT '1',
  PRIMARY KEY (`Id_user`),
  UNIQUE KEY `Id_user_UNIQUE` (`Id_user`),
  KEY `pendidikan_idx` (`id_pendidikan`),
  KEY `pekerjaan_idx` (`id_pekerjaan`),
  CONSTRAINT `pekerjaan` FOREIGN KEY (`id_pekerjaan`) REFERENCES `tb_pekerjaan` (`Id_pekerjaan`),
  CONSTRAINT `pendidikan` FOREIGN KEY (`id_pendidikan`) REFERENCES `tb_pendidikan` (`Id_pendidikan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_users`
--

LOCK TABLES `tb_users` WRITE;
/*!40000 ALTER TABLE `tb_users` DISABLE KEYS */;
INSERT INTO `tb_users` VALUES ('1df8bed1-c7f3-4848-9e96-d2abe1af7d25','kim','321321321','1990-01-01','f704c058-797d-11eb-b583-2cfda12ae845','ed2344d7-798c-11eb-b583-2cfda12ae845',1),('54326ee3-797f-11eb-b583-2cfda12ae845','usman','321324552525','1997-05-11','75abe938-7960-11eb-b583-2cfda12ae845','eb24c61f-797d-11eb-b583-2cfda12ae845',1),('671add11-6e15-407a-a035-1ec171ab931f','ozil','321324321','1999-08-08','75abe938-7960-11eb-b583-2cfda12ae845','e206a586-7960-11eb-b583-2cfda12ae845',1),('c90e6894-795f-11eb-b583-2cfda12ae845','alfin','32132456788','1998-08-10','75abe938-7960-11eb-b583-2cfda12ae845','e206a586-7960-11eb-b583-2cfda12ae845',1);
/*!40000 ALTER TABLE `tb_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'testmekar'
--

--
-- Dumping routines for database 'testmekar'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-02-28 15:12:13
