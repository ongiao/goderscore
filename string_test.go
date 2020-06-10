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

		assertEqual(t, got, want, "KebabCase failed")
	})
}

//func TestLowerCase(t *testing.T) {
//	t.Run("Test lower case", func(t *testing.T) {
//		str1 := "--Foo-Bar--"
//
//		got := LowerCase(str1)
//		want := "foo bar"
//
//		assertEqual(t, got, want, "KebabCase failed")
//
//		str2 := "fooBar"
//
//		got = LowerCase(str2)
//		want = "foo bar"
//
//		assertEqual(t, got, want, "KebabCase failed")
//
//		str3 := "__FOO_BAR__"
//
//		got = LowerCase(str3)
//		want = "foo bar"
//
//		fmt.Println(got, want)
//
//		assertEqual(t, got, want, "KebabCase failed")
//	})
//}

func assertEqual(t *testing.T, got, want interface{}, errMsg string) {
	t.Helper()

	if got != want {
		t.Error(errMsg)
	}
}
