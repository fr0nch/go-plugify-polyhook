package polyhook

/*
#include "setters.h"
#cgo noescape SetArgumentBool
#cgo noescape SetArgumentInt8
#cgo noescape SetArgumentUInt8
#cgo noescape SetArgumentInt16
#cgo noescape SetArgumentUInt16
#cgo noescape SetArgumentInt32
#cgo noescape SetArgumentUInt32
#cgo noescape SetArgumentInt64
#cgo noescape SetArgumentUInt64
#cgo noescape SetArgumentFloat
#cgo noescape SetArgumentDouble
#cgo noescape SetArgumentPointer
#cgo noescape SetArgumentString
#cgo noescape SetArgument
#cgo noescape SetReturnBool
#cgo noescape SetReturnInt8
#cgo noescape SetReturnUInt8
#cgo noescape SetReturnInt16
#cgo noescape SetReturnUInt16
#cgo noescape SetReturnInt32
#cgo noescape SetReturnUInt32
#cgo noescape SetReturnInt64
#cgo noescape SetReturnUInt64
#cgo noescape SetReturnFloat
#cgo noescape SetReturnDouble
#cgo noescape SetReturnPointer
#cgo noescape SetReturnString
#cgo noescape SetReturn
#cgo noescape SetRegisterBool
#cgo noescape SetRegisterInt8
#cgo noescape SetRegisterUInt8
#cgo noescape SetRegisterInt16
#cgo noescape SetRegisterUInt16
#cgo noescape SetRegisterInt32
#cgo noescape SetRegisterUInt32
#cgo noescape SetRegisterInt64
#cgo noescape SetRegisterUInt64
#cgo noescape SetRegisterFloat
#cgo noescape SetRegisterDouble
#cgo noescape SetRegisterPointer
#cgo noescape SetRegisterString
#cgo noescape SetRegister
*/
import "C"
import (
	"errors"
	"reflect"
	"runtime"
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = errors.New("")
var _ = reflect.TypeOf(0)
var _ = runtime.GOOS
var _ = unsafe.Sizeof(0)
var _ = plugify.Plugin.Loaded

// Generated from polyhook (group: setters)

// SetArgumentBool 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentBool(params ParametersHandle, index uint64, value bool) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.bool(value)
	C.SetArgumentBool(__params, __index, __value)
}

// SetArgumentInt8 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentInt8(params ParametersHandle, index uint64, value int8) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.int8_t(value)
	C.SetArgumentInt8(__params, __index, __value)
}

// SetArgumentUInt8 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentUInt8(params ParametersHandle, index uint64, value uint8) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uint8_t(value)
	C.SetArgumentUInt8(__params, __index, __value)
}

// SetArgumentInt16 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentInt16(params ParametersHandle, index uint64, value int16) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.int16_t(value)
	C.SetArgumentInt16(__params, __index, __value)
}

// SetArgumentUInt16 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentUInt16(params ParametersHandle, index uint64, value uint16) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uint16_t(value)
	C.SetArgumentUInt16(__params, __index, __value)
}

// SetArgumentInt32 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentInt32(params ParametersHandle, index uint64, value int32) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.int32_t(value)
	C.SetArgumentInt32(__params, __index, __value)
}

// SetArgumentUInt32 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentUInt32(params ParametersHandle, index uint64, value uint32) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uint32_t(value)
	C.SetArgumentUInt32(__params, __index, __value)
}

// SetArgumentInt64 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentInt64(params ParametersHandle, index uint64, value int64) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.int64_t(value)
	C.SetArgumentInt64(__params, __index, __value)
}

// SetArgumentUInt64 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentUInt64(params ParametersHandle, index uint64, value uint64) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uint64_t(value)
	C.SetArgumentUInt64(__params, __index, __value)
}

// SetArgumentFloat 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentFloat(params ParametersHandle, index uint64, value float32) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.float(value)
	C.SetArgumentFloat(__params, __index, __value)
}

