package dockerutil

import (
	"bytes"
	"os/exec"
	"strings"
)

func ImageExists(image string) bool {
	cmd := exec.Command("docker", "image", "inspect", image)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil && strings.Contains(stderr.String(), "No such image") {
		return false
	}
	return true
}
