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
	CanUserAccessThisPost(user *models.User, post *models.Post) error
	GetPosts(opt *models.GetPostsOption) (*dto.GetPostsRes, error)
	GetPost(post *models.Post) (*models.Post, error)
	AddActivity(user *models.User, post *models.Post) (*models.UserPostActivities, error)
	UpdateActivity(user *models.User, post *models.Post, actReq *dto.ActivityReq) (*models.UserPostActivities, error)
	GetPostBySlug(slug string) (*models.Post, error)
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

func (serv *postService) CanUserAccessThisPost(user *models.User, post *models.Post) error {
	if post.PostTierID != 1 {
		unlock := &models.PostUnlock{
			UserID: user.ID,
			PostID: post.ID,
		}
		_, fetchErr := serv.postRepository.FindUnlock(unlock)
		return fetchErr
	}
	return nil
}

func (serv *postService) GetPosts(opt *models.GetPostsOption) (*dto.GetPostsRes, error) {
	postsRes, err := serv.postRepository.FindPosts(opt)
	if err != nil {
		return nil, err
	}
	return postsRes, nil
}

func (serv *postService) GetPost(post *models.Post) (*models.Post, error) {
	return serv.postRepository.FindPost(post)
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
	updatedAct, updateErr := serv.postRepository.IncreaseViewsActivity(fetchedAct)
	return updatedAct, updateErr
}

func (serv *postService) UpdateActivity(user *models.User, post *models.Post, actReq *dto.ActivityReq) (*models.UserPostActivities, error) {
	act := &models.UserPostActivities{
		UserID: user.ID,
		PostID: post.ID,
	}
	fetchedAct, fetchErr := serv.postRepository.FindActivity(act)
	if fetchErr != nil {
		fetchedAct, _ = serv.postRepository.SaveActivity(act)
	}
	fetchedAct.IsLiked = actReq.IsLiked
	fetchedAct.IsShared = actReq.IsShared

	updatedAct, updateErr := serv.postRepository.UpdateActivity(fetchedAct)
	return updatedAct, updateErr
}

func (serv *postService) GetPostBySlug(slug string) (*models.Post, error) {
	fetchedPost, err := serv.postRepository.FindPostBySlug(slug)
	if err != nil {
		return nil, err
	}

	//Check post tier
	//if fetchedPost.PostTierID != 1 {
	//
	//}

	return fetchedPost, nil
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
