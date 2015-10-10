package cb5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFrame(t *testing.T) {

	f := NewFrame()
	assert.NotNil(t, f)

	for x := uint8(0); x < 5; x++ {
		for y := uint8(0); y < 5; y++ {
			for z := uint8(0); z < 5; z++ {
				bit, _ := f.Get(x, y, z)
				assert.False(t, bit, "Initialization not correct")
			}
		}
	}
}

func TestFrameOOB(t *testing.T) {

	f := NewFrame()

	_, err := f.Get(0, 0, 42)
	assert.Error(t, err, "Should have returned an Out Of Bounds error :/")

	_, err = f.Get(7, 0, 0)
	assert.Error(t, err, "Should have returned an Out Of Bounds error :/")

	_, err = f.Get(0, 8, 0)
	assert.Error(t, err, "Should have returned an Out Of Bounds error :/")
}

func TestFrameGetSet(t *testing.T) {

	f := NewFrame()

	f.Set(0, 0, 0, true)
	bit, _ := f.Get(0, 0, 0)
	assert.True(t, bit, "False negative")

	bit, _ = f.Get(0, 0, 1)
	assert.False(t, bit, "False positive")
}

func TestClz(t *testing.T) {
	bits := uint8(0xFF)
	for i := uint8(0); i < 8; i++ {
		assert.Equal(t, i, clz(bits), "Count is wrong.")
		bits >>= 1
	}
}

func TestBrightness(t *testing.T) {
	f := NewFrame()
	assert.Equal(t, f.GetBrightness(), uint8(5), "Wrong initialised brightness.")

	for i := uint8(0); i <= 5; i++ {
		f.SetBrightness(i)
		assert.Equal(t, f.GetBrightness(), i, "Wrong brightness.")
	}

}

func TestEndFlag(t *testing.T) {
	f := NewFrame()
	assert.False(t, f.IsEnd(), "Wrong initialised endframe.")

	f.SetEnd(true)
	assert.True(t, f.IsEnd(), "Wrong endframe.")
	f.SetEnd(false)
	assert.False(t, f.IsEnd(), "Wrong endframe.")
}

func TestSceneEndFlag(t *testing.T) {
	f := NewFrame()
	assert.False(t, f.IsSceneEnd(), "Wrong initialised scene end flag.")

	f.SetSceneEnd(true)
	assert.True(t, f.IsSceneEnd(), "Wrong scene end flag.")
	f.SetSceneEnd(false)
	assert.False(t, f.IsSceneEnd(), "Wrong scene end flag.")
}

func TestSceneStartFlag(t *testing.T) {
	f := NewFrame()
	assert.False(t, f.IsSceneStart(), "Wrong initialised scene start flag.")

	f.SetSceneStart(true)
	assert.True(t, f.IsSceneStart(), "Wrong scene start flag.")
	f.SetSceneStart(false)
	assert.False(t, f.IsSceneStart(), "Wrong scene start flag.")
}