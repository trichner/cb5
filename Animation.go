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

type Animation struct {
	frames []Frame
}

func NewAnimation() *Animation {
	return &Animation{
		frames: make([]Frame, 50),
	}
}

// Set sets an LED on or off
func (a *Animation) Append(Frame f) {
	a.frames = append(a.frames, f)
}

func (a *Animation) Bytes() []byte {
	if len(a.frames) == 0 {
		a.Append(Frame{})
	}
	// set flags
	a.frames[0].SetSceneStart(true)
	a.frames[len(a.frames)-1].SetEnd(true)
	a.frames[len(a.frames)-1].SetSceneEnd(true)

}
