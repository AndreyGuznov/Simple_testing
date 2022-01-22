package example2

import "testing"

func average(numb int) bool {
	return numb%2 == 0
}
func TestFuncIsEven(t *testing.T) {
	x := 3
	res := IsEven(x)
	average := average(x)
	switch {
	case x == 1 || x == (-1):
		t.Error("Our programm do not work with -1 or 1")
	case res != average:
		t.Error("Programm dont work")
	}
}
