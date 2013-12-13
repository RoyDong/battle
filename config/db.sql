DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
      `id` bigint(11) NOT NULL AUTO_INCREMENT,
      `name` varchar(25) COLLATE utf8_unicode_ci DEFAULT NULL,
      `salt` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
      `passwd` varchar(128) COLLATE utf8_unicode_ci NOT NULL,
      `email` varchar(60) COLLATE utf8_unicode_ci NOT NULL,
      `created_at` bigint(20) NOT NULL DEFAULT '0',
      `updated_at` bigint(20) NOT NULL DEFAULT '0',
      PRIMARY KEY (`id`),
      UNIQUE KEY `UNIQ_8D93D649E7927C74` (`email`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `topic`;

CREATE TABLE `topic` (
      `id` bigint(11) NOT NULL AUTO_INCREMENT,
      `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
      `content` text CHARACTER SET utf8,
      `state` int(11) NOT NULL DEFAULT '0',
      `created_at` bigint(20) NOT NULL DEFAULT '0',
      `updated_at` bigint(20) NOT NULL DEFAULT '0',
      PRIMARY KEY (`id`),
      KEY `title` (`title`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
      `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
      `tid` bigint(11) unsigned NOT NULL DEFAULT '0',
      `uid` bigint(11) unsigned NOT NULL DEFAULT '0',
      `content` text,
      `state` int(11) NOT NULL DEFAULT '0',
      `created_at` bigint(20) NOT NULL DEFAULT '0',
      `updated_at` bigint(20) NOT NULL DEFAULT '0',
      PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
