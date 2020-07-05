// Copyright 2020 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package clarkns_test

import (
	"encoding/xml"
	"errors"
	"strconv"
	"testing"

	"mellium.im/clarkns"
)

var nsTests = [...]struct {
	s    string
	name xml.Name
	err  error
}{
	0: {},
	1: {
		s:    "{foo}bar",
		name: xml.Name{Space: "foo", Local: "bar"},
	},
	2: {
		s:    "bar",
		name: xml.Name{Local: "bar"},
	},
	3: {
		s:   "{foo",
		err: clarkns.ErrUnclosedNS,
	},
}

func TestNS(t *testing.T) {
	for i, tc := range nsTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			name, err := clarkns.Name(tc.s)
			if !errors.Is(err, tc.err) {
				t.Errorf("wrong error: want=%v, got=%v", tc.err, err)
			}

			if tc.name != name {
				t.Errorf("wrong namespace: want=%v, got=%v", tc.name, name)
			}
		})
	}
}
