package bignum

import "testing"

func TestAdd(t *testing.T) {

	for i := uint8(0); i < 128; i++ {
		for j := uint8(0); j < 128; j++ {
			result := Add(uint8(i), uint8(j))
			if result != i+j {
				t.Errorf("Mismatch %v + %v: expected %v, got %v", i, j, i+j, result)
			}
		}
	}
}
