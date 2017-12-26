package hellogo

import "testing"


func TestCalcFailed(t *testing.T) {
	result := Calc(4)
	if result != 10 {
		t.Fatalf("Calc with 4 should return 10. but return %f.", result)
	}
}
