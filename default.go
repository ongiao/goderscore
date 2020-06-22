package goderscore

import "fmt"

type T interface {}

type GoSlice []T

type Iteratee func(interface{}) interface{}

type Comparator func(T, T) bool

//type Predicate func(map[string]bool) bool
//
//type Predicate struct {
//	predicate
//}

func (g *GoSlice) Map(iteratee Iteratee) GoSlice {
	res := GoSlice{}

	for _, n := range *g {
		res = append(res, iteratee(n))
	}

	return res
}

func CatchPanic() {
	if r := recover(); r != nil {
		fmt.Printf("Caught error：%s\n", r)
	}
}

