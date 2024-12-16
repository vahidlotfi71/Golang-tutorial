package employee

func ExtraSalaryCalculate(baseSalary, extraHours float64) float64 {
	return (baseSalary * extraHours / 30 / 7 * 1.4) // math.Ceil رند می کند

}
