package dto

import "git.garena.com/sea-labs-id/batch-01/bagus-tambunan/final-project-backend/models"

type GetPostsRes struct {
	Count     int           `json:"count"`
	Limit     int           `json:"limit"`
	Page      int           `json:"page"`
	TotalPage int           `json:"totalPage"`
	TotalData int           `json:"totalData"`
	Posts     []*GetPostRes `json:"posts"`
}

func (_ *GetPostsRes) FromPosts(posts []*models.Post) *GetPostsRes {
	var postsRes = make([]*GetPostRes, 0)
	for _, p := range posts {
		postsRes = append(postsRes, new(GetPostRes).FromPost(p))
	}
	return &GetPostsRes{
		Posts: postsRes,
	}
}

func (res *GetPostsRes) SetValues(count int, limit int, page int, totalPage int, totalData int) {
	res.Count = count
	res.Limit = limit
	res.Page = page
	res.TotalPage = totalPage
	res.TotalData = totalData
}
