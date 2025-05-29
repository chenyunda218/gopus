package goopus

import "time"

type Streaming struct {
	frames  []*Frame
	buffers []byte
}

func (s *Streaming) AppendBytes(bs ...[]byte) *Streaming {
	for _, b := range bs {
		s.buffers = append(s.buffers, b...)
	}
	return s
}

func (s *Streaming) AppendFrameWithBytes(bs ...[]byte) *Streaming {
	for _, b := range bs {
		s.frames = append(s.frames, New(WithBytes(b)))
	}
	return s
}

func (s *Streaming) AppendFrame(frame ...*Frame) *Streaming {
	for _, f := range frame {
		if f != nil {
			s.frames = append(s.frames, f)
		}
	}
	return s
}

func (s *Streaming) Frames() []*Frame {
	return s.frames
}

type Option func(*Frame) *Frame

func WithBytes(bytes []byte) func(*Frame) *Frame {
	return func(f *Frame) *Frame {
		f.bytes = bytes
		return f
	}
}

func New(option Option) *Frame {
	var f Frame
	if option != nil {
		return option(&f)
	} else {
		return &f
	}
}

type Frame struct {
	bytes []byte
}

func (f *Frame) Bytes() []byte {
	if f.bytes == nil {
		return []byte{}
	}
	return f.bytes
}

func (f *Frame) Config() int {
	return int(f.bytes[0] & 0b11111)
}

func (f *Frame) Channels() int {
	return int(f.bytes[0] >> 5 & 1)
}

func (f *Frame) Bandwidth() Bandwidth {
	config := f.Config()
	switch {
	case config <= 3:
		return NB
	case config <= 7:
		return MB
	case config <= 11:
		return WB
	case config <= 13:
		return SWB
	case config <= 15:
		return FB
	case config <= 19:
		return NB
	case config <= 23:
		return WB
	case config <= 27:
		return SWB
	default:
		return FB
	}
}

func (f *Frame) Duration() time.Duration {
	config := f.Config()
	switch config {
	case 0:
		return 10000 * time.Microsecond
	case 1:
		return 20000 * time.Microsecond
	case 2:
		return 40000 * time.Microsecond
	case 3:
		return 60000 * time.Microsecond
	case 4:
		return 10000 * time.Microsecond
	case 5:
		return 20000 * time.Microsecond
	case 6:
		return 40000 * time.Microsecond
	case 7:
		return 60000 * time.Microsecond
	case 8:
		return 10000 * time.Microsecond
	case 9:
		return 20000 * time.Microsecond
	case 10:
		return 40000 * time.Microsecond
	case 11:
		return 60000 * time.Microsecond
	case 12:
		return 10000 * time.Microsecond
	case 13:
		return 20000 * time.Microsecond
	case 14:
		return 10000 * time.Microsecond
	case 15:
		return 20000 * time.Microsecond
	case 16:
		return 2500 * time.Microsecond
	case 17:
		return 5000 * time.Microsecond
	case 18:
		return 10000 * time.Microsecond
	case 19:
		return 20000 * time.Microsecond
	case 20:
		return 2500 * time.Microsecond
	case 21:
		return 5000 * time.Microsecond
	case 22:
		return 10000 * time.Microsecond
	case 23:
		return 20000 * time.Microsecond
	case 24:
		return 2500 * time.Microsecond
	case 25:
		return 5000 * time.Microsecond
	case 26:
		return 10000 * time.Microsecond
	case 27:
		return 20000 * time.Microsecond
	case 28:
		return 2500 * time.Microsecond
	case 29:
		return 5000 * time.Microsecond
	case 30:
		return 10000 * time.Microsecond
	case 31:
		return 20000 * time.Microsecond
	}
	panic("Duration not implemented for this config")
}

func (f *Frame) Contents() int {
	return int(f.bytes[0] >> 6)
}
