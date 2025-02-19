// Copyright (c) 2020 Siemens AG
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
// Author(s): Jonas Plum

package server

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/forensicanalysis/forensicstore"
)

func PrintAny(w io.Writer, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return PrintJSON(w, b)
}

func PrintJSONList(w io.Writer, count int64, elements []forensicstore.JSONElement) error {
	_, err := w.Write([]byte(fmt.Sprintf("{\"count\": %d, \"elements\": [", count)))
	if err != nil {
		return err
	}

	for i, element := range elements {
		if i != 0 {
			_, err := w.Write([]byte(","))
			if err != nil {
				return err
			}
		}
		_, err := w.Write(element)
		if err != nil {
			return err
		}
	}

	_, err = w.Write([]byte("]}\n"))
	return err
}

func PrintJSON(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("\n"))

	return err
}
