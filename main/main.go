package main

import (
    "fmt"
    "log"
    "runtime"
    "github.com/go-gl/gl/v2.1/gl" // OpenGL antiguo
    "github.com/go-gl/glfw/v3.2/glfw"
)

const (
    tituloVentana = "03_Golang, usando entradas"
    anchoVentana = 640
    altoVentana = 480
)

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

    window.SetKeyCallback(onTecla)
    window.SetMouseButtonCallback(onRaton)

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

// **************** TECLADO *****************************
func onTecla(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

    fmt.Print("------------->onTecla01\n")

    if key == glfw.KeyA && action == glfw.Press {
        //w.SetShouldClose(true)
        fmt.Print("aaaaaaaaaa\n")
    }

    if key == glfw.KeyY && action == glfw.Press {
        //w.SetShouldClose(true)
        fmt.Print("yyyyyyyyy\n")
    }
}

// ***************** RATON *******************************
func onRaton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
    fmt.Print("------------->onRaton01\n")
    if button == glfw.MouseButtonLeft {
        fmt.Print("glfw.BUTTON_LEFT\n")
    }

    if button == glfw.MouseButtonRight {
        fmt.Print("glfw.BUTTON_RIGHT\n")
    }

    if button == glfw.MouseButtonMiddle {
        fmt.Print("glfw.BUTTON_MIDDL\n")
    }

    if action == glfw.Press {
        fmt.Print("glfw.MOUSEBUTTONDOWN\n")
    }

    if action == glfw.Release {
        fmt.Print("glfw.MOUSEBUTTONUP\n")
    }
}
