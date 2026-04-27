package helm

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFile(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "prometheus-0.0.0.tgz")
	assert.NoError(t, os.WriteFile(filePath, []byte("archive contents"), 0o644))

	err := removeFile(filePath)
	assert.NoError(t, err)

	_, err = os.Stat(filePath)
	assert.True(t, os.IsNotExist(err))
}

func TestRemoveFileNotExist(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "prometheus-0.0.0.tgz")

	err := removeFile(filePath)
	assert.Error(t, err)
	assert.True(t, os.IsNotExist(err))
}
