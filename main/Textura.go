package main

import (
    "fmt"

    "image"
    "image/draw"
    _ "image/png"


    "github.com/go-gl/gl/v2.1/gl" // OpenGL antiguo


    "os"
)

type Textura struct {
    texturaID uint32
    ancho     int32
    alto      int32
}

func (t *Textura) getIdTextura() uint32 {
    return t.texturaID
}

// Recibe un string con el nombre del fichero gráfico para usarlo como textura devuelve ID de textura
func (t *Textura) nuevaTextura(file string)  {
    ficheroImagen, err := os.Open(file) // Abre fichero
    if err != nil {               // Si hay fallo devolver 0 e información

        fmt.Errorf("textura %q no encontrada en disco: %v", file, err)
    }
    img, _, err := image.Decode(ficheroImagen) // Decodifica imagen png y devuelve una interface rectángulo img
    if err != nil {                      // Devolver 0 y error si falla

        fmt.Errorf("Fallo en Decode", file, err)
    }

    // El método Bounds() es de la interface img que devuelve el tipo Rectangle,
    // NewRGBA devuelve un tipo RGBA de imagen
    rgba := image.NewRGBA(img.Bounds())
    if rgba.Stride != rgba.Rect.Size().X*4 { // Si no coincide stride devolver 0 e información
        fmt.Errorf("stride no soportado")
    }
    // (imagen destino, rectángulo, imágen fuente, coordenada inicial, operación)
    draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

    //var texturaID uint32
    gl.GenTextures(1, &t.texturaID)            // Crea una textura vacía llamada texturaID
    gl.ActiveTexture(gl.TEXTURE0)          // Que será la textura 0
    gl.BindTexture(gl.TEXTURE_2D, t.texturaID) // Enlazar texturaID al punto TEXTURE_2D de OpenGL
    // Parámetros del tipo de textura
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

    t.ancho = (int32)(rgba.Rect.Size().X) // Ancho
    t.alto =  (int32)(rgba.Rect.Size().Y)  // Alto

    gl.TexImage2D( // Especifica una textura de imagen 2D
        gl.TEXTURE_2D,             // Tipo
        0,                         // Nivel de mipmap
        gl.RGBA,                   // Formato RGBA
        t.ancho, // Ancho
        t.alto, // Alto
        0,                // Borde
        gl.RGBA,          // Formato
        gl.UNSIGNED_BYTE, // Tipo de los datos
        gl.Ptr(rgba.Pix)) // Puntero a los datos de imagen en memoria


}

func (t *Textura) enlazaTextura() {
    gl.BindTexture(gl.TEXTURE_2D, t.texturaID)
}