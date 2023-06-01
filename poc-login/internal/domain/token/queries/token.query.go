package queries

import (
	"context"
	"time"

	"github.com/hasemeneh/poc-material/poc-login/internal/constants"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/database"
	"github.com/jmoiron/sqlx"
)

type resource struct {
	db *database.Store
}
type Resource interface {
	StartTransaction(ctx context.Context) (*sqlx.Tx, error)
	FetchtToBeInvalidRefreshTokenByUserID(ctx context.Context, user_id int64) ([]*models.RefreshToken, error)
	InsertRefreshToken(ctx context.Context, dbTx database.DBTX, userID int64, refreshToken string, expiration *time.Time, status constants.RefreshTokenStatus) (int64, error)
	UpdateRefreshToken(ctx context.Context, dbTx database.DBTX, status constants.RefreshTokenStatus, id int64) error
	FetchRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error)
}

func NewQuery(db *database.Store) Resource {
	return &resource{
		db: db,
	}
}

func (r *resource) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return r.db.GetMaster().BeginTxx(ctx, nil)

}

const fetchRefreshTokenQuery = `
SELECT
	rt.id,
	rt.token,
	rt.user_id,
	rt.status,
	rt.expires_at
FROM rbac_refresh_token rt
WHERE rt.token = ?
LIMIT 1
`

func (r *resource) FetchRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	resp := &models.RefreshToken{}
	err := r.db.GetSlave().GetContext(ctx, resp, fetchRefreshTokenQuery, token)
	return resp, err
}

const fetchInvalidRefreshTokenQuery = `
SELECT
	rt.id,
	rt.token,
	rt.user_id,
	rt.status,
	rt.expires_at
FROM rbac_refresh_token rt
WHERE rt.user_id = ? AND rt.expires_at > NOW() AND rt.status = ?
ORDER BY rt.expires_at DESC
LIMIT 18446744073709551615
OFFSET ?
`

// using big nuber according to this https://dev.mysql.com/doc/refman/8.0/en/select.html#id4651990
func (r *resource) FetchtToBeInvalidRefreshTokenByUserID(ctx context.Context, user_id int64) ([]*models.RefreshToken, error) {
	resp := make([]*models.RefreshToken, 0)
	err := r.db.GetSlave().SelectContext(ctx, &resp, fetchInvalidRefreshTokenQuery, user_id, constants.RefreshTokenStatusActive, constants.RefreshTokenQuantityLimit)
	return resp, err
}

const insertQuery = "INSERT INTO `rbac_refresh_token` (`user_id`, `token`, `status`, `expires_at`) VALUES (?, ?, ?, ?);"

func (r *resource) InsertRefreshToken(ctx context.Context, dbTx database.DBTX, userID int64, refreshToken string, expiration *time.Time, status constants.RefreshTokenStatus) (int64, error) {
	result, err := r.db.GetExecContext(dbTx).ExecContext(ctx, insertQuery, userID, refreshToken, status, expiration)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return ID, nil
}

const updateQuery = "UPDATE `rbac_refresh_token` SET  `status` = ? WHERE `rbac_refresh_token`.`id` = ?;"

func (r *resource) UpdateRefreshToken(ctx context.Context, dbTx database.DBTX, status constants.RefreshTokenStatus, id int64) error {
	_, err := r.db.GetExecContext(dbTx).ExecContext(ctx, updateQuery, status, id)
	return err
}
