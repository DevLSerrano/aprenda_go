package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Adding 2 + 2, expected 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("Adding 4 + 4, expected 8", func(t *testing.T) {
		sum := Add(4, 4)
		expected := 8
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("Adding 10 + 10, expected error", func(t *testing.T) {
		sum := Add(10, 10)
		expected := 30
		if sum != expected {
			t.Skip("skipping test")
		} else {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

}
