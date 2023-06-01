package queries

import (
	"context"
	"time"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
)

const insertQuery = "INSERT INTO `rbac_permission` (`name`, `code`, `status`) VALUES(?, ?, ?);"

func (r *resource) InsertPermission(ctx context.Context, dbTx database.DBTX, m *models.Permission) (int64, error) {
	result, err := r.db.GetExecContext(dbTx).ExecContext(ctx, insertQuery, m.Name, m.Code, m.Status)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return ID, nil
}

const updateQuery = "UPDATE `rbac_permission` SET `name` = ?, `code` = ?, `status` = ?, `deleted_at` = ? WHERE `rbac_permission`.`id` = ?;"

func (r *resource) UpdatePermission(ctx context.Context, dbTx database.DBTX, m *models.Permission) error {
	var deletedAt *time.Time
	if m.Status == constants.PermissionStatusInactive {
		temp := time.Now()
		deletedAt = &temp
	}
	_, err := r.db.GetExecContext(dbTx).ExecContext(ctx, updateQuery, m.Name, m.Code, m.Status, deletedAt, m.ID)
	if err != nil {
		return err
	}

	return nil
}

const insertRolePermissionQuery = "INSERT INTO `rbac_role_permission` (`role_id`, `permission_id`, `status`) VALUES( ?, ?, ?);"

func (r *resource) InsertRolePermission(ctx context.Context, dbTx database.DBTX, m *models.InternalRolePermission) (int64, error) {
	result, err := r.db.GetExecContext(dbTx).ExecContext(ctx, insertRolePermissionQuery, m.RoleID, m.PermissionID, m.Status)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return ID, nil
}

const changeStatusRolePermissionQuery = "UPDATE `rbac_role_permission` SET `status` = ? WHERE `rbac_role_permission`.`role_id` = ? AND `rbac_role_permission`.`permission_id` = ?;"

func (r *resource) ChangeStatusRolePermission(ctx context.Context, dbTx database.DBTX, status constants.RolePermissionStatus, RoleID, PermissionID int64) error {
	_, err := r.db.GetExecContext(dbTx).ExecContext(ctx, changeStatusRolePermissionQuery, status, RoleID, PermissionID)
	if err != nil {
		return err
	}

	return nil
}
