package selpg

type Selpg struct {
	Begin int
	End int
	/* false for static line number */
	PageType bool
	Length int
	Destination string
	Src string
	data []string
}