package goderscore

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	t.Run("Test camelcase 1", func(t *testing.T) {
		str := "Foo Bar"

		got := CamelCase(str)
		want := "fooBar"

		assertEqual(t, got, want, "CamelCase failed")
	})

	t.Run("Test camelcase 2", func(t *testing.T) {
		str := "--foo-bar--"

		got := CamelCase(str)
		want := "fooBar"

		assertEqual(t, got, want, "CamelCase failed")
	})
}

func TestCapitalize(t *testing.T) {
	t.Run("Test capitalize", func(t *testing.T) {
		str := "FRED"

		got := Capitalize(str)
		want := "Fred"

		assertEqual(t, got, want, "Capitalize failed")
	})
}

func TestEndWith(t *testing.T) {
	t.Run("Test end with", func(t *testing.T) {
		str := "abc"

		got := EndWith(str, "c")
		want := true

		assertEqual(t, got, want, "Capitalize failed")
	})
}

func TestEndWithPosition(t *testing.T) {
	t.Run("Test end with position", func(t *testing.T) {
		str := "abc"

		got := EndWithPosition(str, "b", 2)
		want := true

		assertEqual(t, got, want, "Capitalize failed")
	})
}

func TestEscape(t *testing.T) {
	t.Run("Test escape", func(t *testing.T) {
		str := "fred, barney, & pebbles"

		got := Escape(str)
		want := "fred, barney, &amp; pebbles"

		assertEqual(t, got, want, "Escape failed")
	})
}

func TestEscapeRegExp(t *testing.T) {
	t.Run("Test escapeRegExp", func(t *testing.T) {
		str := "[lodash](https://lodash.com/)"

		got := EscapeRegExp(str)
		want := `\[lodash\]\(https://lodash\.com/\)`

		assertEqual(t, got, want, "EscapeRegExp failed")
	})
}

func TestKebabCase(t *testing.T) {
	t.Run("Test kebabCase", func(t *testing.T) {
		str1 := "Foo Bar"

		got := KebabCase(str1)
		want := "foo-bar"

		assertEqual(t, got, want, "KebabCase failed")

		str2 := "fooBar"

		got = KebabCase(str2)
		want = "foo-bar"

		assertEqual(t, got, want, "KebabCase failed")

		str3 := "Get_foo-BarTar"

		got = KebabCase(str3)
		want = "get-foo-bar-tar"

		str4 := "__FOO_BAR__"
		got = KebabCase(str4)
		want = "foo-bar"

		assertEqual(t, got, want, "KebabCase failed")
	})
}

func TestLowerCase(t *testing.T) {
	t.Run("Test lower case", func(t *testing.T) {
		str1 := "--Foo-Bar--"

		got := LowerCase(str1)
		want := "foo bar"

		assertEqual(t, got, want, "LowerCase failed")

		str2 := "fooBar"

		got = LowerCase(str2)
		want = "foo bar"

		assertEqual(t, got, want, "LowerCase failed")

		str3 := "__FOO_BAR__"

		got = LowerCase(str3)
		want = "foo bar"

		assertEqual(t, got, want, "LowerCase failed")
	})
}

func TestLowerFirst(t *testing.T) {
	t.Run("Test lower first", func(t *testing.T) {
		str1 := "Fred"

		got := LowerFirst(str1)
		want := "fred"

		assertEqual(t, got, want, "LowerFirst failed")

		str2 := "FRED"

		got = LowerFirst(str2)
		want = "fRED"

		assertEqual(t, got, want, "LowerFirst failed")
	})
}

func TestPad(t *testing.T) {
	t.Run("Test pad", func(t *testing.T) {
		str1 := "abc"

		got := Pad(str1, 8, " ")
		want := "  abc   "

		assertEqual(t, got, want, "Pad failed")

		str2 := "abc"

		got = Pad(str2, 8, "_-")
		want = "_-abc_-_"

		assertEqual(t, got, want, "Pad failed")

		str3 := "abc"

		got = Pad(str3, 3, "")
		want = "abc"

		assertEqual(t, got, want, "Pad failed")
	})
}

func TestPadEnd(t *testing.T) {
	t.Run("Test pad end", func(t *testing.T) {
		str1 := "abc"

		got := PadEnd(str1, 6, " ")
		want := "abc   "

		assertEqual(t, got, want, "Pad end failed")

		str2 := "abc"

		got = PadEnd(str2, 6, "_-")
		want = "abc_-_"

		assertEqual(t, got, want, "Pad end failed")

		str3 := "abc"

		got = PadEnd(str3, 3, "")
		want = "abc"

		assertEqual(t, got, want, "Pad end failed")
	})
}

func TestPadStart(t *testing.T) {
	t.Run("Test pad start", func(t *testing.T) {
		str1 := "abc"

		got := PadStart(str1, 6, " ")
		want := "   abc"

		assertEqual(t, got, want, "Pad start failed")

		str2 := "abc"

		got = PadStart(str2, 6, "_-")
		want = "_-_abc"

		assertEqual(t, got, want, "Pad start failed")

		str3 := "abc"

		got = PadStart(str3, 3, "")
		want = "abc"

		assertEqual(t, got, want, "Pad start failed")
	})
}

func TestParseInt(t *testing.T) {
	t.Run("Test parse int", func(t *testing.T) {
		str1 := "08"

		got := ParseInt(str1, 10)
		want := 8

		assertEqual(t, got, want, "ParseInt failed")

		str2 := "15"

		got = ParseInt(str2, 16)
		want = 21

		assertEqual(t, got, want, "ParseInt failed")
	})
}

