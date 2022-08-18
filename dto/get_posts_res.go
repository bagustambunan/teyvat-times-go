package dto

import "final-project-backend/models"

type GetPostsRes struct {
	Posts []*GetPostRes `json:"posts"`
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
