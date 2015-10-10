package cb5

import (
	"bytes"
)

type Animation struct {
	frames []Frame
}

func NewAnimation() *Animation {
	return &Animation{
		frames: make([]Frame, 50),
	}
}

// Set sets an LED on or off
func (a *Animation) Append(f Frame) {
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

	buf := new(bytes.Buffer)
	for _, f := range a.frames {
		buf.Write(f.Bytes())
	}

	return buf.Bytes()
}
