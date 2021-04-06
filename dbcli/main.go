package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h bool

	s   string
	d   string
	u   string
	p   string
	t   string
	f   string
	tab string
)

func init() {
	flag.BoolVar(&h, "h", false, "this `help`")
	flag.StringVar(&s, "s", "", "this is `db server`")
	flag.StringVar(&d, "d", "", "this is `db name`")
	flag.StringVar(&u, "u", "", "this is `db user`")
	flag.StringVar(&p, "p", "", "this is `db password`")
	flag.StringVar(&t, "t", "", "this is `table name`")
	flag.StringVar(&f, "f", "", "this is `file name` for excel export")
	flag.StringVar(&tab, "tab", "", "this is `sheet name` for excel")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}
	

}
func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}