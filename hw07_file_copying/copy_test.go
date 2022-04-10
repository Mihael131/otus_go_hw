package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	fromPath := "testdata/input.txt"
	toPath := "out.txt"
	defer os.Remove(toPath)
	t.Run("offset 0 limit 0", func(t *testing.T) {
		_ = Copy(fromPath, toPath, 0, 0)

		template, _ := ioutil.ReadFile("testdata/out_offset0_limit0.txt")
		out, _ := ioutil.ReadFile(toPath)

		require.Equal(t, string(template), string(out))
	})

	t.Run("offset 100 limit 1000", func(t *testing.T) {
		_ = Copy(fromPath, toPath, 100, 1000)

		template, _ := ioutil.ReadFile("testdata/out_offset100_limit1000.txt")
		out, _ := ioutil.ReadFile(toPath)

		require.Equal(t, string(template), string(out))
	})

	t.Run("offset > limit", func(t *testing.T) {
		err := Copy(fromPath, toPath, 7000, 0)

		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", toPath, 0, 0)

		require.Equal(t, ErrUnsupportedFile, err)
	})
}
