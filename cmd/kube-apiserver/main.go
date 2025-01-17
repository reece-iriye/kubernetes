package main

import (
	"log"

	execkubeapiserver "github.com/reece-iriye/kubernetes/internal/executables/kube-apiserver"
)

func main() {
	err := execkubeapiserver.ExecKubeApiServer()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
