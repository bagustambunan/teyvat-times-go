package models

import (
	"strconv"
)

type ReadHistoryOption struct {
	Limit int
	Page  int
}

func NewReadHistoryOption(params map[string][]string) (*ReadHistoryOption, error) {
	limitVal := 10
	pageVal := 1
	var err error

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

	return &ReadHistoryOption{
		Limit: limitVal,
		Page:  pageVal,
	}, nil
}
