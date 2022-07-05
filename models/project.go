package models

type Project struct {
	ID     int
	Base   Frame
	Enable bool
	All    []Frame
}

type Frame struct {
	ID      int
	Title   map[Language]string
	Image   string
	Content map[Language]string
}
