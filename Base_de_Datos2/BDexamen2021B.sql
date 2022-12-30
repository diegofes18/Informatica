-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Servidor: localhost:8889
-- Tiempo de generación: 24-12-2022 a las 00:30:16
-- Versión del servidor: 5.7.34
-- Versión de PHP: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `examen2021B`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `CATEGORIA`
--

CREATE TABLE `CATEGORIA` (
  `idCat` int(11) NOT NULL,
  `RISCCAT` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Volcado de datos para la tabla `CATEGORIA`
--

INSERT INTO `CATEGORIA` (`idCat`, `RISCCAT`) VALUES
(1, 10),
(2, 20),
(3, 30);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `DOLENCIES`
--

CREATE TABLE `DOLENCIES` (
  `idper` int(11) DEFAULT NULL,
  `idmal` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Volcado de datos para la tabla `DOLENCIES`
--

INSERT INTO `DOLENCIES` (`idper`, `idmal`) VALUES
(1, 2),
(1, 3),
(2, 1);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `MALATIA`
--

CREATE TABLE `MALATIA` (
  `idMal` int(11) NOT NULL,
  `riscMal` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Volcado de datos para la tabla `MALATIA`
--

INSERT INTO `MALATIA` (`idMal`, `riscMal`) VALUES
(1, 5),
(2, 10),
(3, 15);

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `PERSONA`
--

CREATE TABLE `PERSONA` (
  `idPer` int(11) NOT NULL,
  `anyPer` int(11) DEFAULT NULL,
  `riscPer` decimal(6,2) DEFAULT NULL,
  `catper` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Volcado de datos para la tabla `PERSONA`
--

INSERT INTO `PERSONA` (`idPer`, `anyPer`, `riscPer`, `catper`) VALUES
(1, 1960, '0.50', 1),
(2, 1970, '0.20', 2),
(3, 1980, '0.80', 3);

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `CATEGORIA`
--
ALTER TABLE `CATEGORIA`
  ADD PRIMARY KEY (`idCat`);

--
-- Indices de la tabla `DOLENCIES`
--
ALTER TABLE `DOLENCIES`
  ADD KEY `idper` (`idper`),
  ADD KEY `idmal` (`idmal`);

--
-- Indices de la tabla `MALATIA`
--
ALTER TABLE `MALATIA`
  ADD PRIMARY KEY (`idMal`);

--
-- Indices de la tabla `PERSONA`
--
ALTER TABLE `PERSONA`
  ADD PRIMARY KEY (`idPer`),
  ADD KEY `catper` (`catper`);

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `DOLENCIES`
--
ALTER TABLE `DOLENCIES`
  ADD CONSTRAINT `dolencies_ibfk_1` FOREIGN KEY (`idper`) REFERENCES `PERSONA` (`idPer`),
  ADD CONSTRAINT `dolencies_ibfk_2` FOREIGN KEY (`idmal`) REFERENCES `MALATIA` (`idMal`);

--
-- Filtros para la tabla `PERSONA`
--
ALTER TABLE `PERSONA`
  ADD CONSTRAINT `persona_ibfk_1` FOREIGN KEY (`catper`) REFERENCES `CATEGORIA` (`idCat`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
