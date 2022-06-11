package main

import (
	"Demo/controller"
	"embed"
	"fmt"

	goweb "github.com/cjie9759/goWeb"
)

//go:embed public/*
var FS embed.FS

func main() {
	fmt.Println(goweb.NewApp(&FS).Get(&controller.Index{}).SetMiddle(goweb.MWLog).Run(":8080"))
}
