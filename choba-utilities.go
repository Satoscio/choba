package choba

import (
	"os"
	"os/exec"
	"strconv"
)

// Thanks to lnmx @ stackoverflow
func TerminalChoba() int {
	var w int
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	out, _ := cmd.Output()

	if out[5] == 10 {
		w, _ = strconv.Atoi(string(out[3]) + string(out[4]))
	} else {
		w, _ = strconv.Atoi(string(out[3]) + string(out[4]) + string(out[5]))
	}

	return w
}

func ChobaCheck(w int, rw int) bool {
	return w >= rw
}
