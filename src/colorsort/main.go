package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

func time_left(next_time uint32) uint32 {
	var now uint32 = sdl.GetTicks()

	if next_time <= now {
		return 0
	} else {
		return next_time - now
	}
}

const (
	winWidth     = 940
	winHeight    = 720
	tickInterval = 30
)

type Color struct {
	red   uint8
	green uint8
	blue  uint8
	alpha uint8
}

type Button struct {
	sortField   string
	left        int32
	right       int32
	top         int32
	bottom      int32
	pressed     bool
	shouldSort  bool
	buttonColor Color
}

type ReverseButton struct {
	left        int32
	right       int32
	top         int32
	bottom      int32
	pressed     bool
	shouldSort  bool
	buttonColor Color
}

var (
	next_time uint32

	err      error
	window   *sdl.Window
	renderer *sdl.Renderer

	quit      bool
	event     sdl.Event
	locationX = 0.0
	locationY = 0.0
	step      = 0.5
	deviation = 2.0

	sorted bool = false

	buttonList []Button

	revb = ReverseButton{380, 430, winHeight - 10, winHeight - 40, false, true, Color{100, 100, 100, 255}}
)

func run() int {
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", err)
		return 1
	}
	defer sdl.Quit()

	if window, err = sdl.CreateWindow("colorsort", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 2
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 3 // don't use os.Exit(3); otherwise, previous deferred calls will never run
	}
	renderer.Clear()
	defer renderer.Destroy()

	myList := NewColorList(winWidth / 2)
	lw := winWidth / len(myList.list)

	buttonList = append(buttonList, Button{"red", 100, 150, winHeight - 10, winHeight - 40, false, true, Color{255, 0, 0, 255}})
	buttonList = append(buttonList, Button{"green", 170, 220, winHeight - 10, winHeight - 40, false, true, Color{0, 255, 0, 255}})
	buttonList = append(buttonList, Button{"blue", 240, 290, winHeight - 10, winHeight - 40, false, true, Color{0, 0, 255, 255}})
	buttonList = append(buttonList, Button{"alpha", 310, 360, winHeight - 10, winHeight - 40, false, true, Color{255, 255, 255, 255}})

	quit = false
	for !quit {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.MouseButtonEvent:
				for i, button := range buttonList {
					if t.Button == 1 && t.State == 1 {
						if t.X > button.left && t.X < button.right && t.Y > button.bottom && t.Y < button.top {
							buttonList[i].pressed = true
						} else {
							buttonList[i].pressed = false
						}
						if t.X > revb.left && t.X < revb.right && t.Y > revb.bottom && t.Y < revb.top {
							revb.pressed = true
						} else {
							revb.pressed = false
						}
					} else if t.Button == 1 && t.State == 0 {
						buttonList[i].pressed = false
						buttonList[i].shouldSort = true
						revb.pressed = false
						revb.shouldSort = true
					}
				}
			}
		}

		renderer.Clear()
		gfx.BoxRGBA(renderer, 0, 0, winWidth, winHeight, 0, 0, 0, 255)

		for i := 0; i < len(myList.list); i++ {
			gfx.BoxRGBA(renderer, int32(lw*i), int32(10), int32(lw*i+lw-2), int32(winHeight-50), myList.list[i].red, myList.list[i].green, myList.list[i].blue, myList.list[i].alpha)
		}

		for i, button := range buttonList {
			if button.pressed {
				gfx.BoxRGBA(renderer, button.left, button.bottom, button.right, button.top, button.buttonColor.red, button.buttonColor.green, button.buttonColor.blue, button.buttonColor.alpha)
				if button.shouldSort {
					myList.sortColor = button.sortField
					sort.Sort(myList)
					buttonList[i].shouldSort = false
				}
			} else {
				gfx.RectangleRGBA(renderer, button.left, button.bottom, button.right, button.top, button.buttonColor.red, button.buttonColor.green, button.buttonColor.blue, button.buttonColor.alpha)
			}
		}

		if revb.pressed {
			gfx.BoxRGBA(renderer, revb.left, revb.bottom, revb.right, revb.top, revb.buttonColor.red, revb.buttonColor.green, revb.buttonColor.blue, revb.buttonColor.alpha)
			if revb.shouldSort {
				myList.reverse()
				revb.shouldSort = false
			}
		} else {
			gfx.RectangleRGBA(renderer, revb.left, revb.bottom, revb.right, revb.top, revb.buttonColor.red, revb.buttonColor.green, revb.buttonColor.blue, revb.buttonColor.alpha)
		}

		renderer.Present()
		sdl.Delay(time_left(next_time))
		next_time += tickInterval
	}
	return 1
}

func main() {
	os.Exit(run())
}
