START TRANSACTION;

TRUNCATE `rbac_permission`;
TRUNCATE `rbac_role`;
TRUNCATE `rbac_role_permission`;

COMMIT;
