package controllers

import (
	"deku/models"
	"deku/services"
	"deku/sources"
	"fmt"
)

type PostsController struct {
	Service services.PostService
}

func (c *PostsController) Get() (res []models.Post) {
	return c.Service.GetAll()
}

func (c *PostsController) GetBy(id int64) (data models.Post, found bool) {
	fmt.Printf("%v\n", sources.Posts[1])
	return c.Service.GetByID(id)
}
