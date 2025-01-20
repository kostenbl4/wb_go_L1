package main

// Set - множество
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet - конструктор множества
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

// Add - добавление элемента в множество
func (s *Set[T]) Add(v T) {
	s.data[v] = struct{}{}
}

// Delete - удаление элемента из множества
func (s *Set[T]) Delete(v T) {
	delete(s.data, v)
}

// Contains - проверка наличия элемента в множестве
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.data[v]
	return ok
}

// Len - количество элементов в множестве
func (s *Set[T]) Len() int {
	return len(s.data)
}

// Intersection - пересечение множеств
func (s *Set[T]) Intersection(other Set[T]) Set[T] {
	res := NewSet[T]()
	for v := range s.data {
		if other.Contains(v) {
			res.Add(v)
		}
	}
	return *res
}

func main() {
	// Создание множеств
	set1 := NewSet[int]()
	set2 := NewSet[int]()
	// Добавление элементов в множества
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	// Пересечение множеств
	intersection := set1.Intersection(*set2)
	// Вывод результата
	for v := range intersection.data {
		println(v)
	}
}
