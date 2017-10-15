package selpg

import (
	"fmt"
	"io"
	"bufio"
	"os"
)

func (selpg *Selpg) Read() {
	if selpg == nil {
		fmt.Fprintf(os.Stderr, "Error: Unknown error.\n")
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
}

func (selpg *Selpg) Write() {
	if selpg == nil {
		fmt.Fprintf(os.Stderr, "Error: Unknown error.\n")
		os.Exit(0)
	}
	for i := 0; i < len(selpg.data); i++ {
		fmt.Fprintln(os.Stdout, selpg.data[i])
	}
}