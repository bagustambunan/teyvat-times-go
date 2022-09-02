package models

import "strconv"

type GetGiftClaimsOption struct {
	UserID   int
	StatusID int
	Limit    int
	Page     int
}

func NewGetGiftClaimsOption(params map[string][]string) (*GetGiftClaimsOption, error) {
	userIDVal := 0
	statusIDVal := 0
	limitVal := 10
	pageVal := 1

	var err error

	if params["userID"] != nil {
		userIDVal, err = strconv.Atoi(params["userID"][0])
		if err != nil {
			return nil, err
		}
	}
	if params["status"] != nil {
		statusIDVal, err = strconv.Atoi(params["status"][0])
		if err != nil {
			return nil, err
		}
	}
	if params["limit"] != nil {
		limitVal, err = strconv.Atoi(params["limit"][0])
		if err != nil {
			return nil, err
		}
		limitVal, err = parseLimit(limitVal)
		if err != nil {
			return nil, err
		}
	}
	if params["page"] != nil {
		pageVal, err = strconv.Atoi(params["page"][0])
		if err != nil {
			return nil, err
		}
		pageVal, err = parsePage(pageVal)
		if err != nil {
			return nil, err
		}
	}

	return &GetGiftClaimsOption{
		UserID:   userIDVal,
		StatusID: statusIDVal,
		Limit:    limitVal,
		Page:     pageVal,
	}, nil
}
