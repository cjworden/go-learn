package euler

import "testing"

func TestOne(t *testing.T) {
	ans := 233168
	if sol := problemOne(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestTwo(t *testing.T) {
	ans := 4613732
	if sol := problemTwo(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestThree(t *testing.T) {
	ans := 6857
	if sol := problemThree(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestFour(t *testing.T) {
	ans := 906609
	if sol := problemFour(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestFive(t *testing.T) {
	ans := 232792560
	if sol := problemFive(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestSix(t *testing.T) {
	ans := 25164150
	if sol := problemSix(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestSeven(t *testing.T) {
	ans := 104743
	if sol := problemSeven(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestEight(t *testing.T) {
	ans := 23514624000
	if sol := problemEight(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestNine(t *testing.T) {
	ans := 31875000
	if sol := problemNine(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}

func TestTen(t *testing.T) {
	ans := 142913828922
	if sol := problemTen(); sol != ans {
		t.Errorf("Expected %v, got %v", ans, sol)
	}
}