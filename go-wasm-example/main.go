package main

import (
	"fmt"
	"os"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	wasmBytes, _ := os.ReadFile("./hello/pkg/hello_bg.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, _ := wasmer.NewModule(store, wasmBytes)

	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	add, _ := instance.Exports.GetFunction("add")
	result, _ := add(1, 2)

	fmt.Println(result)
}
