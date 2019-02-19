package randnum

import "testing"

func TestIntSlice(t *testing.T) {
	intSlice := RandomIntSlice(20)
	n := len(intSlice)

	for i := 0; i < n; i++ {
		if intSlice[i] < 0 || intSlice[i] > 100 {
			t.Error("[Error] scope over flow!")
			return
		}
	}
}

func TestDistinctIntSlice(t *testing.T) {
	s := RandomDistinctIntSlice(30)
	t.Log(s)
	for i := len(s)-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	t.Log(s)
	for i := 0; i < len(s)-1; i++ {
		if s[i] >= s[i+1] {
			t.Error("[error]not distinct random number")
		}
	}
	t.Log("generate distinct random number succeed!")
}
