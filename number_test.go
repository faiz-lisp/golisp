// Copyright 2014 SteelSeries ApS.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package implements a basic LISP interpretor for embedding in a go program for scripting.
// This file tests the integer type.

package golisp

import (
    . "launchpad.net/gocheck"
)

type IntegerAtomSuite struct {
    n *Data
}

var _ = Suite(&IntegerAtomSuite{})

func (s *IntegerAtomSuite) SetUpTest(c *C) {
    s.n = IntegerWithValue(5)
}

func (s *IntegerAtomSuite) TestIntegerValue(c *C) {
    c.Assert(IntegerValue(s.n), Equals, int32(5))
}

func (s *IntegerAtomSuite) TestString(c *C) {
    c.Assert(String(s.n), Equals, "5")
}

func (s *IntegerAtomSuite) TestEval(c *C) {
    n, err := Eval(s.n, Global)
    c.Assert(err, IsNil)
    c.Assert(n, Equals, s.n)
}

func (s *IntegerAtomSuite) TestBooleanValue(c *C) {
    c.Assert(BooleanValue(s.n), Equals, true)
}
