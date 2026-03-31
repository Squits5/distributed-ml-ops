package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Squits5/distributed-ml-ops/pkg/orchestrator"
)

func main() {
	fmt.Println("Starting Distributed MLOps Engine...")
	orchestrator.Init()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
