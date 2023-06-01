package queries

import (
	"context"
	"time"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/constants"
	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
)

const insertQuery = "INSERT INTO `rbac_role` (`name`, `status`) VALUES(?, ?);"

func (r *resource) InsertRole(ctx context.Context, dbTx database.DBTX, m *models.InternalRole) (int64, error) {
	result, err := r.db.GetExecContext(dbTx).ExecContext(ctx, insertQuery, m.Name, m.Status)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return ID, nil
}

const updateQuery = "UPDATE `rbac_role` SET `name` = ?, `status` = ?, `deleted_at` = ? WHERE `rbac_role`.`id` = ?;"

func (r *resource) UpdateRole(ctx context.Context, dbTx database.DBTX, m *models.InternalRole) error {
	var deletedAt *time.Time
	if m.Status == constants.RoleStatusInactive {
		temp := time.Now()
		deletedAt = &temp
	}
	_, err := r.db.GetExecContext(dbTx).ExecContext(ctx, updateQuery, m.Name, m.Status, deletedAt, m.ID)
	if err != nil {
		return err
	}

	return nil
}

const fetchRoleByIDInternalQuery = `
SELECT
	ua.role_id,
	ua.user_id,
	r.name,
	ua.status
FROM rbac_user_access ua
INNER JOIN rbac_role r ON ua.role_id = r.id 
WHERE ua.user_id = ?
`

func (r *resource) FetchUserAccessByIDInternal(ctx context.Context, userID int64) ([]*models.UserAccess, error) {
	resp := make([]*models.UserAccess, 0)
	err := r.db.GetSlave().SelectContext(ctx, &resp, fetchRoleByIDInternalQuery, userID)
	return resp, err
}

const insertUserAccessQuery = "INSERT INTO `rbac_user_access` (`user_id`, `role_id`, `status`) VALUES ( ?, ?, ?);"

func (r *resource) InsertUserAccess(ctx context.Context, dbTx database.DBTX, m *models.UserAccess) (int64, error) {
	result, err := r.db.GetExecContext(dbTx).ExecContext(ctx, insertUserAccessQuery, m.UserID, m.RoleID, m.Status)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return ID, nil
}

const changeStatusUserAccessQuery = "UPDATE `rbac_user_access` SET `status` = ? WHERE `rbac_user_access`.`role_id` = ? AND `rbac_user_access`.`user_id` = ?;"

func (r *resource) ChangeStatusUserAccess(ctx context.Context, dbTx database.DBTX, status constants.UserAccessStatus, RoleID, UserID int64) error {
	_, err := r.db.GetExecContext(dbTx).ExecContext(ctx, changeStatusUserAccessQuery, status, RoleID, UserID)
	if err != nil {
		return err
	}

	return nil
}
