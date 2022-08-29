package services

import (
	"final-project-backend/dto"
	"final-project-backend/httperror"
	"final-project-backend/models"
	"final-project-backend/repositories"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type PostService interface {
	CanUserAccessThisPost(user *models.User, post *models.Post) error
	UnlockAPost(user *models.User, post *models.Post) (*models.PostUnlock, error)
	GetPosts(opt *models.GetPostsOption) (*dto.GetPostsRes, error)
	GetPost(post *models.Post) (*models.Post, error)
	GetReadingHistory(user *models.User, opt *models.ReadHistoryOption) (*dto.GetPostsRes, error)
	GetActivity(user *models.User, post *models.Post) (*models.UserPostActivities, error)
	AddActivity(user *models.User, post *models.Post) (*models.UserPostActivities, error)
	UpdateActivity(user *models.User, post *models.Post, actReq *dto.ActivityReq) (*models.UserPostActivities, error)
	GetPostBySlug(slug string) (*models.Post, error)
	AddPost(post *models.Post) (*dto.GetPostRes, error)
	UpdatePost(post *models.Post) (*dto.GetPostRes, error)
	DeletePost(post *models.Post) error
	GetTiers() ([]*models.PostTier, error)
	GetCategories() ([]*models.PostCategory, error)
	GetTier(tier *models.PostTier) (*models.PostTier, error)
	GetCategory(category *models.PostCategory) (*models.PostCategory, error)
	AddCategory(req *dto.CategoryReq) (*models.PostCategory, error)
	UpdateCategory(category *models.PostCategory) (*models.PostCategory, error)
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
	//if user.RoleID == 1 {
	//	return nil
	//}
	if post.PostTierID != 1 {
		latestUs, usErr := serv.postRepository.FindUserLatestSubscription(user)
		if usErr != nil {
			return httperror.BadRequestError("You have no active subscription", "NO_ACTIVE_SUBSCRIPTION")
		}
		latestUsEnded, _ := time.Parse("2006-01-02T00:00:00Z", latestUs.DateEnded)
		if latestUsEnded.Before(time.Now()) {
			return httperror.BadRequestError("You have no active subscription", "NO_ACTIVE_SUBSCRIPTION")
		}
		unlock := &models.PostUnlock{
			UserID: user.ID,
			PostID: post.ID,
		}
		_, fetchErr := serv.postRepository.FindUnlock(unlock)
		if fetchErr != nil {
			return httperror.BadRequestError("Post is locked", "UNLOCKED_POST")
		}
	}
	return nil
}

func (serv *postService) DoesUserHaveAnActiveSubscription(user *models.User) error {
	latestUs, usErr := serv.postRepository.FindUserLatestSubscription(user)
	if usErr != nil {
		return httperror.BadRequestError("You have no active subscription", "NO_ACTIVE_SUBSCRIPTION")
	}
	latestUsEnded, _ := time.Parse("2006-01-02T00:00:00Z", latestUs.DateEnded)
	if latestUsEnded.Before(time.Now()) {
		return httperror.BadRequestError("You have no active subscription", "NO_ACTIVE_SUBSCRIPTION")
	}
	return nil
}

func (serv *postService) UnlockAPost(user *models.User, post *models.Post) (*models.PostUnlock, error) {
	subErr := serv.DoesUserHaveAnActiveSubscription(user)
	if subErr != nil {
		return nil, subErr
	}

	unlock := &models.PostUnlock{
		UserID: user.ID,
		PostID: post.ID,
	}
	_, fetchErr := serv.postRepository.FindUnlock(unlock)
	if fetchErr == nil {
		return nil, httperror.BadRequestError("Post already unlocked", "POST_ALREADY_UNLOCKED")
	}
	if post.PostTierID == 1 {
		return nil, httperror.BadRequestError("Cannot unlock free tier post", "INVALID_UNLOCK")
	}

	if user.Mora < post.GetMoraRequired() {
		return nil, httperror.BadRequestError("Not enough mora", "COINS_NOT_ENOUGH")
	}
	return serv.postRepository.SaveUnlock(unlock)
}

func (serv *postService) GetPosts(opt *models.GetPostsOption) (*dto.GetPostsRes, error) {
	postsRes, err := serv.postRepository.FindPosts(opt)
	if err != nil {
		return nil, err
	}
	return postsRes, nil
}

func (serv *postService) GetReadingHistory(user *models.User, opt *models.ReadHistoryOption) (*dto.GetPostsRes, error) {
	postsRes, err := serv.postRepository.FindReadingHistory(user, opt)
	if err != nil {
		return nil, err
	}
	return postsRes, nil
}

func (serv *postService) GetPost(post *models.Post) (*models.Post, error) {
	return serv.postRepository.FindPost(post)
}

func (serv *postService) GetActivity(user *models.User, post *models.Post) (*models.UserPostActivities, error) {
	act := &models.UserPostActivities{
		UserID: user.ID,
		PostID: post.ID,
	}
	return serv.postRepository.FindActivity(act)
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

func (serv *postService) UpdatePost(post *models.Post) (*dto.GetPostRes, error) {
	updatedPost, err := serv.postRepository.UpdatePost(post)
	if err != nil {
		return nil, err
	}
	return new(dto.GetPostRes).FromPost(updatedPost), nil
}

func (serv *postService) DeletePost(post *models.Post) error {
	return serv.postRepository.DeletePost(post)
}

func (serv *postService) GetTiers() ([]*models.PostTier, error) {
	return serv.postRepository.FindTiers()
}

func (serv *postService) GetCategories() ([]*models.PostCategory, error) {
	return serv.postRepository.FindCategories()
}

func (serv *postService) GetTier(tier *models.PostTier) (*models.PostTier, error) {
	return serv.postRepository.FindTier(tier)
}

func (serv *postService) GetCategory(category *models.PostCategory) (*models.PostCategory, error) {
	return serv.postRepository.FindCategory(category)
}

func (serv *postService) AddCategory(req *dto.CategoryReq) (*models.PostCategory, error) {
	category, err := serv.postRepository.SaveCategory(&models.PostCategory{
		Name:  req.Name,
		Color: req.Color,
	})
	return category, err
}

func (serv *postService) UpdateCategory(category *models.PostCategory) (*models.PostCategory, error) {
	cat, err := serv.postRepository.UpdateCategory(category)
	return cat, err
}
