package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Concept) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Category":
			z.Category, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Version":
			err = dc.ReadExactBytes(z.Version[:])
			if err != nil {
				return
			}
		case "Doc":
			z.Doc, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Keywords":
			z.Keywords, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Status":
			z.Status, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Published":
			z.Published, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Parameters":
			var msz uint32
			msz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Parameters == nil && msz > 0 {
				z.Parameters = make(map[string]*Parameter, msz)
			} else if len(z.Parameters) > 0 {
				for key, _ := range z.Parameters {
					delete(z.Parameters, key)
				}
			}
			for msz > 0 {
				msz--
				var bzg string
				var bai *Parameter
				bzg, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					bai = nil
				} else {
					if bai == nil {
						bai = new(Parameter)
					}
					err = bai.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Parameters[bzg] = bai
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Concept) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "ID"
	err = en.Append(0x8a, 0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ID)
	if err != nil {
		return
	}
	// write "Category"
	err = en.Append(0xa8, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Category)
	if err != nil {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Version"
	err = en.Append(0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Version[:])
	if err != nil {
		return
	}
	// write "Doc"
	err = en.Append(0xa3, 0x44, 0x6f, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Doc)
	if err != nil {
		return
	}
	// write "Keywords"
	err = en.Append(0xa8, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Keywords)
	if err != nil {
		return
	}
	// write "Status"
	err = en.Append(0xa6, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Status)
	if err != nil {
		return
	}
	// write "Published"
	err = en.Append(0xa9, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Published)
	if err != nil {
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Type)
	if err != nil {
		return
	}
	// write "Parameters"
	err = en.Append(0xaa, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Parameters)))
	if err != nil {
		return
	}
	for bzg, bai := range z.Parameters {
		err = en.WriteString(bzg)
		if err != nil {
			return
		}
		if bai == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = bai.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Concept) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "ID"
	o = append(o, 0x8a, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "Category"
	o = append(o, 0xa8, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79)
	o = msgp.AppendString(o, z.Category)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Version"
	o = append(o, 0xa7, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	o = msgp.AppendArrayHeader(o, 3)
	for xvk := range z.Version {
		o = msgp.AppendUint8(o, z.Version[xvk])
	}
	// string "Doc"
	o = append(o, 0xa3, 0x44, 0x6f, 0x63)
	o = msgp.AppendString(o, z.Doc)
	// string "Keywords"
	o = append(o, 0xa8, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73)
	o = msgp.AppendString(o, z.Keywords)
	// string "Status"
	o = append(o, 0xa6, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendString(o, z.Status)
	// string "Published"
	o = append(o, 0xa9, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Published)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Parameters"
	o = append(o, 0xaa, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Parameters)))
	for bzg, bai := range z.Parameters {
		o = msgp.AppendString(o, bzg)
		if bai == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = bai.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Concept) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Category":
			z.Category, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Version":
			var asz uint32
			asz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if asz != 3 {
				err = msgp.ArrayError{Wanted: 3, Got: asz}
				return
			}
			for xvk := range z.Version {
				z.Version[xvk], bts, err = msgp.ReadUint8Bytes(bts)
				if err != nil {
					return
				}
			}
		case "Doc":
			z.Doc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Keywords":
			z.Keywords, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Status":
			z.Status, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Published":
			z.Published, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Parameters":
			var msz uint32
			msz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Parameters == nil && msz > 0 {
				z.Parameters = make(map[string]*Parameter, msz)
			} else if len(z.Parameters) > 0 {
				for key, _ := range z.Parameters {
					delete(z.Parameters, key)
				}
			}
			for msz > 0 {
				var bzg string
				var bai *Parameter
				msz--
				bzg, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					bai = nil
				} else {
					if bai == nil {
						bai = new(Parameter)
					}
					bts, err = bai.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
				z.Parameters[bzg] = bai
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Concept) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.ID) + 9 + msgp.StringPrefixSize + len(z.Category) + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.ArrayHeaderSize + (3 * (msgp.Uint8Size)) + 4 + msgp.StringPrefixSize + len(z.Doc) + 9 + msgp.StringPrefixSize + len(z.Keywords) + 7 + msgp.StringPrefixSize + len(z.Status) + 10 + msgp.BoolSize + 5 + msgp.StringPrefixSize + len(z.Type) + 11 + msgp.MapHeaderSize
	if z.Parameters != nil {
		for bzg, bai := range z.Parameters {
			_ = bai
			s += msgp.StringPrefixSize + len(bzg)
			if bai == nil {
				s += msgp.NilSize
			} else {
				s += bai.Msgsize()
			}
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Domain) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Doc":
			z.Doc, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Domain) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Name"
	err = en.Append(0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Doc"
	err = en.Append(0xa3, 0x44, 0x6f, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Doc)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Domain) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Name"
	o = append(o, 0x82, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Doc"
	o = append(o, 0xa3, 0x44, 0x6f, 0x63)
	o = msgp.AppendString(o, z.Doc)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Domain) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Doc":
			z.Doc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z Domain) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 4 + msgp.StringPrefixSize + len(z.Doc)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Operator) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Doc":
			z.Doc, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Multiple":
			z.Multiple, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Operator) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "Name"
	err = en.Append(0x84, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Doc"
	err = en.Append(0xa3, 0x44, 0x6f, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Doc)
	if err != nil {
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Type)
	if err != nil {
		return
	}
	// write "Multiple"
	err = en.Append(0xa8, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Multiple)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Operator) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Name"
	o = append(o, 0x84, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Doc"
	o = append(o, 0xa3, 0x44, 0x6f, 0x63)
	o = msgp.AppendString(o, z.Doc)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Multiple"
	o = append(o, 0xa8, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65)
	o = msgp.AppendBool(o, z.Multiple)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Operator) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Doc":
			z.Doc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Multiple":
			z.Multiple, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Operator) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 4 + msgp.StringPrefixSize + len(z.Doc) + 5 + msgp.StringPrefixSize + len(z.Type) + 9 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Parameter) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Doc":
			z.Doc, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Type":
			z.Type, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Operators":
			var msz uint32
			msz, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Operators == nil && msz > 0 {
				z.Operators = make(map[string]*Operator, msz)
			} else if len(z.Operators) > 0 {
				for key, _ := range z.Operators {
					delete(z.Operators, key)
				}
			}
			for msz > 0 {
				msz--
				var cmr string
				var ajw *Operator
				cmr, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					ajw = nil
				} else {
					if ajw == nil {
						ajw = new(Operator)
					}
					err = ajw.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Operators[cmr] = ajw
			}
		case "Required":
			z.Required, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Parameter) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "Name"
	err = en.Append(0x85, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Doc"
	err = en.Append(0xa3, 0x44, 0x6f, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Doc)
	if err != nil {
		return
	}
	// write "Type"
	err = en.Append(0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Type)
	if err != nil {
		return
	}
	// write "Operators"
	err = en.Append(0xa9, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteMapHeader(uint32(len(z.Operators)))
	if err != nil {
		return
	}
	for cmr, ajw := range z.Operators {
		err = en.WriteString(cmr)
		if err != nil {
			return
		}
		if ajw == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = ajw.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	// write "Required"
	err = en.Append(0xa8, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Required)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Parameter) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Name"
	o = append(o, 0x85, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Doc"
	o = append(o, 0xa3, 0x44, 0x6f, 0x63)
	o = msgp.AppendString(o, z.Doc)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	// string "Operators"
	o = append(o, 0xa9, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Operators)))
	for cmr, ajw := range z.Operators {
		o = msgp.AppendString(o, cmr)
		if ajw == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = ajw.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Required"
	o = append(o, 0xa8, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Required)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Parameter) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Doc":
			z.Doc, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Operators":
			var msz uint32
			msz, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.Operators == nil && msz > 0 {
				z.Operators = make(map[string]*Operator, msz)
			} else if len(z.Operators) > 0 {
				for key, _ := range z.Operators {
					delete(z.Operators, key)
				}
			}
			for msz > 0 {
				var cmr string
				var ajw *Operator
				msz--
				cmr, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					ajw = nil
				} else {
					if ajw == nil {
						ajw = new(Operator)
					}
					bts, err = ajw.UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
				z.Operators[cmr] = ajw
			}
		case "Required":
			z.Required, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Parameter) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 4 + msgp.StringPrefixSize + len(z.Doc) + 5 + msgp.StringPrefixSize + len(z.Type) + 10 + msgp.MapHeaderSize
	if z.Operators != nil {
		for cmr, ajw := range z.Operators {
			_ = ajw
			s += msgp.StringPrefixSize + len(cmr)
			if ajw == nil {
				s += msgp.NilSize
			} else {
				s += ajw.Msgsize()
			}
		}
	}
	s += 9 + msgp.BoolSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Version) DecodeMsg(dc *msgp.Reader) (err error) {
	err = dc.ReadExactBytes(z[:])
	if err != nil {
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Version) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteBytes(z[:])
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Version) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, 3)
	for wht := range z {
		o = msgp.AppendUint8(o, z[wht])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Version) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var asz uint32
	asz, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if asz != 3 {
		err = msgp.ArrayError{Wanted: 3, Got: asz}
		return
	}
	for wht := range z {
		z[wht], bts, err = msgp.ReadUint8Bytes(bts)
		if err != nil {
			return
		}
	}
	o = bts
	return
}

func (z *Version) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (3 * (msgp.Uint8Size))
	return
}
