package cmdutils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecCmd(name string, args ...string) error {
	//fmt.Println("\n[EXEC]", name, strings.Join(args, " "))
	defer fmt.Println()

	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		outScanner := bufio.NewScanner(stdout)
		for outScanner.Scan() {
			fmt.Println(outScanner.Text())
		}
	}()

	var errmsg bytes.Buffer
	go func() {
		errScanner := bufio.NewScanner(stderr)
		for errScanner.Scan() {
			fmt.Fprintln(os.Stderr, errScanner.Text())
			errmsg.Write(errScanner.Bytes())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func ExecCmdWithWorkspace(workspace, name string, args ...string) error {
	//fmt.Println("\n[EXEC]", name, strings.Join(args, " "))
	defer fmt.Println()

	cmd := exec.Command(name, args...)
	cmd.Dir = workspace

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	go func() {
		outScanner := bufio.NewScanner(stdout)
		for outScanner.Scan() {
			fmt.Println(outScanner.Text())
		}
	}()

	var errmsg bytes.Buffer
	go func() {
		errScanner := bufio.NewScanner(stderr)
		for errScanner.Scan() {
			fmt.Fprintln(os.Stderr, errScanner.Text())
			errmsg.Write(errScanner.Bytes())
		}
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

func ExecCmdSilently(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
