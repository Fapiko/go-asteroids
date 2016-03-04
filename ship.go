package main

import (
	"math"

	"github.com/fapiko/go-asteroids/camera"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	RotationSpeed = 4
	MoveSpeed     = .25
)

type Ship struct {
	triangle *Triangle2D
	rotation float32
}

func NewShip() *Ship {

	triangle := NewTriangle2D(-1.5, 0, 1, 1, 1, -1)
	triangle.camera = camera.NewOrtho(50, 50, 50)

	ship := &Ship{
		triangle: triangle,
	}

	window := glfw.GetCurrentContext()
	window.SetKeyCallback(ship.KeyCallback)

	return ship

}

func (ship *Ship) Render() {
	ship.triangle.Render()
}

func (ship *Ship) Close() {
	ship.triangle.Close()
}

func (ship *Ship) KeyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	// TODO: Replace these with some sort of deltaTime rate limiters to ensure we're not basing our speed off keyboard
	// repeat event speed
	if key == glfw.KeyRight && (action == glfw.Repeat || action == glfw.Press) {
		ship.Rotate(-RotationSpeed)
	} else if key == glfw.KeyLeft && (action == glfw.Repeat || action == glfw.Press) {
		ship.Rotate(RotationSpeed)
	} else if key == glfw.KeyUp && (action == glfw.Repeat || action == glfw.Press) {

		// Get a forward vector based on current rotation
		forwardVector := mgl32.Vec3{
			float32(math.Cos(float64(mgl32.DegToRad(ship.rotation)))),
			float32(math.Sin(float64(mgl32.DegToRad(ship.rotation)))),
			1}

		ship.triangle.Move(forwardVector, MoveSpeed)

	}

}

func (ship *Ship) Rotate(angle float32) {
	ship.rotation += angle
	ship.triangle.Rotate(angle)
}
