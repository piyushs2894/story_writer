DROP DATABASE IF EXISTS `story_db`;
Create Database IF NOT EXISTS `story_db`;
CREATE USER 'verloop'@'localhost' IDENTIFIED BY 'verloop';
USE `story_db`;
GRANT ALL PRIVILEGES ON story_db.* TO 'verloop'@'localhost';

CREATE TABLE `words` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `word` varchar(255) NOT NULL,
  `sentence_id` int(11),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `sentences` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `paragraph_id` int(11),
  `length` int(11),
  `status` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `paragraphs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `story_id` int(11),
  `length` int(11),
  `status` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `stories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `length` int(11),
  `title_length` int(11),
  `status` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `paragraphs` ADD CONSTRAINT `fk_stories_paragraph` FOREIGN KEY (`story_id`) REFERENCES `stories` (`id`) ON DELETE CASCADE;
ALTER TABLE `sentences` ADD CONSTRAINT `fk_paragraphs_sentence` FOREIGN KEY (`paragraph_id`) REFERENCES `paragraphs` (`id`) ON DELETE CASCADE;
ALTER TABLE `words` ADD CONSTRAINT `fk_sentences_words` FOREIGN KEY (`sentence_id`) REFERENCES `sentences` (`id`) ON DELETE CASCADE;