package role

import "github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"

var (
	RolePageSpec = pagespecifier.PageSpecification{
		Filters: pagespecifier.FilterGroupDefinitions{
			[]pagespecifier.FilterDefinition{
				{
					Label:    "status",
					Operator: pagespecifier.ComparisonOperatorEqualTo,
					Data:     "status",
					Variable: "status",
					Type:     pagespecifier.FilterTypeText,
					IsShown:  true,
				},
				{
					Label:               "name",
					Operator:            pagespecifier.ComparisonOperatorLike,
					Data:                "name",
					FilterToQueryParser: pagespecifier.FilterToQueryGlobalMatch,
					Variable:            "name",
					Type:                pagespecifier.FilterTypeText,
					IsShown:             true,
				},
			},
		},
	}
)
