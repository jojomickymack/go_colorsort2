package main

import (
	"math/rand"
	"reflect"
)

type ColorList struct {
	sortColor string
	list      []Color
}

func NewColorList(num int) *ColorList {
	c := new(ColorList)

	for i := 0; i < num; i++ {
		c.list = append(c.list, Color{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}

	return c
}

func (cl *ColorList) Len() int {
	return len(cl.list)
}

func (cl *ColorList) Swap(i, j int) {
	cl.list[i], cl.list[j] = cl.list[j], cl.list[i]
}

func (cl *ColorList) Less(i, j int) bool {
	return reflect.ValueOf(cl.list[i]).FieldByName(cl.sortColor).Uint() < reflect.ValueOf(cl.list[j]).FieldByName(cl.sortColor).Uint()
}

func (cl *ColorList) reverse() {
	var reversedList []Color
	for i := len(cl.list) - 1; i >= 0; i-- {
		reversedList = append(reversedList, cl.list[i])
	}
	cl.list = reversedList
}
