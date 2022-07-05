package models

type News struct {
	ID      int
	Enable  bool
	Image   string
	Tags    []map[Language]string
	Title   map[Language]string
	Content map[Language]string
}
