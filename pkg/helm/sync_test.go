package helm

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFile(t *testing.T) {

	// File path to create
	var filePath string = "prometheus-0.0.0.tgz"

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()

	err = removeFile(filePath)
	assert.NoError(t, err)
}

func TestRemoveFileError(t *testing.T) {

	// File path
	var filePath string = "prometheus-0.0.0.tgz"

	err := removeFile(filePath)
	assert.Error(t, err)
}
