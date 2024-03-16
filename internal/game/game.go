package game

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Actor struct {
	W     int32
	H     int32
	Top   float64
	Left  float64
	Color sdl.Color

	OnUpdate func(deltaTime time.Duration)

	Children []*Actor
	Parent   *Actor
}

func NewActor(left float64, top float64, w int32, h int32, color sdl.Color) *Actor {
	return &Actor{
		Left:     left,
		Top:      top,
		W:        w,
		H:        h,
		Color:    color,
		OnUpdate: func(deltaTime time.Duration) {},
		Children: make([]*Actor, 0),
		Parent:   nil,
	}
}

func (actor *Actor) AddChild(child *Actor) {
	child.Parent = actor
	actor.Children = append(actor.Children, child)
}

func (actor *Actor) GetPosition() (left float64, top float64) {
	if actor.Parent != nil {
		pLeft := actor.Parent.Left
		pTop := actor.Parent.Top

		return pLeft + actor.Left, pTop + actor.Top
	}

	return actor.Left, actor.Top
}

func (actor *Actor) Update(deltaTime time.Duration) {
	actor.OnUpdate(deltaTime)
	for _, child := range actor.Children {
		child.Update(deltaTime)
	}
}

func (actor *Actor) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(
		actor.Color.R,
		actor.Color.G,
		actor.Color.B,
		actor.Color.A,
	)

	left, top := actor.GetPosition()

	rect := &sdl.Rect{X: int32(left), Y: int32(top), W: actor.W, H: actor.H}

	renderer.FillRect(rect)

	for _, child := range actor.Children {
		child.Draw(renderer)
	}
}
