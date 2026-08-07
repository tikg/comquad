package deps

import (
	"fmt"
	"os/exec"
	"strings"
)

// EnsureRequirements verifies that the required external tools
// are available in PATH.
func EnsureRequirements() error {
	required := []string{
		"podman",
		"podlet",
	}

	var missing []string

	for _, tool := range required {
		if _, err := exec.LookPath(tool); err != nil {
			missing = append(missing, tool)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf(
			"missing required tools: %s",
			strings.Join(missing, ", "),
		)
	}

	return nil
}