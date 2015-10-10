package cb5

import (
	"bytes"
	//"fmt"
)

type Animation struct {
	frames []Frame
}

func NewAnimation() *Animation {
	return &Animation{
		frames: make([]Frame, 0, 50), // should maybe use pointers to save memory (premature optimization)
	}
}

func (a *Animation) Append(f Frame) {
	a.frames = append(a.frames, f)
}

func (a *Animation) Len() uint32 {
	return uint32(len(a.frames))
}

func (a *Animation) Get(i uint32) *Frame {
	if i >= a.Len() {
		panic("out of bounds")
	}
	return &a.frames[i]
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
