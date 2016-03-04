package main

import (
	"runtime"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func main() {

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	go renderRoutine(waitGroup)

	waitGroup.Wait()

	// Render ship - Done

	// Move ship

	// Render asteroid

	// Collision detect ship -> asteroid

	// Fire bullets from ship

	// Collision detect bullets -> asteroid

	// Render asteroid split into smaller asteroid

	// Render scoreboard

	// Render lives

}

func (triangle *Triangle2D) Close() {
	gl.DeleteProgram(triangle.programId)
	gl.DeleteVertexArrays(1, &triangle.vertexArrayObjectId)
	gl.DeleteBuffers(1, &triangle.vertexBufferObjectId)
}

func renderRoutine(waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		log.Panic("Failed to initialize GLFW: ", err)
		return
	}
	defer glfw.Terminate()

	window, err := createWindow(1024, 768, "Go Asteroids")
	if err != nil {
		log.Panic(err)
		return
	}

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		panic(err)
	}

	window.SetInputMode(glfw.StickyKeysMode, gl.TRUE)

	gl.ClearColor(0.0, 0.0, 0.0, 0.0)

	ship := NewShip()
	defer ship.Close()

	window.SetKeyCallback(ship.KeyCallback)

	for window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)

		ship.Render()

		window.SwapBuffers()
		glfw.PollEvents()

	}

}

func createWindow(width int, height int, title string) (*glfw.Window, error) {

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	return window, nil

}
