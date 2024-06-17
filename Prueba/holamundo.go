package main

import (
	"fmt"
)

func init() {
	// HOLA MUNDO -----------------------
	var numero int
	numero = 25
	fmt.Println(numero)
	numero = 40
	fmt.Println(numero)
	//Asignacion logica de go con el tipo de datos que se le asigana go define que es string
	//al ves que esta en comillas y los := es para que ponga el tipo automaticamente
	nombre := "Julian"
	fmt.Println(nombre)
	//Asignacion de varias variables al tiempo
	nombre, numero = "Jose", 30
	fmt.Println(nombre, numero)

	// CONVERTIR -------
	var entero32 = 25
	var entero64 = 85
	fmt.Println(entero32, int32(entero64))

	//SI tienes variables que definis y no usas go revienta para que las uses o las quites
	// var note string

	// Typos -----------------
	var entero8 int8   // (-128 to 127 )
	var entero16 int16 // (-32768 to  32767)
	var entero8 int32  // (-2147483648 to 2147483647)
	var entero32 int64 // ( - 9223372036854775808 to  9223372036854775807)

	var uentero8 uint8   // (0 to 127 )
	var uentero16 uint16 // (0 to  32767)
	var uentero8 uint32  // (0 to 2147483647)
	var uentero32 uint64 // (0 to  9223372036854775807)

	var enteroByte byte // Alias para uint8
	var enteroRune rune // alias para int32

	var enteroUint uint // 32 a 64 bits
	var enteroInt int   //32 o 64 bits

	var enteroUintptr uintptr //Entero sin signo lo suficiente mente grante para contener un puntero
}
