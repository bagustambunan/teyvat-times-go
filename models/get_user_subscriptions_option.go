package models

import (
	"strconv"
)

type GetUserSubscriptionsOption struct {
	UserID  int
	IsEnded int
	Limit   int
	Page    int
}

func NewGetUserSubscriptionsOption(params map[string][]string) (*GetUserSubscriptionsOption, error) {
	userIDVal := 0
	isEndedVal := 0
	limitVal := 10
	pageVal := 1

	var err error

	if params["userID"] != nil {
		userIDVal, err = strconv.Atoi(params["userID"][0])
		if err != nil {
			return nil, err
		}
	}
	if params["isEnded"] != nil {
		isEndedVal, err = strconv.Atoi(params["isEnded"][0])
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

	return &GetUserSubscriptionsOption{
		UserID:  userIDVal,
		IsEnded: isEndedVal,
		Limit:   limitVal,
		Page:    pageVal,
	}, nil
}
