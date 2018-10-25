package main

import (
	"fmt"

	"github.com/kingreatwill/go_study/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
