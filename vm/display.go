package vm

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Pixel uint8

const Width = 320
const Height = 160

const vram_size = Width * Height

const vram_start = 0x1000

const pallete = 0x500

func DisplayPixel(s *State, pix Pixel, xoff, yoff int32, win *sdl.Window, ren *sdl.Renderer) {
	r, g, b := lookup_colour(s, pix)
	ren.SetDrawColor(r, g, b, 255)
	ren.DrawPoint(xoff, yoff)
	ren.SetDrawColor(0, 0, 0, 0)
}

func Display(s *State, win *sdl.Window, ren *sdl.Renderer) error {
	for i := 0; i < vram_size; i++ {
		pix, err := s.r.PeekB(Address(vram_start + i))
		if err != nil {
			return err
		}
		xoff, yoff := off_to_coords(int32(i))
		DisplayPixel(s, Pixel(pix), xoff, yoff, win, ren)
	}
	ren.Present()
	return nil
}

func off_to_coords(off int32) (int32, int32) {
	xoff := off % Width
	yoff := off / Width
	return xoff, yoff
}

func lookup_colour(s *State, pix Pixel) (uint8, uint8, uint8) {
	r, err := s.r.PeekB(pallete + Address(pix)*3)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	g, err := s.r.PeekB(pallete + Address(pix)*3 + 1)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	b, err := s.r.PeekB(pallete + Address(pix)*3 + 2)
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	return uint8(r), uint8(g), uint8(b)
}
