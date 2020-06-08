package goderscore

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"
)

func TestChunk(t *testing.T) {
	t.Run("Should return {{1,2,3}, {4,5,6}, {7,8,9}}", func(t *testing.T) {
		sliceSample := GoSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}

		got := sliceSample.Chunk(3)
		want := []GoSlice {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

		assertDeepEqual(t, got, want, "Chunk array failed")
	})

	t.Run("Should return {{1,2,3}, {4,5,6}, {7,8}}", func(t *testing.T) {
		sliceSample := GoSlice{1, 2, 3, 4, 5, 6, 7, 8}

		got := sliceSample.Chunk(3)
		want := []GoSlice {{1, 2, 3}, {4, 5, 6}, {7, 8}}

		assertDeepEqual(t, got, want, "Chunk array failed")
	})

	t.Run("Should return {{'a', 'b'}, {'c', 'd'}}", func(t *testing.T) {
		sliceSample := GoSlice{'a', 'b', 'c', 'd'}

		got := sliceSample.Chunk(2)
		want := []GoSlice {{'a', 'b'}, {'c', 'd'}}

		assertDeepEqual(t, got, want, "Chunk array failed")
	})
}

func TestCompact(t *testing.T) {
	t.Run("Should return [0, 1, 2, 3, 4]", func(t *testing.T) {
		sliceSample := GoSlice{0, 1, false, 2, "", 3, nil, 4}

		got := sliceSample.Compact()
		want := GoSlice{0, 1, 2, 3, 4}

		assertDeepEqual(t, got, want, "Compact array failed")
	})
}

func TestConcat(t *testing.T) {
	t.Run("Test concat slice", func(t *testing.T) {
		slice := GoSlice{2,3,4}
		subSlice := GoSlice{5,6,7}

		got := slice.Concat(subSlice)
		want := GoSlice{2,3,4,5,6,7}

		assertDeepEqual(t, got, want, "Concat failed")
	})

	t.Run("Test concat numbers", func(t *testing.T) {
		slice := GoSlice{2,3,4}

		got := slice.Concat(5,6,7)
		want := GoSlice{2,3,4,5,6,7}

		assertDeepEqual(t, got, want, "Concat failed")
	})

	t.Run("Test concat numbers and slice", func(t *testing.T) {
		slice := GoSlice{2,3,4}

		got := slice.Concat(5,6,7, GoSlice{8,9,10})
		want := GoSlice{2,3,4,5,6,7, 8, 9, 10}

		assertDeepEqual(t, got, want, "Concat failed")
	})
}

func TestDifference(t *testing.T) {
	t.Run("Test difference", func(t *testing.T) {
		slice := GoSlice{3,2,1}
		otherSlice := GoSlice{4,2}

		got := slice.Difference(otherSlice)
		want := GoSlice{3,1}

		assertDeepEqual(t, got, want, "Difference failed")
	})
}

func TestDifferenceBy(t *testing.T) {
	t.Run("Test difference by string", func(t *testing.T) {
		slice := GoSlice{"junxing", "jack", "marnie"}
		otherSlice := GoSlice{"jack", "marnie"}


		got := slice.DifferenceBy(otherSlice, func(i interface{}) interface{} {
			return strings.ToUpper(i.(string))
		})
		want := GoSlice{"junxing"}

		assertDeepEqual(t, got, want, "Difference by failed")
	})

	t.Run("Test difference by int", func(t *testing.T) {
		slice := GoSlice{3.1, 2.2, 1.3}
		otherSlice := GoSlice{4.4, 2.5}


		got := slice.DifferenceBy(otherSlice, func(i interface{}) interface{} {
			return math.Floor(i.(float64))
		})
		want := GoSlice{3.1, 1.3}

		assertDeepEqual(t, got, want, "Difference by failed")
	})

	//t.Run("Test difference by slicemap", func(t *testing.T) {
	//	slice := GoSlice{map[string]int{"x": 2}, map[string]int{"x": 1}}
	//	otherSlice := GoSlice{map[string]int{"x": 1}}
	//
	//
	//	got := slice.DifferenceBy(otherSlice, func(i interface{}) interface{} {
	//		return math.Floor(i.(float64))
	//	})
	//	want := GoSlice{map[string]int{"x": 2}}
	//
	//	assertDeepEqual(t, got, want, "Difference by failed")
	//})
}

