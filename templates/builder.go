package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

// Templates is used for generating a bunch of redundant code that makes it easier on developers
// using this library.  For example, creates all the int, int8, int16, int32, and int64 assertions,
// which as all very similar code except for type checking.

func main() {
	var directory string
	if len(os.Args) < 2 {
		directory = "templates"
	} else {
		directory = os.Args[1]
	}

	templates, err := template.ParseGlob(directory + "/*.tmpl")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to parse %s for template files: %s\n", directory, err)
		os.Exit(1)
	}

	// Integers
	ints := templates.Lookup("ints.tmpl")
	if ints == nil {
		_, _ = fmt.Fprintln(os.Stderr, "ints.tmpl not found")
		os.Exit(1)
	}

	uints := templates.Lookup("uints.tmpl")
	if uints == nil {
		_, _ = fmt.Fprintln(os.Stderr, "uints.tmpl not found")
		os.Exit(1)
	}

	floats := templates.Lookup("floats.tmpl")
	if floats == nil {
		_, _ = fmt.Fprintln(os.Stderr, "floats.tmpl not found")
		os.Exit(1)
	}

	for _, bits := range []int{0, 8, 16, 32, 64} {
		var filename string
		if bits == 0 {
			filename = "int.go"
		} else {
			filename = fmt.Sprintf("int%d.go", bits)
		}

		var str string
		if bits != 0 {
			str = strconv.Itoa(bits)
		}

		// Integers
		out, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to open %s: %s\n", filename, err)
			os.Exit(1)
		}

		if err := ints.Execute(out, struct{Bits string}{str}); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to generate %s: %s\n", filename, err)
			os.Exit(1)
		}

		// Unsigned integers
		filename = "u" + filename
		out, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to open %s: %s\n", filename, err)
			os.Exit(1)
		}

		if err := ints.Execute(out, struct{Bits string}{str}); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to generate %s: %s\n", filename, err)
			os.Exit(1)
		}

		// Floats
		if bits > 16 {
			filename = fmt.Sprintf("float%d.go", bits)
			out, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Unable to open %s: %s\n", filename, err)
				os.Exit(1)
			}

			if err := floats.Execute(out, struct{Bits string}{str}); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Unable to generate %s: %s\n", filename, err)
				os.Exit(1)
			}
		}
	}
}
