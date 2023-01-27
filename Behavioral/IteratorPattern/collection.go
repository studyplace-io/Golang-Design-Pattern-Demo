package IteratorPattern

type Collection interface {
	createIterator() Iterator
}
