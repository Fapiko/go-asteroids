package main

import (
	"github.com/fapiko/go-asteroids/camera"
	"github.com/fapiko/go-learn-gl/opengl-tutorial/common"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type GlRenderable interface {
	Render()
}

type Triangle2D struct {
	vertices                  []float32
	vertexArrayObjectId       uint32
	vertexBufferObjectId      uint32
	programId                 uint32
	vertexPositionAttributeId uint32
	mvMatrixId                int32
	camera                    camera.Camera
}

func NewTriangle2D(x1 float32, y1 float32, x2 float32, y2 float32, x3 float32, y3 float32) *Triangle2D {

	vertices := []float32{
		x1, y1, 0,
		x2, y2, 0,
		x3, y3, 0,
	}

	programId := common.LoadShaders("shaders/SingleColor.vertexshader",
		"shaders/SingleColor.fragmentshader")

	var vertexBufferObjectId uint32
	gl.GenBuffers(1, &vertexBufferObjectId)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObjectId)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	// Populate vertex data
	vertexPositionAttributeId := uint32(gl.GetAttribLocation(programId, gl.Str("vertexPosition_modelspace\x00")))

	// Populate color data
	colorData := []float32{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}

	// Populate color data
	vertexColorAttributeId := uint32(gl.GetAttribLocation(programId, gl.Str("vertexColor\x00")))

	var colorBufferId uint32
	gl.GenBuffers(1, &colorBufferId)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorBufferId)
	gl.BufferData(gl.ARRAY_BUFFER, len(colorData)*4, gl.Ptr(colorData), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	mvMatrixId := gl.GetUniformLocation(programId, gl.Str("MV\x00"))

	// Create vertex array object and do all the things
	var vertexArrayObjectId uint32
	gl.GenVertexArrays(1, &vertexArrayObjectId)
	gl.BindVertexArray(vertexArrayObjectId)
	defer gl.BindVertexArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBufferObjectId)
	gl.EnableVertexArrayAttrib(vertexArrayObjectId, vertexPositionAttributeId)
	gl.VertexAttribPointer(vertexPositionAttributeId, 3, gl.FLOAT, false, 0, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, colorBufferId)
	gl.EnableVertexArrayAttrib(vertexArrayObjectId, vertexColorAttributeId)
	gl.VertexAttribPointer(vertexColorAttributeId, 3, gl.FLOAT, false, 0, nil)

	gl.BindVertexArray(0)

	triangle := &Triangle2D{
		vertices:                  vertices,
		vertexArrayObjectId:       vertexArrayObjectId,
		vertexBufferObjectId:      vertexBufferObjectId,
		vertexPositionAttributeId: vertexPositionAttributeId,
		programId:                 programId,
		mvMatrixId:                mvMatrixId,
	}

	return triangle

}

func (triangle *Triangle2D) Render() {

	gl.UseProgram(triangle.programId)

	mvp := triangle.camera.GetMVP()
	gl.UniformMatrix4fv(triangle.mvMatrixId, 1, false, &mvp[0])

	gl.BindVertexArray(triangle.vertexArrayObjectId)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindVertexArray(0)

}

func (triangle *Triangle2D) Rotate(angle float32) {

	vertices := triangle.vertices

	var newVertices []float32

	radians := mgl32.DegToRad(angle)

	// Start by retrieving our center coords so that we can transform to/from the origin in order to perform the rotation
	centerX := (triangle.vertices[0] + vertices[3] + vertices[6]) / 3.0
	centerY := (triangle.vertices[1] + vertices[4] + vertices[7]) / 3.0

	rotationMatrix := mgl32.HomogRotate3DZ(radians)

	for i := 0; i < len(vertices); i += 3 {

		// Get a vector which will take us to the origin based on the object's center
		originVector := mgl32.Vec3{-centerX, -centerY, 0}

		originTranslationMatrix := mgl32.Translate3D(originVector[0], originVector[1], 0)

		// Form our full transformation matrix:
		// 1. Move triangle back to origin
		// 2. Perform rotation
		// 3. Move triangle back to its original position
		transformationMatrix := originTranslationMatrix.Inv().Mul4(rotationMatrix).Mul4(originTranslationMatrix)

		// Perform transformation
		vertex := transformationMatrix.Mul4x1(mgl32.Vec4{vertices[i], vertices[i+1], 0, 1})
		newVertices = append(newVertices, vertex[0], vertex[1], vertex[2])

	}

	triangle.vertices = newVertices

	gl.BindBuffer(gl.ARRAY_BUFFER, triangle.vertexBufferObjectId)
	gl.BufferData(gl.ARRAY_BUFFER, len(triangle.vertices)*4, gl.Ptr(triangle.vertices), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

}

func (triangle *Triangle2D) Move(vector mgl32.Vec3, amount float32) {

	translationMatrix := mgl32.Translate3D(vector[0], vector[1], 1)

	var newVertices []float32

	vertices := triangle.vertices

	for i := 0; i < len(vertices); i += 3 {

		vertex := translationMatrix.Mul4x1(mgl32.Vec4{vertices[i], vertices[i+1], vertices[i+2], -amount})

		newVertices = append(newVertices, vertex[0], vertex[1], vertex[2])

	}

	triangle.vertices = newVertices

	gl.BindBuffer(gl.ARRAY_BUFFER, triangle.vertexBufferObjectId)
	gl.BufferData(gl.ARRAY_BUFFER, len(triangle.vertices)*4, gl.Ptr(triangle.vertices), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

}
