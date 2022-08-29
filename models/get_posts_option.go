package models

import (
	"final-project-backend/httperror"
	"strconv"
)

type GetPostsOption struct {
	S         string
	Category  int
	Tier      int
	SortBy    string
	SortOrder string
	Limit     int
	Page      int
}

func NewGetPostsOption(params map[string][]string) (*GetPostsOption, error) {
	sVal := ""
	categoryVal := 0
	tierVal := 0
	sortByVal := "created_at"
	sortOrderVal := "desc"
	limitVal := 10
	pageVal := 1
	var err error

	if params["s"] != nil {
		sVal = params["s"][0]
	}
	if params["category"] != nil {
		categoryVal, err = strconv.Atoi(params["category"][0])
		if err != nil {
			return nil, err
		}
		categoryVal, err = parseCategory(categoryVal)
		if err != nil {
			return nil, err
		}
	}
	if params["tier"] != nil {
		tierVal, err = strconv.Atoi(params["tier"][0])
		if err != nil {
			return nil, err
		}
		tierVal, err = parseTier(tierVal)
		if err != nil {
			return nil, err
		}
	}
	if params["sortBy"] != nil {
		sortByVal, err = parseSortBy(params["sortBy"][0])
		if err != nil {
			return nil, err
		}
	}
	if params["sortOrder"] != nil {
		sortOrderVal, err = parseSortOrder(params["sortOrder"][0])
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

	return &GetPostsOption{
		S:         sVal,
		Category:  categoryVal,
		Tier:      tierVal,
		SortBy:    sortByVal,
		SortOrder: sortOrderVal,
		Limit:     limitVal,
		Page:      pageVal,
	}, nil
}

func parseCategory(value int) (int, error) {
	if value < 0 {
		return 0, httperror.BadRequestError("Invalid category parameter", "INVALID_QUERY_PARAMETER")
	}
	return value, nil
}

func parseTier(value int) (int, error) {
	if value < 0 {
		return 0, httperror.BadRequestError("Invalid tier parameter", "INVALID_QUERY_PARAMETER")
	}
	return value, nil
}

func parseSortBy(value string) (string, error) {
	var validSortBy = map[string]string{}
	validSortBy["date"] = "created_at"
	if validSortBy[value] == "" {
		return "", httperror.BadRequestError("Invalid sortBy parameter", "INVALID_QUERY_PARAMETER")
	}
	return validSortBy[value], nil
}

func parseSortOrder(value string) (string, error) {
	var validSortOrder = map[string]string{}
	validSortOrder["asc"] = "asc"
	validSortOrder["desc"] = "desc"
	if validSortOrder[value] == "" {
		return "", httperror.BadRequestError("Invalid sortOrder parameter", "INVALID_QUERY_PARAMETER")
	}
	return validSortOrder[value], nil
}

func parseLimit(value int) (int, error) {
	if value < 1 {
		return 0, httperror.BadRequestError("Invalid limit parameter", "INVALID_QUERY_PARAMETER")
	}
	return value, nil
}

func parsePage(value int) (int, error) {
	if value < 1 {
		return 0, httperror.BadRequestError("Invalid page parameter", "INVALID_QUERY_PARAMETER")
	}
	return value, nil
}
