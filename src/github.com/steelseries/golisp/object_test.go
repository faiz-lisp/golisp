// Copyright 2013 SteelSeries ApS.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package impliments a basic LISP interpretor for embedding in a go program for scripting.
// This file tests the Object data type
package golisp

import (
    . "launchpad.net/gocheck"
    "unsafe"
)

type ObjectAtomSuite struct {
    o *Data
}

var _ = Suite(&ObjectAtomSuite{})

type TestStruct struct {
    D int
}

func (s *ObjectAtomSuite) TestObject(c *C) {
    obj := &TestStruct{D: 5}
    s.o = ObjectWithTypeAndValue("TestStruct", unsafe.Pointer(obj))
    c.Assert(s.o.ObjType, Equals, "TestStruct")
    c.Assert((*TestStruct)(ObjectValue(s.o)), Equals, obj)
    c.Assert((*TestStruct)(ObjectValue(s.o)).D, Equals, 5)
}
