-- phpMyAdmin SQL Dump
-- version 6.0.0-dev+20230516.e7d1ecbeae
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 09, 2023 at 04:26 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `online_shop`
--

DELIMITER $$
--
-- Functions
--
CREATE DEFINER=`root`@`localhost` FUNCTION `DeleteTransactionId` (`id_transaction` INT) RETURNS TINYINT(1)  BEGIN
    DECLARE affected_rows INT$$

CREATE DEFINER=`root`@`localhost` FUNCTION `UpdateQty` (`id_transaction` INT) RETURNS TINYINT(1)  BEGIN
	DECLARE totalQty INT$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Budi', '2023-09-07 12:06:29', '2023-09-07 12:06:29'),
(2, 'Reski', '2023-09-07 12:06:29', '2023-09-07 12:06:29'),
(3, 'Siti', '2023-09-07 12:06:29', '2023-09-07 12:06:29'),
(4, 'Adam', '2023-09-07 12:06:29', '2023-09-07 12:06:29'),
(5, 'Andi', '2023-09-07 12:06:29', '2023-09-07 12:06:29');

-- --------------------------------------------------------

--
-- Table structure for table `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 'BRI', 1, '2023-09-07 15:47:17', '2023-09-07 15:47:17'),
(2, 'CASH', 1, '2023-09-07 15:47:17', '2023-09-07 15:47:17'),
(3, 'DANA', 0, '2023-09-07 15:47:17', '2023-09-07 15:47:17');

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `id` int(11) NOT NULL,
  `product_type_id` int(11) DEFAULT NULL,
  `operator_id` int(11) DEFAULT NULL,
  `code` varchar(50) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES
(3, 2, 1, 'A03', 'Taro', 0, '2023-09-07 15:26:14', '2023-09-07 15:26:14'),
(4, 2, 1, 'A04', 'Doritos', 1, '2023-09-07 15:26:14', '2023-09-07 15:26:14'),
(5, 2, 1, 'A05', 'Roma Kelapa', 2, '2023-09-07 15:26:14', '2023-09-07 15:26:14'),
(6, 3, 4, 'A06', 'Monti', 1, '2023-09-07 15:28:42', '2023-09-07 15:28:42'),
(7, 3, 4, 'A07', 'Teh Kotak', 0, '2023-09-07 15:28:42', '2023-09-07 15:28:42'),
(8, 3, 4, 'A08', 'Cocacola', 2, '2023-09-07 15:28:42', '2023-09-07 15:28:42');

-- --------------------------------------------------------

--
-- Table structure for table `product_description`
--

CREATE TABLE `product_description` (
  `id` int(11) NOT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_description`
--

INSERT INTO `product_description` (`id`, `description`, `created_at`, `updated_at`) VALUES
(3, 'Snack yang berbentuk jaring', '2023-09-07 15:42:38', '2023-09-07 15:42:38'),
(4, 'Snack keripik tortilla dengan bahan jagung', '2023-09-07 15:42:38', '2023-09-07 15:42:38'),
(5, 'Biskuit yang terbuat dari kelapa segar pilihan', '2023-09-07 15:42:38', '2023-09-07 15:42:38'),
(6, 'Minuman dalam kemasan gelas yang terbuat dari teh alami', '2023-09-07 15:42:38', '2023-09-07 15:42:38'),
(7, 'Minuman teh yang berbentuk kotak', '2023-09-07 15:42:38', '2023-09-07 15:42:38'),
(8, 'Minuman ringan yang berkarbonasi yang tebuat dari daun koka', '2023-09-07 15:42:38', '2023-09-07 15:42:38');

-- --------------------------------------------------------

--
-- Table structure for table `product_type`
--

CREATE TABLE `product_type` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_type`
--

INSERT INTO `product_type` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Sabun', '2023-09-07 12:23:45', '2023-09-07 12:23:45'),
(2, 'Cemilan', '2023-09-07 12:23:45', '2023-09-07 12:23:45'),
(3, 'Minuman', '2023-09-07 12:23:45', '2023-09-07 12:23:45');

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

CREATE TABLE `transaction` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `payment_method_id` int(11) DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `total_qty` int(11) DEFAULT NULL,
  `total_price` decimal(25,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction`
--

