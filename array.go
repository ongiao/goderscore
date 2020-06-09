package goderscore

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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

func (g *GoSlice) Contains(item interface{}) (bool, error) {
	for _, n := range *g {
		if n == item {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

func (g *GoSlice) Difference(slice GoSlice) GoSlice {
	res := GoSlice{}

	for _, n := range *g {
		if _, ok := slice.IndexOf(n, 0); !ok{
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
		if _, ok := newOtherSlice.IndexOf(n, 0); !ok {
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

func (g *GoSlice) Flatten() GoSlice {
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

func (g *GoSlice) FlattenDeep() GoSlice {
	res := GoSlice{}

	for _, n := range *g {
		if reflect.TypeOf(n).Kind() == reflect.Slice {
			sub, _ := n.(GoSlice)
			res = append(res, sub.FlattenDeep()...)
		} else {
			res = append(res, n)
		}
	}

	return res
}

func (g *GoSlice) FlattenDepth(depth int) GoSlice {
	if depth == 0 {
		return *g
	}

	res := GoSlice{}

	for _, n := range *g {
		if reflect.TypeOf(n).Kind() == reflect.Slice {
			sub, _ := n.(GoSlice)
			depth -= 1
			res = append(res, sub.FlattenDepth(depth)...)
		} else {
			res = append(res, n)
		}
	}

	return res
}

//func (g *GoSlice) FromPairs() map[interface{}]interface{} {}

func (g *GoSlice) Head() interface{} {
	if len(*g) > 0 {
		return (*g)[0]
	}

	return nil
}

func (g *GoSlice) IndexOf(item interface{}, fromIndex int) (int, bool) {
	if fromIndex < 0 {
		for i := len(*g) - 1; i >=0; i-- {
			if (*g)[i] == item {
				return i, true
			}
		}
	} else if fromIndex >=0 && fromIndex < len(*g) {
		for i := fromIndex; i < len(*g); i++ {
			if (*g)[i] == item {
				return i, true
			}
		}
	}

	return -1, false
}

func (g *GoSlice) Initial() GoSlice {
	res := GoSlice{}

	if len(*g) > 0 {
		res = append(res, (*g)[:len(*g)-1]...)
	}

	return res
}

func (g *GoSlice) Intersection(arrays ...GoSlice) GoSlice {
	res := GoSlice{}
	recordMap := make(map[interface{}]int)
	correctCount := len(arrays) + 1

	for _, n := range *g {
		if recordMap[n] == 0 {
			recordMap[n] = 1
		} else {
			recordMap[n]++
		}
	}

	for _, item := range arrays {
		itemType := reflect.TypeOf(item)

		if itemType.Kind() == reflect.Slice {
			values := item.Uniq()

			for index := 0; index < len(values); index++ {
				val := values[index]
				if recordMap[val] == 0 {
					recordMap[val] = 1
				} else {
					recordMap[val]++
				}
			}
		}
	}

	for k, v := range recordMap {
		if v == correctCount {
			res = append(res, k)
		}
	}

	return res
}

//func (g *GoSlice) IntersectionBy() {}
//func (g *GoSlice) IntersectionWith() {}

func (g *GoSlice) Join(separator string) string {
	stringRes := []string{}

	for _, n := range *g {
		stringRes = append(stringRes, fmt.Sprintf("%v", n))
	}

	return strings.Join(stringRes, separator)
}

func (g *GoSlice) Last() interface{} {
	return (*g)[len(*g)-1]
}

func (g *GoSlice) LastIndexOf(item interface{}, fromIndex int) (int, bool) {
	if fromIndex < 0 {
		for i := 0; i < len(*g); i++ {
			if (*g)[i] == item {
				return i, true
			}
		}
	} else if fromIndex >=0 && fromIndex < len(*g) {
		for i := fromIndex; i >=0; i-- {
			if (*g)[i] == item {
				return i, true
			}
		}
	}

	return -1, false
}

func (g *GoSlice) Nth(n int) interface{} {
	if n >= len(*g) || n < -len(*g) {
		return nil
	} else if n >= 0 {
		return (*g)[n]
	} else {
		return (*g)[len(*g)+n]
	}
}

func (g *GoSlice) Pull(items ...interface{}) GoSlice {
	res := GoSlice{}
	deleteNums := GoSlice{}

	for _, n := range items {
		deleteNums = append(deleteNums, n)
	}

	for _, n := range *g {
		if ok, _ := deleteNums.Contains(n); !ok {
			res = append(res, n)
		}
	}

	return res
}

func (g *GoSlice) PullAll(values GoSlice) GoSlice {
	return g.Pull(values...)
}

//func (g *GoSlice) PullBy() {}
//func (g *GoSlice) PullWith() {}

func (g *GoSlice) PullAt(index ...int) GoSlice {
	res := GoSlice{}

	for _, i := range index {
		if i < 0 || i >= len(*g) {
			errors.New("Given index is out of slice's bound")
		} else {
			res = append(res, (*g)[i])
		}
	}

	removeIndexes := GoSlice{}
	for _, n := range index {
		removeIndexes = append(removeIndexes, n)
	}

	newG := GoSlice{}
	for i, n := range *g {
		if ok, _ := removeIndexes.Contains(i); !ok {
			newG = append(newG, n)
		}
	}

	*g = newG

	return res
}

//func (g *GoSlice) Remove() {}

func (g *GoSlice) Reverse() GoSlice {
	res := make(GoSlice, len(*g))
	copy(res, *g)

	i := 0
	j := len(*g) - 1

	for ; i < j; {
		tmp := res[i]
		res[i] = res[j]
		res[j] = tmp
		i++
		j--
	}

	return res
}

func (g *GoSlice) Slice(start, end int) GoSlice {
	if start > end || start < 0 || end >= len(*g) {
		return nil
	}

	return (*g)[start:end]
}

func (g *GoSlice) Tail() GoSlice {
	return (*g)[1:]
}

func (g *GoSlice) Take(n int) GoSlice {
	return (*g)[:n]
}

func (g *GoSlice) TakeRight(n int) GoSlice {
	return (*g)[len(*g)-n:]
}

//func (g *GoSlice) TakeRightWhile() {}

func Union(arrays ...GoSlice) GoSlice {
	res := GoSlice{}

	for _, arr := range arrays {
		res = append(res, arr...)
	}

	return res.Uniq()
}

//func UnionBy() GoSlice {}
//func UnionWith() GoSlice {}

func (g *GoSlice) Uniq() GoSlice {
	res := GoSlice{}

	for _, n := range *g {
		if index, _ := res.IndexOf(n, 0); index == -1 {
			res = append(res, n)
		}
	}

	return res
}