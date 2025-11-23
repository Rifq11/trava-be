-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Nov 23, 2025 at 02:41 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `trava`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin_profiles`
--

CREATE TABLE `admin_profiles` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `birth_date` date DEFAULT NULL,
  `user_photo` varchar(500) DEFAULT NULL,
  `is_completed` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `admin_profiles`
--

INSERT INTO `admin_profiles` (`id`, `user_id`, `phone`, `address`, `birth_date`, `user_photo`, `is_completed`, `created_at`, `updated_at`) VALUES
(1, 2, '123456789', '123 Main', '1990-01-01', '/uploads/134560386_p4_master1200-1763723836039396723-801969638.jpg', 1, '2025-11-21 18:17:16', '2025-11-21 18:17:16');

-- --------------------------------------------------------

--
-- Table structure for table `bookings`
--

CREATE TABLE `bookings` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `destination_id` int(11) NOT NULL,
  `transportation_id` int(11) NOT NULL,
  `payment_method_id` int(11) NOT NULL,
  `status_id` int(11) NOT NULL,
  `people_count` int(11) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `transport_price` int(11) NOT NULL,
  `destination_price` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `bookings`
--

INSERT INTO `bookings` (`id`, `user_id`, `destination_id`, `transportation_id`, `payment_method_id`, `status_id`, `people_count`, `start_date`, `end_date`, `transport_price`, `destination_price`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 9, 5, 1, 3, 3, 12, '2025-11-21 12:29:47', '2025-11-21 12:29:47', 123123, 123123, 123213213, '2025-11-21 18:30:11', '2025-11-21 18:30:11'),
(3, 11, 5, 1, 3, 5, 3, '2025-11-22 07:00:00', '2025-11-22 08:00:00', 235, 1500000, 1500235, '2025-11-22 23:37:35', '2025-11-22 23:37:35'),
(10, 14, 5, 1, 1, 4, 6, '2025-11-23 07:00:00', '2025-11-24 07:00:00', 235, 3000000, 3000235, '2025-11-23 01:05:49', '2025-11-23 01:05:49'),
(11, 14, 5, 1, 1, 1, 1, '2025-11-23 07:00:00', '2025-11-24 07:00:00', 235, 500000, 500235, '2025-11-23 01:21:33', '2025-11-23 01:21:33');

-- --------------------------------------------------------

--
-- Table structure for table `booking_status`
--

CREATE TABLE `booking_status` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `booking_status`
--

INSERT INTO `booking_status` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'pending', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(2, 'approved', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(3, 'rejected', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(4, 'canceled', '2025-11-22 23:52:10', '2025-11-22 23:52:10'),
(5, 'completed', '2025-11-23 00:28:52', '2025-11-23 00:28:52');

-- --------------------------------------------------------

--
-- Table structure for table `destinations`
--

CREATE TABLE `destinations` (
  `id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `created_by` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `price_per_person` int(11) NOT NULL,
  `image` text DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `destinations`
--

INSERT INTO `destinations` (`id`, `category_id`, `created_by`, `name`, `description`, `location`, `price_per_person`, `image`, `created_at`, `updated_at`) VALUES
(5, 1, 1, 'Beautiful Beach', 'A beautiful beach destination', 'Bali, Indonesia', 500000, '/uploads/134560386_p4_master1200-1763690463870747906-238966758.jpg', '2025-11-21 09:01:03', '2025-11-21 09:01:03'),
(14, 3, 1, 'Updated Beach', 'Updated description', 'Bali, Indonesia', 600000, '/uploads/family_card-1763830771770997439-923798235.jpg', '2025-11-22 22:38:33', '2025-11-22 22:38:33'),
(15, 3, 1, 'Updated Beach', 'Updated description', 'Bali, Indonesia', 600000, '/uploads/family_card-1763830773686987695-912992250.jpg', '2025-11-22 22:59:45', '2025-11-22 22:59:45'),
(16, 3, 1, 'Updated Beach', 'Updated description', 'Bali, Indonesia', 600000, '/uploads/family_card-1763830775263304291-918171010.jpg', '2025-11-22 22:59:47', '2025-11-22 22:59:47'),
(17, 3, 1, 'Updated Beach', 'Updated description', 'Bali, Indonesia', 600000, '/uploads/family_card-1763830776612280372-897138437.jpg', '2025-11-22 22:59:49', '2025-11-22 22:59:49'),
(18, 3, 1, 'Updated Beach', 'Updated description', 'Bali, Indonesia', 600000, '/uploads/family_card-1763830778089986846-34926329.jpg', '2025-11-22 22:59:51', '2025-11-22 22:59:51');

