package employee

import "testing"

// baseSalary -> 10000000
//extra -> base / 30 / 7 *1.4 * hours
// karaneh -> 5000000
// insurance -> b + e * 0.03
//tax -> b + e * 0.10

func TestExtraSalaryCalculate(t *testing.T) {
	// Arrange
	baseSalary := 10000000.0
	extraHours := 12.0
	want := 80000.0
	// Act
	got := ExtraSalaryCalculate(baseSalary, extraHours)

	//Assert
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
