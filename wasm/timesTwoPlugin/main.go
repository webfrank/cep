package main

import (
	"encoding/binary"
	"math"

	"github.com/extism/go-pdk"
)

//go:export run
func Run() int32 {
	bits := binary.LittleEndian.Uint32(pdk.Input())
	value := math.Float32frombits(bits)

	floatBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(floatBuffer, math.Float32bits(value*2.0))

	pdk.Output(floatBuffer)
	return 0
}

func main() {}
