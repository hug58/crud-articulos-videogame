package handlers

import (
	"go-postgres/internal/data"
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()
	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	art := &ArticleRouter{
		Repository: &data.ArticlesRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())
	r.Mount("/articles-video-game", art.Routes())

	return r
}
