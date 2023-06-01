package pagespecifier

import (
	"fmt"
	"strconv"
	"strings"
)

type Sort struct {
	Column    string
	SortOrder string
}

type QueryFilter struct {
	Filter []string
	Args   []interface{}
}

func (qf *QueryFilter) JoinFilter() string {
	return strings.Join(qf.Filter, " AND ")
}

func (qf *QueryFilter) GetWhereClause() string {
	if len(qf.Filter) == 0 {
		return ""
	}
	return fmt.Sprintf("WHERE %s", qf.JoinFilter())
}

type FilterQuery struct {
	Column     string
	Operator   string
	Value      interface{}
	ColumnHash string
	ValueHash  interface{}
}

type QuerySortParserFunc func(s Sort) string

type QuerySort struct {
	Sort []string
}

func (qs *QuerySort) JoinSort() string {
	return strings.Join(qs.Sort, ",")
}

func (qs *QuerySort) GetOrderClause() string {
	if len(qs.Sort) == 0 {
		return ""
	}
	return fmt.Sprintf("ORDER BY %s", qs.JoinSort())
}
func GetQuerySortWithParser(sort []Sort, sp QuerySortParserFunc) *QuerySort {
	qs := &QuerySort{
		Sort: make([]string, 0),
	}
	for _, s := range sort {
		sort := sp(s)
		if sort == "" {
			continue
		}
		qs.Sort = append(qs.Sort, sort)
	}
	return qs
}

func DefaultQuerySortParser(s Sort) string {
	return fmt.Sprintf("%s %s", s.Column, s.SortOrder)
}

func ParseSortV2(_sort string, m map[string]string) []Sort {
	var result = make([]Sort, 0)
	if _sort == "" {
		return result
	}
	s := strings.Split(_sort, ";")
	for i := range s {
		var sort Sort
		key := s[i]
		order := "ASC"
		if strings.HasPrefix(key, "-") {
			order = "DESC"
			key = key[1:]
		}
		column, ok := m[key]
		if ok {
			sort = Sort{
				Column:    column,
				SortOrder: order,
			}
			result = append(result, sort)
		}
	}
	return result
}

func parseInt64(s string) int64 {
	result, _ := strconv.ParseInt(s, 10, 64)
	return result
}
