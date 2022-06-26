package models

type News struct {
	ID      int
	Enable  bool
	Image   string
	Tags    []string
	Title   string
	Content string
}
