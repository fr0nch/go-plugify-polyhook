package polyhook

/*
#include "getters.h"
#cgo noescape GetFunctionAddr
#cgo noescape GetOriginalAddr
#cgo noescape GetArgumentBool
#cgo noescape GetArgumentInt8
#cgo noescape GetArgumentUInt8
#cgo noescape GetArgumentInt16
#cgo noescape GetArgumentUInt16
#cgo noescape GetArgumentInt32
#cgo noescape GetArgumentUInt32
#cgo noescape GetArgumentInt64
#cgo noescape GetArgumentUInt64
#cgo noescape GetArgumentFloat
#cgo noescape GetArgumentDouble
#cgo noescape GetArgumentPointer
#cgo noescape GetArgumentString
#cgo noescape GetArgument
#cgo noescape GetReturnBool
#cgo noescape GetReturnInt8
#cgo noescape GetReturnUInt8
#cgo noescape GetReturnInt16
#cgo noescape GetReturnUInt16
#cgo noescape GetReturnInt32
#cgo noescape GetReturnUInt32
#cgo noescape GetReturnInt64
#cgo noescape GetReturnUInt64
#cgo noescape GetReturnFloat
#cgo noescape GetReturnDouble
#cgo noescape GetReturnPointer
#cgo noescape GetReturnString
#cgo noescape GetReturn
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

// Generated from polyhook (group: getters)

// GetFunctionAddr 
//  @brief Get generated function address
//
//  @param hook: Hook pointer
//
//  @return Returns jit generated function pointer
func GetFunctionAddr(hook uintptr) uintptr {
	var __retVal uintptr
	__hook := C.uintptr_t(hook)
	__retVal = uintptr(C.GetFunctionAddr(__hook))
	return __retVal
}

// GetOriginalAddr 
//  @brief Get original function address
//
//  @param hook: Hook pointer
//
//  @return Returns original function pointer
func GetOriginalAddr(hook uintptr) uintptr {
	var __retVal uintptr
	__hook := C.uintptr_t(hook)
	__retVal = uintptr(C.GetOriginalAddr(__hook))
	return __retVal
}

// GetArgumentBool 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentBool(params uintptr, index uint64) bool {
	var __retVal bool
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = bool(C.GetArgumentBool(__params, __index))
	return __retVal
}

// GetArgumentInt8 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentInt8(params uintptr, index uint64) int8 {
	var __retVal int8
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = int8(C.GetArgumentInt8(__params, __index))
	return __retVal
}

// GetArgumentUInt8 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentUInt8(params uintptr, index uint64) uint8 {
	var __retVal uint8
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uint8(C.GetArgumentUInt8(__params, __index))
	return __retVal
}

// GetArgumentInt16 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentInt16(params uintptr, index uint64) int16 {
	var __retVal int16
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = int16(C.GetArgumentInt16(__params, __index))
	return __retVal
}

// GetArgumentUInt16 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentUInt16(params uintptr, index uint64) uint16 {
	var __retVal uint16
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uint16(C.GetArgumentUInt16(__params, __index))
	return __retVal
}

// GetArgumentInt32 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentInt32(params uintptr, index uint64) int32 {
	var __retVal int32
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = int32(C.GetArgumentInt32(__params, __index))
	return __retVal
}

// GetArgumentUInt32 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentUInt32(params uintptr, index uint64) uint32 {
	var __retVal uint32
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uint32(C.GetArgumentUInt32(__params, __index))
	return __retVal
}

// GetArgumentInt64 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentInt64(params uintptr, index uint64) int64 {
	var __retVal int64
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = int64(C.GetArgumentInt64(__params, __index))
	return __retVal
}

// GetArgumentUInt64 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentUInt64(params uintptr, index uint64) uint64 {
	var __retVal uint64
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uint64(C.GetArgumentUInt64(__params, __index))
	return __retVal
}

// GetArgumentFloat 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentFloat(params uintptr, index uint64) float32 {
	var __retVal float32
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = float32(C.GetArgumentFloat(__params, __index))
	return __retVal
}

// GetArgumentDouble 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentDouble(params uintptr, index uint64) float64 {
	var __retVal float64
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = float64(C.GetArgumentDouble(__params, __index))
	return __retVal
}

// GetArgumentPointer 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentPointer(params uintptr, index uint64) uintptr {
	var __retVal uintptr
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uintptr(C.GetArgumentPointer(__params, __index))
	return __retVal
}

// GetArgumentString 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgumentString(params uintptr, index uint64) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	plugify.Block {
		Try: func() {
			__native := C.GetArgumentString(__params, __index)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetArgument 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Value to set
func GetArgument(params uintptr, index uint64) any {
	var __retVal any
	var __retVal_native plugify.PlgVariant
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	plugify.Block {
		Try: func() {
			__native := C.GetArgument(__params, __index)
			__retVal_native = *(*plugify.PlgVariant)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVariantData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetReturnBool 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnBool(ret uintptr) bool {
	var __retVal bool
	__ret := C.uintptr_t(ret)
	__retVal = bool(C.GetReturnBool(__ret))
	return __retVal
}

// GetReturnInt8 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnInt8(ret uintptr) int8 {
	var __retVal int8
	__ret := C.uintptr_t(ret)
	__retVal = int8(C.GetReturnInt8(__ret))
	return __retVal
}

// GetReturnUInt8 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnUInt8(ret uintptr) uint8 {
	var __retVal uint8
	__ret := C.uintptr_t(ret)
	__retVal = uint8(C.GetReturnUInt8(__ret))
	return __retVal
}

// GetReturnInt16 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnInt16(ret uintptr) int16 {
	var __retVal int16
	__ret := C.uintptr_t(ret)
	__retVal = int16(C.GetReturnInt16(__ret))
	return __retVal
}

// GetReturnUInt16 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnUInt16(ret uintptr) uint16 {
	var __retVal uint16
	__ret := C.uintptr_t(ret)
	__retVal = uint16(C.GetReturnUInt16(__ret))
	return __retVal
}

// GetReturnInt32 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnInt32(ret uintptr) int32 {
	var __retVal int32
	__ret := C.uintptr_t(ret)
	__retVal = int32(C.GetReturnInt32(__ret))
	return __retVal
}

// GetReturnUInt32 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnUInt32(ret uintptr) uint32 {
	var __retVal uint32
	__ret := C.uintptr_t(ret)
	__retVal = uint32(C.GetReturnUInt32(__ret))
	return __retVal
}

// GetReturnInt64 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnInt64(ret uintptr) int64 {
	var __retVal int64
	__ret := C.uintptr_t(ret)
	__retVal = int64(C.GetReturnInt64(__ret))
	return __retVal
}

// GetReturnUInt64 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnUInt64(ret uintptr) uint64 {
	var __retVal uint64
	__ret := C.uintptr_t(ret)
	__retVal = uint64(C.GetReturnUInt64(__ret))
	return __retVal
}

// GetReturnFloat 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnFloat(ret uintptr) float32 {
	var __retVal float32
	__ret := C.uintptr_t(ret)
	__retVal = float32(C.GetReturnFloat(__ret))
	return __retVal
}

// GetReturnDouble 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnDouble(ret uintptr) float64 {
	var __retVal float64
	__ret := C.uintptr_t(ret)
	__retVal = float64(C.GetReturnDouble(__ret))
	return __retVal
}

// GetReturnPointer 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnPointer(ret uintptr) uintptr {
	var __retVal uintptr
	__ret := C.uintptr_t(ret)
	__retVal = uintptr(C.GetReturnPointer(__ret))
	return __retVal
}

// GetReturnString 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnString(ret uintptr) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__ret := C.uintptr_t(ret)
	plugify.Block {
		Try: func() {
			__native := C.GetReturnString(__ret)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetReturn 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturn(ret uintptr) any {
	var __retVal any
	var __retVal_native plugify.PlgVariant
	__ret := C.uintptr_t(ret)
	plugify.Block {
		Try: func() {
			__native := C.GetReturn(__ret)
			__retVal_native = *(*plugify.PlgVariant)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetVariantData(&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVariant(&__retVal_native)
		},
	}.Do()
	return __retVal
}

