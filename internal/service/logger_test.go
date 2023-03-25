package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigureCallerFormatter(t *testing.T) {
	testCases := []struct {
		res  string
		file string
		line int
	}{
		{"   main.go:22", "main.go", 22},
		{"  main.go:222", "main.go", 222},
		{" main.go:2222", "main.go", 2222},
		{" main.go:222*", "main.go", 22222},
		{" storage.go:2", "storage.go", 2},
		{"storage.go:22", "storage.go", 22},
		{"storage.*:222", "storage.go", 222},
		{"storage*:2222", "storage.go", 2222},
		{"storage*:222*", "storage.go", 22222},
	}

	for _, testCase := range testCases {
		res := configureCaller(0, testCase.file, testCase.line)
		require.Equal(t, testCase.res, res)
	}
}
