package user

import "github.com/hasemeneh/poc-material/poc-login/pkg/pagespecifier"

var (
	UserPageSpec = pagespecifier.PageSpecification{
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
					Label:               "auth_id",
					Operator:            pagespecifier.ComparisonOperatorLike,
					Data:                "auth_id",
					FilterToQueryParser: pagespecifier.FilterToQueryGlobalMatch,
					Variable:            "auth_id",
					Type:                pagespecifier.FilterTypeText,
					IsShown:             true,
				},
			},
		},
	}
)
