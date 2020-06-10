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

func TestContains(t *testing.T) {
	t.Run("Test contains 1", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got, _ := slice.Contains(3)
		want := true

		if got != want {
			t.Error("Contains failed")
		}
	})

	t.Run("Test contains 2", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got, _ := slice.Contains(8)
		want := false

		if got != want {
			t.Error("Contains failed")
		}
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

	t.Run("Test flatten 3", func(t *testing.T) {
		slice := GoSlice{}

		got := slice.FlattenDeep()
		want := GoSlice{}

		assertDeepEqual(t, got, want, "Flatten failed")
	})
}

func TestFlattenDeep(t *testing.T) {
	t.Run("Test flatten deep 1", func(t *testing.T) {
		slice := GoSlice{1, GoSlice{2, GoSlice{3, GoSlice{4}}}, 5}

		got := slice.FlattenDeep()
		want := GoSlice{1, 2, 3, 4, 5}

		assertDeepEqual(t, got, want, "FlattenDeep failed")
	})

	t.Run("Test flatten deep 2", func(t *testing.T) {
		slice := GoSlice{}

		got := slice.FlattenDeep()
		want := GoSlice{}

		assertDeepEqual(t, got, want, "FlattenDeep failed")
	})
}

func TestFlattenDepth(t *testing.T) {
	t.Run("Test flatten depth", func(t *testing.T) {
		slice := GoSlice{1, GoSlice{2, GoSlice{3, GoSlice{4}}}, 5}

		got := slice.FlattenDepth(2)
		want := GoSlice{1, 2, 3, GoSlice{4}, 5}

		assertDeepEqual(t, got, want, "FlattenDeep failed")
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("Test index of 1", func(t *testing.T) {
		slice := GoSlice{1, 2, 3, 4, 5}
		searchNum := 3

		got1, _ := slice.IndexOf(searchNum, 3)
		want1 := -1

		if got1 != want1 {
			t.Error("IndexOf failed")
		}

		got2, _ := slice.IndexOf(searchNum, 0)
		want2 := 2

		if got2 != want2 {
			t.Error("IndexOf failed")
		}

		otherSlice := GoSlice{2,3,4,5,3,1}

		got3, _ := otherSlice.IndexOf(searchNum, -1)
		want3 := 4

		if got3 != want3 {
			t.Error("IndexOf failed")
		}
	})
}

func TestInitial(t *testing.T) {
	t.Run("Test initial", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got := slice.Initial()
		want := GoSlice{1,2,3,4}

		assertDeepEqual(t, got, want, "Initial failed")
	})
}

func TestIntersection(t *testing.T) {
	t.Run("Test intersection", func(t *testing.T) {
		a := GoSlice{2,1}
		b := GoSlice{4,2,1}
		c := GoSlice{1,2}
		d := GoSlice{3,1,2}

		got := a.Intersection(b, c, d)
		want := GoSlice{1,2}

		assertSliceHasSameElems(t, got, want, "Intersection failed")
	})
}

func TestJoin(t *testing.T) {
	t.Run("Test join 1", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got := slice.Join("")
		want := "12345"

		if got != want {
			t.Error("Join failed")
		}
	})

	t.Run("Test join 2", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got := slice.Join("-")
		want := "1-2-3-4-5"

		if got != want {
			t.Error("Join failed")
		}
	})

	t.Run("Test join 3", func(t *testing.T) {
		slice := GoSlice{"Jack","is","Junxing"}

		got := slice.Join("-")
		want := "Jack-is-Junxing"

		if got != want {
			t.Error("Join failed")
		}
	})
}

func TestLast(t *testing.T) {
	t.Run("Test last 1", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got := slice.Last()
		want := 5

		if got != want {
			t.Error("Last failed")
		}
	})

	t.Run("Test last 2", func(t *testing.T) {
		slice := GoSlice{"Jack", "is", "Junxing"}

		got := slice.Last()
		want := "Junxing"

		if got != want {
			t.Error("Last failed")
		}
	})
}

func TestLastIndexOf(t *testing.T) {
	t.Run("Test last index of 1", func(t *testing.T) {
		slice := GoSlice{1,2,1,2}

		got, _ := slice.LastIndexOf(2, len(slice) - 1)
		want := 3

		if got != want {
			t.Error("Last failed")
		}
	})

	t.Run("Test last index of 2", func(t *testing.T) {
		slice := GoSlice{1,2,1,2}

		got, _ := slice.LastIndexOf(2, 2)
		want := 1

		if got != want {
			t.Error("Last failed")
		}
	})

	t.Run("Test last index of 2", func(t *testing.T) {
		slice := GoSlice{1,2,1,2}

		got, _ := slice.LastIndexOf(2, -1)
		want := 1

		if got != want {
			t.Error("Last failed")
		}
	})
}

func TestNth(t *testing.T) {
	t.Run("Test nth", func(t *testing.T) {
		slice := GoSlice{"a","b","c","d"}

		got := slice.Nth(1)
		want := "b"

		if got != want {
			t.Error("Nth failed")
		}
	})
}

func TestUniq(t *testing.T) {
	t.Run("Test uniq", func(t *testing.T) {
		slice := GoSlice{2,3,4,5,3,2,1,9}

		got := slice.Uniq()
		want := GoSlice{2,3,4,5,1,9}

		assertDeepEqual(t, got, want, "Uniq failed")
	})
}

