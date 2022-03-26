package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/andrew-d/go-termutil"
)

var (
	input *bufio.Reader
	re1   = regexp.MustCompile("make(\\[[0-9]+\\])?:")
)

const (
	bold  = "\033[1m"
	reset = "\033[0m"

	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
)

func first(args ...interface{}) interface{} {
	return args[0]
}

func Colorize(line string) string {
	tokens := strings.Split(line, " ")
	if cap(tokens) < 2 {
		return line
	}

	// TODO: すごい書きかけです
	cline := line
	switch {
	case (cap(tokens) > 2) && (tokens[1] == "undefined"):
		cline = bold + tokens[0] + " " + red + strings.Join(tokens[1:], " ") + reset

	case (cap(tokens) > 2) && (tokens[1] == "error:"):
		cline = "\n"
		cline += bold + green
		cline += "❗️🐈< " + tokens[0] + " \n"
		cline += "　😸< " + tokens[1] + " " + strings.TrimRight(strings.Join(tokens[2:], " "), "\n") + " meow!\n" + reset

	case (cap(tokens) > 2) && (tokens[1] == "note:"):
		cline = "\n"
		cline += bold + blue
		cline += "❗️🐈< " + tokens[0] + " \n"
		cline += "💡😼< " + tokens[1] + " " + strings.TrimRight(strings.Join(tokens[2:], " "), "\n") + " meow!\n" + reset

	case (cap(tokens) > 3) && (tokens[1] == "fatal") && (tokens[1] == "error:"):
		cline = bold + strings.Join(tokens[0:3], " ") + " " + red + strings.Join(tokens[3:], " ") + reset

	case (cap(tokens) > 2) && (tokens[1] == "warning:"):
		cline = bold + strings.Join(tokens[0:2], " ") + " " + yellow + strings.Join(tokens[2:], " ") + reset

	case (cap(tokens) > 3) && (tokens[1] == "In") && (tokens[2] == "function"):
		cline = strings.Join(tokens[0:2], " ") + " " + magenta + strings.Join(tokens[3:], " ") + reset

	case (cap(tokens) > 4) && (tokens[1] == "In") && (tokens[2] == "member") && (tokens[3] == "function"):
		cline = strings.Join(tokens[0:3], " ") + " " + magenta + strings.Join(tokens[4:], " ") + reset

	case (cap(tokens) > 2) && re1.MatchString(tokens[0]) && (tokens[1] == "***"):
		cline = bold + red + strings.Join(tokens[:cap(tokens)], " ") + reset
	}

	return cline
}

func main() {
	flag.Parse()
	fname := flag.Arg(0)

	if fname == "" {
		input = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(fname)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		input = bufio.NewReader(f)
	}

	fmt.Print(yellow + "\n")
	fmt.Print("ฅ(=･ω･=)ฅ\nCat started code reading... >\n\n")
	fmt.Print("ヾ(=`･ω･´)∩ nyannnn...\nHe is very focused so let's keep an eye on him for a bit ... >\n\n") // 「彼はとても集中しているので、少し見守ってあげましょう」みたいなことが書きたかった
	fmt.Print(reset)

	tty := termutil.Isatty(os.Stdout.Fd())
	for {
		line, err := input.ReadString('\n')
		if tty {
			fmt.Print(Colorize(line))
		} else {
			fmt.Print(line)
		}
		if err != nil {
			break
		}
	}
}
