package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	ln := 0
	for sc.Scan() {
	    ln++
		l := sc.Bytes()

		stderr, err := checkJSON(l)
		if err == nil {
		    continue
		}

		errStr := string(stderr.Bytes())
		fmt.Printf("LINE %d:\n%s\nERROR[jq] %s\n", ln, string(l), errStr)
	}
}

func checkJSON(l []byte) (bytes.Buffer, error) {
	var stdout, stderr bytes.Buffer

	jq := exec.Command("jq")
	jq.Stdin = bytes.NewReader(l)
	jq.Stdout = &stdout
	jq.Stderr = &stderr

	err := jq.Run()

	return stderr, err
}