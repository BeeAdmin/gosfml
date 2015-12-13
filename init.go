package sf

import (
	"github.com/go-gl/gl"
	glfw "/home/malek/Go/src/github.com/go-gl/glfw3/v3.1/glfw"
)

func init() {
	if !glfw.Init() {
		panic("Can't init glfw!")
	}
	gl.Init()
}
