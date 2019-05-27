
-- --------------------------------------------------------

--
-- Table structure for table `banwiredash02graph02`
--

CREATE TABLE `banwiredash02graph02` (
  `id_record` bigint(20) NOT NULL,
  `numsecuecial_grupoddatos` bigint(20) DEFAULT NULL,
  `numsecuecial` bigint(20) DEFAULT NULL,
  `nombrecolumna` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `valorcolumna` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `valoramount` bigint(20) DEFAULT NULL,
  `graphnbr` varchar(10) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `banwiredash02graph02`


--

INSERT INTO `banwiredash02graph02` (`id_record`, `numsecuecial_grupoddatos`, `numsecuecial`, `nombrecolumna`, `created_at`, `valorcolumna`, `valoramount`, `graphnbr`) 
VALUES
(1, 10, 1, 'tag 1', '2019-03-21 17:34:45', '25', 25, '22'),

(2, 10, 2, 'tag B', '2019-03-21 17:34:45', '10', 10, '22'),

(3, 10, 3, 'tag C', '2019-03-21 17:34:45', '22', 22, '22'),

(4, 10, 4, 'tag D', '2019-03-21 17:34:45', '27', 27, '22'),

(5, 20, 1, 'tag 1', '2019-03-21 17:34:45', '65', 65, '22'),

(6, 20, 2, 'tag B', '2019-03-21 17:34:45', '60', 60, '22'),

(7, 20, 3, 'tag C', '2019-03-21 17:34:45', '62', 62, '22'),

(8, 20, 4, 'tag D', '2019-03-21 17:34:45', '67', 67, '22');



--
-- Indexes for dumped tables
--

--
-- Indexes for table `banwiredash02graph02`
--
ALTER TABLE `banwiredash02graph02`
  ADD PRIMARY KEY (`id_record`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `banwiredash02graph02`
--
ALTER TABLE `banwiredash02graph02`
  MODIFY `id_record` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;




--


INSERT INTO `banwiredash02graph02` (`numsecuecial_grupoddatos`, `numsecuecial`, `nombrecolumna`, `created_at`, `valorcolumna`, `valoramount`, `graphnbr`) 
VALUES

( 10, 1, 'PC', '2019-03-21 17:34:45', '25', 25, '11'),

( 10, 2, 'Comercio', '2019-03-21 17:34:45', '10', 10, '11'),

( 10, 3, 'Mobile', '2019-03-21 17:34:45', '22', 22, '11'),

( 10, 4, 'Otros', '2019-03-21 17:34:45', '27', 27, '11'),

( 20, 1, 'PC', '2019-03-21 17:34:45', '65', 65, '11'),

( 20, 2, 'Comercio', '2019-03-21 17:34:45', '60', 60, '11'),

( 20, 3, 'Mobile', '2019-03-21 17:34:45', '62', 62, '11'),

( 20, 4, 'Otros', '2019-03-21 17:34:45', '67', 67, '11'),
( 30, 1, 'PC', '2019-03-21 17:34:45', '15', 15, '11'),

( 30, 2, 'Comercio', '2019-03-21 17:34:45', '20', 20, '11'),

( 30, 3, 'Mobile', '2019-03-21 17:34:45', '42', 42, '11'),

( 30, 4, 'Otros', '2019-03-21 17:34:45', '17', 17, '11');


INSERT INTO `banwiredash02graph02` (`numsecuecial_grupoddatos`, `numsecuecial`, `nombrecolumna`, `created_at`, `valorcolumna`, `valoramount`, `graphnbr`) 
VALUES

( 20, 1, 'Electronic', '2019-03-21 17:34:45', '35', 35, '12'),

( 20, 2, 'Web', '2019-03-21 17:34:45', '20', 20, '12'),

( 20, 3, 'Service', '2019-03-21 17:34:45', '32', 32, '12'),

( 20, 4, 'Otros', '2019-03-21 17:34:45', '17', 17, '12'),

( 30, 1, 'Electronic', '2019-03-21 17:34:45', '65', 65, '12'),

( 30, 2, 'Web', '2019-03-21 17:34:45', '10', 10, '12'),

( 30, 3, 'Service', '2019-03-21 17:34:45', '12', 12, '12'),

( 30, 4, 'Otros', '2019-03-21 17:34:45', '20', 20, '12'),
( 50, 1, 'Electronic', '2019-03-21 17:34:45', '15', 15, '12'),

( 50, 2, 'Web', '2019-03-21 17:34:45', '20', 20, '12'),

( 50, 3, 'Service', '2019-03-21 17:34:45', '32', 32, '12'),

( 50, 4, 'Otros', '2019-03-21 17:34:45', '57', 57, '12');


INSERT INTO `banwiredash02graph02` (`numsecuecial_grupoddatos`, `numsecuecial`, `nombrecolumna`, `created_at`, `valorcolumna`, `valoramount`, `graphnbr`) 
VALUES

( 60, 1, 'Pagos', '2019-03-21 17:34:45', '25', 25, '13'),

( 60, 2, 'Confirmados', '2019-03-21 17:34:45', '40', 40, '13'),

( 60, 3, 'Reversados', '2019-03-21 17:34:45', '22', 22, '13'),

( 60, 4, 'Aplicados', '2019-03-21 17:34:45', '27', 27, '13'),

( 70, 1, 'Pagos', '2019-03-21 17:34:45', '55', 55, '13'),

( 70, 2, 'Confirmados', '2019-03-21 17:34:45', '10', 10, '13'),

( 70, 3, 'Reversados', '2019-03-21 17:34:45', '22', 22, '13'),

( 70, 4, 'Aplicados', '2019-03-21 17:34:45', '34', 34, '13'),
( 80, 1, 'Pagos', '2019-03-21 17:34:45', '25', 25, '13'),

( 80, 2, 'Confirmados', '2019-03-21 17:34:45', '27', 27, '13'),

( 80, 3, 'Reversados', '2019-03-21 17:34:45', '42', 42, '13'),

( 80, 4, 'Aplicados', '2019-03-21 17:34:45', '19', 19, '13');




//dash 3 graphs 1 and 2

INSERT INTO `banwiredash02graph02` (`numsecuecial_grupoddatos`, `numsecuecial`, `nombrecolumna`, `created_at`, `valorcolumna`, `valoramount`, `graphnbr`) 
VALUES

( 60, 1, 'Centro', '2019-03-21 17:34:45', '25', 25, '31'),

( 60, 2, 'Norte', '2019-03-21 17:34:45', '30', 30, '31'),

( 60, 3, 'Occidente', '2019-03-21 17:34:45', '22', 22, '31'),

( 60, 4, 'Peninsula', '2019-03-21 17:34:45', '27', 27, '31'),

( 70, 1, 'Centro', '2019-03-21 17:34:45', '35', 35, '31'),

( 70, 2, 'Norte', '2019-03-21 17:34:45', '10', 10, '31'),

( 70, 3, 'Occidente', '2019-03-21 17:34:45', '22', 22, '31'),

( 70, 4, 'Peninsula', '2019-03-21 17:34:45', '24', 24, '31'),
( 80, 1, 'Centro', '2019-03-21 17:34:45', '25', 25, '31'),

( 80, 2, 'Norte', '2019-03-21 17:34:45', '27', 27, '31'),

( 80, 3, 'Occidente', '2019-03-21 17:34:45', '22', 22, '31'),

( 80, 4, 'Peninsula', '2019-03-21 17:34:45', '19', 19, '31');

INSERT INTO `banwiredash02graph02` (`numsecuecial_grupoddatos`, `numsecuecial`, `nombrecolumna`, `created_at`, `valorcolumna`, `valoramount`, `graphnbr`) 
VALUES

( 20, 1, 'Electronic', '2019-03-21 17:34:45', '35', 35, '32'),

( 20, 2, 'Web', '2019-03-21 17:34:45', '20', 20, '32'),

( 20, 3, 'Service', '2019-03-21 17:34:45', '32', 32, '32'),

( 20, 4, 'Otros', '2019-03-21 17:34:45', '17', 17, '32'),

( 30, 1, 'Electronic', '2019-03-21 17:34:45', '65', 65, '32'),

( 30, 2, 'Web', '2019-03-21 17:34:45', '10', 10, '32'),

( 30, 3, 'Service', '2019-03-21 17:34:45', '12', 12, '32'),

( 30, 4, 'Otros', '2019-03-21 17:34:45', '20', 20, '32'),
( 50, 1, 'Electronic', '2019-03-21 17:34:45', '15', 15, '32'),

( 50, 2, 'Web', '2019-03-21 17:34:45', '20', 20, '32'),

( 50, 3, 'Service', '2019-03-21 17:34:45', '32', 32, '32'),

( 50, 4, 'Otros', '2019-03-21 17:34:45', '57', 57, '32');
