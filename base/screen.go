package base

import "unsafe"

type Screen struct {
	buffer []byte
	Width  int
	Height int
}

func NewScreen(width, height int) *Screen {
	return &Screen{
		buffer: make([]byte, width*height*4),
		Width:  width,
		Height: height,
	}
}

func (s *Screen) SetPixel(x, y int, color Color) {
	index := (y*int(s.Width) + x) * 4
	if index > s.Width*s.Height*4-4 || index < 0 {
		return
	}
	s.buffer[index+3] = color.R
	s.buffer[index+2] = color.G
	s.buffer[index+1] = color.B
}
func (s *Screen) UnsafePointer() unsafe.Pointer {
	return unsafe.Pointer(&s.buffer[0])
}

func (s *Screen) Clear() {
	for i := range s.buffer {
		s.buffer[i] = 0
	}
}
