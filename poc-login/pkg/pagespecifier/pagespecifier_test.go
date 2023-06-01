package pagespecifier

import (
	"net/url"
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetResultFilter(t *testing.T) {
	g := NewGomegaWithT(t)

	CategoryPageSpec := PageSpecification{
		Filters: FilterGroupDefinitions{
			[]FilterDefinition{
				{
					Label:    "sample Select",
					Operator: ComparisonOperatorEqualTo,
					Data:     "c.parent_id", //base on database
					Variable: "parent_id",   //base on url query
					Type:     FilterTypeSelect,
				},
				{
					Label:    "Sample Date",
					Operator: ComparisonOperatorEqualTo, //for FilterTypeDateRange this one will be ignored and overridden
					Data:     "c.created_at",            //base on database
					Variable: "created_at",              //base on url query
					Type:     FilterTypeDateRange,
				},
			},
		},
		Columns: []Columns{
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
		ActionButtons: Actions{
			Top: []Action{
				{
					Id:          "download",
					Name:        "Download",
					GroupAccess: []string{"put Group Access here"},
					ElementID:   "Put ElementID here",
					Class:       "btn btn-sm btn-custom w-md waves-effect waves-light mr-10",
				},
			},
			Middle: []Action{
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
	filter := url.Values{}
	filter.Add("parent_id", "1")
	filter.Add("created_at_start", "2020-10-10 01:01:01")
	filter.Add("created_at_end", "2020-10-10 01:01:01")
	filter.Add("limit", "1")
	filter.Add("page", "0")
	rf := CategoryPageSpec.GetResultFilter(filter)
	g.Expect(rf.Limit).Should(Equal(int64(1)))
	g.Expect(rf.Page).Should(Equal(int64(1)))
	g.Expect(rf.Filter).Should(HaveLen(3))
	qf := GetQueryFilterWithParser(rf.Filter)
	g.Expect(qf.Args).Should(HaveLen(3))
	g.Expect(qf.Filter).Should(HaveLen(3))
}

func TestGetFilterData(t *testing.T) {

	CategoryPageSpec := PageSpecification{
		Filters: FilterGroupDefinitions{
			[]FilterDefinition{
				{
					Label:    "sample Select",
					Operator: ComparisonOperatorEqualTo,
					Data:     "c.parent_id", //base on database
					Variable: "parent_id",   //base on url query
					Type:     FilterTypeSelect,
				},
				{
					Label:    "Sample Date",
					Operator: ComparisonOperatorEqualTo, //for FilterTypeDateRange this one will be ignored and overridden
					Data:     "c.created_at",            //base on database
					Variable: "created_at",              //base on url query
					Type:     FilterTypeDateRange,
				},
			},
		},
		Columns: []Columns{
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
		ActionButtons: Actions{
			Top: []Action{
				{
					Id:          "download",
					Name:        "Download",
					GroupAccess: []string{"put Group Access here"},
					ElementID:   "Put ElementID here",
					Class:       "btn btn-sm btn-custom w-md waves-effect waves-light mr-10",
				},
			},
			Middle: []Action{
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
	opts := make(map[string][]SelectOption)
	opts["parent_id"] = []SelectOption{
		{Label: "Parent1", Value: "1"},
		{Label: "Parent2", Value: "2"},
	}
	CategoryPageSpec.GetFilterData(opts)

}
