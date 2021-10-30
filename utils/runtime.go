package utils

import (
	"fmt"
	"runtime"
	"strings"
)

type Runtime struct {
	Packages []string
}

func (rt *Runtime) NotImplemented() error {
	frame := rt.GetFrame(3)
	fileName := rt.NormalizeFileName(frame.File)
	function := rt.NormalizeFunction(frame.Function)

	return fmt.Errorf("NOT IMPLEMENTED: %s:%d %s", fileName, frame.Line, function)
}

func (rt *Runtime) Errorf(format string, a ...interface{}) error {
	frame := rt.GetFrame(3)
	fileName := rt.NormalizeFileName(frame.File)
	function := rt.NormalizeFunction(frame.Function)

	s := fmt.Sprintf(format, a...)

	return fmt.Errorf("%s %s:%d - %s", function, fileName, frame.Line, s)
}

func (rt *Runtime) GetFrame(depth int) runtime.Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(depth, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return frame
}

func (rt *Runtime) NormalizeFileName(fileName string) string {
	results := fileName

	for _, pkgName := range rt.Packages {
		idx := strings.Index(results, pkgName)
		if idx != -1 {
			results = results[idx:]
		}
	}

	return results
}

func (rt *Runtime) NormalizeFunction(functionName string) string {
	results := rt.NormalizeFileName(functionName)
	idx := strings.Index(results, ".")
	if idx != -1 {
		results = results[idx+1:]
	}
	return results
}
