package jvm

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
	m.Info, next = bytes(b, next, int(m.AttributeLength))
	return
}
