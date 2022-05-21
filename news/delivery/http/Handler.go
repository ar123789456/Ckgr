package http

import "cgr/news"

type Handler struct {
	usecase news.UseCase
}
