package main

import "flag"
import "os"
import "github.com/txzdream/serviceCourse/selpg/lib/selpg"
import "fmt"

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of seplg:\n")
		fmt.Printf("seplg -s num1 -e num2 [-f -l num3 -d str1 file]\n")
		flag.PrintDefaults()
	}
	start := flag.Int("s", -1, "Start of the page")
	end := flag.Int("e", -1, "End of the page")
	pagetype := flag.Bool("f", false, "If the page has static number of lines")
	length := flag.Int("l", 72, "the number of lines of every page")
	destination := flag.String("d", "", "the destination to send")
	flag.Parse()

	if *start <= 0 || *end <= 0 || *end < *start || *length < 0 {
		fmt.Fprintf(os.Stderr, "Error: Invalid start, end page or line number. Use selpg -help to know more.\n")
		os.Exit(0)
	}
	if *pagetype == false && *length != 72 {
		fmt.Fprintln(os.Stderr, "Error: Conflict flags -f and -l")
	}
	var src string
	if flag.NArg() == 1 {
		src = flag.Args()[0]
	} else if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "Error: Too much argument. Use selpg -help to know more.\n")
		os.Exit(0)
	} else {
		src = ""
	}

	data := selpg.Selpg{
		Begin: *start,
		End: *end,
		PageType: *pagetype,
		Length: *length,
		Destination: *destination,
		Src: src,
	}
	
	data.Read()
	data.Write()
}
