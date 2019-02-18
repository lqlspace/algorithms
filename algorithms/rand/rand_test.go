package randnum

import "testing"

func TestIntSlice(t *testing.T) {
	intSlice := RandomIntSlice(20)
	n := len(intSlice)

	for i := 0; i < n; i++ {
		if intSlice[i] < 30 || intSlice[i] > 100 {
			t.Error("[Error] scope over flow!")
			return
		}
	}
}
