package strs

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

const (
	ldigits = "0123456789abcdef"
	udigits = "0123456789ABCDEF"
)

var (
	space        = []byte(" ")
	doubleSpace  = []byte("  ")
	dot          = []byte(".")
	newLine      = []byte("\n")
	unsignedChar = []byte("unsigned char ")
	unsignedInt  = []byte("};\nunsigned int ")
	lenEquals    = []byte("_len = ")
	brackets     = []byte("[] = {")
	asterisk     = []byte("*")
	commaSpace   = []byte(", ")
	comma        = []byte(",")
	semiColonNl  = []byte(";\n")
	bar          = []byte("|")
)

func cfmtEncode(dst, src []byte, hextable string) {
	b := src[0]
	dst[3] = hextable[b&0x0f]
	dst[2] = hextable[b>>4]
	dst[1] = 'x'
	dst[0] = '0'
}

func xxd(r io.Reader, w io.Writer, fname string) error {
	var (
		cols      int
		octs      int
		caps      = ldigits
		doCHeader = true
		doCEnd    bool
		// enough room for "unsigned char NAME_FORMAT[] = {"
		varDeclChar = make([]byte, 14+len(fname)+6)
		// enough room for "unsigned int NAME_FORMAT = "
		varDeclInt = make([]byte, 16+len(fname)+7)
	)

	// Generate the first and last line in the -i output:
	// e.g. unsigned char foo_txt[] = { and unsigned int foo_txt_len =
	// copy over "unnsigned char " and "unsigned int"
	_ = copy(varDeclChar[0:14], unsignedChar[:])
	_ = copy(varDeclInt[0:16], unsignedInt[:])

	for i := 0; i < len(fname); i++ {
		if fname[i] != '.' {
			varDeclChar[14+i] = fname[i]
			varDeclInt[16+i] = fname[i]
		} else {
			varDeclChar[14+i] = '_'
			varDeclInt[16+i] = '_'
		}
	}
	// copy over "[] = {" and "_len = "
	_ = copy(varDeclChar[14+len(fname):], brackets[:])
	_ = copy(varDeclInt[16+len(fname):], lenEquals[:])

	cols = 12
	octs = 4

	// These are bumped down from the beginning of the function in order to
	// allow for their sizes to be allocated based on the user's speficiations
	var (
		line = make([]byte, cols)
		char = make([]byte, octs)
	)

	c := int64(0) // number of characters
	r = bufio.NewReader(r)

	for {
		n, err := io.ReadFull(r, line)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return err
		}

		if n == 0 {
			doCEnd = true
		}

		if doCHeader {
			w.Write(varDeclChar)
			w.Write(newLine)
			doCHeader = false
		}

		// C values
		if !doCEnd {
			w.Write(doubleSpace)
		}
		for i := 0; i < n; i++ {
			cfmtEncode(char, line[i:i+1], caps)
			w.Write(char)
			c++
			// don't add spaces to EOL
			if i != n-1 {
				w.Write(commaSpace)
			} else if n == cols {
				w.Write(comma)
			}
		}

		if doCEnd {
			w.Write(varDeclInt)
			w.Write([]byte(strconv.FormatInt(c, 10)))
			w.Write(semiColonNl)
			return nil
		}

		w.Write(newLine)
	}
	return nil
}

func XXD(in string, out string) error {
	inFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()

	outwr := bufio.NewWriter(outFile)
	defer outwr.Flush()

	return xxd(inFile, outwr, filepath.Base(in))
}
