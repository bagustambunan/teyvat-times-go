package services

import (
	"final-project-backend/dto"
	"final-project-backend/models"
	"final-project-backend/repositories"
	"log"
	"math/rand"
	"regexp"
	"strings"
)

type PostService interface {
	GetPosts(opt *models.GetPostsOption) (*dto.GetPostsRes, error)
	GetPost(post *models.Post) (*dto.GetPostRes, error)
	AddActivity(user *models.User, post *models.Post) (*models.UserPostActivities, error)
	GetPostBySlug(user *models.User, slug string) (*dto.GetPostRes, error)
	AddPost(post *models.Post) (*dto.GetPostRes, error)
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

func (serv *postService) generateSlug(text string) string {
	re, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	text = re.ReplaceAllString(text, " ")
	text = strings.Trim(text, " ")
	text = strings.ReplaceAll(text, " ", "-")
	text = strings.ToLower(text)
	return text
}

func (serv *postService) generatePrefix(size int) string {
	alpha := "wxyzghjklqrstabcdefmnpuv"
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

func (serv *postService) GetPosts(opt *models.GetPostsOption) (*dto.GetPostsRes, error) {
	postsRes, err := serv.postRepository.FindPosts(opt)
	if err != nil {
		return nil, err
	}
	return postsRes, nil
}

func (serv *postService) GetPost(post *models.Post) (*dto.GetPostRes, error) {
	fetchedPost, err := serv.postRepository.FindPost(post)
	if err != nil {
		return nil, err
	}
	return new(dto.GetPostRes).FromPost(fetchedPost), nil
}

func (serv *postService) AddActivity(user *models.User, post *models.Post) (*models.UserPostActivities, error) {
	act := &models.UserPostActivities{
		UserID: user.ID,
		PostID: post.ID,
	}
	fetchedAct, fetchErr := serv.postRepository.FindActivity(act)
	if fetchErr != nil {
		fetchedAct, _ = serv.postRepository.SaveActivity(act)
	}
	updatedAct, updateErr := serv.postRepository.UpdateActivity(fetchedAct)
	return updatedAct, updateErr
}

func (serv *postService) GetPostBySlug(user *models.User, slug string) (*dto.GetPostRes, error) {
	fetchedPost, err := serv.postRepository.FindPostBySlug(slug)
	if err != nil {
		return nil, err
	}

	//Check post tier
	//if fetchedPost.PostTierID != 1 {
	//
	//}

	_, err2 := serv.AddActivity(user, fetchedPost)
	return new(dto.GetPostRes).FromPost(fetchedPost), err2
}

func (serv *postService) AddPost(post *models.Post) (*dto.GetPostRes, error) {
	newSlug := serv.generateSlug(post.Title)
	for {
		if _, err := serv.postRepository.FindPostBySlug(newSlug); err != nil {
			post.Slug = newSlug
			break
		}
		newSlug += "-" + serv.generatePrefix(3)
	}

	insertedPost, _, err := serv.postRepository.Save(post)
	if err != nil {
		return nil, err
	}
	return new(dto.GetPostRes).FromPost(insertedPost), nil
}
