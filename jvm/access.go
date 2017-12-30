package jvm

const (
	_accPublic    = 0x0001 // 0000000000000001
	_accFinal     = 0x0010 // 0000000000010000
	_accSuper     = 0x0020 // 0000000000100000
	_accInterface = 0x0200 // 0000001000000000
	_accAbstract  = 0x0400 // 0000010000000000
	_accSynthetic = 0x1000 // 0001000000000000

	_public       = 0x0001 // 0000000000000001
	_private      = 0x0002 // 0000000000000010
	_protected    = 0x0004 // 0000000000000100
	_static       = 0x0008 // 0000000000001000
	_final        = 0x0010 // 0000000000010000
	_synchronized = 0x0020 // 0000000000100000
	_bridge       = 0x0040 // 0000000001000000
	_varargs      = 0x0080 // 0000000010000000
	_native       = 0x0100 // 0000000100000000
	_abstract     = 0x0400 // 0000010000000000
	_strict       = 0x0800 // 0000100000000000
	_synthetic    = 0x1000 // 0001000000000000
)

var (
	_accFm = map[uint16]string{
		_accPublic:    "ACC_PUBLIC",
		_accFinal:     "ACC_FINAL",
		_accSuper:     "ACC_SUPER",
		_accInterface: "ACC_INTERFACE",
		_accAbstract:  "ACC_ABSTRACT",
		_accSynthetic: "ACC_SYNTHETIC",
	}
)

func ParseAccessFlags(accessFlags uint16) (flags []uint16) {
	flags = make([]uint16, 0)
	if accessFlags&_accPublic != 0 {
		flags = append(flags, _accPublic)
	}

	if accessFlags&_accFinal != 0 {
		flags = append(flags, _accFinal)
	}

	if accessFlags&_accSuper != 0 {
		flags = append(flags, _accSuper)
	}

	if accessFlags&_accInterface != 0 {
		flags = append(flags, _accInterface)
	}

	if accessFlags&_accAbstract != 0 {
		flags = append(flags, _accAbstract)
	}

	if accessFlags&_accSynthetic != 0 {
		flags = append(flags, _accSynthetic)
	}
	return
}

func ParseFlags(accflags uint16) (flags []uint16) {
	flags = make([]uint16, 0)
	if accflags&_public == 1 {
		flags = append(flags, _public)
	}

	if accflags&_private == 1 {
		flags = append(flags, _private)
	}

	if accflags&_protected == 1 {
		flags = append(flags, _protected)
	}

	if accflags&_static == 1 {
		flags = append(flags, _static)
	}

	if accflags&_final == 1 {
		flags = append(flags, _final)
	}

	if accflags&_synchronized == 1 {
		flags = append(flags, _synchronized)
	}

	if accflags&_bridge == 1 {
		flags = append(flags, _bridge)
	}

	if accflags&_varargs == 1 {
		flags = append(flags, _varargs)
	}

	if accflags&_native == 1 {
		flags = append(flags, _native)
	}

	if accflags&_abstract == 1 {
		flags = append(flags, _abstract)
	}

	if accflags&_strict == 1 {
		flags = append(flags, _strict)
	}

	if accflags&_synthetic == 1 {
		flags = append(flags, _synthetic)
	}
	return
}
