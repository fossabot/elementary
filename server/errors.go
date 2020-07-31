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
	"io"
	"net/http"

	"github.com/spf13/pflag"

	"github.com/forensicanalysis/forensicstore"
)

func ErrorsCommand() *Command {
	return &Command{
		Name:   "listErrors",
		Route:  "/errors",
		Method: http.MethodGet,
		Handler: func(w io.Writer, _ io.Reader, flags *pflag.FlagSet) error {
			storeName := flags.Args()[0]
			store, teardown, err := forensicstore.Open(storeName)
			if err != nil {
				return err
			}
			defer teardown()

			elements, err := store.Query("SELECT json FROM elements WHERE json_extract(json, '$.errors') != ''")
			if err != nil {
				return err
			}

			return PrintJSONList(w, int64(len(elements)), elements)
		},
	}
}
