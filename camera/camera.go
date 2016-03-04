package camera

import "github.com/go-gl/mathgl/mgl32"

type Camera interface {
	GetMVP() mgl32.Mat4
}

type OrthoCamera struct {
	modelViewMatrix mgl32.Mat4
}

func NewOrtho(width int, height int, depth int) *OrthoCamera {

	projection := mgl32.Ortho(
		float32(-(width / 2)), float32(width/2),
		float32(-(height / 2)), float32(height/2),
		0.1, float32(depth))

	view := mgl32.LookAt(0, 0, 3, 0, 0, 0, 0, 1, 0)

	model := mgl32.Ident4()

	mvp := projection.Mul4(view).Mul4(model)

	camera := &OrthoCamera{
		modelViewMatrix: mvp,
	}

	return camera

}

func (camera *OrthoCamera) GetMVP() mgl32.Mat4 {
	return camera.modelViewMatrix
}
