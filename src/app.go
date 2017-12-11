package main

import (
	"fmt"
	"github.com/mmagsino/ingestor"
)

func main() {
	fmt.Println("Begin Diagnostic By Id")
	ingestor.DiagnosticGroupById(2)
	fmt.Println("Begin Diagnostic By Groups")
	ingestor.DiagnosticGroups()
}