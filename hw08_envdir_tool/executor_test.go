package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("command success", func(t *testing.T) {
		cmd := []string{
			"/bin/bash",
			"testdata/echo.sh",
		}

		code := RunCmd(cmd, nil)

		require.Equal(t, code, 0)
	})

	t.Run("command not found", func(t *testing.T) {
		cmd := []string{
			"/bin/bash",
			"testdata/echo2.sh",
		}

		code := RunCmd(cmd, nil)

		require.Equal(t, code, 127)
	})
}
