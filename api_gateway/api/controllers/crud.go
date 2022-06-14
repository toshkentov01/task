package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	errorhandler "github.com/toshkentov01/task/api_gateway/api/error_handler"
	"github.com/toshkentov01/task/api_gateway/api/models"
	"github.com/toshkentov01/task/api_gateway/config"

	crudPb "github.com/toshkentov01/task/api_gateway/genproto/crud_service"
	client "github.com/toshkentov01/task/api_gateway/grpc_client"
)

// ListPosts ...
// @Description
// @Summary Get list of posts
// @Tags post
// @Accept json
// @Produce json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} models.ListPostsModel
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/post/list/ [get]
func ListPosts(c *fiber.Ctx) error {
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		return errorhandler.HandleBadRequestErrWithMessage(c, err, "Failed to parse limit param")
	}

	page, err := ParsePageQueryParam(c)
	if err != nil {
		return errorhandler.HandleBadRequestErrWithMessage(c, err, "Failed to parse page param")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Config().CtxTimeout))
	defer cancel()

	result, err := client.CrudService().ListPosts(ctx, &crudPb.ListPostsRequest{
		Limit: uint32(limit),
		Page:  uint32(page),
	})

	if err != nil {
		return errorhandler.HandleGrpcErrWithMessage(c, err, "Failed to list posts")
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetPost ...
// @Description
// @Summary Get a single post
// @Tags post
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 200 {object} models.PostModel
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/post/get/{post_id}/ [get]
func GetPost(c *fiber.Ctx) error {
	postIDStr := c.Params("post_id")
	postID, err := strconv.Atoi(postIDStr)

	if err != nil {
		return errorhandler.HandleBadRequestErrWithMessage(c, err, "Invalid post id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Config().CtxTimeout))
	defer cancel()

	result, err := client.CrudService().GetPost(ctx, &crudPb.GetPostRequest{
		PostId: int64(postID),
	})

	if err != nil {
		return errorhandler.HandleGrpcErrWithMessage(c, err, "Failed to get post")
	}

	return c.Status(http.StatusOK).JSON(result)
}

// UpdatePost ...
// @Description UpdatePost API used for updating a post
// @Summary Update a post
// @Tags post
// @Accept json
// @Produce json
// @Param updatePostRequest body models.UpdatePostModel true "Update Post Reuqest Model"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/post/update/ [put]
func UpdatePost(c *fiber.Ctx) error {
	var (
		body models.UpdatePostModel
	)

	err := c.BodyParser(&body)
	if err != nil {
		return errorhandler.HandleBadRequestErrWithMessage(c, err, "Failed to parse body, check body again")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Config().CtxTimeout))
	defer cancel()

	_, err = client.CrudService().UpdatePost(ctx, &crudPb.UpdatePostRequest{
		PostId: int64(body.PostID),
		Title:  body.Title,
		Body:   body.Body,
	})

	if err != nil {
		return errorhandler.HandleGrpcErrWithMessage(c, err, "Failed to update a post")
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Error: false,
		Data: models.Success{
			Success: true,
		},
	})
}

// DeletePost ...
// @Description DeletePost API used for deleting a post
// @Summary DeletePost a post
// @Tags post
// @Accept json
// @Produce json
// @Param post_id path string true "Post ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /v1/post/delete/{post_id}/ [delete]
func DeletePost(c *fiber.Ctx) error {
	postIDstr := c.Params("post_id")
	postID, err := strconv.Atoi(postIDstr)

	if err != nil {
		return errorhandler.HandleBadRequestErrWithMessage(c, err, "Invalid post id")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Config().CtxTimeout))
	defer cancel()

	_, err = client.CrudService().DeletePost(ctx, &crudPb.DeletePostRequest{
		PostId: int64(postID),
	})

	if err != nil {
		return errorhandler.HandleGrpcErrWithMessage(c, err, "Failed to update a post")
	}

	return c.Status(http.StatusOK).JSON(models.Response{
		Error: false,
		Data: models.Success{
			Success: true,
		},
	})
}
