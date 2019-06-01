package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

func main() {
	formatFuncSignatures(os.Stdin)
}

func formatFuncSignatures(r io.Reader) {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	lines := []int{}
	line := 0
	for o, b := range src {
		if line >= 0 {
			lines = append(lines, line)
		}
		line = -1
		if b == '\n' {
			line = o + 1
		}
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(node ast.Node) bool {
		if node == nil {
			return true
		}

		switch n := node.(type) {
		case *ast.FuncDecl:
			if len(n.Type.Params.List) < 3 {
				return true
			}

			file := fset.File(n.Pos())
			lineNo := file.Line(n.Pos())
			newlines := []int{}
			prevLineNo := lineNo
			lastPos := token.NoPos
			for _, p := range n.Type.Params.List {
				lastPos = p.Pos()
				currLineNo := file.Line(lastPos)
				if currLineNo == prevLineNo {
					newlines = append(newlines, int(lastPos-1))
				}
				prevLineNo = currLineNo
			}

			if lastPos != token.NoPos && prevLineNo >= file.Line(n.Type.Params.Closing) {
				if file.Line(n.Body.Lbrace) == file.Line(n.Body.Rbrace) {
					newlines = append(newlines, int(n.Body.Lbrace-1))
				}
				n.Type.Params.Closing = n.Body.Lbrace + 2
			}

			if len(newlines) > 0 {
				tmplines := make([]int, 0, len(lines)+len(newlines))
				tmplines = append(tmplines, lines[:lineNo]...)
				tmplines = append(tmplines, newlines...)
				tmplines = append(tmplines, lines[lineNo:]...)
				sort.Ints(tmplines)
				lines = tmplines
				file.SetLines(lines)
			}
		}

		return true
	})

	var buf bytes.Buffer
	format.Node(&buf, fset, f)

	fmt.Print(buf.String())
}
