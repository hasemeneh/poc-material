package role

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-pricing/internal/models"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/database"
	"github.com/hasemeneh/poc-material/poc-pricing/pkg/pagespecifier"
)

func (r *repo) GetRoleForInternal(ctx context.Context, queries url.Values) (*pagespecifier.IntoolsResponse, error) {
	rf := RolePageSpec.GetResultFilter(queries)
	data, total, err := r.queries.FetchRoleForInternal(ctx, rf)
	if err != nil {
		return nil, err
	}
	var totalPage int64
	if rf.Limit > 0 {
		totalPage = (total / rf.Limit)
	}

	return &pagespecifier.IntoolsResponse{
		Table: pagespecifier.IntoolsTable{
			Limit:     rf.Limit,
			Page:      rf.Page,
			TotalPage: totalPage + 1,
			Rows:      data,
			Total:     total,
		},
	}, nil
}

func (r *repo) AddRole(ctx context.Context, dbTx database.DBTX, req *models.InternalRole) (int64, error) {
	return r.queries.InsertRole(ctx, dbTx, req)
}

func (r *repo) UpdateRole(ctx context.Context, req *models.InternalRole) error {
	return r.queries.UpdateRole(ctx, nil, req)
}

func (r *repo) FetchUserAccessByIDInternal(ctx context.Context, userID int64) ([]*models.UserAccess, error) {
	return r.queries.FetchUserAccessByIDInternal(ctx, userID)
}

func (r *repo) AddUserAccess(ctx context.Context, req *models.UserAccess) error {
	_, err := r.queries.InsertUserAccess(ctx, nil, req)
	return err
}

func (r *repo) ChangeStatusUserAccess(ctx context.Context, req *models.UserAccess) error {
	return r.queries.ChangeStatusUserAccess(ctx, nil, req.Status, req.RoleID, req.UserID)
}
