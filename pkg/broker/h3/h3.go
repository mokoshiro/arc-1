package h3

import (
	"fmt"
	"os/exec"
	"strings"
)

func InvokeH3(resolution int, longitude, latitude float64) (string, error) {
	args := []string{
		"--resolution", fmt.Sprintf("%d", resolution),
		"--longitude", fmt.Sprintf("%f", longitude),
		"--latitude", fmt.Sprintf("%f", latitude),
	}
	command := exec.Command("geoToH3", args...)
	b, err := command.CombinedOutput()
	if err != nil {
		return "", err
	}
	result := strings.TrimRight(string(b), "\n")
	return result, err
}
