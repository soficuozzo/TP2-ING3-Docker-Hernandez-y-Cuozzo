CREATE DATABASE  IF NOT EXISTS `users-api` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `users-api`;
-- MySQL dump 10.13  Distrib 8.0.38, for macos14 (arm64)
--
-- Host: 127.0.0.1    Database: users-api
-- ------------------------------------------------------
-- Server version	9.1.0

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
-- Table structure for table `inscripcions`
--

DROP TABLE IF EXISTS `inscripcions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inscripcions` (
  `usuario_id` bigint DEFAULT NULL,
  `curso_id` longtext
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inscripcions`
--

LOCK TABLES `inscripcions` WRITE;
/*!40000 ALTER TABLE `inscripcions` DISABLE KEYS */;
INSERT INTO `inscripcions` VALUES (1,'672fa54884420a3ffcb4ad3a'),(2,'672fa54884420a3ffcb4ad3a'),(6,'6738b9dc1fa859cd31afd2de'),(6,'6738babf1fa859cd31afd2df'),(6,'6738f22ed226609c883f4fe9'),(9,'673cccf71d66b021455be3df'),(9,'673ccd831d66b021455be3e0'),(9,'673cce031d66b021455be3e1'),(9,'673cd0bc1d66b021455be3e4'),(9,'67850e6ebd9b18d7cc9ef3bc');
/*!40000 ALTER TABLE `inscripcions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usuarios` (
  `usuario_id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) DEFAULT NULL,
  `apellido` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `passwordhash` varchar(255) DEFAULT NULL,
  `tipo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`usuario_id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuarios`
--

LOCK TABLES `usuarios` WRITE;
/*!40000 ALTER TABLE `usuarios` DISABLE KEYS */;
INSERT INTO `usuarios` VALUES (1,'Luna','Hueda','lunahueda@gmail.com','4d186321c1a7f0f354b297e8914ab240','Estudiante'),(2,'Sofia','Tula','sofitula@gmail.com','ba77a5448b1208afe6effd5194c2a8b6','Estudiante'),(3,'Silvia','Morales','silviamorales@gmail.com','4d186321c1a7f0f354b297e8914ab240','Estudiante'),(4,'Lola','Tula','lolita@gmail.com','dfb1e5ea201d8027c27fdd56d1d28f36','estudiante'),(5,'Dillom','Dillom','lolita@gmail.com','c96f7c00e46923ffbe344d497a7531ed','estudiante'),(6,'Rose','Apate','rose@gmail.com','5c967b958b1993032a02c006919fd1ab','estudiante'),(7,'Sofia','Cuozzo','sofic@gmail.com','10230dd88250bf7238dc727271320afd','estudiante'),(8,'Luciana','Hueda','lucianahueda13@gmail.com','41362191c32d0ed346d68464fef8d0ad','estudiante'),(9,'lala','lala','lala@gmail.com','81dc9bdb52d04dc20036dbd8313ed055','estudiante'),(10,'admin','admin','admin@gmail.com','0192023a7bbd73250516f069df18b500','admin'),(11,'admin 2','admin 2','admin2@gmail.com','0192023a7bbd73250516f069df18b500','admin'),(12,'Luciana','Hueda','lucianahueda14@gmail.com','2e3817293fc275dbee74bd71ce6eb056','estudiante'),(13,'lala','lala','lala123@gmail.com','2e3817293fc275dbee74bd71ce6eb056','estudiante'),(14,'Luna','Hueda','luna123@gmail.com','e2fb30683ccecdcac77c95380c0e541b','estudiante'),(15,'Sofia','Tula','sofitula123@gmail.com','202cb962ac59075b964b07152d234b70','estudiante'),(16,'Christian','Luna','christian@gmail.com','f7050fa5b63ca3f9c663f606edd93f15','estudiante'),(17,'Beyonce','Beyonce','beyonce@gmail.com','202cb962ac59075b964b07152d234b70','admin'),(18,'Martina','Hueda','marti@gmail.com','202cb962ac59075b964b07152d234b70','estudiante'),(19,'Jose','Jose','jose123@gmail.com','202cb962ac59075b964b07152d234b70','estudiante'),(20,'po','lkj','po@gmail.com','202cb962ac59075b964b07152d234b70','estudiante');
/*!40000 ALTER TABLE `usuarios` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'users-api'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-01-22 19:03:39