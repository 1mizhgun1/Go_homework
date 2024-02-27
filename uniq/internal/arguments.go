package uniq

import (
	"flag"
)

type Arguments struct {
	c      bool
	d      bool
	u      bool
	i      bool
	num    int
	chars  int
	input  string
	output string
}

func (args *Arguments) IsValid() bool {
	return !(args.c && args.d || args.c && args.u || args.d && args.u) && args.num >= 0 && args.chars >= 0
}

func (args *Arguments) Parse() {
	flag.BoolVar(&args.c, "c", false, "Count the number of occurrences of each line")
	flag.BoolVar(&args.d, "d", false, "Only print duplicate lines")
	flag.BoolVar(&args.u, "u", false, "Only print unique lines")
	flag.BoolVar(&args.i, "i", false, "Ignore case differences")
	flag.IntVar(&args.num, "f", 0, "Skip num fields before checking for uniqueness")
	flag.IntVar(&args.chars, "s", 0, "Skip chars characters before checking for uniqueness")

	flag.Parse()

	args.input = flag.Arg(0)
	args.output = flag.Arg(1)
}
