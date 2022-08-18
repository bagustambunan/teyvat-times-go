package services

import (
	"final-project-backend/dto"
	"final-project-backend/repositories"
)

type PostService interface {
	GetPosts() (*dto.GetPostsRes, error)
	GetPostBySlug(slug string) (*dto.GetPostRes, error)
}

type postService struct {
	postRepository repositories.PostRepository
}

type PSConfig struct {
	PostRepository repositories.PostRepository
}

func NewPostService(conf *PSConfig) PostService {
	return &postService{
		postRepository: conf.PostRepository,
	}
}

func (serv *postService) GetPosts() (*dto.GetPostsRes, error) {
	posts, err := serv.postRepository.FindPosts()
	if err != nil {
		return nil, err
	}
	return new(dto.GetPostsRes).FromPosts(posts), nil
}

func (serv *postService) GetPostBySlug(slug string) (*dto.GetPostRes, error) {
	fetchedPost, err := serv.postRepository.FindPostBySlug(slug)
	if err != nil {
		return nil, err
	}
	return new(dto.GetPostRes).FromPost(fetchedPost), nil
}
