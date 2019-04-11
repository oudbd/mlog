package mlog

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestTrace(t *testing.T) {
	Start(LevelTrace, "")

	Trace("trace log")
	Info("info log")
	Warning("warning log")

	err := errors.New("error log")
	Error(err)

	// Fatalf("fatalf log")
	if err = Stop(); err != nil {
		fmt.Println(err)
	}
}

func TestInfo(t *testing.T) {
	Start(LevelInfo, "")

	Trace("trace log")
	Info("info log")
	Warning("warning log")

	err := errors.New("error log")
	Error(err)

	// Fatalf("fatalf log")
	if err = Stop(); err != nil {
		fmt.Println(err)
	}
}

func TestWarning(t *testing.T) {
	Start(LevelWarn, "")

	Trace("trace log")
	Info("info log")
	Warning("warning log")

	err := errors.New("error log")
	Error(err)

	// Fatalf("fatalf log")
	if err = Stop(); err != nil {
		fmt.Println(err)
	}
}

func TestError(t *testing.T) {
	Start(LevelError, "")

	Trace("trace log")
	Info("info log")
	Warning("warning log")

	err := errors.New("error log")
	Error(err)

	// Fatalf("fatalf log")
	if err = Stop(); err != nil {
		fmt.Println(err)
	}
}

func TestStartEx(t *testing.T) {
	path := "./test"
	os.RemoveAll(path)

	if err := os.Mkdir(path, 0777); err != nil {
		fmt.Println(err)
	}
	fileName := path + "/startex"

	StartEx(LevelInfo, fileName, 10, 2, true)

	Info("Test 1")
	Info("Test 2")

	if _, err := os.Stat(fileName + ".1"); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(fileName + ".2"); err == nil {
		t.Fatal(err)
	}

	Info("Test 3")

	if _, err := os.Stat(fileName + ".2"); err != nil {
		t.Fatal(err)
	}

	Info("Test 4")

	if _, err := os.Stat(fileName + ".3"); err == nil {
		t.Fatal(err)
	}

	if err := Stop(); err != nil {
		fmt.Println(err)
	}

	if err := os.RemoveAll(path); err != nil {
		fmt.Println(err)
	}
}

func TestRotatingFileHandler(t *testing.T) {
	path := "./test_log"
	if err := os.RemoveAll(path); err != nil {
		fmt.Println(err)
	}

	if err := os.Mkdir(path, 0777); err != nil {
		fmt.Println(err)
	}
	fileName := path + "/test"

	h, err := NewRotatingFileHandler(fileName, 10, 2)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 10)

	if _, err := h.Write(buf); err != nil {
		fmt.Println(err)
	}

	if _, err := h.Write(buf); err != nil {
		fmt.Println(err)
	}

	if _, err := os.Stat(fileName + ".1"); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(fileName + ".2"); err == nil {
		t.Fatal(err)
	}

	if _, err := h.Write(buf); err != nil {
		fmt.Println(err)
	}
	if _, err := os.Stat(fileName + ".2"); err != nil {
		t.Fatal(err)
	}

	h.Close()

	if err := os.RemoveAll(path); err != nil {
		fmt.Println(err)
	}
}
