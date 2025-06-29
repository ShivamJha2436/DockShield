package output

import (
	"os"
)

// WriteToFile saves scan report to a file
func WriteToFile(data []byte, path string) error {
	return os.WriteFile(path, data, 0644)
}
