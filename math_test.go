package math

import (
	"errors"
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

		v2 := v
		t.Run(v2.name, func(t *testing.T) {
			t.Parallel()

			res := MultipleGcd(v2.val)

			if diff := cmp.Diff(v2.want, res); diff != "" {
				t.Errorf("MultipleGcd() mismatch (-want +got):\n%s", diff)
			}
		})
	}

}

func TestChineseRemainderTheorem(t *testing.T) {
	t.Parallel()

	for _, v := range []struct {
		name                 string
		m1, m2, b1, b2, want int
		err                  error
	}{
		{
			name: "success 1",
			m1:   5,
			m2:   6,
			b1:   2,
			b2:   3,
			want: 27,
		},
		{
			name: "success 2",
			m1:   100,
			m2:   101,
			b1:   10,
			b2:   11,
			want: 10010,
		},
		{
			name: "success 2",
			m1:   100,
			m2:   101,
			b1:   10,
			b2:   11,
			want: 10010,
		},
		{
			name: "error",
			m1:   4,
			m2:   8,
			b1:   1,
			b2:   12,
			err:  ErrChineseRemainderTheorem,
		},
	} {
		v2 := v
		t.Run(v2.name, func(t *testing.T) {
			t.Parallel()

			res, err := ChineseRemainderTheorem(v2.m1, v2.m2, v2.b1, v2.b2)
			if diff := cmp.Diff(v2.want, res); diff != "" {
				t.Errorf("ChineseRemainderTheorem() mismatch (-want +got):\n%s", diff)
			}
			if v2.err != nil {
				if !errors.Is(err, v2.err) {
					t.Error(err)
				}
			}
		})
	}
}
