INSERT INTO `rbac_user`(`id`,`auth_id`,`password`,`status`,`created_at`,`updated_at`,`deleted_at`) VALUES ( '1', 'admin', '$2a$10$naaIc3zefo3JSRC8JL/d8uum15rVJn2P4GSLhQwgZKXEAMWoHYTwS', 'ACTIVE', '2022-12-18 14:31:59', NULL, NULL );
INSERT INTO `rbac_user`(`id`,`auth_id`,`password`,`status`,`created_at`,`updated_at`,`deleted_at`) VALUES ( '2', 'M1-ABDI-11', '$2a$10$naaIc3zefo3JSRC8JL/d8uum15rVJn2P4GSLhQwgZKXEAMWoHYTwS', 'ACTIVE', '2022-12-18 14:31:59', NULL, NULL );
INSERT INTO `rbac_user`(`id`,`auth_id`,`password`,`status`,`created_at`,`updated_at`,`deleted_at`) VALUES ( '3', 'M1-PUMS-11', '$2a$10$naaIc3zefo3JSRC8JL/d8uum15rVJn2P4GSLhQwgZKXEAMWoHYTwS', 'ACTIVE', '2022-12-18 14:31:59', NULL, NULL );
INSERT INTO `rbac_user`(`id`,`auth_id`,`password`,`status`,`created_at`,`updated_at`,`deleted_at`) VALUES ( '4', 'M1-STQI-11', '$2a$10$naaIc3zefo3JSRC8JL/d8uum15rVJn2P4GSLhQwgZKXEAMWoHYTwS', 'ACTIVE', '2022-12-18 14:31:59', NULL, NULL );
INSERT INTO `rbac_user`(`id`,`auth_id`,`password`,`status`,`created_at`,`updated_at`,`deleted_at`) VALUES ( '5', 'M1-SUGP-11', '$2a$10$naaIc3zefo3JSRC8JL/d8uum15rVJn2P4GSLhQwgZKXEAMWoHYTwS', 'ACTIVE', '2022-12-18 14:31:59', NULL, NULL );
INSERT INTO `rbac_user`(`id`,`auth_id`,`password`,`status`,`created_at`,`updated_at`,`deleted_at`) VALUES ( '6', 'M1-SURA-11', '$2a$10$naaIc3zefo3JSRC8JL/d8uum15rVJn2P4GSLhQwgZKXEAMWoHYTwS', 'ACTIVE', '2022-12-18 14:31:59', NULL, NULL );
-- ---------------------------------------------------------


INSERT INTO `rbac_user_access`(`id`,`user_id`,`role_id`,`status`,`created_at`,`updated_at`,`deleted_at`) VALUES ( '1', '1', '6', 'ACTIVE', '2022-12-18 14:33:41', NULL, NULL );
