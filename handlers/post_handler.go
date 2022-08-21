package handlers

import (
	"final-project-backend/dto"
	"final-project-backend/helpers"
	"final-project-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetPosts(c *gin.Context) {
	opt, parsingErr := models.NewGetPostsOption(c.Request.URL.Query())
	if parsingErr != nil {
		_ = c.Error(parsingErr)
		return
	}

	postsRes, fetchErr := h.postService.GetPosts(opt)
	if fetchErr != nil {
		_ = c.Error(fetchErr)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, postsRes)
}

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
		CreatedByID:    postReq.CreatedByID,
		UpdatedById:    postReq.UpdatedByID,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	helpers.StandardResponse(c, http.StatusOK, postRes)
}
