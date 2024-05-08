package math

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSieveEratosthenes(t *testing.T) {
	t.Parallel()

	for _, v := range []struct {
		name string
		val  int
		want []bool
	}{
		{
			name: "minus",
			val:  -10,
			want: nil,
		},
		{
			name: "0",
			val:  0,
			want: []bool{true},
		},
		{
			name: "1",
			val:  1,
			want: []bool{true, true},
		},
		{
			name: "small even number",
			val:  2,
			want: []bool{true, true, false},
		},
		{
			name: "small odd number",
			val:  3,
			want: []bool{true, true, false, false},
		},
		{
			name: "10",
			val:  10,
			want: []bool{true, true, false, false, true, false, true, false, true, true, true},
		},
	} {
		v2 := v
		t.Run(v2.name, func(t *testing.T) {
			t.Parallel()

			se := SieveEratosthenes(v2.val)
			if diff := cmp.Diff(v2.want, se); diff != "" {
				t.Errorf("SieveEratosthenes() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	t.Parallel()

	for _, v := range []struct {
		val  int64
		want bool
	}{
		{
			val:  -5,
			want: false,
		},
		{
			val:  0,
			want: false,
		},
		{
			val:  1,
			want: false,
		},
		{
			val:  2,
			want: true,
		},
		{
			val:  3,
			want: true,
		},
		{
			val:  4,
			want: false,
		},
		{
			val:  5,
			want: true,
		},
		{
			val:  6,
			want: false,
		},
		{
			val:  7,
			want: true,
		},
		{
			val:  8,
			want: false,
		},
		{
			val:  9,
			want: false,
		},
		{
			val:  10,
			want: false,
		},
	} {
		v2 := v
		t.Run(fmt.Sprintf("case %d", v2.val), func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(v2.want, IsPrime(v2.val)); diff != "" {
				t.Errorf("IsPrime() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
