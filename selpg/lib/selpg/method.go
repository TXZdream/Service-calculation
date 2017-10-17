package selpg

import (
	"os/exec"
	"fmt"
	"io"
	"bufio"
	"os"
)

func (selpg *Selpg) Read(Logfile *os.File) {
	if selpg == nil {
		fmt.Fprintf(os.Stderr, "Error: Unknown error.\n")
		Logfile.WriteString("[error] Use null object\n")
		os.Exit(0)
	}
	var in io.Reader

	if selpg.Src == "" {
		in = os.Stdin
	} else {
		var err error
		in, err = os.Open(selpg.Src)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: No such file found. Please pass right path.\n")
			Logfile.WriteString("[error] Unknown file to be read\n")
			os.Exit(0)
		}
	}
	scanner := bufio.NewScanner(in)
	if selpg.PageType == false {
		cnt := 0
		for scanner.Scan() {
			line := scanner.Text()
			if cnt / selpg.Length + 1 >= selpg.Begin && cnt / selpg.Length + 1 <= selpg.End {
				selpg.data = append(selpg.data, line)
			}
			cnt++
		}
	} else {
		cnt := 1
		onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			for i := 0; i < len(data); i++ {
				if data[i] == '\f' {
					return i + 1, data[:i], nil
				}
			}
			if atEOF {
				return 0, data, bufio.ErrFinalToken
			} else {
				return 0, nil, nil
			}
		}
		scanner.Split(onComma)
		for scanner.Scan() {
			line := scanner.Text()
			if cnt >= selpg.Begin && cnt <= selpg.End {
				selpg.data = append(selpg.data, line)
			}
			cnt++
		}
	}
	Logfile.WriteString("[info]  Read data finished\n")
}

func (selpg *Selpg) Write(Logfile *os.File) {
	if selpg == nil {
		fmt.Fprintf(os.Stderr, "Error: Unknown error.\n")
		Logfile.WriteString("[error] Use null object\n")
		os.Exit(0)
	}
	for i := 0; i < len(selpg.data); i++ {
		fmt.Fprintln(os.Stdout, selpg.data[i])
	}
	Logfile.WriteString("[info]  Write data finished\n")
}

func (selpg *Selpg) Print(Logfile *os.File) {
	if selpg.Destination != "" {
		lp := exec.Command("lp", fmt.Sprintf("-d %s", selpg.Destination))
		// lp := exec.Command("go")
		stdout, err := lp.StdoutPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Pipe to stdout failed.")
			Logfile.WriteString("[error] Can not pipe stdout to new process\n")			
		}
		stdin, err := lp.StdinPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Pipe to stdin failed.")
			Logfile.WriteString("[error] Can not pipe stdin to new process\n")
		}
		stderr, err := lp.StderrPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Pipe to stderr failed.")
			Logfile.WriteString("[error] Can not pipe stderr to new process\n")
		}
		for i := 0; i < len(selpg.data); i++ {
			fmt.Fprintf(stdin, "%s", selpg.data[i])
		}
		err = lp.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			Logfile.WriteString("[error] Open new process failed\n")
			os.Exit(0)
		}
		r := bufio.NewScanner(stdout)
		for r.Scan() {
			fmt.Fprintln(os.Stdout, r.Text())
		}
		r = bufio.NewScanner(stderr)
		for r.Scan() {
			fmt.Fprintln(os.Stdout, r.Text())
		}
	}
	Logfile.WriteString("[info]  Print data finished\n")
}