//+build js,wasm

package main

import (
	"strconv"

	"syscall/js"
)

func init() {
	// we have to declare our functions in an init func otherwise they aren't
	// available in JS land at the call time.
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))

}

//func add(this js.Value, args []js.Value) interface{} {
//}

func add(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
	return nil
}

func subtract(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)

	return nil
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")

	<-c
}
