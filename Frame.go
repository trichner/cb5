package cb5

import (
	"errors"
)

const BIT_OFFSET = 7
const SIDE = 5
const BRIGHTNESS_Z = 1
const BRIGHTNESS_MAX = 5
const BRIGHTNESS_MASK = 0x7F

const SCENE_START_Z = 0
const SCENE_START_OFFSET = 0

const SCENE_END_Z = 4
const SCENE_END_OFFSET = 1

const END_Z = 4
const END_OFFSET = 0

type Frame struct {
	cube [5]uint32
}

func NewFrame() *Frame {
	return new(Frame)
}

// Set sets an LED on or off
func (f *Frame) Set(x uint8, y uint8, z uint8, on bool) error {
	if !inRanges(x, y, z) {
		return errors.New("out of bounds")
	}
	plane := f.cube[z]
	shift := (SIDE*y + x + BIT_OFFSET)
	plane &= ^(0x01 << shift)
	plane |= (btoi(on) << shift)
	f.cube[z] = plane
	return nil

}

// Get retrieves the state of an LED, error if out of bounds
func (f *Frame) Get(x uint8, y uint8, z uint8) (bool, error) {
	if !inRanges(x, y, z) {
		return false, errors.New("out of bounds")
	}
	plane := f.cube[z]
	bit := (plane >> (SIDE*y + x + BIT_OFFSET)) != 0
	return bit, nil
}

// SetBrightness sets the brightness, values between [0,5] where 0 is off and 5 is the maximum
// error if out of bounds
func (f *Frame) SetBrightness(b uint8) error {
	if !validBrightness(b) {
		return errors.New("out of bounds")
	}
	plane := f.cube[BRIGHTNESS_Z]
	plane &= ^uint32(BRIGHTNESS_MASK)
	if b != BRIGHTNESS_MAX {
		shift := (BRIGHTNESS_MAX - 1) - b
		plane |= (0x01 << shift)
	}
	f.cube[BRIGHTNESS_Z] = plane
	return nil
}

// GetBrightness gets the brightness, values between [0,5] where 0 is off and 5 is the maximum
func (f *Frame) GetBrightness() uint8 {
	plane := f.cube[BRIGHTNESS_Z]
	bits := uint8(plane) & BRIGHTNESS_MASK
	bits <<= 3
	bits |= 0x07
	return clz(bits)
}

func (f *Frame) SetSceneStart(s bool) {
	f.cube[SCENE_START_Z] = setBit(SCENE_START_OFFSET, f.cube[SCENE_START_Z], s)
}

func (f *Frame) IsSceneStart() bool {
	return isBit(SCENE_START_OFFSET, f.cube[SCENE_START_Z])
}

func (f *Frame) SetSceneEnd(s bool) {
	f.cube[SCENE_END_Z] = setBit(SCENE_END_OFFSET, f.cube[SCENE_END_Z], s)
}

func (f *Frame) IsSceneEnd() bool {
	return isBit(SCENE_END_OFFSET, f.cube[SCENE_END_Z])
}

func (f *Frame) SetEnd(s bool) {
	f.cube[END_Z] = setBit(END_OFFSET, f.cube[END_Z], s)
}

func (f *Frame) IsEnd() bool {
	return isBit(END_OFFSET, f.cube[END_Z])
}

//==== Helpers

func getBits(mask uint32, offset uint32, value uint32) uint32 {
	value >>= offset
	value &= mask
	return value
}

func setBits(mask uint32, offset uint32, value uint32, bits uint32) uint32 {
	bits &= mask
	value &= ^(mask << offset)
	value |= (bits << offset)
	return value
}

func isBit(offset uint32, value uint32) bool {
	return getBits(0x01, offset, value) != 0
}

func setBit(offset uint32, value uint32, bit bool) uint32 {
	return setBits(0x01, offset, value, btoi(bit))
}

func validBrightness(b uint8) bool {
	return (b >= 0) && (b <= BRIGHTNESS_MAX)
}

func inRanges(x uint8, y uint8, z uint8) bool {
	return (inRange(x) && inRange(y) && inRange(z))
}

func inRange(i uint8) bool {
	return i >= 0 && i < 5
}

func btoi(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// clz counts leading zeros
func clz(x uint8) uint8 {
	var n = uint8(1)

	if (x >> (4)) == 0 {
		n = n + 4
		x = x << 4
	}

	if (x >> (4 + 2)) == 0 {
		n = n + 2
		x = x << 2
	}

	n = n - (x >> 7)
	return n
}
