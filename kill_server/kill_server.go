package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)
// NOTE: must create the `server.pid` file to debug since its removed
// after each session

func main() {
	err := KillServer("server.pid")
	if err != nil {
		fmt.Println("ERROR:", err)
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("not found")
		}

		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf(">%s\n",e)
		}
	}

}

func KillServer(pidFile string) error {
	file, err := os.Open((pidFile))
	if err != nil {
		return err
	}
	defer file.Close()
	// called when function exits regardless of outcome
	// works at function level
	// defer executed in reverse order
	// Idiom: try to acquire a resource, check for error, defer release

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid: %w", pidFile, err)
	}
	slog.Info("killing", "pid", pid)

	if err := os.Remove(pidFile); err != nil { 
		slog.Warn("delete", "file", pidFile, "error", err)
	}

	return nil 
}
