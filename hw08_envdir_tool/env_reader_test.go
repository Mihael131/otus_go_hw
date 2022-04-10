package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	fromDir := "testdata/env/"

	t.Run("dir not found", func(t *testing.T) {
		_, err := ReadDir(fromDir + "xxx/")
		require.Equal(t, err.Error(), "open testdata/env/xxx/: no such file or directory")
	})

	t.Run("env", func(t *testing.T) {
		env, err := ReadDir(fromDir)

		require.NoError(t, err)

		require.Equal(t, env["BAR"].Value, "bar")
		require.False(t, env["BAR"].NeedRemove)

		require.Empty(t, env["EMPTY"].Value)
		require.True(t, env["EMPTY"].NeedRemove)

		require.Equal(t, env["FOO"].Value, "   foo\nwith new line")
		require.False(t, env["FOO"].NeedRemove)

		require.Equal(t, env["HELLO"].Value, "\"hello\"")
		require.False(t, env["HELLO"].NeedRemove)

		require.Empty(t, env["UNSET"].Value)
		require.True(t, env["UNSET"].NeedRemove)

		_, ok := env["SKIP="]
		require.False(t, ok)
	})
}
