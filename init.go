package sf

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	glfw "github.com/go-gl/glfw3/v3.1/glfw"
)

func init() {
	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	gl.Init()
}
