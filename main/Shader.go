package main

import (
    "fmt"
    "github.com/go-gl/gl/v2.1/gl" // OpenGL antiguo
    "io/ioutil"
)

type Shader struct {
    idPrograma uint32
    idVS       uint32
    idFS       uint32
}

func (s *Shader) Inicializar(filename string) {
    var err error

    s.idPrograma = gl.CreateProgram()

    // VERTEX SHADER *********************************************************
    s.idVS = gl.CreateShader(gl.VERTEX_SHADER) // Crea y lee VS

    verTxt, err := ioutil.ReadFile("./shaders/" + filename + ".vs.glsl") // Lee Vertex Shader
    if err != nil {
        fmt.Printf("Error leyendo Vertex Shader ----->\n")
    }
    vsources, free := gl.Strs(string(verTxt)+ "\x00") // verTxt es tipo byte

    gl.ShaderSource(s.idVS, 1, vsources, nil)
    free()
    gl.CompileShader(s.idVS)
    var estado int32
    // Prueba errores de compilación de Vertex Shader
    gl.GetShaderiv(s.idVS, gl.COMPILE_STATUS, &estado)
    if estado == gl.FALSE {
        fmt.Printf("Error en compilación de Vertex Shader ----->\n")
        var logLength int32
        gl.GetShaderiv(s.idVS, gl.INFO_LOG_LENGTH, &logLength)
    }

    // FRAGMENT SHADER *********************************************************
    s.idFS = gl.CreateShader(gl.FRAGMENT_SHADER) // Crea y lee FS

    fraTxt, err := ioutil.ReadFile("./shaders/" + filename + ".fs.glsl") // Lee Fragment Shader
    if err != nil {
        fmt.Printf("Error leyendo Fragment Shader ----->\n")
    }
    //fsources := string(fraTxt) + "\x00" // fraTxt es tipo byte
    fsources, free := gl.Strs(string(fraTxt)+ "\x00") // fraTxt es tipo byte

    gl.ShaderSource(s.idFS, 1, fsources, nil)
    free()
    gl.CompileShader(s.idFS)
    /// Prueba errores de compilación de Fragment Shader
    gl.GetShaderiv(s.idFS, gl.COMPILE_STATUS, &estado)
    if estado == gl.FALSE {
        fmt.Printf("Error en compilación de Fragment Shader ----->\n")
        var logLength int32
        gl.GetShaderiv(s.idFS, gl.INFO_LOG_LENGTH, &logLength)
    }

    // PROGRAM SHADER ***********************************************************
    gl.AttachShader(s.idPrograma, s.idVS) // Enlaza VS
    gl.AttachShader(s.idPrograma, s.idFS) // Enlaza FS

    gl.BindAttribLocation(s.idPrograma, 0, gl.Str("vertices\x00")) // Localización 0 a "vertices" en programa shader

    gl.LinkProgram(s.idPrograma) // "linka" el programa
    // Prueba errores de "linkado"
    gl.GetProgramiv(s.idPrograma, gl.LINK_STATUS, &estado)
    if estado == gl.FALSE {
        var logLength int32
        gl.GetProgramiv(s.idPrograma, gl.INFO_LOG_LENGTH, &logLength)

        fmt.Println("Error en linkado de programa Shader\n")
    }

    gl.DeleteShader(s.idVS)
    gl.DeleteShader(s.idFS)
}

// Dar valor a una variable uniform del programa shader
func (s *Shader) setUniform(name string, value int32) {
    location:=gl.GetUniformLocation(s.idPrograma, gl.Str(name + "\x00"))
    if location != -1 { // Si existe ese nombre de variable
        gl.Uniform1i(location, value)
    }
}

func (s *Shader) enlazar() {
    gl.UseProgram(s.idPrograma)
}
