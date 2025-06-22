package output

import (
	"os"
)

func WriteToFile(data []byte, filename string) error {
	return os.WriteFile(filename, data, 0644)
}
