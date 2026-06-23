package polyhook

/*
#include "getters.h"
#cgo noescape GetError
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
#cgo noescape GetRegisterBool
#cgo noescape GetRegisterInt8
#cgo noescape GetRegisterUInt8
#cgo noescape GetRegisterInt16
#cgo noescape GetRegisterUInt16
#cgo noescape GetRegisterInt32
#cgo noescape GetRegisterUInt32
#cgo noescape GetRegisterInt64
#cgo noescape GetRegisterUInt64
#cgo noescape GetRegisterFloat
#cgo noescape GetRegisterDouble
#cgo noescape GetRegisterPointer
#cgo noescape GetRegisterString
#cgo noescape GetRegister
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
var _ = plugify.ApiVersion

// Generated from polyhook (group: getters)

var _GetError = func() string {
	var __retVal string
	var __retVal_native plugify.PlgString
	plugify.Block {
		Try: func() {
			__native := C.GetError()
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetError 
//  @brief Get last error
//
//
//  @return Returns last error
func GetError() string {
	return _GetError()
}

var _GetFunctionAddr = func(hook HookHandle) uintptr {
	var __retVal uintptr
	__hook := C.uintptr_t(hook)
	__retVal = uintptr(C.GetFunctionAddr(__hook))
	return __retVal
}

// GetFunctionAddr 
//  @brief Get generated function address
//
//  @param hook: Hook pointer
//
//  @return Returns jit generated function pointer
func GetFunctionAddr(hook HookHandle) uintptr {
	return _GetFunctionAddr(hook)
}

var _GetOriginalAddr = func(hook HookHandle) uintptr {
	var __retVal uintptr
	__hook := C.uintptr_t(hook)
	__retVal = uintptr(C.GetOriginalAddr(__hook))
	return __retVal
}

// GetOriginalAddr 
//  @brief Get original function address
//
//  @param hook: Hook pointer
//
//  @return Returns original function pointer
func GetOriginalAddr(hook HookHandle) uintptr {
	return _GetOriginalAddr(hook)
}

var _GetArgumentBool = func(params ParametersHandle, index uint64) bool {
	var __retVal bool
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = bool(C.GetArgumentBool(__params, __index))
	return __retVal
}

// GetArgumentBool 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentBool(params ParametersHandle, index uint64) bool {
	return _GetArgumentBool(params, index)
}

var _GetArgumentInt8 = func(params ParametersHandle, index uint64) int8 {
	var __retVal int8
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = int8(C.GetArgumentInt8(__params, __index))
	return __retVal
}

// GetArgumentInt8 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentInt8(params ParametersHandle, index uint64) int8 {
	return _GetArgumentInt8(params, index)
}

var _GetArgumentUInt8 = func(params ParametersHandle, index uint64) uint8 {
	var __retVal uint8
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uint8(C.GetArgumentUInt8(__params, __index))
	return __retVal
}

// GetArgumentUInt8 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentUInt8(params ParametersHandle, index uint64) uint8 {
	return _GetArgumentUInt8(params, index)
}

var _GetArgumentInt16 = func(params ParametersHandle, index uint64) int16 {
	var __retVal int16
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = int16(C.GetArgumentInt16(__params, __index))
	return __retVal
}

// GetArgumentInt16 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentInt16(params ParametersHandle, index uint64) int16 {
	return _GetArgumentInt16(params, index)
}

var _GetArgumentUInt16 = func(params ParametersHandle, index uint64) uint16 {
	var __retVal uint16
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uint16(C.GetArgumentUInt16(__params, __index))
	return __retVal
}

// GetArgumentUInt16 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentUInt16(params ParametersHandle, index uint64) uint16 {
	return _GetArgumentUInt16(params, index)
}

var _GetArgumentInt32 = func(params ParametersHandle, index uint64) int32 {
	var __retVal int32
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = int32(C.GetArgumentInt32(__params, __index))
	return __retVal
}

// GetArgumentInt32 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentInt32(params ParametersHandle, index uint64) int32 {
	return _GetArgumentInt32(params, index)
}

var _GetArgumentUInt32 = func(params ParametersHandle, index uint64) uint32 {
	var __retVal uint32
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uint32(C.GetArgumentUInt32(__params, __index))
	return __retVal
}

// GetArgumentUInt32 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentUInt32(params ParametersHandle, index uint64) uint32 {
	return _GetArgumentUInt32(params, index)
}

var _GetArgumentInt64 = func(params ParametersHandle, index uint64) int64 {
	var __retVal int64
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = int64(C.GetArgumentInt64(__params, __index))
	return __retVal
}

// GetArgumentInt64 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentInt64(params ParametersHandle, index uint64) int64 {
	return _GetArgumentInt64(params, index)
}

var _GetArgumentUInt64 = func(params ParametersHandle, index uint64) uint64 {
	var __retVal uint64
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uint64(C.GetArgumentUInt64(__params, __index))
	return __retVal
}

// GetArgumentUInt64 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentUInt64(params ParametersHandle, index uint64) uint64 {
	return _GetArgumentUInt64(params, index)
}

var _GetArgumentFloat = func(params ParametersHandle, index uint64) float32 {
	var __retVal float32
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = float32(C.GetArgumentFloat(__params, __index))
	return __retVal
}

// GetArgumentFloat 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentFloat(params ParametersHandle, index uint64) float32 {
	return _GetArgumentFloat(params, index)
}

var _GetArgumentDouble = func(params ParametersHandle, index uint64) float64 {
	var __retVal float64
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = float64(C.GetArgumentDouble(__params, __index))
	return __retVal
}

// GetArgumentDouble 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentDouble(params ParametersHandle, index uint64) float64 {
	return _GetArgumentDouble(params, index)
}

var _GetArgumentPointer = func(params ParametersHandle, index uint64) uintptr {
	var __retVal uintptr
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	__retVal = uintptr(C.GetArgumentPointer(__params, __index))
	return __retVal
}

// GetArgumentPointer 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentPointer(params ParametersHandle, index uint64) uintptr {
	return _GetArgumentPointer(params, index)
}

var _GetArgumentString = func(params ParametersHandle, index uint64) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__params := C.uintptr_t(params)
	__index := C.uint64_t(index)
	plugify.Block {
		Try: func() {
			__native := C.GetArgumentString(__params, __index)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetArgumentString 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgumentString(params ParametersHandle, index uint64) string {
	return _GetArgumentString(params, index)
}

var _GetArgument = func(params ParametersHandle, index uint64) any {
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

// GetArgument 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func GetArgument(params ParametersHandle, index uint64) any {
	return _GetArgument(params, index)
}

var _GetReturnBool = func(ret ReturnHandle) bool {
	var __retVal bool
	__ret := C.uintptr_t(ret)
	__retVal = bool(C.GetReturnBool(__ret))
	return __retVal
}

// GetReturnBool 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnBool(ret ReturnHandle) bool {
	return _GetReturnBool(ret)
}

var _GetReturnInt8 = func(ret ReturnHandle) int8 {
	var __retVal int8
	__ret := C.uintptr_t(ret)
	__retVal = int8(C.GetReturnInt8(__ret))
	return __retVal
}

// GetReturnInt8 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnInt8(ret ReturnHandle) int8 {
	return _GetReturnInt8(ret)
}

var _GetReturnUInt8 = func(ret ReturnHandle) uint8 {
	var __retVal uint8
	__ret := C.uintptr_t(ret)
	__retVal = uint8(C.GetReturnUInt8(__ret))
	return __retVal
}

// GetReturnUInt8 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnUInt8(ret ReturnHandle) uint8 {
	return _GetReturnUInt8(ret)
}

var _GetReturnInt16 = func(ret ReturnHandle) int16 {
	var __retVal int16
	__ret := C.uintptr_t(ret)
	__retVal = int16(C.GetReturnInt16(__ret))
	return __retVal
}

// GetReturnInt16 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnInt16(ret ReturnHandle) int16 {
	return _GetReturnInt16(ret)
}

var _GetReturnUInt16 = func(ret ReturnHandle) uint16 {
	var __retVal uint16
	__ret := C.uintptr_t(ret)
	__retVal = uint16(C.GetReturnUInt16(__ret))
	return __retVal
}

// GetReturnUInt16 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnUInt16(ret ReturnHandle) uint16 {
	return _GetReturnUInt16(ret)
}

var _GetReturnInt32 = func(ret ReturnHandle) int32 {
	var __retVal int32
	__ret := C.uintptr_t(ret)
	__retVal = int32(C.GetReturnInt32(__ret))
	return __retVal
}

// GetReturnInt32 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnInt32(ret ReturnHandle) int32 {
	return _GetReturnInt32(ret)
}

var _GetReturnUInt32 = func(ret ReturnHandle) uint32 {
	var __retVal uint32
	__ret := C.uintptr_t(ret)
	__retVal = uint32(C.GetReturnUInt32(__ret))
	return __retVal
}

// GetReturnUInt32 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnUInt32(ret ReturnHandle) uint32 {
	return _GetReturnUInt32(ret)
}

var _GetReturnInt64 = func(ret ReturnHandle) int64 {
	var __retVal int64
	__ret := C.uintptr_t(ret)
	__retVal = int64(C.GetReturnInt64(__ret))
	return __retVal
}

// GetReturnInt64 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnInt64(ret ReturnHandle) int64 {
	return _GetReturnInt64(ret)
}

var _GetReturnUInt64 = func(ret ReturnHandle) uint64 {
	var __retVal uint64
	__ret := C.uintptr_t(ret)
	__retVal = uint64(C.GetReturnUInt64(__ret))
	return __retVal
}

// GetReturnUInt64 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnUInt64(ret ReturnHandle) uint64 {
	return _GetReturnUInt64(ret)
}

var _GetReturnFloat = func(ret ReturnHandle) float32 {
	var __retVal float32
	__ret := C.uintptr_t(ret)
	__retVal = float32(C.GetReturnFloat(__ret))
	return __retVal
}

// GetReturnFloat 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnFloat(ret ReturnHandle) float32 {
	return _GetReturnFloat(ret)
}

var _GetReturnDouble = func(ret ReturnHandle) float64 {
	var __retVal float64
	__ret := C.uintptr_t(ret)
	__retVal = float64(C.GetReturnDouble(__ret))
	return __retVal
}

// GetReturnDouble 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnDouble(ret ReturnHandle) float64 {
	return _GetReturnDouble(ret)
}

var _GetReturnPointer = func(ret ReturnHandle) uintptr {
	var __retVal uintptr
	__ret := C.uintptr_t(ret)
	__retVal = uintptr(C.GetReturnPointer(__ret))
	return __retVal
}

// GetReturnPointer 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnPointer(ret ReturnHandle) uintptr {
	return _GetReturnPointer(ret)
}

var _GetReturnString = func(ret ReturnHandle) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__ret := C.uintptr_t(ret)
	plugify.Block {
		Try: func() {
			__native := C.GetReturnString(__ret)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetReturnString 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturnString(ret ReturnHandle) string {
	return _GetReturnString(ret)
}

var _GetReturn = func(ret ReturnHandle) any {
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

// GetReturn 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func GetReturn(ret ReturnHandle) any {
	return _GetReturn(ret)
}

var _GetRegisterBool = func(registers RegistersHandle, reg RegisterType) bool {
	var __retVal bool
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = bool(C.GetRegisterBool(__registers, __reg))
	return __retVal
}

// GetRegisterBool 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterBool(registers RegistersHandle, reg RegisterType) bool {
	return _GetRegisterBool(registers, reg)
}

var _GetRegisterInt8 = func(registers RegistersHandle, reg RegisterType) int8 {
	var __retVal int8
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = int8(C.GetRegisterInt8(__registers, __reg))
	return __retVal
}

// GetRegisterInt8 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterInt8(registers RegistersHandle, reg RegisterType) int8 {
	return _GetRegisterInt8(registers, reg)
}

var _GetRegisterUInt8 = func(registers RegistersHandle, reg RegisterType) uint8 {
	var __retVal uint8
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = uint8(C.GetRegisterUInt8(__registers, __reg))
	return __retVal
}

// GetRegisterUInt8 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterUInt8(registers RegistersHandle, reg RegisterType) uint8 {
	return _GetRegisterUInt8(registers, reg)
}

var _GetRegisterInt16 = func(registers RegistersHandle, reg RegisterType) int16 {
	var __retVal int16
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = int16(C.GetRegisterInt16(__registers, __reg))
	return __retVal
}

// GetRegisterInt16 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterInt16(registers RegistersHandle, reg RegisterType) int16 {
	return _GetRegisterInt16(registers, reg)
}

var _GetRegisterUInt16 = func(registers RegistersHandle, reg RegisterType) uint16 {
	var __retVal uint16
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = uint16(C.GetRegisterUInt16(__registers, __reg))
	return __retVal
}

// GetRegisterUInt16 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterUInt16(registers RegistersHandle, reg RegisterType) uint16 {
	return _GetRegisterUInt16(registers, reg)
}

var _GetRegisterInt32 = func(registers RegistersHandle, reg RegisterType) int32 {
	var __retVal int32
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = int32(C.GetRegisterInt32(__registers, __reg))
	return __retVal
}

// GetRegisterInt32 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterInt32(registers RegistersHandle, reg RegisterType) int32 {
	return _GetRegisterInt32(registers, reg)
}

var _GetRegisterUInt32 = func(registers RegistersHandle, reg RegisterType) uint32 {
	var __retVal uint32
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = uint32(C.GetRegisterUInt32(__registers, __reg))
	return __retVal
}

// GetRegisterUInt32 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterUInt32(registers RegistersHandle, reg RegisterType) uint32 {
	return _GetRegisterUInt32(registers, reg)
}

var _GetRegisterInt64 = func(registers RegistersHandle, reg RegisterType) int64 {
	var __retVal int64
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = int64(C.GetRegisterInt64(__registers, __reg))
	return __retVal
}

// GetRegisterInt64 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterInt64(registers RegistersHandle, reg RegisterType) int64 {
	return _GetRegisterInt64(registers, reg)
}

var _GetRegisterUInt64 = func(registers RegistersHandle, reg RegisterType) uint64 {
	var __retVal uint64
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = uint64(C.GetRegisterUInt64(__registers, __reg))
	return __retVal
}

// GetRegisterUInt64 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterUInt64(registers RegistersHandle, reg RegisterType) uint64 {
	return _GetRegisterUInt64(registers, reg)
}

var _GetRegisterFloat = func(registers RegistersHandle, reg RegisterType) float32 {
	var __retVal float32
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = float32(C.GetRegisterFloat(__registers, __reg))
	return __retVal
}

// GetRegisterFloat 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterFloat(registers RegistersHandle, reg RegisterType) float32 {
	return _GetRegisterFloat(registers, reg)
}

var _GetRegisterDouble = func(registers RegistersHandle, reg RegisterType) float64 {
	var __retVal float64
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = float64(C.GetRegisterDouble(__registers, __reg))
	return __retVal
}

// GetRegisterDouble 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterDouble(registers RegistersHandle, reg RegisterType) float64 {
	return _GetRegisterDouble(registers, reg)
}

var _GetRegisterPointer = func(registers RegistersHandle, reg RegisterType) uintptr {
	var __retVal uintptr
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	__retVal = uintptr(C.GetRegisterPointer(__registers, __reg))
	return __retVal
}

// GetRegisterPointer 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterPointer(registers RegistersHandle, reg RegisterType) uintptr {
	return _GetRegisterPointer(registers, reg)
}

var _GetRegisterString = func(registers RegistersHandle, reg RegisterType) string {
	var __retVal string
	var __retVal_native plugify.PlgString
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	plugify.Block {
		Try: func() {
			__native := C.GetRegisterString(__registers, __reg)
			__retVal_native = *(*plugify.PlgString)(unsafe.Pointer(&__native))
			// Unmarshal - Convert native data to managed data.
			__retVal = plugify.GetStringData[string](&__retVal_native)
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyString(&__retVal_native)
		},
	}.Do()
	return __retVal
}

// GetRegisterString 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegisterString(registers RegistersHandle, reg RegisterType) string {
	return _GetRegisterString(registers, reg)
}

var _GetRegister = func(registers RegistersHandle, reg RegisterType) any {
	var __retVal any
	var __retVal_native plugify.PlgVariant
	__registers := C.uintptr_t(registers)
	__reg := C.uint64_t(reg)
	plugify.Block {
		Try: func() {
			__native := C.GetRegister(__registers, __reg)
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

// GetRegister 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func GetRegister(registers RegistersHandle, reg RegisterType) any {
	return _GetRegister(registers, reg)
}