func TestDrop(t *testing.T) {
	t.Run("Test drop", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got1, _ := slice.Drop(2)
		want1 := GoSlice{3,4,5}

		assertDeepEqual(t, got1, want1, "Drop failed")

		got2, _ := slice.Drop(0)
		want2 := GoSlice{1,2,3,4,5}

		assertDeepEqual(t, got2, want2, "Drop failed")

		got3, _ := slice.Drop(len(slice))
		want3 := GoSlice{}

		assertDeepEqual(t, got3, want3, "Drop failed")
	})

	t.Run("Test drop size error", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got1, _ := slice.Drop(6)
		var want1 GoSlice = nil

		assertDeepEqual(t, got1, want1, "Drop failed")

		got2, _ := slice.Drop(-2)
		var want2 GoSlice = nil

		assertDeepEqual(t, got2, want2, "Drop failed")
	})
}

func TestDropRIght(t *testing.T) {
	t.Run("Test drop right", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got, _ := slice.DropRight(2)
		want := GoSlice{1,2,3}

		assertDeepEqual(t, got, want, "Drop failed")
	})
}

//func TestDropRightWhile(t *testing.T) {
//	t.Run("Test drop right while", func(t *testing.T) {
//		slice := GoSlice{
//			map[string]interface{}{"user": "barney",  "active": true},
//			map[string]interface{}{"user": "fred",  "active": false},
//			map[string]interface{}{"user": "pebbles",  "active": false},
//		}
//
//		got, ok := slice.DropRightWhile(func(i map[string]bool) bool {
//			return !i["active"]
//		})
//	})
//}

func TestFill(t *testing.T) {
	t.Run("Test fill 1", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got := slice.Fill("a", nil,nil, nil)
		want := GoSlice{"a","a","a","a","a"}

		assertDeepEqual(t, got, want, "Fill failed")
	})

	t.Run("Test fill 2", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got := slice.Fill("a", nil, 1, 4)
		want := GoSlice{1,"a","a","a", 5}

		assertDeepEqual(t, got, want, "Fill failed")
	})

	t.Run("Test fill 3", func(t *testing.T) {
		slice := GoSlice{}

		got := slice.Fill("a", 5,nil, nil)
		want := GoSlice{"a","a","a","a", "a"}

		assertDeepEqual(t, got, want, "Fill failed")
	})
}

func TestHead(t *testing.T) {
	t.Run("Test head", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got := slice.Head();
		want := 1

		assertDeepEqual(t, got, want, "Get head failed")
	})

	t.Run("Test head is empty", func(t *testing.T) {
		slice := GoSlice{}

		got := slice.Head()

		assertDeepEqual(t, got, nil, "Get head failed")
	})

	t.Run("Test head is map", func(t *testing.T) {
		slice := GoSlice{
			map[string]interface{}{"user": "barney",  "active": true},
			map[string]interface{}{"user": "fred",  "active": false},
			map[string]interface{}{"user": "pebbles",  "active": false},
		}

		got := slice.Head()
		want := map[string]interface{}{"user": "barney",  "active": true}

		assertDeepEqual(t, got, want, "Get head failed")
	})
}

func TestFlatten(t *testing.T) {
	t.Run("Test flatten 1", func(t *testing.T) {
		slice := GoSlice{1, GoSlice{2, GoSlice{3, GoSlice{4}}}, 5}

		got := slice.Flatten()
		want := GoSlice{1, 2, GoSlice{3, GoSlice{4}}, 5}

		assertDeepEqual(t, got, want, "Flatten failed")
	})

	t.Run("Test flatten 2", func(t *testing.T) {
		slice := GoSlice{1, 2, 3, 4, 5}

		got := slice.Flatten()
		want := GoSlice{1, 2, 3, 4, 5}

		assertDeepEqual(t, got, want, "Flatten failed")
	})
}

//func TestFlattenDeep(t *testing.T) {
//	t.Run("Test flatten deep", func(t *testing.T) {
//		slice := GoSlice{1, GoSlice{2, GoSlice{3, GoSlice{4}}}, 5}
//
//		got := slice.Flatten()
//		want := GoSlice{1, 2, 3, 4, 5}
//
//		assertDeepEqual(t, got, want, "Flatten failed")
//	})
//}

func assertDeepEqual(t *testing.T, got, want interface{}, errMsg string) {
	t.Helper()

	if !(reflect.DeepEqual(got, want)) {
		fmt.Println(got, want)
		t.Error(errMsg)
	}
}