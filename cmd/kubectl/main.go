package main

import (
	"log"

	execkubectl "github.com/reece-iriye/kubernetes/internal/executables/kubectl"
)

func main() {
	err := execkubectl.ExecKubectl()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
