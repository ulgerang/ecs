package ecs

import "testing"

func TestSimpleContainer(t *testing.T) {

	sc := NewSimpleContainer(4)

	sc.Add(1)
	sc.Add(2)
	sc.Add(3)

	sum := 0

	sc.Each(func(idx int, v interface{}) {
		sum += v.(int)
	})

	if sum != 6 {
		t.Fatal("wrong sum1")
	}

	sc.RemoveIdx(1)

	sum = 0
	sc.Each(func(idx int, v interface{}) {
		sum += v.(int)
	})

	if sum != 4 {
		t.Fatal("wrong sum2")
	}

	sc.Add(5)

	sum = 0
	sc.Each(func(idx int, v interface{}) {
		sum += v.(int)
	})

	if sum != 9 {
		t.Fatal("wrong sum3")
	}

	if len(sc.items) != 3 || sc.cnt != 3 {
		t.Fatal("wrong cnt")
	}

	sc.Add(6)

	sum = 0
	sc.Each(func(idx int, v interface{}) {
		sum += v.(int)
	})

	if sum != 15 {
		t.Fatal("wrong sum4")
	}

	if len(sc.items) != 4 || sc.cnt != 4 {
		t.Fatal("wrong cnt5")
	}

}
