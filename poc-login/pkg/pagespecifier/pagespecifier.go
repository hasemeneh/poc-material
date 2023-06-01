package pagespecifier

import (
	"fmt"
	"net/url"
)

/*
PageSpecification Will Specify some page detail to front end

avoid using exact keyword sort,page,limit on filter variabel as it is reserved for the functional purpose
if you need it so badly then you have to find other way or don't use this feature make your own.

example :
	CategoryPageSpec = pagespecifier.PageSpecification{
		Filters: pagespecifier.FilterGroupDefinitions{
			[]pagespecifier.FilterDefinition{
				{
					Label:    "Parent Category",
					Operator: pagespecifier.ComparisonOperatorEqualTo,
					Data:     "c.parent_id", //base on database
					Variable: "parent_id", //base on url query
					Type:     pagespecifier.FilterTypeSelect,
				},
			},
		},
		Columns: []pagespecifier.Columns{
			{Header: "Column1", Data: "i_am_the_json_data1", Orderable: true},
			{Header: "Column2", Data: "i_am_the_json_data2", Orderable: false},
			{Header: "Column3", Data: "i_am_the_json_data3", Orderable: true},
			{Header: "Column4", Data: "i_am_the_json_data4", Orderable: false},
			{Header: "Column5", Data: "i_am_the_json_data5", Orderable: true},
			{Header: "Column6", Data: "i_am_the_json_data6", Orderable: true},
			{Header: ""}, //Put Blank for action column
		},
		UserActions: map[string]string{
			"searchUrl": "/order",
		},
		ActionButtons: pagespecifier.Actions{
			Top: []pagespecifier.Action{
				{
					Id:          "download",
					Name:        "Download",
					GroupAccess: []string{"put Group Access here"},
					ElementID:   "Put ElementID here",
					Class:       "btn btn-sm btn-custom w-md waves-effect waves-light mr-10",
				},
			},
			Middle: []pagespecifier.Action{
				{
					Id:          "search",
					Name:        "Search",
					Type:        "submit",
					GroupAccess: []string{"put Group Access here"},
					ElementID:   "Put ElementID here",
					Class:       "btn btn-sm btn-search btn-primary w-md  waves-effect waves-light mr-10",
				},
				{
					Id:          "reset",
					Name:        "Reset",
					Type:        "reset",
					GroupAccess: []string{"ftinins-product", "fitnins-business"},
					ElementID:   "Put ElementID here",
					Class:       "btn btn-sm btn-default btn-custom w-md waves-effect waves-light mr-10",
				},
			},
		},
	}
*/
type PageSpecification struct {
	Filters       FilterGroupDefinitions
	Columns       []Columns
	UserActions   map[string]string
	ActionButtons Actions
}

func (ps *PageSpecification) GetResultFilter(input url.Values) ResultFilter {
	orderMap := map[string]string{}
	for _, val := range ps.Columns {
		if val.Orderable {
			orderMap[val.Variable] = val.Data
		}
	}
	rf := ResultFilter{
		Filter: ps.Parsefilter(input),
		Sort:   ParseSortV2(input.Get("sort"), orderMap),
		Page:   parseInt64(input.Get("page")),
		Limit:  parseInt64(input.Get("limit")),
	}
	if rf.Page == 0 {
		rf.Page = 1
	}
	return rf
}

/*Parsefilter equivalent of parser.ParseFilterRaw but with input of url.values
but do not support 'in' clause yet will add it later :D */
func (ps *PageSpecification) Parsefilter(input url.Values) []FilterValue {
	filters := make([]FilterValue, 0)
	for _, val := range ps.Filters {
		for _, filter := range val {
			if filter.Type == FilterTypeDateRange {
				if data := input.Get(fmt.Sprintf("%s_start", filter.Variable)); data != "" {
					filters = append(filters, FilterValue{
						Column:              filter.Data,
						Operator:            ComparisonOperatorGreaterThanEqualTo.String(),
						FilterToQueryParser: filter.FilterToQueryParser,
						Value:               data,
					})
				}
				if data := input.Get(fmt.Sprintf("%s_end", filter.Variable)); data != "" {
					filters = append(filters, FilterValue{
						Column:              filter.Data,
						Operator:            ComparisonOperatorLessThan.String(),
						FilterToQueryParser: filter.FilterToQueryParser,
						Value:               data,
					})
				}
			} else {
				if data := input.Get(filter.Variable); data != "" {
					filterValue := FilterValue{
						Column:              filter.Data,
						Operator:            filter.Operator.String(),
						FilterToQueryParser: filter.FilterToQueryParser,
						Value:               data,
					}
					filters = append(filters, filterValue)
				}
			}
		}
	}

	return filters
}
func (ps *PageSpecification) GetFilterData(options map[string][]SelectOption) *FilterData {
	return &FilterData{
		GroupMap:      ps.Filters.GetFilterGroup(options),
		Columns:       ps.Columns,
		UserActions:   ps.UserActions,
		ActionButtons: ps.ActionButtons,
	}
}

type FilterGroupDefinitions [][]FilterDefinition

func (f *FilterGroupDefinitions) GetFilterGroup(options map[string][]SelectOption) []map[string][]FilterGroup { //need to revise this
	groupMap := []map[string][]FilterGroup{}
	for _, group := range *f {
		filters := []FilterGroup{}
		for _, value := range group {
			filterGroup := FilterGroup{
				Label: value.Label,
			}
			if value.Type == FilterTypeDateRange {
				filterGroup.Class = "filter-input form-inline"
				filterGroup.Filters = []Filter{
					{
						Label:      "Date Start",
						Type:       value.Type,
						TypeString: value.Type.String(),
						ID:         fmt.Sprintf("%s_start", value.Variable),
					},
					{
						Label:      "Date End",
						Type:       value.Type,
						TypeString: value.Type.String(),
						ID:         fmt.Sprintf("%s_end", value.Variable),
					},
				}
			} else {
				filterGroup.Class = "filter-input form-group"
				filterGroup.Filters = []Filter{
					{
						Label:      fmt.Sprintf("[%s]", value.Label),
						Type:       value.Type,
						TypeString: value.Type.String(),
						ID:         value.Variable,
					},
				}
				if value.Type == FilterTypeSelect {
					filterGroup.Filters[0].Label = "select"
					filterGroup.Filters[0].MetaData = MetaData{
						Options: options[value.Variable],
					}
				}
			}
			filters = append(filters, filterGroup)
		}
		groupMap = append(groupMap, map[string][]FilterGroup{
			"col-sm-6": filters,
		})
	}
	return groupMap
}
