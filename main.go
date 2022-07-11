package main

import (
	"github.com/jplanckeel/scope/internal"
)

func main() {

	internal.Sync("helm3", "example.yaml", "98545421.ecr.aws.com")
}
