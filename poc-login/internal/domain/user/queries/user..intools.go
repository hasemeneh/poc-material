package queries

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hasemeneh/poc-material/poc-login/internal/models"
	"github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"
)

const fetchInternalQuery = `
SELECT
	id,
	auth_id,
	status
FROM rbac_user
`

func (u *resource) FetchUserForInternal(ctx context.Context, rf pagespecifier.ResultFilter) ([]*models.User, int64, error) {
	qf := pagespecifier.GetQueryFilterWithParser(rf.Filter)
	args := qf.Args
	whereClause := qf.GetWhereClause()
	orderClause := pagespecifier.GetQuerySortWithParser(rf.Sort, pagespecifier.DefaultQuerySortParser).GetOrderClause()

	var limitClause string
	if rf.Limit > 0 {
		var offset int64
		if rf.Page > 0 {
			offset = (rf.Page - 1) * rf.Limit
		}
		limitClause = fmt.Sprintf("LIMIT %d OFFSET %d", rf.Limit, offset)
	}
	query := fmt.Sprintf(`%s %s %s %s`, fetchInternalQuery, whereClause, orderClause, limitClause)
	p := make([]*models.User, 0)
	var total int64

	dbTx, err := u.db.GetSlave().BeginTxx(ctx, nil)
	if err != nil {
		return p, total, err
	}
	defer dbTx.Rollback()

	err = dbTx.SelectContext(ctx, &p, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, 0, nil
		}
		return p, 0, err
	}
	row := dbTx.QueryRow("SELECT FOUND_ROWS()")
	err = row.Scan(&total)
	if err != nil {
		return p, 0, err
	}

	return p, total, nil
}
