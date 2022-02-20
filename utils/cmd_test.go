package utils_test

import (
	"os/exec"
	"testing"

	"github.com/karmek-k/ssgbuild/utils"
	"github.com/stretchr/testify/assert"
)

func Test_StringToCmdNoArgs(t *testing.T) {
	cmd := utils.StringToCmd("ls")

	assert.Equal(t, exec.Command("ls"), cmd)
}

func Test_StringToCmdWithArgs(t *testing.T) {
	cmd := utils.StringToCmd("npm run build")

	assert.Equal(t, exec.Command("npm", "run", "build"), cmd)
}

func Test_StringToCmdEmptyString(t *testing.T) {
	cmd := utils.StringToCmd("")

	assert.Equal(t, exec.Command(""), cmd)
}

func Test_StringToCmdWithFlags(t *testing.T) {
	actual := utils.StringToCmd("npm install -D prettier eslint")
	expected := exec.Command("npm", "install", "-D", "prettier", "eslint")

	assert.Equal(t, expected, actual)
}