func TestRepeat(t *testing.T) {
	t.Run("Test repeat", func(t *testing.T) {
		str1 := "*"

		got := Repeat(str1, 3)
		want := "***"

		assertEqual(t, got, want, "Pad start failed")

		str2 := "abc"

		got = Repeat(str2, 2)
		want = "abcabc"

		assertEqual(t, got, want, "Pad start failed")

		str3 := "abc"

		got = Repeat(str3, 0)
		want = ""

		assertEqual(t, got, want, "Pad start failed")
	})
}

func TestReplace(t *testing.T) {
	t.Run("Test replace", func(t *testing.T) {
		str1 := "Hi Fred"

		got := Replace(str1, "Hi", "Hello", 2)
		want := "Hello Fred"

		assertEqual(t, got, want, "Replace failed")

		str2 := "Hi Fred, Hi World"

		got = Replace(str2, "Hi", "Hello", 2)
		want = "Hello Fred, Hello World"

		assertEqual(t, got, want, "Replace failed")
	})
}

func TestSnakeCase(t *testing.T) {
	t.Run("Test lower case", func(t *testing.T) {
		str1 := "Foo Bar"

		got := SnakeCase(str1)
		want := "foo_bar"

		assertEqual(t, got, want, "SnakeCase failed")

		str2 := "fooBar"

		got = SnakeCase(str2)
		want = "foo_bar"

		assertEqual(t, got, want, "SnakeCase failed")

		str3 := "--FOO-BAR--"

		got = SnakeCase(str3)
		want = "foo_bar"


		assertEqual(t, got, want, "SnakeCase failed")
	})
}

func TestSplit(t *testing.T) {
	t.Run("Test replace", func(t *testing.T) {
		str1 := "a-b-c"

		got := Split(str1, "-")
		want := []string{"a", "b", "c"}

		assertDeepEqual(t, got, want, "Split failed")
	})
}

func TestSplitLimit(t *testing.T) {
	t.Run("Test splitLimit", func(t *testing.T) {
		str1 := "a-b-c"

		got := SplitLimit(str1, "-", 2)
		want := []string{"a", "b"}

		assertDeepEqual(t, got, want, "SplitLimit failed")
	})
}

func TestStartCase(t *testing.T) {
	t.Run("Test start case", func(t *testing.T) {
		str1 := "--foo-bar--"

		got := StartCase(str1)
		want := "Foo Bar"

		assertEqual(t, got, want, "StartCase failed")

		str2 := "fooBar"

		got = StartCase(str2)
		want = "Foo Bar"

		assertEqual(t, got, want, "StartCase failed")

		str3 := "__FOO_BAR__"

		got = StartCase(str3)
		want = "FOO BAR"


		assertEqual(t, got, want, "StartCase failed")
	})
}

func TestStartsWith(t *testing.T) {
	t.Run("Test starts with", func(t *testing.T) {
		str1 := "abc"

		got := StartsWith(str1, "a", 0)
		want := true

		assertEqual(t, got, want, "StartsWith failed")

		str2 := "abc"

		got = StartsWith(str2, "b", 0)
		want = false

		assertEqual(t, got, want, "StartsWith failed")

		str3 := "abc"

		got = StartsWith(str3, "b", 1)
		want = true


		assertEqual(t, got, want, "StartsWith failed")
	})
}

func TestToLower(t *testing.T) {
	t.Run("Test to lower", func(t *testing.T) {
		str1 := "--Foo-Bar--"

		got := ToLower(str1)
		want := "--foo-bar--"

		assertEqual(t, got, want, "ToLower failed")

		str2 := "fooBar"

		got = ToLower(str2)
		want = "foobar"

		assertEqual(t, got, want, "ToLower failed")

		str3 := "__FOO_BAR__"

		got = ToLower(str3)
		want = "__foo_bar__"


		assertEqual(t, got, want, "ToLower failed")
	})
}

func TestToUpper(t *testing.T) {
	t.Run("Test to upper", func(t *testing.T) {
		str1 := "--Foo-Bar--"

		got := ToUpper(str1)
		want := "--FOO-BAR--"

		assertEqual(t, got, want, "ToUpper failed")

		str2 := "fooBar"

		got = ToUpper(str2)
		want = "FOOBAR"

		assertEqual(t, got, want, "ToUpper failed")

		str3 := "__FOO_BAR__"

		got = ToUpper(str3)
		want = "__FOO_BAR__"


		assertEqual(t, got, want, "ToUpper failed")
	})
}

func TestTrim(t *testing.T) {
	t.Run("Test trim", func(t *testing.T) {
		str1 := "  abc  "

		got := Trim(str1, " ")
		want := "abc"

		assertEqual(t, got, want, "Trim failed")

		str2 := "-_-abc-_-"

		got = Trim(str2, "_-")
		want = "abc"

		assertEqual(t, got, want, "Trim failed")
	})
}

func TestTrimEnd(t *testing.T) {
	t.Run("Test trim end", func(t *testing.T) {
		str1 := "  abc  "

		got := TrimEnd(str1, " ")
		want := "  abc"

		assertEqual(t, got, want, "TrimEnd failed")

		str2 := "-_-abc-_-"

		got = TrimEnd(str2, "_-")
		want = "-_-abc"

		assertEqual(t, got, want, "TrimEnd failed")
	})
}

func TestTrimStart(t *testing.T) {
	t.Run("Test trim start", func(t *testing.T) {
		str1 := "  abc  "

		got := TrimStart(str1, " ")
		want := "abc  "

		assertEqual(t, got, want, "TrimStart failed")

		str2 := "-_-abc-_-"

		got = TrimStart(str2, "_-")
		want = "abc-_-"

		assertEqual(t, got, want, "TrimStart failed")
	})
}

func assertEqual(t *testing.T, got, want interface{}, errMsg string) {
	t.Helper()

	if got != want {
		t.Error(errMsg)
	}
}
