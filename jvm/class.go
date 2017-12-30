package jvm

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// ClassFile class file.
type ClassFile struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	CpInfo            []ConstantInfo
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	InterfacesCount   uint16
	Interfaces        []*ClassInfo
	FieldsCount       uint16
	Fields            []*FieldInfo
	MethodsCount      uint16
	Methods           []*MethodInfo
	AttributesCount   uint16
	Attributes        []*AttributeInfo
}

func (m *ClassFile) Format() (res string, err error) {
	b := bytes.Buffer{}
	b.WriteString(fmt.Sprintf("magic: %x\n", m.Magic))
	b.WriteString(fmt.Sprintf("minor version: %d\n", m.MinorVersion))
	b.WriteString(fmt.Sprintf("major version: %d\n", m.MajorVersion))

	// flags
	accessFlags := ParseAccessFlags(m.AccessFlags)
	b.WriteString("flags: ")
	for i, flag := range accessFlags {
		b.WriteString(_accFm[flag])
		if i != len(accessFlags)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("\n")
	b.WriteString("constant pool:\n")
	l := 19
	c := len(strconv.FormatInt(int64(m.ConstantPoolCount), 10))
	for i := 1; i < int(m.ConstantPoolCount); i++ {
		if m.CpInfo[i] == nil {
			continue
		}
		c2 := len(strconv.FormatInt(int64(i), 10))
		b.WriteString(strings.Repeat(" ", c-c2))
		b.WriteString(fmt.Sprintf("#%d = ", i))
		switch u := m.CpInfo[i].(type) {
		case *Utf8Info:
			n := "utf8"
			b.WriteString(n)
			var rs []rune
			rs, err = DecodeRunes(u.Bytes)
			if err != nil {
				return
			}
			b.WriteString(strings.Repeat(" ", l-len(n)))
			b.WriteString(string(rs))
		case *NameAndType:
			n := "NameAndType"
			b.WriteString(n)
			b.WriteString(strings.Repeat(" ", l-len(n)))
			b.WriteString(fmt.Sprintf("#%d:#%d", u.NameIndex, u.DescriptorIndex))
		case *MethodRefInfo:
			n := "Methodref"
			b.WriteString(n)
			b.WriteString(strings.Repeat(" ", l-len(n)))
			b.WriteString(fmt.Sprintf("#%d:#%d", u.ClassIndex, u.NameAndTypeIndex))
		case *FieldRefInfo:
			n := "Fieldref"
			b.WriteString(n)
			b.WriteString(strings.Repeat(" ", l-len(n)))
			b.WriteString(fmt.Sprintf("#%d:#%d", u.ClassIndex, u.NameAndTypeIndex))
		case *StringInfo:
			n := "String"
			b.WriteString(n)
			b.WriteString(strings.Repeat(" ", l-len(n)))
			b.WriteString(fmt.Sprintf("#%d", u.StringIndex))
		case *ClassInfo:
			n := "Class"
			b.WriteString(n)
			b.WriteString(strings.Repeat(" ", l-len(n)))
			b.WriteString(fmt.Sprintf("#%d", u.NameIndex))
		}
		b.WriteString("\n")
	}
	return b.String(), nil
}

// FieldInfo field info.
type FieldInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []*AttributeInfo
}

func (m *FieldInfo) Read(b []byte, s int) (next int) {
	m.AccessFlags, next = u16(b, s)
	m.NameIndex, next = u16(b, next)
	m.DescriptorIndex, next = u16(b, next)
	m.AttributesCount, next = u16(b, next)
	m.Attributes = make([]*AttributeInfo, m.AttributesCount)
	for i := 0; i < int(m.AttributesCount); i++ {
		m.Attributes[i] = new(AttributeInfo)
		next = m.Attributes[i].Read(b, next)
	}
	return
}

// MethodInfo method info.
type MethodInfo struct {
	FieldInfo
}

// AttributeInfo attribute info.
type AttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               []byte
}

func (m *AttributeInfo) Read(b []byte, s int) (next int) {
	m.AttributeNameIndex, next = u16(b, s)
	m.AttributeLength, next = u32(b, next)
	m.Info, next = bs(b, next, int(m.AttributeLength))
	return
}
