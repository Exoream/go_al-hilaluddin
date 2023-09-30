package calculate

import "testing"

func TestAddition(t *testing.T) {
	result := Addition(3, 4)
	expected := 7
	if result != expected {
		t.Errorf("Error in Addition function")
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(5, 2)
	expected := 3
	if result != expected {
		t.Errorf("Error in Subtraction function")
	}
}

func TestDivision(t *testing.T) {
	result := Division(8, 2)
	expected := 4
	if result != expected {
		t.Errorf("Error in Division function")
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(4, 3)
	expected := 12
	if result != expected {
		t.Errorf("Error in Multiplication function")
	}
}