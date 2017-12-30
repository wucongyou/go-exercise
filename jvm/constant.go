package jvm

const (
	_class              = 7
	_fieldRef           = 9
	_methodRef          = 10
	_interfaceMethodRef = 11
	_string             = 8
	_inteager           = 3
	_float              = 4
	_long               = 5
	_double             = 6
	_nameAndType        = 12
	_utf8               = 1
	_methodHandle       = 15
	_methodType         = 16
	_invokeDynamic      = 18
)

// ConstantInfo constant info, each constant info holds tag.
type ConstantInfo struct {
	Tag uint8
}

// ClassInfo CONSTANT_Class_info.
type ClassInfo struct {
	ConstantInfo
	NameIndex uint16
}

// FiledRefInfo CONSTANT_Fieldref_info.
type FieldRefInfo struct {
	ConstantInfo
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

// MethodRefInfo CONSTANT_Methodref_info.
type MethodRefInfo struct {
	FieldRefInfo
}

// InterfaceMethodRefInfo CONSTANT_InterfaceMethodref_info.
type InterfaceMethodRefInfo struct {
	FieldRefInfo
}

// StringInfo CONSTANT_String_info.
type StringInfo struct {
	ConstantInfo
	StringIndex uint16
}

// IntegerInfo CONSTANT_Integer_info.
type IntegerInfo struct {
	ConstantInfo
	Bytes uint32
}

// FloatInfo CONSTANT_Float_info.
type FloatInfo struct {
	IntegerInfo
}

// LongInfo CONSTANT_Long_info.
type LongInfo struct {
	ConstantInfo
	HighBytes uint32
	LowBytes  uint32
}

// DoubleInfo CONSTANT_Double_info.
type DoubleInfo struct {
	LongInfo
}

// NameAndTypeInfo CONSTANT_NameAndType_info.
type NameAndType struct {
	ConstantInfo
	NameIndex       uint16
	DescriptorIndex uint16
}

// Utf8Info CONSTANT_Utf8_info.
type Utf8Info struct {
	ConstantInfo
	Length uint16
	Bytes  []byte
}

// MethodHandler CONSTANT_MethodHandle_info.
type MethodHandler struct {
	ConstantInfo
	ReferenceKind  uint8
	ReferenceIndex uint16
}

// MethodTypeInfo CONSTANT_MethodType_info.
type MethodTypeInfo struct {
	ConstantInfo
	DescriptorIndex uint16
}

// InvokeDynamicInfo CONSTANT_InvokeDynamic_info.
type InvokeDynamicInfo struct {
	ConstantInfo
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}
