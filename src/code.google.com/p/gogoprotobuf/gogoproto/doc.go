// Extensions for Protocol Buffers to create more go like structures.
//
// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf/gogoproto
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

/*
Package gogoproto provides extensions for protocol buffers to achieve:

  - fast marshalling and unmarshalling.
  - peace of mind by optionally generating test and benchmark code.
  - more canonical Go structures.
  - less typing by optionally generating extra helper code.
  - goprotobuf compatibility

More Canonical Go Structures

A lot of time working with a goprotobuf struct will lead you to a place where you create another struct that is easier to work with and then have a function to copy the values between the two structs.
You might also find that basic structs that started their life as part of an API need to be sent over the wire. With gob, you could just send it. With goprotobuf, you need to make a parallel struct.
Gogoprotobuf tries to fix these problems with the nullable, embed and customtype field extensions.

  - nullable, if false, a field is generated without a pointer (see warning below).
  - embed, if true, the field is generated as an embedded field.
  - customtype, It works with the Marshal and Unmarshal methods, to allow you to have your own types in your struct, but marshal to bytes. For example, custom.Uuid or custom.Fixed128

Warning about nullable: According to the Protocol Buffer specification, you should be able to tell whether a field is set or unset. With the option nullable=false this feature is lost, since your non-nullable fields will always be set. It can be seen as a layer on top of Protocol Buffers, where before and after marshalling all non-nullable fields are set and they cannot be unset.

Let us look at:

	code.google.com/p/gogoprotobuf/test/example/example.proto

for a quicker overview.

The following message:

	message A {
		optional string Description = 1 [(gogoproto.nullable) = false];
		optional int64 Number = 2 [(gogoproto.nullable) = false];
		optional bytes Id = 3 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uuid", (gogoproto.nullable) = false];
	}

Will generate a go struct which looks a lot like this:

	type A struct {
		Description string
		Number      int64
		Id          code_google_com_p_gogoprotobuf_test_custom.Uuid
	}

You will see there are no pointers, since all fields are non-nullable.
You will also see a custom type which marshals to a string.

Next we will embed the message A in message B.

	message B {
		optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
		repeated bytes G = 2 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uint128", (gogoproto.nullable) = false];
	}

See below that A is embedded in B.

	type B struct {
		A
		G []code_google_com_p_gogoprotobuf_test_custom.Uint128
	}

Also see the repeated custom type.

	type Uint128 [2]uint64

Gogoprotobuf also has some more subtle changes, these could be changed back:

  - the generated package name for imports do not have the extra /filename.pb,
  but are actually the imports specified in the .proto file.

Gogoprotobuf also has lost some features which should be brought back with time:

  - Marshalling and unmarshalling with reflect and without the unsafe package,
  this requires work in pointer_reflect.go

Why does nullable break protocol buffer specifications:

The protocol buffer specification states, somewhere, that you should be able to tell whether a
field is set or unset.  With the option nullable=false this feature is lost,
since your non-nullable fields will always be set.  It can be seen as a layer on top of
protocol buffers, where before and after marshalling all non-nullable fields are set
and they cannot be unset.

Goprotobuf Compatibility:

Gogoprotobuf is compatible with Goprotobuf, because it is compatible with protocol buffers.
Gogoprotobuf generates the same code as goprotobuf if no extensions are used.
The enumprefix, getters and msgstringmethod extensions can be used to remove some of the unnecessary code generated by goprotobuf:

  - goproto_enum_prefix, if false, generates the enum constant names without the messagetype prefix
  - goproto_enum_stringer (experimental), if false, the enum is generated without the default string method, this is useful for rather using enum_stringer, or allowing you to write your own string method.
  - goproto_getters, if false, the message is generated without get methods, this is useful when you would rather want to use face
  - goproto_stringer, if false, the message is generated without the default string method, this is useful for rather using stringer, or allowing you to write your own string method.

Less Typing and Peace of Mind is explained in their specific plugin folders godoc:

	- code.google.com/p/gogoprotobuf/plugin/<extension_name>

If you do not use any of these extension the code that is generated
will be the same as if goprotobuf has generated it.

The most complete way to see examples is to look at

	code.google.com/p/gogoprotobuf/test/thetest.proto

Gogoprototest is a seperate project,
because we want to keep gogoprotobuf independant of goprotobuf,
but we still want to test it thoroughly.

*/
package gogoproto
