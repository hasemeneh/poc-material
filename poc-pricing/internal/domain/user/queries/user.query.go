package queries

import (
	"context"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/pagespecifier"
)

type resource struct {
	db *database.Store
}
type Resource interface {
	Insert(ctx context.Context, dbTx database.DBTX, m *models.User) (int64, error)
	FetchUserByAuthID(ctx context.Context, dbTx database.DBTX, authID string) (*models.User, error)
	FetchUserForInternal(ctx context.Context, rf pagespecifier.ResultFilter) ([]*models.User, int64, error)
	FetchUserByID(ctx context.Context, dbTx database.DBTX, userID int64) (*models.User, error)
}

func NewUserQuery(db *database.Store) Resource {
	return &resource{
		db: db,
	}
}

const insertQuery = "INSERT INTO `rbac_user` (`auth_id`, `status`,`password`) VALUES (?, ?, ?);"

func (r *resource) Insert(ctx context.Context, dbTx database.DBTX, m *models.User) (int64, error) {
	result, err := r.db.GetExecContext(dbTx).ExecContext(ctx, insertQuery, m.AuthID, m.Status, m.Password)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return ID, nil
}

const selectUserByAuthID = "SELECT id,auth_id,password,status FROM rbac_user WHERE auth_id = ?"

func (r *resource) FetchUserByAuthID(ctx context.Context, dbTx database.DBTX, authID string) (*models.User, error) {
	resp := &models.User{}
	err := r.db.GetSelectContext(dbTx).GetContext(ctx, resp, selectUserByAuthID, authID)
	return resp, err
}

const selectUserID = "SELECT id,auth_id,password,status FROM rbac_user WHERE id = ?"

func (r *resource) FetchUserByID(ctx context.Context, dbTx database.DBTX, userID int64) (*models.User, error) {
	resp := &models.User{}
	err := r.db.GetSelectContext(dbTx).GetContext(ctx, resp, selectUserID, userID)
	return resp, err
}
