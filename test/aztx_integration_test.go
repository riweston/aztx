package aztx

import (
	"os/exec"
	"testing"
)

func TestAzBinary(t *testing.T) {
	_, err := exec.LookPath("az")

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFzfBinary(t *testing.T) {
	_, err := exec.LookPath("fzf")

	if err != nil {
		t.Errorf(err.Error())
	}
}
