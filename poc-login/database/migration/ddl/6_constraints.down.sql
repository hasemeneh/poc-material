BEGIN;
ALTER TABLE `rbac_permission`
  DROP PRIMARY KEY (`id`),
  DROP UNIQUE KEY `code` (`code`);

ALTER TABLE `rbac_role`
  DROP PRIMARY KEY (`id`);

ALTER TABLE `rbac_role_permission`
  DROP PRIMARY KEY (`id`),
  DROP KEY `role_id` (`role_id`),
  DROP KEY `permission_id` (`permission_id`);

ALTER TABLE `rbac_user`
  DROP PRIMARY KEY (`id`),
  DROP UNIQUE KEY `auth_id` (`auth_id`);

ALTER TABLE `rbac_user_access`
  DROP PRIMARY KEY (`id`),
  DROP KEY `user_id` (`user_id`),
  DROP KEY `role_id` (`role_id`);


ALTER TABLE `rbac_permission`
  MODIFY `id` int(10) UNSIGNED NOT NULL;

ALTER TABLE `rbac_role`
  MODIFY `id` int(10) UNSIGNED NOT NULL;

ALTER TABLE `rbac_role_permission`
  MODIFY `id` int(10) UNSIGNED NOT NULL;

ALTER TABLE `rbac_user`
  MODIFY `id` int(10) UNSIGNED NOT NULL;

ALTER TABLE `rbac_user_access`
  MODIFY `id` int(10) UNSIGNED NOT NULL;


ALTER TABLE `rbac_role_permission`
  DROP CONSTRAINT `rbac_role_permission_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `rbac_role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  DROP CONSTRAINT `rbac_role_permission_ibfk_2` FOREIGN KEY (`permission_id`) REFERENCES `rbac_permission` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE `rbac_user_access`
  DROP CONSTRAINT `rbac_user_access_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `rbac_user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  DROP CONSTRAINT `rbac_user_access_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `rbac_role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

COMMIT;