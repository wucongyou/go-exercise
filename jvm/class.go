package jvm

// ClassFile class file.
type ClassFile struct {
	Magic             uint32
	MinorVersion      uint16
	MajorVersion      uint16
	ConstantPoolCount uint16
	CpInfo            []*ConstantInfo
	AccessFlags       uint16
	ThisClass         uint16
	SuperClass        uint16
	InterfacesCount   uint16
	Interfaces        []*ClassInfo
	FieldsCount       uint16
	Fields            []*FieldInfo
	MethodCount       uint16
	Methods           []*MethodInfo
	AttributeCount    uint16
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

// MethodInfo method info.
type MethodInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []*AttributeInfo
}

// AttributeInfo attribute info.
type AttributeInfo struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               []byte
}
