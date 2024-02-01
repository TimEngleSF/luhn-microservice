package main

import (
	"testing"
)

func TestLuhnAlgo(t *testing.T) {
	subtests := []struct {
		d int
		r bool
	}{
		{
			d: 1008,
			r: true,
		},
		{
			d: 1016,
			r: true,
		},
		{
			d: 1024,
			r: true,
		},
		{
			d: 1032,
			r: true,
		},
		{
			d: 1040,
			r: true,
		},
		{
			d: 1009,
			r: false,
		},
		{
			d: 1017,
			r: false,
		},
		{
			d: 1025,
			r: false,
		},
		{
			d: 1031,
			r: false,
		},
		{
			d: 1043,
			r: false,
		},
	}

	for _, st := range subtests {
		if b := validateCard(st.d); b != st.r {
			t.Errorf("wanted %v (%d), got %v", st.r, st.d, b)
		}
	}
}
