package math

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGcd(t *testing.T) {
	t.Parallel()

	res := Gcd(8, 12)

	if diff := cmp.Diff(4, res); diff != "" {
		t.Errorf("gcd() mismatch (-want +got):\n%s", diff)
	}
}

func TestMultipleGcd(t *testing.T) {
	t.Parallel()

	for _, v := range []struct {
		name string
		val  []int
		want int
	}{
		{
			name: "case1",
			val:  []int{2, 6, 8},
			want: 2,
		},
		{
			name: "case2",
			val:  []int{4, 7, 14, 35, 60},
			want: 1,
		},
	} {
		res := MultipleGcd(v.val)

		if diff := cmp.Diff(v.want, res); diff != "" {
			t.Errorf("MultipleGcd() mismatch (-want +got):\n%s", diff)
		}
	}

}
