package models

type Project struct {
	ID     int
	Base   Frame
	Enable bool
	All    []Frame
}

type Frame struct {
	ID      int
	Title   string
	Image   string
	Content string
}
