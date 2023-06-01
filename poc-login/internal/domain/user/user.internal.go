package user

import (
	"context"
	"net/url"

	"github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"
)

func (r *repo) GetUserForInternal(ctx context.Context, queries url.Values) (*pagespecifier.IntoolsResponse, error) {
	rf := UserPageSpec.GetResultFilter(queries)
	data, total, err := r.queries.FetchUserForInternal(ctx, rf)
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
