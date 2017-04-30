package main

import (
    "fmt"
    "log"
    "runtime"
    "github.com/go-gl/gl/v2.1/gl" // OpenGL antiguo
    "github.com/go-gl/glfw/v3.2/glfw"
)

const (
    tituloVentana = "02_Golang, un cuadrilátero"
    anchoVentana = 640
    altoVentana = 480
)

//*********************************************************************************************************
// Función que se llama después de haber evaluado todas las declaraciones de variables del paquete y antes que main

//func init() {
// manejo de eventos GLFW se debe ejecutar en el hilo principal
//    runtime.LockOSThread()
//}

//*********************************************************************************************************

func main() {

    runtime.LockOSThread()

    if err := glfw.Init(); err != nil {
        // Inicializa GLFW
        log.Fatalln("Fallo al inicializar glfw:", err) // Si hay error informar
    }
    defer glfw.Terminate() // Al acabar el programa cerrar GLFW

    // Valores que tendrá la ventana WindowHint(variable, valor que tendrá esa variable)
    glfw.WindowHint(glfw.Resizable, glfw.False)                 // No cambia de tamaño
    glfw.WindowHint(glfw.ContextVersionMajor, 2)                // Versión mayor permitida
    glfw.WindowHint(glfw.ContextVersionMinor, 1)                // Versión menor permitida
    //glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // Usar el profile OpenGL Core
    //glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(// Crear ventana
        anchoVentana,
        altoVentana,
        tituloVentana,
        nil,
        nil)
    if err != nil {
        panic(err)
    }
    window.MakeContextCurrent()

    // Inicializar Glow
    if err := gl.Init(); err != nil {
        panic(err)
    }

    version := gl.GoStr(gl.GetString(gl.VERSION))
    fmt.Println("OpenGL versión", version) // Imprime versión de OpenGL del sistema

    gl.ClearColor(.5, 1, 0, 0.0) // Especifica valores de color de limpieza

    // -------------> BUCLE PRINCIPAL
    for !window.ShouldClose() {

        dibuja()

        // Mantenimiento
        window.SwapBuffers() // Intercambia buffers para presenter en pantalla
        glfw.PollEvents()
    } // -----------> FIN DE BUCLE PRINCIPAL
}

func dibuja() {
    gl.Clear(gl.COLOR_BUFFER_BIT)

    gl.Begin(gl.QUADS) // Cada coordenada con su color

    gl.Color4f(1, 0, 0, 0)
    gl.Vertex2f(-0.5, 0.5)

    gl.Color4f(0, 1, 0, 0)
    gl.Vertex2f(0.5, 0.5)

    gl.Color4f(0, 0, 1, 0)
    gl.Vertex2f(0.5, -0.5)

    gl.Color4f(1, 1, 1, 0)
    gl.Vertex2f(-0.5, -0.5)

    gl.End()
}
