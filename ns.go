// Copyright 2020 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

// Package clarkns implements Clark Notation for XML.
//
// For more information see http://www.jclark.com/xml/xmlns.htm
package clarkns // import "mellium.im/clarkns"

import (
	"encoding/xml"
	"errors"
	"strings"
)

// Errors returned by methods and functions in this package.
var (
	ErrUnclosedNS = errors.New("clarkns: no namespace closing bracket found")
)

// Name parses the provided XML name from Clark notation.
// For example the string "{http://www.cars.com/xml}part" would result in a
// local name of "part" and a namespace of "http://www.cars.com/xml"
//
// Name does not enforce XML name constraints and there is no guarantee that the
// returned xml.Name is valid.
func Name(s string) (xml.Name, error) {
	var name xml.Name
	if s == "" {
		return name, nil
	}

	// Parse out the namespace.
	if s[0] == '{' {
		s = s[1:]
		idx := strings.IndexByte(s, '}')
		if idx == -1 {
			return name, ErrUnclosedNS
		}
		name.Space = s[:idx]
		s = s[idx+1:]
	}

	name.Local = s
	return name, nil
}
