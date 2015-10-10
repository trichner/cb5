package cb5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnimation(t *testing.T) {

	a := NewAnimation()
	assert.NotNil(t, a)

	f := NewFrame()
	assert.NotNil(t, f)

	for x := uint8(0); x < 5; x += 2 {
		for y := uint8(0); y < 5; y += 2 {
			for z := uint8(0); z < 5; z += 2 {
				f.Set(x, y, z, true)
			}
		}
	}

	a.Append(*f)

	assert.Equal(t, uint32(1), a.Len(), "Not long enough :/")

	f = a.Get(0)

	for x := uint8(0); x < 5; x += 2 {
		for y := uint8(0); y < 5; y += 2 {
			for z := uint8(0); z < 5; z += 2 {
				assert.True(t, f.Get(x, y, z))
			}
		}
	}

	for x := uint8(1); x < 5; x += 2 {
		for y := uint8(1); y < 5; y += 2 {
			for z := uint8(1); z < 5; z += 2 {
				assert.False(t, f.Get(x, y, z))
			}
		}
	}
}
