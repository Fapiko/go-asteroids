package main

import (
	"runtime"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/fapiko/go-learn-gl/opengl-tutorial/common"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

func main() {

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	go renderRoutine(waitGroup)

	waitGroup.Wait()

	// Render ship

	// Render asteroid

	// Move ship

	// Collision detect ship -> asteroid

	// Fire bullets from ship

	// Collision detect bullets -> asteroid

	// Render asteroid split into smaller asteroid

	// Render scoreboard

	// Render lives

}

type GlRenderable interface {
	Render()
}

type Triangle2D struct {
	vertices                  []float32
	vertexArrayObjectId       uint32
	vertexBufferObjectId      uint32
	programId                 uint32
	vertexPositionAttributeId uint32
}

func NewTriangle2D(x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32) *Triangle2D {

	vertices := []float32{
		x1, y1, 0,
		x2, y2, 0,
		x3, y3, 0,
	}

	programId := common.LoadShaders("shaders/SimpleVertexShader.vertexshader",
		"shaders/SimpleFragmentShader.fragmentshader")

	var vertexArrayObjectId uint32
	gl.GenVertexArrays(1, &vertexArrayObjectId)
	gl.BindVertexArray(vertexArrayObjectId)
	defer gl.BindVertexArray(0)

	var vertexBufferObjectId uint32
	gl.GenBuffers(1, &vertexBufferObjectId)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObjectId)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	vertexPositionAttributeId := uint32(gl.GetAttribLocation(programId, gl.Str("vertexPosition_modelspace\x00")))
	gl.EnableVertexArrayAttrib(vertexArrayObjectId, vertexPositionAttributeId)

	triangle := &Triangle2D{
		vertices:                  vertices,
		vertexArrayObjectId:       vertexArrayObjectId,
		vertexBufferObjectId:      vertexBufferObjectId,
		vertexPositionAttributeId: vertexPositionAttributeId,
		programId:                 programId,
	}

	return triangle

}

func (triangle *Triangle2D) Render() {

	gl.UseProgram(triangle.programId)
	gl.BindVertexArray(triangle.vertexArrayObjectId)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindVertexArray(0)

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

	triangle := NewTriangle2D(-1, -1, 1, -1, 0, 1)

	for window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.UseProgram(triangle.programId)

		triangle.Render()

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
