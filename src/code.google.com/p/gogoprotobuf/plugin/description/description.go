// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf
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
The description (experimental) plugin generates a Description method for each message.
The Description method returns a populated google_protobuf.FileDescriptorSet struct.
This contains the description of the files used to generate this message.

It is enabled by the following extensions:

  - description
  - description_all

The description plugin also generates a test given it is enabled using one of the following extensions:

  - testgen
  - testgen_all

Let us look at:

  code.google.com/p/gogoprotobuf/test/example/example.proto

Btw all the output can be seen at:

  code.google.com/p/gogoprotobuf/test/example/*

The following message:

  message B {
	option (gogoproto.description) = true;
	optional A A = 1 [(gogoproto.nullable) = false, (gogoproto.embed) = true];
	repeated bytes G = 2 [(gogoproto.customtype) = "code.google.com/p/gogoprotobuf/test/custom.Uint128", (gogoproto.nullable) = false];
  }

given to the description plugin, will generate the following code:

  func (this *B) Description() (desc *google_protobuf.FileDescriptorSet) {
	return ExampleDescription()
  }

and the following test code:

  func TestDescription(t *testing9.T) {
	ExampleDescription()
  }

The hope is to use this struct in some way instead of reflect.
This package is subject to change, since a use has not been figured out yet.

*/
package description

import (
	"code.google.com/p/gogoprotobuf/gogoproto"
	"code.google.com/p/gogoprotobuf/protoc-gen-gogo/generator"
	"fmt"
)

type plugin struct {
	*generator.Generator
	used bool
}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "description"
}

func (p *plugin) Init(g *generator.Generator) {
	p.used = false
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	localName := generator.FileName(file)
	for _, message := range file.Messages() {
		if !gogoproto.HasDescription(file.FileDescriptorProto, message.DescriptorProto) {
			continue
		}
		p.used = true
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func (this *`, ccTypeName, `) Description() (desc *google_protobuf.FileDescriptorSet) {`)
		p.In()
		p.P(`return `, localName, `Description()`)
		p.Out()
		p.P(`}`)
	}

	if p.used {

		p.P(`func `, localName, `Description() (desc *google_protobuf.FileDescriptorSet) {`)
		p.In()
		s := fmt.Sprintf("%#v", p.Generator.AllFiles())
		p.P(`return `, s)
		p.Out()
		p.P(`}`)
	}
}

func (this *plugin) GenerateImports(file *generator.FileDescriptor) {
	if this.used {
		this.P(`import google_protobuf "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"`)
	}
}

func init() {
	generator.RegisterPlugin(NewPlugin())
}
