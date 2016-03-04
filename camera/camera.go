package camera

import "github.com/go-gl/mathgl/mgl32"

type Camera interface {
	GetMVP() mgl32.Mat4
}

type OrthoCamera struct {
	modelViewMatrix mgl32.Mat4
}

func NewOrtho() *OrthoCamera {

	projection := mgl32.Ortho(-40, 40, -40, 40, 0.1, 100)

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
