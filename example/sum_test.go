package modoct

import "testing"

//go test or go test -v
func TestSum(t *testing.T) {
	//t.Error("fail")
	t.Run("5+5 should be 10", func(t *testing.T) {
		total := Sum(5, 5)

		if total != 10 {
			t.Errorf("sum was incorrect,got %d want %d", total, 10)
		}
	})
}
