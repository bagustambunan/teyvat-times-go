package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetPost(c *gin.Context) {
	postID, idErr := strconv.Atoi(c.Param("postID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}

	fetchedPost, fetchErr := h.postService.GetPost(&models.Post{ID: postID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	postRes := new(dto.GetPostRes).FromPost(fetchedPost)
	helpers.StandardResponse(c, http.StatusOK, postRes)
}

func (h *Handler) AddPost(c *gin.Context) {
	user := h.GetUserFromToken(c)
	payload, _ := c.Get("payload")
	postReq, _ := payload.(*dto.PostReq)

	postRes, err := h.postService.AddPost(&models.Post{
		PostTierID:     postReq.PostTierID,
		PostCategoryID: postReq.PostCategoryID,
		Title:          postReq.Title,
		Content:        postReq.Content,
		Summary:        postReq.Summary,
		ImgThumbnailID: 2,
		ImgContentID:   3,
		CreatedByID:    user.ID,
		UpdatedByID:    user.ID,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusCreated, postRes)
}

func (h *Handler) UpdatePost(c *gin.Context) {
	user := h.GetUserFromToken(c)
	payload, _ := c.Get("payload")
	postReq, _ := payload.(*dto.PostReq)

	postID, idErr := strconv.Atoi(c.Param("postID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}

	postRes, err := h.postService.UpdatePost(&models.Post{
		ID:             postID,
		PostTierID:     postReq.PostTierID,
		PostCategoryID: postReq.PostCategoryID,
		Title:          postReq.Title,
		Content:        postReq.Content,
		Summary:        postReq.Summary,
		UpdatedByID:    user.ID,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, postRes)
}

func (h *Handler) DeletePost(c *gin.Context) {
	postID, idErr := strconv.Atoi(c.Param("postID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	err := h.postService.DeletePost(&models.Post{ID: postID})
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, "")
}

func (h *Handler) GetTiers(c *gin.Context) {
	tiers, fetchErr := h.postService.GetTiers()
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, tiers)
}

func (h *Handler) GetCategories(c *gin.Context) {
	categories, fetchErr := h.postService.GetCategories()
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, categories)
}

func (h *Handler) GetTier(c *gin.Context) {
	postTierID, idErr := strconv.Atoi(c.Param("postTierID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	tier, fetchErr := h.postService.GetTier(&models.PostTier{ID: postTierID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, tier)
}

func (h *Handler) GetCategory(c *gin.Context) {
	postCategoryID, idErr := strconv.Atoi(c.Param("postCategoryID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	category, fetchErr := h.postService.GetCategory(&models.PostCategory{ID: postCategoryID})
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, category)
}

func (h *Handler) AddCategory(c *gin.Context) {
	payload, _ := c.Get("payload")
	req, _ := payload.(*dto.CategoryReq)
	category, err := h.postService.AddCategory(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusCreated, category)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	postCategoryID, idErr := strconv.Atoi(c.Param("postCategoryID"))
	if idErr != nil {
		_ = c.Error(idErr)
		return
	}
	payload, _ := c.Get("payload")
	req, _ := payload.(*dto.CategoryReq)
	category, err := h.postService.UpdateCategory(&models.PostCategory{
		ID:    postCategoryID,
		Name:  req.Name,
		Color: req.Color,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, category)
}
