-- phpMyAdmin SQL Dump
-- version 4.6.4
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: May 05, 2017 at 10:19 AM
-- Server version: 5.6.28
-- PHP Version: 7.0.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `apidb`
--

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(100) NOT NULL,
  `pwd` varchar(200) NOT NULL,
  `email` varchar(200) NOT NULL,
  `fullname` varchar(200) NOT NULL,
  `address` varchar(250) NOT NULL,
  `telephone` varchar(50) NOT NULL,
  `longitude` decimal(11,8) NOT NULL,
  `latitude` decimal(10,8) NOT NULL,
  `googleacc` varchar(250) NOT NULL,
  `resetcode` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `username`, `pwd`, `email`, `fullname`, `address`, `telephone`, `longitude`, `latitude`, `googleacc`, `resetcode`) VALUES
(3, 'te', '5a211a1761528b977faf747368c601245b3349ff', '', 'jesse junsay', 'purok 5 mahayahay st, tugbok', '+639185903558', '0.00000000', '0.00000000', '', ''),
(4, 'c', '84a516841ba77a5b4648de2cd0dfcb30ea46dbb4', '', 'jesse junsay', 'purok 5 mahayahay st, tugbok', '', '0.00000000', '0.00000000', '', ''),
(5, 'x', '11f6ad8ec52a2984abaafd7c3b516503785c2072', '', 'jesse junsay', 'purok 5 mahayahay st, tugbok', '', '0.00000000', '0.00000000', '', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
