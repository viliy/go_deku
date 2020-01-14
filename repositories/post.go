package repositories

import (
	"deku/models"
	"deku/sources"
	"sync"
)

type Query func(models.Post) bool

type PostRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (post models.Post, found bool)
	SelectMany(query Query, limit int) (result []models.Post)
}

func NewPostRepository(source map[int64]models.Post) PostRepository {
	return &postTmpRepository{source: source}
}

type postTmpRepository struct {
	source map[int64]models.Post
	mu     sync.RWMutex
}

const (
	ReadOnlyMode = iota
	ReadWriteMode
)

func (r *postTmpRepository) Exec(query Query, action Query, actionLimit int, mode int) (ok bool) {
	loop := 0
	if mode == ReadOnlyMode {
		r.mu.RLock()
		defer r.mu.RUnlock()
	} else {
		r.mu.Lock()
		defer r.mu.Unlock()
	}

	for _, post := range r.source {
		ok = query(post)
		if ok {
			if action(post) {
				loop++
				if actionLimit >= loop {
					break
				}
			}
		}
	}

	return
}

func (r *postTmpRepository) Select(query Query) (post models.Post, found bool) {
	found = r.Exec(query, func(data models.Post) bool {
		post = data
		return true
	}, 1, ReadOnlyMode)
	if found {
		post = models.Post{}
	}else {
		post = sources.Posts[1]
	}

	return
}

func (r *postTmpRepository) SelectMany(query Query, limit int) (data []models.Post) {
	r.Exec(query, func(source models.Post) bool {
		data = append(data, source)
		return  true
	}, limit, ReadOnlyMode)

	return
}