func TestPull(t *testing.T) {
	t.Run("Test pull 1", func(t *testing.T) {
		slice := GoSlice{1,2,3,1,2,3}

		got := slice.Pull(2, 3)
		want := GoSlice{1,1}

		assertDeepEqual(t, slice, want, "Pull failed")
		assertDeepEqual(t, got, want, "Pull failed")
	})

	t.Run("Test pull 2", func(t *testing.T) {
		slice := GoSlice{1,2,3,1,2,3,4,5}

		got := slice.Pull(7,8)
		want := GoSlice{1,2,3,1,2,3,4,5}

		assertDeepEqual(t, slice, want, "Pull failed")
		assertDeepEqual(t, got, want, "Pull failed")
	})
}

func TestPullAll(t *testing.T) {
	t.Run("Test pull all", func(t *testing.T) {
		slice := GoSlice{1,2,3,1,2,3}
		values := GoSlice{2,3}

		got := slice.PullAll(values)
		want := GoSlice{1,1}

		assertDeepEqual(t, got, want, "Pull all failed")
	})
}

func TestPullAt(t *testing.T) {
	t.Run("Test pull at", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5}

		got := slice.PullAt(1,3)
		want1 := GoSlice{2,4}
		want2 := GoSlice{1,3,5}

		assertDeepEqual(t, got, want1, "Pull at failed")
		assertDeepEqual(t, slice, want2, "Pull at failed")
	})
}

func TestReverse(t *testing.T) {
	t.Run("Test reverse", func(t *testing.T) {
		slice := GoSlice{1,2,3,4}

		got := slice.Reverse()
		want := GoSlice{4,3,2,1}

		assertDeepEqual(t, got, want, "Reverse failed")
	})
}

func TestSlice(t *testing.T) {
	t.Run("Test slice", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5,6}

		got := slice.Slice(1, 5)
		want := GoSlice{2,3,4,5}

		assertDeepEqual(t, got, want, "Slice failed")
	})
}

func TestTail(t *testing.T) {
	t.Run("Test tail", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5,6}

		got := slice.Tail()
		want := GoSlice{2,3,4,5,6}

		assertDeepEqual(t, got, want, "Tail failed")
	})
}

func TestTake(t *testing.T) {
	t.Run("Test take", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5,6}

		got := slice.Take(3)
		want := GoSlice{1,2,3}

		assertDeepEqual(t, got, want, "Take failed")
	})
}

func TestTakeRight(t *testing.T) {
	t.Run("Test take right", func(t *testing.T) {
		slice := GoSlice{1,2,3,4,5,6}

		got := slice.TakeRight(3)
		want := GoSlice{4,5,6}

		assertDeepEqual(t, got, want, "Take right failed")
	})
}

func TestUnion(t *testing.T) {
	t.Run("Test union", func(t *testing.T) {
		slice1 := GoSlice{2}
		slice2 := GoSlice{2,1}
		slice3 := GoSlice{2,4,7}

		got := Union(slice1, slice2, slice3)
		want := GoSlice{2,1,4,7}

		assertDeepEqual(t, got, want, "Take union failed")
	})
}

func TestUnzip(t *testing.T) {
	t.Run("Test unzip", func(t *testing.T) {
		slice1 := GoSlice{"fred","barney"}
		slice2 := GoSlice{30,40}
		slice3 := GoSlice{true,false}

		got := Unzip(slice1, slice2, slice3)
		want := GoSlice{GoSlice{"fred",30,true}, GoSlice{"barney",40,false}}

		assertDeepEqual(t, got, want, "Unzip failed")
	})
}

func TestWithout(t *testing.T) {
	t.Run("Test without", func(t *testing.T) {
		slice := GoSlice{1,2,3,1,2,3}

		got := slice.Pull(2, 3)
		want := GoSlice{1,1}

		assertDeepEqual(t, got, want, "Without failed")
	})
}

func TestXor(t *testing.T) {
	t.Run("Test xor", func(t *testing.T) {
		slice1 := GoSlice{2,1}
		slice2 := GoSlice{2,3}

		got := Xor(slice1, slice2)
		want := GoSlice{1,3}

		assertDeepEqual(t, got, want, "Xor failed")
	})
}

func TestZip(t *testing.T) {
	t.Run("Test zip", func(t *testing.T) {
		slice1 := GoSlice{"a",1,true}
		slice2 := GoSlice{"b",2,false}

		got := Zip(slice1, slice2)
		want := GoSlice{GoSlice{"a","b"},GoSlice{1,2},GoSlice{true,false}}

		assertDeepEqual(t, got, want, "Zip failed")
	})
}

func assertSliceHasSameElems(t *testing.T, got, want interface{}, errMsg string) {
	mapA := make(map[interface{}]int)
	mapB := make(map[interface{}]int)

	for _, n := range got.(GoSlice) {
		if mapA[n] == 0 {
			mapA[n] = 1
		} else {
			mapA[n]++
		}
	}

	for _, n := range want.(GoSlice) {
		if mapB[n] == 0 {
			mapB[n] = 1
		} else {
			mapB[n]++
		}
	}

	assertDeepEqual(t, mapA, mapB, errMsg)
}

func assertDeepEqual(t *testing.T, got, want interface{}, errMsg string) {
	t.Helper()

	if !(reflect.DeepEqual(got, want)) {
		fmt.Println(got, want)
		t.Error(errMsg)
	}
}