package main

import (
    "github.com/go-gl/gl/v2.1/gl" // OpenGL antiguo
)

type Modelo struct {
    contadorDib int32
    v_id uint32 // ID de vertices
    t_id uint32 // ID de textura
    i_id uint32 // ID de índices
}

func (m *Modelo) Inicializar(vertices []float32, tex_coords []float32) {

    m.contadorDib = (int32) (len(indices))

    gl.GenBuffers(1, &m.v_id); // Crea un VBO para coordenadas de vértice
    gl.BindBuffer(gl.ARRAY_BUFFER, m.v_id); // Conecta VBO de posición
    gl.BufferData(gl.ARRAY_BUFFER, len (vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW); // Conecta datos

    gl.GenBuffers(1, &m.t_id); // Crea un VBO para coordenadas de textura
    gl.BindBuffer(gl.ARRAY_BUFFER, m.t_id); // Conecta VBO de textura
    gl.BufferData(gl.ARRAY_BUFFER, len(tex_coords)*4, gl.Ptr(tex_coords), gl.STATIC_DRAW); // Conecta datos

    gl.GenBuffers(1, &m.i_id); // Crea un VBO para índices
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.i_id); // Conecta VBO de índices
    gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0); // Desconecta
    gl.BindBuffer(gl.ARRAY_BUFFER, 0); // Desconecta
}

func (m *Modelo) dibujar(vertices []float32) {

    gl.ClearColor(.5, 1, 0, 0.0) // Especifica valores de color de limpieza
    gl.Clear(gl.COLOR_BUFFER_BIT)

    gl.EnableClientState(gl.VERTEX_ARRAY); // Activa VERTEX_ARRAY
    gl.EnableClientState(gl.TEXTURE_COORD_ARRAY); // Activa TEXTURE_COORD_ARRAY

    gl.BindBuffer(gl.ARRAY_BUFFER, m.v_id); // Conecta VBO de posición
    gl.VertexPointer(3, gl.FLOAT, 0, nil); // 3 = x,y,z

    gl.BindBuffer(gl.ARRAY_BUFFER, m.t_id); // Conecta VBO de textura
    gl.TexCoordPointer(2, gl.FLOAT, 0, nil); // 2 = u,v

    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.i_id); // Conecta VBO de índices

    gl.DrawElements(gl.TRIANGLES, m.contadorDib, gl.UNSIGNED_INT, nil); // Ahora son índices)

    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0) // Desconecta
    gl.BindBuffer(gl.ARRAY_BUFFER, 0); // Desconecta

    gl.DisableClientState(gl.VERTEX_ARRAY); // Desactiva VERTEX_ARRAY
    gl.DisableClientState(gl.TEXTURE_COORD_ARRAY); // Desactiva TEXTURE_COORD_ARRAY
}
