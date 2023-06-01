START TRANSACTION;

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'permission.view', 'permission.view', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 'permission.add', 'permission.add', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(3, 'permission.update', 'permission.update', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(4, 'role.view', 'role.view', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(5, 'role.add', 'role.add', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(6, 'role.update', 'role.update', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(7, 'role_permission.view', 'role_permission.view', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(8, 'role_permission.add', 'role_permission.add', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(9, 'role_permission.update', 'role_permission.update', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(10, 'user.view', 'user.view', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(11, 'user_access.view', 'user_access.view', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(12, 'user_access.add', 'user_access.add', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_permission` (`id`, `name`, `code`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(13, 'user_access.update', 'user_access.update', 'ACTIVE', '2021-09-09 10:29:24', NULL, NULL);

INSERT INTO `rbac_role` (`id`, `name`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'permission', 'ACTIVE', '2021-09-20 01:13:14', NULL, NULL);
INSERT INTO `rbac_role` (`id`, `name`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 'role', 'ACTIVE', '2021-09-20 01:13:14', NULL, NULL);
INSERT INTO `rbac_role` (`id`, `name`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(3, 'role_permission', 'ACTIVE', '2021-09-20 01:13:14', NULL, NULL);
INSERT INTO `rbac_role` (`id`, `name`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(4, 'user', 'ACTIVE', '2021-09-20 01:13:14', NULL, NULL);
INSERT INTO `rbac_role` (`id`, `name`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(5, 'user_access', 'ACTIVE', '2021-09-20 01:13:14', NULL, NULL);
INSERT INTO `rbac_role` (`id`, `name`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(6, 'super_admin', 'ACTIVE', '2021-09-20 01:13:14', NULL, NULL);

INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 1, 1, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 1, 2, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 1, 3, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 2, 4, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 2, 5, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 2, 6, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 3, 7, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 3, 8, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 3, 9, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 4, 10, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 5, 11, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 5, 12, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 5, 13, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);

INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 1, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 2, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 3, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 4, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 5, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 6, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 7, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 8, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 9, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 10, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 11, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 12, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);
INSERT INTO `rbac_role_permission` (`id`, `role_id`, `permission_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(NULL, 6, 13, 'ACTIVE', '2021-09-20 01:13:29', NULL, NULL);


COMMIT;
