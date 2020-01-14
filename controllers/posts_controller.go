package controllers

import (
	"deku/models"
	"deku/services"
	"deku/sources"
)

type PostsController struct {
	service services.PostService
}

func (c *PostsController) Get() (data []models.Post) {
	return c.service.GetAll()
}

func (c *PostsController) GetBy(id int64) (data models.Post, found bool) {
	res, found := c.service.GetByID(id)
	if found {
		return res, found
	}else {
		return sources.Posts[1], false
	}
}
