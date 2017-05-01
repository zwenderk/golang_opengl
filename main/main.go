package main

import (
    "fmt"
    "image"
    "image/draw"
    _ "image/png"
    "log"
    "runtime"
    "github.com/go-gl/gl/v2.1/gl" // OpenGL antiguo
    "github.com/go-gl/glfw/v3.2/glfw"
    "os"
)

const (
    tituloVentana = "04_Golang, usando texturas"
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
    textura, err := nuevaTextura("Barcelona.png") // Carga textura

    if err != nil {
        fmt.Print("Error con textura\n")
    }

    // -------------> BUCLE PRINCIPAL
    for !window.ShouldClose() {
        enlazarTextura(textura)
        dibuja()

        // Mantenimiento
        window.SwapBuffers() // Intercambia buffers para presenter en pantalla
        glfw.PollEvents()
    } // -----------> FIN DE BUCLE PRINCIPAL
}

func dibuja() {
    gl.Clear(gl.COLOR_BUFFER_BIT)

    gl.Begin(gl.QUADS) // Cada coordenada con su color

    gl.TexCoord2f(0, 0)
    gl.Vertex2f(-0.5, 0.5)

    gl.TexCoord2f(1, 0)
    gl.Vertex2f(0.5, 0.5)

    gl.TexCoord2f(1, 1)
    gl.Vertex2f(0.5, -0.5)

    gl.TexCoord2f(0, 1)
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

func enlazarTextura(idTextura uint32 ) {
    gl.BindTexture(gl.TEXTURE_2D, idTextura)
}

// Recibe un string con el nombre del fichero gráfico para usarlo como textura devuelve ID de textura
func nuevaTextura(file string) (uint32, error) {
    ficheroImagen, err := os.Open(file) // Abre fichero
    if err != nil {               // Si hay fallo devolver 0 e información
        return 0, fmt.Errorf("textura %q no encontrada en disco: %v", file, err)
    }
    img, _, err := image.Decode(ficheroImagen) // Decodifica imagen png y devuelve una interface rectángulo img
    if err != nil {                      // Devolver 0 y error si falla
        return 0, err
    }

    // El método Bounds() es de la interface img que devuelve el tipo Rectangle,
    // NewRGBA devuelve un tipo RGBA de imagen
    rgba := image.NewRGBA(img.Bounds())
    if rgba.Stride != rgba.Rect.Size().X*4 { // Si no coincide stride devolver 0 e información
        fmt.Errorf("stride no soportado")
    }
    // (imagen destino, rectángulo, imágen fuente, coordenada inicial, operación)
    draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

    var texturaID uint32
    gl.GenTextures(1, &texturaID)            // Crea una textura vacía llamada texturaID
    gl.ActiveTexture(gl.TEXTURE0)          // Que será la textura 0
    gl.BindTexture(gl.TEXTURE_2D, texturaID) // Enlazar texturaID al punto TEXTURE_2D de OpenGL
    // Parámetros del tipo de textura
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
    gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
    gl.TexImage2D( // Especifica una textura de imagen 2D
        gl.TEXTURE_2D,             // Tipo
        0,                         // Nivel de mipmap
        gl.RGBA,                   // Formato RGBA
        int32(rgba.Rect.Size().X), // Ancho
        int32(rgba.Rect.Size().Y), // Alto
        0,                // Borde
        gl.RGBA,          // Formato
        gl.UNSIGNED_BYTE, // Tipo de los datos
        gl.Ptr(rgba.Pix)) // Puntero a los datos de imagen en memoria

    return texturaID,	nil
}