// SetArgumentDouble 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentDouble(params ParametersHandle, index uint64, value float64) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.double(value)
	C.SetArgumentDouble(__params, __index, __value)
}

// SetArgumentPointer 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentPointer(params ParametersHandle, index uint64, value uintptr) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uintptr_t(value)
	C.SetArgumentPointer(__params, __index, __value)
}

// SetArgumentString 
//  @brief Set argument value
//
//  @param hook: Hook pointer
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func SetArgumentString(hook HookHandle, params ParametersHandle, index uint64, value string) {
	__hook := C.uintptr_t(hook)
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			C.SetArgumentString(__hook, __params, __index, (*C.String)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetArgument 
//  @brief Set argument value
//
//  @param hook: Hook pointer
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgument(hook HookHandle, params ParametersHandle, index uint64, value any) {
	__hook := C.uintptr_t(hook)
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := plugify.ConstructVariant(value)
	plugify.Block {
		Try: func() {
			C.SetArgument(__hook, __params, __index, (*C.Variant)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__value)
		},
	}.Do()
}

// SetReturnBool 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnBool(ret ReturnHandle, value bool) {
	__ret := C.uintptr_t(ret)
	__value := C.bool(value)
	C.SetReturnBool(__ret, __value)
}

// SetReturnInt8 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnInt8(ret ReturnHandle, value int8) {
	__ret := C.uintptr_t(ret)
	__value := C.int8_t(value)
	C.SetReturnInt8(__ret, __value)
}

// SetReturnUInt8 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnUInt8(ret ReturnHandle, value uint8) {
	__ret := C.uintptr_t(ret)
	__value := C.uint8_t(value)
	C.SetReturnUInt8(__ret, __value)
}

// SetReturnInt16 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnInt16(ret ReturnHandle, value int16) {
	__ret := C.uintptr_t(ret)
	__value := C.int16_t(value)
	C.SetReturnInt16(__ret, __value)
}

// SetReturnUInt16 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnUInt16(ret ReturnHandle, value uint16) {
	__ret := C.uintptr_t(ret)
	__value := C.uint16_t(value)
	C.SetReturnUInt16(__ret, __value)
}

// SetReturnInt32 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnInt32(ret ReturnHandle, value int32) {
	__ret := C.uintptr_t(ret)
	__value := C.int32_t(value)
	C.SetReturnInt32(__ret, __value)
}

// SetReturnUInt32 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnUInt32(ret ReturnHandle, value uint32) {
	__ret := C.uintptr_t(ret)
	__value := C.uint32_t(value)
	C.SetReturnUInt32(__ret, __value)
}

// SetReturnInt64 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnInt64(ret ReturnHandle, value int64) {
	__ret := C.uintptr_t(ret)
	__value := C.int64_t(value)
	C.SetReturnInt64(__ret, __value)
}

// SetReturnUInt64 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnUInt64(ret ReturnHandle, value uint64) {
	__ret := C.uintptr_t(ret)
	__value := C.uint64_t(value)
	C.SetReturnUInt64(__ret, __value)
}

// SetReturnFloat 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnFloat(ret ReturnHandle, value float32) {
	__ret := C.uintptr_t(ret)
	__value := C.float(value)
	C.SetReturnFloat(__ret, __value)
}

// SetReturnDouble 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnDouble(ret ReturnHandle, value float64) {
	__ret := C.uintptr_t(ret)
	__value := C.double(value)
	C.SetReturnDouble(__ret, __value)
}

// SetReturnPointer 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnPointer(ret ReturnHandle, value uintptr) {
	__ret := C.uintptr_t(ret)
	__value := C.uintptr_t(value)
	C.SetReturnPointer(__ret, __value)
}

