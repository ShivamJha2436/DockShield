package dockerutil

import (
	"os/exec"
)

// ImageExists checks if the Docker image is present locally
func ImageExists(image string) bool {
	cmd := exec.Command("docker", "image", "inspect", image)
	err := cmd.Run()
	return err == nil
}
