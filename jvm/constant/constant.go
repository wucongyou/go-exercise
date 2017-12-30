package constant

const (
	_class              = 7
	_fiedlRef           = 9
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

// ClassInfo CONSTANT_Class_info.
type ClassInfo struct {
	Tag       uint8
	NameIndex uint16
}

// FiledRefInfo CONSTANT_Fieldref_info.
type FieldRefInfo struct {
	Tag              uint8
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
	Tag         uint8
	StringIndex uint16
}

// IntegerInfo CONSTANT_Integer_info.
type IntegerInfo struct {
	Tag   uint8
	Bytes uint32
}

// FloatInfo CONSTANT_Float_info.
type FloatInfo struct {
	IntegerInfo
}

// LongInfo CONSTANT_Long_info.
type LongInfo struct {
	Tag       uint8
	HighBytes uint32
	LowBytes  uint32
}

// DoubleInfo CONSTANT_Double_info.
type DoubleInfo struct {
	LongInfo
}

// NameAndTypeInfo CONSTANT_NameAndType_info.
type NameAndType struct {
	Tag             uint8
	NameIndex       uint16
	DescriptorIndex uint16
}

// Utf8Info CONSTANT_Utf8_info.
type Utf8Info struct {
	Tag    uint8
	Length uint16
	Bytes  []byte
}

// MethodHandler CONSTANT_MethodHandle_info.
type MethodHandler struct {
	Tag            uint8
	ReferenceKind  uint8
	ReferenceIndex uint16
}

// MethodTypeInfo CONSTANT_MethodType_info.
type MethodTypeInfo struct {
	Tag             uint8
	DescriptorIndex uint16
}

// InvokeDynamicInfo CONSTANT_InvokeDynamic_info.
type InvokeDynamicInfo struct {
	Tag                      uint8
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}
