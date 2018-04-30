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

package proto

import (
	"reflect"
)

// Encode a reference to bool pointer.
func (o *Buffer) enc_ref_bool(p *Properties, base structPointer) error {
	v := structPointer_RefBool(base, p.field)
	if v == nil {
		return ErrNil
	}
	x := 0
	if *v {
		x = 1
	}
	o.buf = append(o.buf, p.tagcode...)
	p.valEnc(o, uint64(x))
	return nil
}

// Encode a reference to int32 pointer.
func (o *Buffer) enc_ref_int32(p *Properties, base structPointer) error {
	v := structPointer_RefWord32(base, p.field)
	if refWord32_IsNil(v) {
		return ErrNil
	}
	x := refWord32_Get(v)
	o.buf = append(o.buf, p.tagcode...)
	p.valEnc(o, uint64(x))
	return nil
}

// Encode a reference to an int64 pointer.
func (o *Buffer) enc_ref_int64(p *Properties, base structPointer) error {
	v := structPointer_RefWord64(base, p.field)
	if refWord64_IsNil(v) {
		return ErrNil
	}
	x := refWord64_Get(v)
	o.buf = append(o.buf, p.tagcode...)
	p.valEnc(o, x)
	return nil
}

// Encode a reference to a string pointer.
func (o *Buffer) enc_ref_string(p *Properties, base structPointer) error {
	v := structPointer_RefString(base, p.field)
	if v == nil {
		return ErrNil
	}
	x := *v
	o.buf = append(o.buf, p.tagcode...)
	o.EncodeStringBytes(x)
	return nil
}

// Encode a reference to a message struct.
func (o *Buffer) enc_ref_struct_message(p *Properties, base structPointer) error {
	structp := structPointer_GetRefStructPointer(base, p.field)
	if structPointer_IsNil(structp) {
		return ErrNil
	}

	// Can the object marshal itself?
	if p.isMarshaler {
		m := structPointer_Interface(structp, p.stype).(Marshaler)
		data, err := m.Marshal()
		if err != nil {
			return err
		}
		o.buf = append(o.buf, p.tagcode...)
		o.EncodeRawBytes(data)
		return nil
	}

	// need the length before we can write out the message itself,
	// so marshal into a separate byte buffer first.
	obuf := o.buf
	o.buf = o.bufalloc()

	err := o.enc_struct(p.stype, p.sprop, structp)

	nbuf := o.buf
	o.buf = obuf
	if err != nil {
		o.buffree(nbuf)
		return err
	}
	o.buf = append(o.buf, p.tagcode...)
	o.EncodeRawBytes(nbuf)
	o.buffree(nbuf)
	return nil
}

// Encode a slice of references to message struct pointers ([]struct).
func (o *Buffer) enc_slice_ref_struct_message(p *Properties, base structPointer) error {
	ss := structPointer_GetStructPointer(base, p.field)
	ss1 := structPointer_GetRefStructPointer(ss, field(0))
	size := p.stype.Size()
	l := structPointer_Len(base, p.field)
	for i := 0; i < l; i++ {
		structp := structPointer_Add(ss1, field(uintptr(i)*size))
		if structPointer_IsNil(structp) {
			return ErrRepeatedHasNil
		}

		// Can the object marshal itself?
		if p.isMarshaler {
			m := structPointer_Interface(structp, p.stype).(Marshaler)
			data, err := m.Marshal()
			if err != nil {
				return err
			}
			o.buf = append(o.buf, p.tagcode...)
			o.EncodeRawBytes(data)
			continue
		}

		obuf := o.buf
		o.buf = o.bufalloc()

		err := o.enc_struct(p.stype, p.sprop, structp)

		nbuf := o.buf
		o.buf = obuf
		if err != nil {
			o.buffree(nbuf)
			if err == ErrNil {
				return ErrRepeatedHasNil
			}
			return err
		}
		o.buf = append(o.buf, p.tagcode...)
		o.EncodeRawBytes(nbuf)

		o.buffree(nbuf)
	}
	return nil
}

func (o *Buffer) enc_custom_bytes(p *Properties, base structPointer) error {
	i := structPointer_InterfaceRef(base, p.field, p.ctype)
	if i == nil {
		return ErrNil
	}
	custom := i.(Marshaler)
	data, err := custom.Marshal()
	if err != nil {
		return err
	}
	if data == nil {
		return ErrNil
	}
	o.buf = append(o.buf, p.tagcode...)
	o.EncodeRawBytes(data)
	return nil
}

func (o *Buffer) enc_custom_ref_bytes(p *Properties, base structPointer) error {
	custom := structPointer_InterfaceAt(base, p.field, p.ctype).(Marshaler)
	data, err := custom.Marshal()
	if err != nil {
		return err
	}
	if data == nil {
		return ErrNil
	}
	o.buf = append(o.buf, p.tagcode...)
	o.EncodeRawBytes(data)
	return nil
}

func (o *Buffer) enc_custom_slice_bytes(p *Properties, base structPointer) error {
	inter := structPointer_InterfaceRef(base, p.field, p.ctype)
	if inter == nil {
		return ErrNil
	}
	slice := reflect.ValueOf(inter)
	l := slice.Len()
	for i := 0; i < l; i++ {
		v := slice.Index(i)
		custom := v.Interface().(Marshaler)
		data, err := custom.Marshal()
		if err != nil {
			return err
		}
		o.buf = append(o.buf, p.tagcode...)
		o.EncodeRawBytes(data)
	}
	return nil
}
