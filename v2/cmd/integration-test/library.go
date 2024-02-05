package main

import (
	"os"

	"github.com/minofmax/naabu/v2/internal/testutils"
	"github.com/minofmax/naabu/v2/pkg/result"
	"github.com/minofmax/naabu/v2/pkg/runner"
)

var libraryTestcases = map[string]testutils.TestCase{
	"naabu as library": &httpxLibrary{},
}

type httpxLibrary struct {
}

func (h *httpxLibrary) Execute() error {
	testFile := "test.txt"
	err := os.WriteFile(testFile, []byte("scanme.sh"), 0644)
	if err != nil {
		return err
	}
	defer os.RemoveAll(testFile)

	options := runner.Options{
		HostsFile: testFile,
		Ports:     "80",
		Passive:   true,
		OnResult:  func(hr *result.HostResult) {},
	}

	naabuRunner, err := runner.NewRunner(&options)
	if err != nil {
		return err
	}
	defer naabuRunner.Close()

	errorMsg, _ := naabuRunner.RunEnumeration()
	return errorMsg
}
