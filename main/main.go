package main

/*
Modificamos ficheros Shader.go, shader.vs.glsl y shader.fs.glsl ,Textura.go y Model.go
para el uso de variables uniform en los shaders
*/

import (
    "fmt"
    "log"
    "runtime"
    "github.com/go-gl/gl/v2.1/gl" // OpenGL antiguo
    "github.com/go-gl/glfw/v3.2/glfw"
)

var vertices = []float32  {
    -0.5,  0.5, 0, // ARRIBA IZQUIERDA     0
    0.5,  0.5, 0, // ARRIBA A LA DERECHA  1
    0.5, -0.5, 0, // ABAJO A LA DERECHA   2
    -0.5, -0.5, 0, // ABAJO A LA IZQUIERDA 3
}

var texturaCoords = []float32 {
    0, 0,
    1, 0,
    1, 1,
    0, 1,
}

var indices = []int32{
    0, 1, 2, // Primer triángulo
    2, 3, 0, // Segundo triángulo
}

const (
    tituloVentana = "08_Golang, usando variables uniform"
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

    window.SetKeyCallback(onTecla) // Cuando se presione una tecla llamar a la función onTecla()
    window.SetMouseButtonCallback(onRaton) // Cuando se use el ratón llamar a la función onRaton()

    window.MakeContextCurrent()

    // Inicializar Glow
    if err := gl.Init(); err != nil {
        panic(err)
    }

    version := gl.GoStr(gl.GetString(gl.VERSION))
    fmt.Println("OpenGL versión", version) // Imprime versión de OpenGL del sistema

    gl.ClearColor(.5, 1, 0, 0.0) // Especifica valores de color de limpieza

    gl.Enable(gl.TEXTURE_2D) // Habilitamos el uso de texturas
    // enable alpha support (prueba)
    gl.Enable(gl.BLEND )
    gl.BlendFunc( gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA )

    modelo := &Modelo{}
    modelo.Inicializar(vertices, texturaCoords)

    shader := &Shader{} // Creamos objeto Shader
    shader.Inicializar("shader")

    textura := &Textura{}                 // Creamos objeto Texture
    textura.nuevaTextura("Barcelona.png") // Carga textura

    if err != nil {
        fmt.Print("Error con textura\n")
    }

    // -------------> BUCLE PRINCIPAL
    for !window.ShouldClose() {
        textura.enlazaTextura()

        shader.enlazar()

        shader.setUniform("sampler", 0); // Pone valor 0 a la variable uniform "sampler"

        modelo.dibujar(vertices)

        // Mantenimiento
        window.SwapBuffers() // Intercambia buffers para presenter en pantalla

        glfw.PollEvents()
    } // -----------> FIN DE BUCLE PRINCIPAL
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

func enlazarTextura(idTextura uint32 ) {
    gl.BindTexture(gl.TEXTURE_2D, idTextura)
}