INSERT INTO `transaction` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(5, 2, 1, 'Done', 27, 45900.10, '2023-09-07 16:18:18', '2023-09-09 01:52:49'),
(6, 2, 1, 'Done', 42, 102900.10, '2023-09-07 16:18:18', '2023-09-09 01:55:08'),
(7, 3, 1, 'Pending', 280, 295000.50, '2023-09-07 16:18:18', '2023-09-09 01:41:20'),
(9, 3, 2, 'Done', 183, 20421.30, '2023-09-07 16:18:18', '2023-09-09 01:51:45'),
(10, 4, 2, 'Done', 11, 25600.00, '2023-09-07 16:18:18', '2023-09-07 16:18:18'),
(11, 4, 2, 'Done', 21, 42132.30, '2023-09-07 16:18:18', '2023-09-07 16:18:18'),
(12, 4, 1, 'Pending', 16, 30212.50, '2023-09-07 16:18:18', '2023-09-09 01:43:47'),
(13, 5, 2, 'Pending', 54, 150300.00, '2023-09-07 16:18:18', '2023-09-07 16:18:18'),
(14, 5, 1, 'Done', 35, 297421.15, '2023-09-07 16:18:18', '2023-09-07 16:18:18'),
(15, 5, 1, 'Done', 7, 39000.00, '2023-09-07 16:18:18', '2023-09-07 16:18:18');

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `transaction_id` int(11) DEFAULT NULL,
  `product_id` int(11) DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `price` decimal(25,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
(7, 8, 'Diproses', 32, 168571.71, '2023-09-09 01:42:29', '2023-09-09 01:42:29'),
(7, 4, 'Diproses', 18, 94821.58, '2023-09-09 01:42:29', '2023-09-09 01:42:29'),
(7, 5, 'Diproses', 6, 31607.19, '2023-09-09 01:42:29', '2023-09-09 01:42:29'),
(10, 8, 'Diterima', 2, 4654.54, '2023-09-09 01:42:42', '2023-09-09 01:42:42'),
(10, 3, 'Diterima', 5, 11636.36, '2023-09-09 01:42:42', '2023-09-09 01:42:42'),
(10, 5, 'Diterima', 4, 9309.09, '2023-09-09 01:42:42', '2023-09-09 01:42:42'),
(11, 5, 'Diterima', 12, 24070.60, '2023-09-09 01:42:46', '2023-09-09 01:42:46'),
(11, 3, 'Diterima', 5, 10031.50, '2023-09-09 01:42:46', '2023-09-09 01:42:46'),
(11, 4, 'Diterima', 4, 8025.20, '2023-09-09 01:42:46', '2023-09-09 01:42:46');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 'Toni', 1, '2022-09-12', 'M', '2023-09-07 15:57:31', '2023-09-08 05:34:11'),
(2, 'Eka', 1, '2023-01-15', 'W', '2023-09-07 15:57:31', '2023-09-08 05:34:13'),
(3, 'Alni', 0, '2021-09-05', 'W', '2023-09-07 15:57:31', '2023-09-08 05:34:15'),
(4, 'Fahmi', 1, '2022-11-10', 'M', '2023-09-07 15:57:31', '2023-09-08 05:34:17'),
(5, 'Bumi', 0, '2023-07-01', 'M', '2023-09-07 15:57:31', '2023-09-08 05:34:18');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_type_id` (`product_type_id`),
  ADD KEY `operator_id` (`operator_id`);

--
-- Indexes for table `product_description`
--
ALTER TABLE `product_description`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- Indexes for table `product_type`
--
ALTER TABLE `product_type`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `payment_method_id` (`payment_method_id`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD KEY `transaction_id` (`transaction_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `product`
--
ALTER TABLE `product`
  ADD CONSTRAINT `product_ibfk_1` FOREIGN KEY (`product_type_id`) REFERENCES `product_type` (`id`),
  ADD CONSTRAINT `product_ibfk_2` FOREIGN KEY (`operator_id`) REFERENCES `operators` (`id`);

--
-- Constraints for table `product_description`
--
ALTER TABLE `product_description`
  ADD CONSTRAINT `product_description_ibfk_1` FOREIGN KEY (`id`) REFERENCES `product` (`id`);

--
-- Constraints for table `transaction`
--
ALTER TABLE `transaction`
  ADD CONSTRAINT `transaction_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `transaction_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`);

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transaction` (`id`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
