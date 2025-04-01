package execute

import "os/exec"

// Command executes a command and returns the output
func Command(inputs ...string) (string, error) {
	cmd := exec.Command(inputs[0], inputs[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
