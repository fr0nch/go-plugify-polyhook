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
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentBool(params uintptr, index uint64, value bool) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.bool(value)
	C.SetArgumentBool(__params, __index, __value)
}

// SetArgumentInt8 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentInt8(params uintptr, index uint64, value int8) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.int8_t(value)
	C.SetArgumentInt8(__params, __index, __value)
}

// SetArgumentUInt8 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentUInt8(params uintptr, index uint64, value uint8) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uint8_t(value)
	C.SetArgumentUInt8(__params, __index, __value)
}

// SetArgumentInt16 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentInt16(params uintptr, index uint64, value int16) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.int16_t(value)
	C.SetArgumentInt16(__params, __index, __value)
}

// SetArgumentUInt16 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentUInt16(params uintptr, index uint64, value uint16) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uint16_t(value)
	C.SetArgumentUInt16(__params, __index, __value)
}

// SetArgumentInt32 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentInt32(params uintptr, index uint64, value int32) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.int32_t(value)
	C.SetArgumentInt32(__params, __index, __value)
}

// SetArgumentUInt32 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentUInt32(params uintptr, index uint64, value uint32) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uint32_t(value)
	C.SetArgumentUInt32(__params, __index, __value)
}

// SetArgumentInt64 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentInt64(params uintptr, index uint64, value int64) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.int64_t(value)
	C.SetArgumentInt64(__params, __index, __value)
}

// SetArgumentUInt64 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentUInt64(params uintptr, index uint64, value uint64) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.uint64_t(value)
	C.SetArgumentUInt64(__params, __index, __value)
}

// SetArgumentFloat 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentFloat(params uintptr, index uint64, value float32) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.float(value)
	C.SetArgumentFloat(__params, __index, __value)
}

// SetArgumentDouble 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentDouble(params uintptr, index uint64, value float64) {
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__value := C.double(value)
	C.SetArgumentDouble(__params, __index, __value)
}

// SetArgumentPointer 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentPointer(params uintptr, index uint64, value uintptr) {
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
//  @param index: Value to set
//  @param value: Value to set
func SetArgumentString(hook uintptr, params uintptr, index uint64, value string) {
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
func SetArgument(hook uintptr, params uintptr, index uint64, value any) {
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
func SetReturnBool(ret uintptr, value bool) {
	__ret := C.uintptr_t(ret)
	__value := C.bool(value)
	C.SetReturnBool(__ret, __value)
}

// SetReturnInt8 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnInt8(ret uintptr, value int8) {
	__ret := C.uintptr_t(ret)
	__value := C.int8_t(value)
	C.SetReturnInt8(__ret, __value)
}

// SetReturnUInt8 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnUInt8(ret uintptr, value uint8) {
	__ret := C.uintptr_t(ret)
	__value := C.uint8_t(value)
	C.SetReturnUInt8(__ret, __value)
}

// SetReturnInt16 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnInt16(ret uintptr, value int16) {
	__ret := C.uintptr_t(ret)
	__value := C.int16_t(value)
	C.SetReturnInt16(__ret, __value)
}

// SetReturnUInt16 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnUInt16(ret uintptr, value uint16) {
	__ret := C.uintptr_t(ret)
	__value := C.uint16_t(value)
	C.SetReturnUInt16(__ret, __value)
}

// SetReturnInt32 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnInt32(ret uintptr, value int32) {
	__ret := C.uintptr_t(ret)
	__value := C.int32_t(value)
	C.SetReturnInt32(__ret, __value)
}

// SetReturnUInt32 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnUInt32(ret uintptr, value uint32) {
	__ret := C.uintptr_t(ret)
	__value := C.uint32_t(value)
	C.SetReturnUInt32(__ret, __value)
}

// SetReturnInt64 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnInt64(ret uintptr, value int64) {
	__ret := C.uintptr_t(ret)
	__value := C.int64_t(value)
	C.SetReturnInt64(__ret, __value)
}

// SetReturnUInt64 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnUInt64(ret uintptr, value uint64) {
	__ret := C.uintptr_t(ret)
	__value := C.uint64_t(value)
	C.SetReturnUInt64(__ret, __value)
}

// SetReturnFloat 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnFloat(ret uintptr, value float32) {
	__ret := C.uintptr_t(ret)
	__value := C.float(value)
	C.SetReturnFloat(__ret, __value)
}

// SetReturnDouble 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnDouble(ret uintptr, value float64) {
	__ret := C.uintptr_t(ret)
	__value := C.double(value)
	C.SetReturnDouble(__ret, __value)
}

// SetReturnPointer 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func SetReturnPointer(ret uintptr, value uintptr) {
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
func SetReturnString(hook uintptr, ret uintptr, value string) {
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
func SetReturn(hook uintptr, ret uintptr, value any) {
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