// SetReturnString 
//  @brief Set return value
//
//  @param hook: Hook pointer
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnString(hook HookHandle, ret ReturnHandle, value string) {
	__hook := C.uintptr_t(hook)
	__ret := C.uintptr_t(ret)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			C.SetReturnString(__hook, __ret, (*C.String)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetReturn 
//  @brief Set return value
//
//  @param hook: Hook pointer
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturn(hook HookHandle, ret ReturnHandle, value any) {
	__hook := C.uintptr_t(hook)
	__ret := C.uintptr_t(ret)
	__value := plugify.ConstructVariant(value)
	plugify.Block {
		Try: func() {
			C.SetReturn(__hook, __ret, (*C.Variant)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__value)
		},
	}.Do()
}

// SetRegisterBool 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterBool(registers RegistersHandle, reg RegisterType, value bool) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.bool(value)
	C.SetRegisterBool(__registers, __reg, __value)
}

// SetRegisterInt8 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterInt8(registers RegistersHandle, reg RegisterType, value int8) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.int8_t(value)
	C.SetRegisterInt8(__registers, __reg, __value)
}

// SetRegisterUInt8 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterUInt8(registers RegistersHandle, reg RegisterType, value uint8) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.uint8_t(value)
	C.SetRegisterUInt8(__registers, __reg, __value)
}

// SetRegisterInt16 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterInt16(registers RegistersHandle, reg RegisterType, value int16) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.int16_t(value)
	C.SetRegisterInt16(__registers, __reg, __value)
}

// SetRegisterUInt16 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterUInt16(registers RegistersHandle, reg RegisterType, value uint16) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.uint16_t(value)
	C.SetRegisterUInt16(__registers, __reg, __value)
}

// SetRegisterInt32 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterInt32(registers RegistersHandle, reg RegisterType, value int32) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.int32_t(value)
	C.SetRegisterInt32(__registers, __reg, __value)
}

// SetRegisterUInt32 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterUInt32(registers RegistersHandle, reg RegisterType, value uint32) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.uint32_t(value)
	C.SetRegisterUInt32(__registers, __reg, __value)
}

// SetRegisterInt64 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterInt64(registers RegistersHandle, reg RegisterType, value int64) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.int64_t(value)
	C.SetRegisterInt64(__registers, __reg, __value)
}

// SetRegisterUInt64 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterUInt64(registers RegistersHandle, reg RegisterType, value uint64) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.uint64_t(value)
	C.SetRegisterUInt64(__registers, __reg, __value)
}

// SetRegisterFloat 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterFloat(registers RegistersHandle, reg RegisterType, value float32) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.float(value)
	C.SetRegisterFloat(__registers, __reg, __value)
}

// SetRegisterDouble 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterDouble(registers RegistersHandle, reg RegisterType, value float64) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.double(value)
	C.SetRegisterDouble(__registers, __reg, __value)
}

// SetRegisterPointer 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterPointer(registers RegistersHandle, reg RegisterType, value uintptr) {
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := C.uintptr_t(value)
	C.SetRegisterPointer(__registers, __reg, __value)
}

// SetRegisterString 
//  @brief Set register value
//
//  @param hook: Hook pointer
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegisterString(hook HookHandle, registers RegistersHandle, reg RegisterType, value string) {
	__hook := C.uintptr_t(hook)
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := plugify.ConstructString(value)
	plugify.Block {
		Try: func() {
			C.SetRegisterString(__hook, __registers, __reg, (*C.String)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__value)
		},
	}.Do()
}

// SetRegister 
//  @brief Set register value
//
//  @param hook: Hook pointer
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func SetRegister(hook HookHandle, registers RegistersHandle, reg RegisterType, value any) {
	__hook := C.uintptr_t(hook)
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__value := plugify.ConstructVariant(value)
	plugify.Block {
		Try: func() {
			C.SetRegister(__hook, __registers, __reg, (*C.Variant)(unsafe.Pointer(&__value)))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__value)
		},
	}.Do()
}

