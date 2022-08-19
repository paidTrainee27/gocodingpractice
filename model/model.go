package model

type List struct {
    root Element
    len  int
}

type Element struct {
	data int
	next *Element
}

func (l *List) AddFront(len int)  {
	l.len = len
}