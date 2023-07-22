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

	t.Run("Subtract 10 - 5, expected 5", func(t *testing.T) {
		sub := Subtract(10, 5)
		expected := 5
		if sub != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sub)
		}
	})

	t.Run("Divide 10 / 5, expected 2", func(t *testing.T) {
		div := Divide(10, 5)
		expected := 2
		if div != expected {
			t.Errorf("expected '%d' but got '%d'", expected, div)
		}
	})

	t.Run("Rest of division 10 % 5, expected 0", func(t *testing.T) {
		rest := RestOfDivision(10, 5)
		expected := 0
		if rest != expected {
			t.Errorf("expected '%d' but got '%d'", expected, rest)
		}
	})

}
