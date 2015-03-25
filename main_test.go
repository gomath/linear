package linear

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestTensor(t *testing.T) {
	x := []float64{1, 2, 3}
	y := []float64{4, 5, 6}
	z := []float64{7, 8, 9}

	T := []float64{
		1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3,
		4, 4, 4, 5, 5, 5, 6, 6, 6, 4, 4, 4, 5, 5, 5, 6, 6, 6, 4, 4, 4, 5, 5, 5, 6, 6, 6,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	}

	assert.Equal(Tensor(x, y, z), T, t)
}