-- --------------------------------------------------------

--
-- Table structure for table `destination_categories`
--

CREATE TABLE `destination_categories` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `destination_categories`
--

INSERT INTO `destination_categories` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Beach', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(2, 'Mountain', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(3, 'City', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(4, 'Nature', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(5, 'Historical', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(6, 'Adventure', '2025-11-18 11:46:32', '2025-11-18 11:46:32');

-- --------------------------------------------------------

--
-- Table structure for table `payments`
--

CREATE TABLE `payments` (
  `id` int(11) NOT NULL,
  `booking_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `payment_status` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `payments`
--

INSERT INTO `payments` (`id`, `booking_id`, `amount`, `payment_status`, `created_at`, `updated_at`) VALUES
(2, 3, 1500235, 'success', '2025-11-22 23:37:35', '2025-11-22 23:37:35'),
(3, 10, 3000235, 'success', '2025-11-23 01:05:49', '2025-11-23 01:05:49'),
(4, 11, 500235, 'success', '2025-11-23 01:21:33', '2025-11-23 01:21:33');

-- --------------------------------------------------------

--
-- Table structure for table `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Credit Card', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(2, 'Debit Card', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(3, 'Bank Transfer', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(4, 'E-Wallet', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(5, 'Cash', '2025-11-18 11:46:32', '2025-11-18 11:46:32');

-- --------------------------------------------------------

--
-- Table structure for table `reviews`
--

CREATE TABLE `reviews` (
  `id` int(11) NOT NULL,
  `booking_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `rating` int(11) NOT NULL,
  `review_text` text DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `reviews`
--

INSERT INTO `reviews` (`id`, `booking_id`, `user_id`, `rating`, `review_text`, `created_at`, `updated_at`) VALUES
(1, 1, 9, 5, 'sadsadsadasdads', '2025-11-21 18:30:35', '2025-11-21 18:30:35'),
(2, 3, 11, 3, '', '2025-11-23 00:38:16', '2025-11-23 00:38:16');

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'admin', '2025-11-18 11:46:32.000', '2025-11-18 11:46:32.000', NULL),
(2, 'user', '2025-11-18 11:46:32.000', '2025-11-18 11:46:32.000', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `transportation`
--

CREATE TABLE `transportation` (
  `id` int(11) NOT NULL,
  `destination_id` int(11) NOT NULL,
  `transport_type_id` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `detail_tranportation` varchar(255) NOT NULL,
  `estimate` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `transportation`
--

INSERT INTO `transportation` (`id`, `destination_id`, `transport_type_id`, `price`, `detail_tranportation`, `estimate`, `created_at`, `updated_at`) VALUES
(1, 5, 1, 235, '', 'asdasdasd', '2025-11-21 18:29:37', '2025-11-21 18:29:37');

-- --------------------------------------------------------

--
-- Table structure for table `transport_types`
--

CREATE TABLE `transport_types` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `transport_types`
--

INSERT INTO `transport_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Plane', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(3, 'Bus', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(4, 'Car', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(5, 'Ship', '2025-11-18 11:46:32', '2025-11-18 11:46:32');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `role_id`, `full_name`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 1, 'Admin User', 'admin@example.com', '$2a$10$3dlSFxAob.iExmzoiZbBYeKQB7MKCTIpB9KV/gTe0IyEaX4DTFxoa', '2025-11-18 11:46:32', '2025-11-18 11:46:32'),
(2, 2, 'Updated User', 'updated@example.com', '$2a$10$hpD/0ZYUpPyDmQ0oDcclcuQa0/tOcznDqXlVpTs0OCXz2H3hxGzt2', '2025-11-18 11:46:32', '2025-11-22 13:21:30'),
(6, 2, 'Test User', 'test1763623918@example.com', '$2a$10$QdAhASWLopFuPnOnosJ2IeEUxKLzZZnCI8uhtqIVbKmDElMWRo/Eq', '2025-11-20 14:31:58', '2025-11-20 14:31:58'),
(9, 2, 'John Doe', 'john@example.com', '$2a$10$w3thRAQuZLjU7e0FwNFBMu8H4zkzcp4hHFZRNDYYQVUVy3xVWNAZW', '2025-11-21 11:09:24', '2025-11-21 11:09:24'),
(11, 2, 'tes', 'tes@example.com', '$2a$10$n1f8U8pO7v3RRfFkwiXT8.Zb/memUbUnhQjADLyV9Av9nGPIfFFsG', '2025-11-22 13:57:48', '2025-11-22 17:19:04'),
(12, 2, 'titu', 'firjd@jdjd.fj', '$2a$10$U1yySIM23jKaZcYkwLAf2.7gcP7xTpOAgIBiGdlU1zn./Ma4pC68a', '2025-11-22 14:07:27', '2025-11-22 14:07:27'),
(13, 2, 'a', 'a@a.a', '$2a$10$dnYu.bLYP6YzkFq5ZT4xmukllD5T4W2ygmKdkZjeRWoOw/HPrtiKe', '2025-11-22 14:16:33', '2025-11-22 14:32:21'),
(14, 2, 'b', 'b@b.b', '$2a$10$h1FrY1Ha.H5bxSTLznCV1.ZzQPnYAFMU44kNoAPF08TNYrBhO/.2S', '2025-11-22 17:47:07', '2025-11-22 18:38:01'),
(15, 2, 'f', 'f@f.f', '$2a$10$biz5KeWHQWUPbYjCgwGmvuaeVrAe3rCCPG.PJPgwZ7EOKveIuuRmW', '2025-11-22 18:38:54', '2025-11-22 18:44:00'),
(16, 2, 'r', 'r@r.r', '$2a$10$hIbzOHMMc7vy456.d56W5O9g089EaTQ.NAyG.CAvgx/9b0cq0liBK', '2025-11-22 18:44:21', '2025-11-22 19:15:40');

-- --------------------------------------------------------

--
-- Table structure for table `user_activity_log`
--

CREATE TABLE `user_activity_log` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `destination_id` int(11) DEFAULT NULL,
  `activity_type` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `user_profiles`
--

CREATE TABLE `user_profiles` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `birth_date` date DEFAULT NULL,
  `user_photo` varchar(500) DEFAULT NULL,
  `is_completed` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `user_profiles`
--

INSERT INTO `user_profiles` (`id`, `user_id`, `phone`, `address`, `birth_date`, `user_photo`, `is_completed`, `created_at`, `updated_at`) VALUES
(1, 2, '123456789', '123 Main Street', '1990-01-01', '/uploads/134560386_p4_master1200-1763723774379784322-636961356.jpg', 1, '2025-11-21 18:10:52', '2025-11-21 18:10:52'),
(2, 13, '4848483', 'ncdndn', '2007-11-27', '/uploads/1000067482-1763821045242532576-88893237.jpg', 1, '2025-11-22 21:17:25', '2025-11-22 21:17:25'),
(3, 11, '8483884', 'ndjd', '2007-11-30', '/uploads/1000068020-1763831944861751272-24240214.jpg', 1, '2025-11-22 21:33:02', '2025-11-22 21:33:02'),
(4, 14, '56889595', 'ncnc', '2007-11-28', '/uploads/1000067996-1763835751369822524-915991829.jpg', 1, '2025-11-23 00:53:17', '2025-11-23 00:53:17'),
(5, 15, '5656', 'djdjdjd', '2007-11-28', '/uploads/1000067996-1763836778742634712-214111679.jpg', 1, '2025-11-23 01:39:38', '2025-11-23 01:39:38'),
(8, 16, '65655', 'hfbf', '2007-11-28', '/uploads/1000068018-1763838940465095928-412471012.jpg', 1, '2025-11-23 02:15:11', '2025-11-23 02:15:11');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin_profiles`
--
ALTER TABLE `admin_profiles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id_idx` (`user_id`);

--
-- Indexes for table `bookings`
--
ALTER TABLE `bookings`
  ADD PRIMARY KEY (`id`),
  ADD KEY `bookings_payment_method_id_payment_methods_id_fk` (`payment_method_id`),
  ADD KEY `user_id_idx` (`user_id`),
  ADD KEY `destination_id_idx` (`destination_id`),
  ADD KEY `transportation_id_idx` (`transportation_id`),
  ADD KEY `status_id_idx` (`status_id`);

--
-- Indexes for table `booking_status`
--
ALTER TABLE `booking_status`
  ADD PRIMARY KEY (`id`),
  ADD KEY `name_idx` (`name`);

--
-- Indexes for table `destinations`
--
ALTER TABLE `destinations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `category_id_idx` (`category_id`),
  ADD KEY `created_by_idx` (`created_by`);

--
-- Indexes for table `destination_categories`
--
ALTER TABLE `destination_categories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `name_idx` (`name`);

--
-- Indexes for table `payments`
--
ALTER TABLE `payments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `booking_id_idx` (`booking_id`);

--
-- Indexes for table `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`),
  ADD KEY `name_idx` (`name`);

--
-- Indexes for table `reviews`
--
ALTER TABLE `reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `booking_id_idx` (`booking_id`),
  ADD KEY `user_id_idx` (`user_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `name_idx` (`name`),
  ADD KEY `idx_roles_name` (`name`),
  ADD KEY `idx_roles_deleted_at` (`deleted_at`);

--
-- Indexes for table `transportation`
--
ALTER TABLE `transportation`
  ADD PRIMARY KEY (`id`),
  ADD KEY `destination_id_idx` (`destination_id`),
  ADD KEY `transport_type_id_idx` (`transport_type_id`);

--
-- Indexes for table `transport_types`
--
ALTER TABLE `transport_types`
  ADD PRIMARY KEY (`id`),
  ADD KEY `name_idx` (`name`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `email_idx` (`email`),
  ADD KEY `role_id_idx` (`role_id`);

--
-- Indexes for table `user_activity_log`
--
ALTER TABLE `user_activity_log`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id_idx` (`user_id`),
  ADD KEY `destination_id_idx` (`destination_id`);

--
-- Indexes for table `user_profiles`
--
ALTER TABLE `user_profiles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id_idx` (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin_profiles`
--
ALTER TABLE `admin_profiles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `bookings`
--
ALTER TABLE `bookings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `booking_status`
--
ALTER TABLE `booking_status`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `destinations`
--
ALTER TABLE `destinations`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `destination_categories`
--
ALTER TABLE `destination_categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `payments`
--
ALTER TABLE `payments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `payment_methods`
--
ALTER TABLE `payment_methods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `reviews`
--
ALTER TABLE `reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `transportation`
--
ALTER TABLE `transportation`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `transport_types`
--
ALTER TABLE `transport_types`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `user_activity_log`
--
ALTER TABLE `user_activity_log`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user_profiles`
--
ALTER TABLE `user_profiles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `admin_profiles`
--
ALTER TABLE `admin_profiles`
  ADD CONSTRAINT `admin_profiles_user_id_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `bookings`
--
ALTER TABLE `bookings`
  ADD CONSTRAINT `bookings_destination_id_destinations_id_fk` FOREIGN KEY (`destination_id`) REFERENCES `destinations` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `bookings_payment_method_id_payment_methods_id_fk` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `bookings_status_id_booking_status_id_fk` FOREIGN KEY (`status_id`) REFERENCES `booking_status` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `bookings_transportation_id_transportation_id_fk` FOREIGN KEY (`transportation_id`) REFERENCES `transportation` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `bookings_user_id_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `destinations`
--
ALTER TABLE `destinations`
  ADD CONSTRAINT `destinations_category_id_destination_categories_id_fk` FOREIGN KEY (`category_id`) REFERENCES `destination_categories` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `destinations_created_by_users_id_fk` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `payments`
--
ALTER TABLE `payments`
  ADD CONSTRAINT `payments_booking_id_bookings_id_fk` FOREIGN KEY (`booking_id`) REFERENCES `bookings` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `reviews`
--
ALTER TABLE `reviews`
  ADD CONSTRAINT `reviews_booking_id_bookings_id_fk` FOREIGN KEY (`booking_id`) REFERENCES `bookings` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `reviews_user_id_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `transportation`
--
ALTER TABLE `transportation`
  ADD CONSTRAINT `transportation_destination_id_destinations_id_fk` FOREIGN KEY (`destination_id`) REFERENCES `destinations` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `transportation_transport_type_id_transport_types_id_fk` FOREIGN KEY (`transport_type_id`) REFERENCES `transport_types` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `users_role_id_roles_id_fk` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `user_activity_log`
--
ALTER TABLE `user_activity_log`
  ADD CONSTRAINT `user_activity_log_destination_id_destinations_id_fk` FOREIGN KEY (`destination_id`) REFERENCES `destinations` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `user_activity_log_user_id_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `user_profiles`
--
ALTER TABLE `user_profiles`
  ADD CONSTRAINT `user_profiles_user_id_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
