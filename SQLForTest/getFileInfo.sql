-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- ホスト: mysql
-- 生成日時: 2022 年 8 月 08 日 10:47
-- サーバのバージョン： 8.0.29
-- PHP のバージョン: 8.0.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- データベース: `myapp`
--

-- --------------------------------------------------------

--
-- テーブルの構造 `UPLOADED_FILE_INFO`
--

CREATE TABLE `UPLOADED_FILE_INFO` (
  `FILE_INFO_ID` bigint NOT NULL,
  `FILE_NAME` text,
  `FILE_CONTENT` mediumtext,
  `MIME_TYPE` text,
  `USER_ID` varchar(5) DEFAULT NULL,
  `POSTED_DATE` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- テーブルのデータのダンプ `UPLOADED_FILE_INFO`
--

INSERT INTO `UPLOADED_FILE_INFO` (`FILE_INFO_ID`, `FILE_NAME`, `FILE_CONTENT`, `MIME_TYPE`, `USER_ID`, `POSTED_DATE`) VALUES
(8, 'test.txt', 'test', 'text/plain', '002', '2022/8/8'),
(9, 'test.txt', 'test', 'text/plain', '002', ''),
(10, 'test.txt', 'test', 'text/plain', '', '2022/8/8'),
(11, 'test.txt', 'test', 'text/plain', '', ''),
(12, 'test.txt', 'test', 'text/plain', '002', '2022/8/9'),
(13, 'test.txt', 'test', 'text/plain', '003', '2022/8/8'),
(14, 'test.txt', 'test', 'text/plain', '003', '2022/8/9');

--
-- ダンプしたテーブルのインデックス
--

--
-- テーブルのインデックス `UPLOADED_FILE_INFO`
--
ALTER TABLE `UPLOADED_FILE_INFO`
  ADD PRIMARY KEY (`FILE_INFO_ID`);

--
-- ダンプしたテーブルの AUTO_INCREMENT
--

--
-- テーブルの AUTO_INCREMENT `UPLOADED_FILE_INFO`
--
ALTER TABLE `UPLOADED_FILE_INFO`
  MODIFY `FILE_INFO_ID` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
