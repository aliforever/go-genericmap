package genericmap

import "sync"

type GenericMap[T any] struct {
	m            sync.Mutex
	mPointer     sync.Mutex
	data         map[string]T
	dataPointers map[string]*T
}

func New[T any]() *GenericMap[T] {
	return &GenericMap[T]{
		data: map[string]T{},
	}
}

func (g *GenericMap[T]) Set(key string, value T) {
	g.m.Lock()
	defer g.m.Unlock()

	if g.data == nil {
		g.data = map[string]T{}
	}

	g.data[key] = value
}

func (g *GenericMap[T]) Get(key string) (value T, exists bool) {
	g.m.Lock()
	defer g.m.Unlock()

	if g.data == nil {
		g.data = map[string]T{}
	}

	value, exists = g.data[key]
	return
}

func (g *GenericMap[T]) GetValue(key string) (value T) {
	g.m.Lock()
	defer g.m.Unlock()

	if g.data == nil {
		g.data = map[string]T{}
	}

	value = g.data[key]

	return
}

func (g *GenericMap[T]) SetPointer(key string, value *T) {
	g.mPointer.Lock()
	defer g.mPointer.Unlock()

	if g.dataPointers == nil {
		g.dataPointers = map[string]*T{}
	}

	g.dataPointers[key] = value
}

func (g *GenericMap[T]) GetPointer(key string) (value *T, exists bool) {
	g.mPointer.Lock()
	defer g.mPointer.Unlock()

	if g.dataPointers == nil {
		g.dataPointers = map[string]*T{}
	}

	value, exists = g.dataPointers[key]
	return
}

func (g *GenericMap[T]) GetPointerValue(key string) (value *T) {
	g.mPointer.Lock()
	defer g.mPointer.Unlock()

	if g.dataPointers == nil {
		g.dataPointers = map[string]*T{}
	}

	value = g.dataPointers[key]
	return
}
