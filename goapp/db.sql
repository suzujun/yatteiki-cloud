
CREATE DATABASE yatteiki;

CREATE USER yatteiki_admin IDENTIFIED BY 'xkMhzFyZWL5ndR0KnACb3KMsKCwbx46n';
GRANT ALL ON `%`.* TO yatteiki_admin;

CREATE TABLE `todos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `completed` bool NOT NULL default 0,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
