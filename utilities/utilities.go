package utilities

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/term"
)

func Clear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func EnableRawMode() (state *term.State) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return oldState
}

func DisableRawMode(state *term.State) {
	term.Restore(int(os.Stdin.Fd()), state)
}
