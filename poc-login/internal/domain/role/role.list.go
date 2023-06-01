package role

import (
	"context"
	"database/sql"
	"log"

	"github.com/hasemeneh/poc-material/poc-login/internal/custerr"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
)

func (r *repo) GetUserValidRole(ctx context.Context, userToken *models.UserToken) ([]*models.Role, error) {

	roles, err := r.cache.GetUserRole(userToken.AuthID)
	if err == nil {
		return roles, nil
	}
	if err != custerr.ErrNoPermissionFound {
		return nil, err
	}
	roles, err = r.queries.FetchValidRoleByAuthID(ctx, userToken.AuthID)
	if err == nil {
		err2 := r.cache.InsertUserRole(userToken.AuthID, userToken.Expiry, roles)
		if err2 != nil {
			log.Println("failed to cache user access because ", err2)
		}
		return roles, nil
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	return roles, nil
}
func (r *repo) RevokeUserRole(ctx context.Context, userToken *models.UserToken) error {
	return r.cache.RemoveUserRole(userToken.AuthID)
}
