package main

import (
	"crypto"
	_ "crypto/sha512"
	"encoding/hex"
	"syscall/js"
)

func main() {
	js.Global().Set("hash", js.FuncOf(hash))
	js.Global().Set("hello", js.FuncOf(hello))
	js.Global().Set("add", js.FuncOf(add))
	select {}
}

func hash(this js.Value, args []js.Value) any {
	h := crypto.SHA512.New()
	h.Write([]byte(args[0].String()))

	return hex.EncodeToString(h.Sum(nil))
}

func hello(this js.Value, args []js.Value) any {
	return "hello " + args[0].String()
}

func add(this js.Value, args []js.Value) any {
	return args[0].Int() + args[1].Int()
}
