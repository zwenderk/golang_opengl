package main

import (
    _ "image/png"
      "github.com/go-gl/mathgl/mgl32"
)

type Camara struct {
    position mgl32.Vec3
    projection mgl32.Mat4
}

func (c *Camara) Inicializar(width int, height int) {
    c.position = mgl32.Vec3{0,0,0}

    c.projection = mgl32.Ortho2D((float32)(-width /2),(float32)(width /2),
        (float32)(-height/2),(float32)(height/2))
    //imprimeMatriz("inicializar Camera",c.projection)
}

func (c *Camara) setPosition(position mgl32.Vec3) {
    c.position = position
}

func (c *Camara) addPosition(position mgl32.Vec3) {
    c.position.Add(position)
}

func (c *Camara) getPosition() (mgl32.Vec3) {
    return c.position
}

func (c *Camara) getProjection() (mgl32.Mat4) {
    target := mgl32.Ident4()
    pos := mgl32.Translate3D(c.position.X(),c.position.Y(),c.position.Z())
    //imprimeMatriz("pos en getProjection",pos)

    target = c.projection.Mul4(pos)
    //imprimeMatriz("target en getProjection",target)
    return target
}


