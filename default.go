package goderscore

type GoSlice []interface{}

type Iteratee func(interface{}) interface{}

type Comparator func(interface{}) bool

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

