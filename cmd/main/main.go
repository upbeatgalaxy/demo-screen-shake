package main

import (
	"fmt"
	"math"
	"time"

	// "github.com/veandco/go-sdl2/mix"
	sdl "github.com/veandco/go-sdl2/sdl"
	// "github.com/veandco/go-sdl2/ttf"
	"simple-sdl2-project/internal/game"
	mog "simple-sdl2-project/internal/mog"
)

const (
	SCREEN_W = 800
	SCREEN_H = 600
)

var screen *game.Actor

// Shake Variables
var Shaking = false
var shakeCounter = time.Duration(0)
var shakeDuration = time.Duration(0)
var shakeIntensity = 0.0
var shakerSwitch = 1.0
var lastShakeDelta = 0.0

func shake(duration time.Duration, intensity float64) {
	screen.Left = 0
	screen.Top = 0
	Shaking = true
	shakeDuration = duration
	shakeIntensity = intensity
	shakerSwitch = 1
	shakeCounter = time.Duration(0)
}

func updateShake(deltaTime time.Duration) {
	if !Shaking {
		return
	}

	if shakeCounter >= shakeDuration {
		Shaking = false
		shakeIntensity = 0
		shakeCounter = time.Duration(0)
		shakeDuration = time.Duration(0)
		screen.Left = 0
		screen.Top = 0
		lastShakeDelta = 0.0
		return
	}

	shakeCounter += deltaTime

	// Used to normalize every 10 milliseconds
	shakeDelta := math.Trunc(float64(shakeCounter) / float64(time.Millisecond*5))

	if lastShakeDelta < shakeDelta {
		shakeX := shakerSwitch * shakeIntensity
		shakeY := shakerSwitch * shakeIntensity

		screen.Left += shakeX
		screen.Top += shakeY

		// This implementation will shake from -45 to +45
		if screen.Top != 0 && screen.Left != 0 {
			shakerSwitch *= -1
		}

		lastShakeDelta = shakeDelta
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to initialize a window: %v", err)
		panic("Failed to initialize a window")
	}

	window, err := sdl.CreateWindow("Simple SDL2 Project",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		SCREEN_W, SCREEN_H,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)

	if err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to create the SDL Window: %v", err)
		panic("Failed to create the SDL Window")
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(mog.MW, "\nFailed to Initialize renderer: %v", err)
		panic("Failed to Initialize renderer")
	}
	defer renderer.Destroy()
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	/** #region Create Game Entities */
	backgroundColor := sdl.Color{R: 128, G: 128, B: 128, A: 255}
	screen = game.NewActor(0, 0, SCREEN_W, SCREEN_H, backgroundColor)
	screen.OnUpdate = func(deltaTime time.Duration) {
		updateShake(deltaTime)
	}

	// Add some trees
	var treeH int32 = 200
	tree1 := game.NewActor(200, SCREEN_H-float64(treeH), 50, treeH, sdl.Color{R: 0, G: 255, B: 0, A: 255})
	tree2 := &game.Actor{}
	*tree2 = *tree1
	tree2.Left = 400

	screen.AddChild(tree1)
	screen.AddChild(tree2)

	/** #region Create Game Entities */

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				sdl.Quit()
				return
			case *sdl.MouseButtonEvent:
				shake(time.Millisecond*300, 1)
			}
		}
		renderer.SetDrawColor(backgroundColor.R, backgroundColor.G, backgroundColor.B, backgroundColor.A)
		renderer.Clear()

		updateAndDraw(renderer)

		renderer.Present()
	}
}

var lastFrameTime = time.Now()

func updateAndDraw(renderer *sdl.Renderer) {
	// nothing
	now := time.Now()
	deltaTime := now.Sub(lastFrameTime)

	screen.Update(deltaTime)
	screen.Draw(renderer)

	lastFrameTime = now
}
