package ecs

//SimpleContainer is very simple container using slice.
//time complexity of add is O(1) when len < capacity
//time complexity of remove using index is O(1)
type SimpleContainer struct {
	items    []interface{}
	cnt      int
	capacity int
}

//NewSimpleContainer make instance of SimpleContainer
func NewSimpleContainer(capa int) *SimpleContainer {

	return &SimpleContainer{
		items:    make([]interface{}, 0, capa),
		cnt:      0,
		capacity: 0,
	}

}

//Add add item
func (s *SimpleContainer) Add(i interface{}) int {

	rIdx := s.cnt
	s.cnt++
	if rIdx == s.capacity {
		s.items = append(s.items, i)
		s.capacity++

		return rIdx
	}
	s.items[rIdx] = i

	return rIdx
}

//RemoveIdx remove item by index
func (s *SimpleContainer) RemoveIdx(idx int) {

	s.cnt--
	s.items[idx] = s.items[s.cnt]
	s.items[s.cnt] = nil
}

//Each iterate items with index and item
func (s *SimpleContainer) Each(f func(int, interface{})) {

	for i := s.cnt - 1; i >= 0; i-- {
		f(i, s.items[i])
	}

}
