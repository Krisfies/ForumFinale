-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : lun. 23 mai 2022 à 21:10
-- Version du serveur : 5.7.36
-- Version de PHP : 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `data-access`
--

-- --------------------------------------------------------

--
-- Structure de la table `data_message`
--

DROP TABLE IF EXISTS `data_message`;
CREATE TABLE IF NOT EXISTS `data_message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(1000) NOT NULL,
  `author` varchar(255) NOT NULL,
  `date` date NOT NULL,
  `Topic_ID` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `Topic_ID` (`Topic_ID`)
) ENGINE=MyISAM AUTO_INCREMENT=511 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `data_message`
--

INSERT INTO `data_message` (`id`, `content`, `author`, `date`, `Topic_ID`) VALUES
(146, 'Impossible de la tuer, une de ses têtes rentre littéralement dans un mur', 'Random', '2022-05-23', 6),
(145, 'Vous savez d\'où ça viens ?', 'Random', '2022-05-23', 5),
(144, 'Les gars j\'arrive pas à les battre donnez moi vos astuces', 'Random', '2022-05-23', 4),
(508, 'Essaie de tuer le gros en deuxième, c\'est lui le plus simple à battre dans la P2.', 'psedo', '2022-05-23', 4),
(509, 'Vraiment je capte pas les gars.', 'Random', '2022-05-23', 12),
(510, 'J\'ai perdu ma sauvegarde en changeant mon compte Steam d\'ordinateur. Est ce que le problème viens de Steam ou du jeu parce que c\'est le seul impacté.', 'Random', '2022-05-23', 14);

-- --------------------------------------------------------

--
-- Structure de la table `data_topic`
--

DROP TABLE IF EXISTS `data_topic`;
CREATE TABLE IF NOT EXISTS `data_topic` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `IsAide` tinyint(1) NOT NULL,
  `IsBug` tinyint(1) NOT NULL,
  `IsBoss` tinyint(1) NOT NULL,
  `IsLore` tinyint(1) NOT NULL,
  `CategorieID` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=15 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `data_topic`
--

INSERT INTO `data_topic` (`ID`, `Name`, `IsAide`, `IsBug`, `IsBoss`, `IsLore`, `CategorieID`) VALUES
(4, 'O&S putain, comment un boss peut être aussi chiant', 1, 0, 1, 0, 1),
(5, 'Ouah j\'ai découvert un truc de fou', 1, 0, 0, 1, 2),
(6, 'Bug au niveau de l\'Hydre', 0, 1, 0, 0, 1),
(12, 'Qui a compris la fin du jeu ?', 0, 0, 0, 1, 2),
(13, 'Perte de sauvegarde', 1, 1, 0, 0, 3),
(14, 'Perte de sauvegarde', 1, 0, 0, 0, 3);

-- --------------------------------------------------------

--
-- Structure de la table `data_user`
--

DROP TABLE IF EXISTS `data_user`;
CREATE TABLE IF NOT EXISTS `data_user` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `Email` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `Admin` varchar(255) NOT NULL,
  `Date` date NOT NULL,
  `Biography` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `data_user`
--

INSERT INTO `data_user` (`ID`, `Name`, `Email`, `Password`, `Admin`, `Date`, `Biography`) VALUES
(1, 'kris', 'kris@gmail.com', '123', 'false', '2022-05-23', 'tesssst'),
(2, 'tets', 'gge', '123', 'false', '2022-05-23', ''),
(3, 'test', 'tets@gmail.com', '123', 'false', '2022-05-23', ''),
(4, 'Psedo', 'psedo@gmail.com', '123', 'false', '2022-05-23', ''),
(5, 'kris', 'kris123@gmail.com', '123', 'false', '2022-05-23', '');
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
