package permission

import (
	"context"
	"database/sql"
	"log"

	"github.com/hasemeneh/poc-material/poc-login/internal/custerr"
	"github.com/hasemeneh/poc-material/poc-login/internal/models"
)

func (r *repo) GetAccessByUserToken(ctx context.Context, userToken *models.UserToken) (map[int64][]*models.Permission, error) {
	permissions, err := r.GetAccessByUserTokenForInternal(ctx, userToken)
	if err != nil {
		return nil, err
	}

	return composePermission(permissions), nil
}

func (r *repo) GetAccessByUserTokenForInternal(ctx context.Context, userToken *models.UserToken) ([]*models.Permission, error) {
	permissions, err := r.cache.GetUserAccess(userToken.AuthID)
	if err == nil {
		return permissions, nil
	}
	if err != custerr.ErrNoPermissionFound {
		return nil, err
	}
	permissions, err = r.queries.FetchValidPermissionByAuthID(ctx, userToken.AuthID)
	if err == nil {
		err2 := r.cache.InsertUserAccess(userToken.AuthID, userToken.Expiry, permissions)
		if err2 != nil {
			log.Println("failed to cache user access because ", err2)
		}
		return permissions, nil
	}
	if err != sql.ErrNoRows {
		return nil, err
	}

	return permissions, nil
}

func composePermission(req []*models.Permission) map[int64][]*models.Permission {
	resp := make(map[int64][]*models.Permission)
	for k := range req {
		resp[req[k].RoleID] = append(resp[req[k].RoleID], req[k])
	}
	return resp
}

func (r *repo) RevokeUserAccess(ctx context.Context, userToken *models.UserToken) error {
	return r.cache.RemoveUserAccess(userToken.AuthID)
}
