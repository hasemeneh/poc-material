package pagespecifier

import (
	"fmt"
)

//GetQueryFilterWithParser added 'like' operator version util.GetQueryFilterWithParser from lending-platform
func GetQueryFilterWithParser(filters []FilterValue) *QueryFilter {
	operator := map[string]string{
		"==":     "=",
		"!=":     "!=",
		">=":     ">=",
		"<=":     "<=",
		">":      ">",
		"<":      "<",
		"in":     "IN",
		"not in": "NOT IN",
		"like":   "LIKE",
	}
	qf := &QueryFilter{
		Filter: make([]string, 0),
		Args:   make([]interface{}, 0),
	}
	for _, f := range filters {
		opr, ok := operator[f.Operator]
		if ok {
			f.Operator = opr
			fp := f.FilterToQueryParser.ParserFunc()
			s, a := fp(FilterQuery{
				Column:     f.Column,
				Operator:   f.Operator,
				Value:      f.Value,
				ColumnHash: f.ColumnHash,
				ValueHash:  f.ValueHash,
			})
			if s == "" {
				continue
			}
			qf.Filter = append(qf.Filter, s)
			arr, ok := a.([]interface{})
			if ok {
				for _, item := range arr {
					qf.Args = append(qf.Args, item)
				}
			} else {
				qf.Args = append(qf.Args, a)
			}
		}
	}
	return qf
}

//GlobalMatchFilterParser use this only with like operator or else it won't work as expected
func GlobalMatchFilterParser(f FilterQuery) (string, interface{}) {
	return fmt.Sprintf(`%s %s %s`, f.Column, f.Operator, "?"), ("%" + fmt.Sprintf("%v", f.Value) + "%")
}

//StartsWithFilterParser use this only with like operator or else it won't work as expected
func StartsWithFilterParser(f FilterQuery) (string, interface{}) {
	return fmt.Sprintf(`%s %s %s`, f.Column, f.Operator, "?"), (fmt.Sprintf("%v", f.Value) + "%")
}

//EndsWithFilterParser use this only with like operator or else it won't work as expected
func EndsWithFilterParser(f FilterQuery) (string, interface{}) {
	return fmt.Sprintf(`%s %s %s`, f.Column, f.Operator, "?"), ("%" + fmt.Sprintf("%v", f.Value))
}
