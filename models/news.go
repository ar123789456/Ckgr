package models

type News struct {
	ID      string
	Enable  bool
	Image   string
	Tags    []string
	Title   string
	Content string
}
