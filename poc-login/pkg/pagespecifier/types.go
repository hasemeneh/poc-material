package pagespecifier

import (
	"fmt"
	"strings"
)

type MetaData struct {
	Options []SelectOption `json:"options,omitempty"`
}

type SelectOption struct {
	Value string
	Label string
}

type Action struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	//GroupAccess need to check this to enterprise if this is not permitted to be wandering around then it need to be hidden but if it's permitted to wander arround it would be nice for front end to use this
	GroupAccess []string `json:"group_access"`
	//ElementID need to check this to enterprise //I think we can use this for better access management integration with enterprise UAM
	ElementID string `json:"element_id"`
	Type      string `json:"type"`
	Class     string `json:"class"`
}

type Actions struct {
	Top    []Action `json:"top"`
	Middle []Action `json:"middle"`
	Bottom []Action `json:"bottom"`
}

type Filter struct {
	ID         string     `json:"id"`
	Label      string     `json:"label"`
	Type       FilterType `json:"type"`
	TypeString string     `json:"type_string"`
	MetaData   MetaData   `json:"meta_data,omitempty"`
	Value      string     `json:"value"`
}

type FilterGroup struct {
	Label   string   `json:"label"`
	Filters []Filter `json:"filters"`
	Class   string   `json:"class"`
}

type Columns struct {
	Header         string `json:"header"`
	Data           string `json:"-"`
	Variable       string `json:"variable"`
	Orderable      bool   `json:"orderable"`
	DefaultContent string `json:"default_content"`
	ClassName      string `json:"class_name"`
}

type FilterData struct {
	GroupMap      []map[string][]FilterGroup `json:"filter_data"`
	Headers       []string                   `json:"headers"`
	IsReverse     bool                       `json:"is_reverse"`
	UserActions   map[string]string          `json:"user_actions"`
	ActionButtons Actions                    `json:"action_buttons"`
	Columns       []Columns                  `json:"columns"`
}

type DetailData struct {
	Title string        `json:"title"`
	Data  []DetailField `json:"data"`
}

type DetailField struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type IntoolsResponse struct {
	Table IntoolsTable  `json:"table"`
	Cards []InToolCards `json:"cards"`
}

type IntoolsTable struct {
	Total     int64       `json:"total"`
	TotalPage int64       `json:"total_page"`
	Page      int64       `json:"page"`
	Limit     int64       `json:"limit"`
	Rows      interface{} `json:"rows"`
}

type InToolCards struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ExportResult struct {
	Status   string `json:"status"`
	Location string `json:"location"`
}

type FilterType int

const (
	FilterTypeDateRange FilterType = iota
	FilterTypeSelect
	FilterTypeNumber
	FilterTypeText
	FilterTypeDate
)

var (
	FilterTypeString = map[FilterType]string{
		FilterTypeDateRange: "date",
		FilterTypeSelect:    "select",
		FilterTypeNumber:    "number",
		FilterTypeText:      "text",
		FilterTypeDate:      "date",
	}
)

func (c FilterType) String() string {
	name, ok := FilterTypeString[c]
	if !ok {
		return "text"
	}
	return name
}

type ComparisonOperator int

const (
	ComparisonOperatorEqualTo ComparisonOperator = iota
	ComparisonOperatorNotEqualTo
	ComparisonOperatorGreaterThan
	ComparisonOperatorLessThan
	ComparisonOperatorGreaterThanEqualTo
	ComparisonOperatorLessThanEqualTo
	ComparisonOperatorLike
)

var (
	ComparisonOperatorString = map[ComparisonOperator]string{
		ComparisonOperatorEqualTo:            "==",
		ComparisonOperatorNotEqualTo:         "!=",
		ComparisonOperatorGreaterThan:        ">",
		ComparisonOperatorLessThan:           "<",
		ComparisonOperatorGreaterThanEqualTo: ">=",
		ComparisonOperatorLessThanEqualTo:    "<=",
		ComparisonOperatorLike:               "like",
	}
)

func (c ComparisonOperator) String() string {
	name, ok := ComparisonOperatorString[c]
	if !ok {
		return "INVALID"
	}
	return name
}

type Hasher interface {
	HashString(value string) string
}
type FilterDefinition struct {
	Label string
	//Variable put url variabel name here
	Variable string
	//Operator please use the constant if you want to add new addition please make adjusment on the code
	Operator            ComparisonOperator
	FilterToQueryParser FilterToQueryParser
	//Data put database variabel name here
	Data string
	//FilterType please use the constant if you want to add new addition please make adjusment on the code
	Type    FilterType
	Hasher  Hasher
	IsShown bool
}

type ResultFilter struct {
	Filter  []FilterValue
	Sort    []Sort
	Page    int64
	Limit   int64
	Finance bool
	Export  bool
}

type FilterValue struct {
	Column              string
	ColumnHash          string
	Operator            string
	FilterToQueryParser FilterToQueryParser
	Value               interface{}
	ValueHash           interface{}
}

type FilterToQueryParser int

const (
	FilterToQueryDefault FilterToQueryParser = iota
	FilterToQueryGlobalMatch
	FilterToQueryStartsWith
	FilterToQueryEndsWith
)

var (
	FilterToQueryParserFunc = map[FilterToQueryParser]QueryFilterParserFunc{
		FilterToQueryDefault:     DefaultQueryFilterParser,
		FilterToQueryGlobalMatch: GlobalMatchFilterParser,
		FilterToQueryStartsWith:  StartsWithFilterParser,
		FilterToQueryEndsWith:    EndsWithFilterParser,
	}
)

func (c FilterToQueryParser) ParserFunc() QueryFilterParserFunc {

	name, ok := FilterToQueryParserFunc[c]
	if !ok {
		return DefaultQueryFilterParser
	}
	return name
}

type QueryFilterParserFunc func(f FilterQuery) (string, interface{})

func DefaultQueryFilterParser(f FilterQuery) (string, interface{}) {
	return parseQueryFilter("", "?", f)
}

func parseQueryFilter(decorator string, token string, f FilterQuery) (string, interface{}) {
	if decorator == "" {
		decorator = "%s"
	}
	if strings.EqualFold(f.Operator, "in") || strings.EqualFold(f.Operator, "not in") {
		l := 1
		arr, ok := f.Value.([]interface{})
		if ok {
			l = len(arr)
		}
		params := make([]string, l)
		for i := range params {
			params[i] = token
		}
		return fmt.Sprintf(`%s %s (%s)`, fmt.Sprintf(decorator, f.Column), f.Operator, strings.Join(params, ",")), f.Value
	}
	return fmt.Sprintf(`%s %s %s`, fmt.Sprintf(decorator, f.Column), f.Operator, token), f.Value
}
