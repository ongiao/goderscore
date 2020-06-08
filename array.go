package goderscore

import (
	"reflect"
)

func (g *GoSlice) Chunk(size int) []GoSlice {
	var res []GoSlice

	i := 0
	sliceSize := len((*g))

	for ;; i += size{
		if i + size > sliceSize - 1 {
			break
		}

		sliceTmp := (*g)[i:i+size]
		res = append(res, sliceTmp)
	}

	if i > sliceSize - 1 {
		i -= size
	}

	res = append(res, (*g)[i:])

	return res
}

func (g *GoSlice) Compact() GoSlice {
	res := GoSlice{}

	for _, item := range *g {
		if item != nil && item != false && item != "" {
			res = append(res, item)
		}
	}

	return res
}

func (g *GoSlice) Concat(args ...interface{}) GoSlice {
	for _, item := range args {
		itemType := reflect.TypeOf(item)

		if itemType.Kind() == reflect.Slice {
			values := reflect.ValueOf(item)

			for index := 0; index < values.Len(); index++ {
				*g = append(*g, values.Index(index).Interface())
			}
		} else {
			*g = append(*g, item)
		}
	}

	return *g
}

func (g *GoSlice) Difference(slice GoSlice) GoSlice {
	res := GoSlice{}

	for _, n := range *g {
		if _, ok := slice.IndexOf(n); !ok{
			res = append(res, n)
		}
	}

	return res
}

// 缺少类型检查
func (g *GoSlice) DifferenceBy(otherSlice GoSlice, iteratee Iteratee) GoSlice {
	newSlice := (*g).Map(iteratee)
	newOtherSlice := otherSlice.Map(iteratee)

	res := GoSlice{}

	for i, n := range newSlice {
		if _, ok := newOtherSlice.IndexOf(n); !ok {
			res = append(res, (*g)[i])
		}
	}

	return res
}

//func (g *GoSlice) DifferenceWith(otherSlice GoSlice, comparator Comparator) GoSlice {
//
//}

func (g *GoSlice) Drop(dropSize int) (GoSlice, bool) {
	res := GoSlice{}

	if dropSize > len(*g) || dropSize < 0 {
		return nil, false
	}

	res = (*g)[dropSize:]

	return res, true
}

func (g *GoSlice) DropRight(dropSize int) (GoSlice, bool) {
	res := GoSlice{}

	if dropSize > len(*g) || dropSize < 0 {
		return nil, false
	}

	res = (*g)[:len(*g)-dropSize]

	return res, true
}

//func (g *GoSlice) DropRightWhile(predicate Predicate) (GoSlice, bool) {
//	res := GoSlice{}
//
//	for i, n := range *g {
//		if predicate(n)
//	}
//
//	return res
//}

func (g *GoSlice) Fill(ch interface{}, initSize, start, end interface{}) GoSlice {
	sliceSize := len(*g)
	res := GoSlice{}


	if init, ok := initSize.(int); ok && sliceSize == 0 && init != 0 && start == nil && end == nil {
		for i:= 0; i < init; i++ {
			res = append(res, ch)
		}
		return res
	} else if sliceSize != 0 && initSize == nil && start == nil && end == nil {
		for i:= 0; i < sliceSize; i++ {
			res = append(res, ch)
		}
		return res
	} else if sliceSize != 0 && initSize == nil && start != nil && end != nil {
		s, startOk := start.(int)
		e, endOk := end.(int)

		if startOk && endOk && s < e && s >= 0 && e <= len(*g) {
			res := GoSlice{}


			for i, _ := range *g {
				if i < s || i >= e {
					res = append(res, (*g)[i])
				} else {
					res = append(res, ch)
				}
			}

			return res
		}
	}

	return nil
}

//func (g *GoSlice) FindIndex()


func (g *GoSlice) Head() interface{} {
	if len(*g) > 0 {
		return (*g)[0]
	}

	return nil
}

func (g *GoSlice) Flatten() GoSlice {
	if len(*g) > 0 {
		res := GoSlice{}

		for _, n := range *g {
			if reflect.TypeOf(n).Kind() == reflect.Slice {
				sub, _ := n.(GoSlice)
				for _, subN := range sub {
					res = append(res, subN)
				}
			} else {
				res = append(res, n)
			}
		}

		return res
	}

	return nil
}

//needtodo
//func (g *GoSlice) FlattenDeep() GoSlice {
//	// use recursive
//	return nil
//}

func (g *GoSlice) IndexOf(item interface{}) (int, bool) {
	for i, n := range *g {
		if n == item {
			return i, true
		}
	}

	return -1, false
}