package goderscore

import "testing"

func TestClamp(t *testing.T) {
	t.Run("Test clamp", func(t *testing.T) {
		got := Clamp(-10, -5, 5)
		want := -5

		assertEqual(t, got, want, "Clamp failed")

		got = Clamp(10, -5, 5)
		want = 5

		assertEqual(t, got, want, "Clamp failed")

		got = Clamp(-4, -5, 5)
		want = -4

		assertEqual(t, got, want, "Clamp failed")
	})
}
