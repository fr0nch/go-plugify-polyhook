package polyhook

/*
#include "core.h"
#cgo noescape HookDetour
#cgo noescape HookDetour2
#cgo noescape HookVirtualTable
#cgo noescape HookVirtualTable2
#cgo noescape HookVirtualFunc
#cgo noescape HookVirtualFunc2
#cgo noescape UnhookDetour
#cgo noescape UnhookVirtualTable
#cgo noescape UnhookVirtualTable2
#cgo noescape UnhookVirtualFunc
#cgo noescape UnhookVirtualFunc2
#cgo noescape UnhookAll
#cgo noescape UnhookAllVirtual
#cgo noescape AddCallback
#cgo noescape AddCallback2
#cgo noescape RemoveCallback
#cgo noescape IsCallbackRegistered
#cgo noescape AreCallbacksRegistered
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

// Generated from polyhook (group: core)

// HookDetour 
//  @brief Sets a detour hook
//
//  @param pFunc: Function address
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func HookDetour(pFunc uintptr, returnType DataType, arguments []DataType, varIndex int32) HookHandle {
	var __retVal HookHandle
	__pFunc := C.uintptr_t(pFunc)
	__returnType := C.uint8_t(returnType)
	__arguments := plugify.ConstructVectorUInt8(arguments)
	__varIndex := C.int32_t(varIndex)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.HookDetour(__pFunc, __returnType, (*C.Vector)(unsafe.Pointer(&__arguments)), __varIndex))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__arguments)
		},
	}.Do()
	return __retVal
}

// HookDetour2 
//  @brief Sets a mid hook
//
//  @param pFunc: Function address
//
//  @return Returns hook pointer
func HookDetour2(pFunc uintptr) HookHandle {
	var __retVal HookHandle
	__pFunc := C.uintptr_t(pFunc)
	__retVal = uintptr(C.HookDetour2(__pFunc))
	return __retVal
}

// HookVirtualTable 
//  @brief Sets a virtual table hook
//
//  @param pClass: Object pointer
//  @param index: Vtable offset
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func HookVirtualTable(pClass uintptr, index int32, returnType DataType, arguments []DataType, varIndex int32) HookHandle {
	var __retVal HookHandle
	__pClass := C.uintptr_t(pClass)
	__index := C.int32_t(index)
	__returnType := C.uint8_t(returnType)
	__arguments := plugify.ConstructVectorUInt8(arguments)
	__varIndex := C.int32_t(varIndex)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.HookVirtualTable(__pClass, __index, __returnType, (*C.Vector)(unsafe.Pointer(&__arguments)), __varIndex))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__arguments)
		},
	}.Do()
	return __retVal
}

// HookVirtualTable2 
//  @brief Sets a virtual table hook by pointer
//
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func HookVirtualTable2(pClass uintptr, pFunc uintptr, returnType DataType, arguments []DataType, varIndex int32) HookHandle {
	var __retVal HookHandle
	__pClass := C.uintptr_t(pClass)
	__pFunc := C.uintptr_t(pFunc)
	__returnType := C.uint8_t(returnType)
	__arguments := plugify.ConstructVectorUInt8(arguments)
	__varIndex := C.int32_t(varIndex)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.HookVirtualTable2(__pClass, __pFunc, __returnType, (*C.Vector)(unsafe.Pointer(&__arguments)), __varIndex))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__arguments)
		},
	}.Do()
	return __retVal
}

// HookVirtualFunc 
//  @brief Sets a virtual function hook
//
//  @param pClass: Object pointer
//  @param index: Vtable offset
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func HookVirtualFunc(pClass uintptr, index int32, returnType DataType, arguments []DataType, varIndex int32) HookHandle {
	var __retVal HookHandle
	__pClass := C.uintptr_t(pClass)
	__index := C.int32_t(index)
	__returnType := C.uint8_t(returnType)
	__arguments := plugify.ConstructVectorUInt8(arguments)
	__varIndex := C.int32_t(varIndex)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.HookVirtualFunc(__pClass, __index, __returnType, (*C.Vector)(unsafe.Pointer(&__arguments)), __varIndex))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__arguments)
		},
	}.Do()
	return __retVal
}

// HookVirtualFunc2 
//  @brief Sets a virtual function hook by pointer
//
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func HookVirtualFunc2(pClass uintptr, pFunc uintptr, returnType DataType, arguments []DataType, varIndex int32) HookHandle {
	var __retVal HookHandle
	__pClass := C.uintptr_t(pClass)
	__pFunc := C.uintptr_t(pFunc)
	__returnType := C.uint8_t(returnType)
	__arguments := plugify.ConstructVectorUInt8(arguments)
	__varIndex := C.int32_t(varIndex)
	plugify.Block {
		Try: func() {
			__retVal = uintptr(C.HookVirtualFunc2(__pClass, __pFunc, __returnType, (*C.Vector)(unsafe.Pointer(&__arguments)), __varIndex))
		},
		Finally: func() {
			// Perform cleanup.
			plugify.DestroyVectorUInt8(&__arguments)
		},
	}.Do()
	return __retVal
}

// UnhookDetour 
//  @brief Removes a detour hook
//
//  @param pFunc: Function address
//
//  @return Returns true on success, false otherwise
func UnhookDetour(pFunc uintptr) bool {
	var __retVal bool
	__pFunc := C.uintptr_t(pFunc)
	__retVal = bool(C.UnhookDetour(__pFunc))
	return __retVal
}

// UnhookVirtualTable 
//  @brief Removes a virtual hook table
//
//  @param pClass: Object pointer
//  @param index: Virtual table index
//
//  @return Returns true on success, false otherwise
func UnhookVirtualTable(pClass uintptr, index int32) bool {
	var __retVal bool
	__pClass := C.uintptr_t(pClass)
	__index := C.int32_t(index)
	__retVal = bool(C.UnhookVirtualTable(__pClass, __index))
	return __retVal
}

// UnhookVirtualTable2 
//  @brief Removes a virtual table hook by pointer
//
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//
//  @return Returns true on success, false otherwise
func UnhookVirtualTable2(pClass uintptr, pFunc uintptr) bool {
	var __retVal bool
	__pClass := C.uintptr_t(pClass)
	__pFunc := C.uintptr_t(pFunc)
	__retVal = bool(C.UnhookVirtualTable2(__pClass, __pFunc))
	return __retVal
}

// UnhookVirtualFunc 
//  @brief Removes a virtual function table
//
//  @param pClass: Object pointer
//  @param index: Virtual table index
//
//  @return Returns true on success, false otherwise
func UnhookVirtualFunc(pClass uintptr, index int32) bool {
	var __retVal bool
	__pClass := C.uintptr_t(pClass)
	__index := C.int32_t(index)
	__retVal = bool(C.UnhookVirtualFunc(__pClass, __index))
	return __retVal
}

// UnhookVirtualFunc2 
//  @brief Removes a virtual function hook by pointer
//
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//
//  @return Returns true on success, false otherwise
func UnhookVirtualFunc2(pClass uintptr, pFunc uintptr) bool {
	var __retVal bool
	__pClass := C.uintptr_t(pClass)
	__pFunc := C.uintptr_t(pFunc)
	__retVal = bool(C.UnhookVirtualFunc2(__pClass, __pFunc))
	return __retVal
}

// UnhookAll 
//  @brief Removes all previously set hooks
//
func UnhookAll() {
	C.UnhookAll()
}

// UnhookAllVirtual 
//  @brief Removes all previously set hooks on the object
//
//  @param pClass: Object pointer
func UnhookAllVirtual(pClass uintptr) {
	__pClass := C.uintptr_t(pClass)
	C.UnhookAllVirtual(__pClass)
}

// AddCallback 
//  @brief Adds a callback to existing hook
//
//  @param hook: Hook pointer
//  @param type_: Whether the hook was in post mode (after processing) or pre mode (before processing).
//  @param handler: Callback function which trigger by hook.
//
//  @return Returns true on success, false otherwise
func AddCallback(hook HookHandle, type_ CallbackType, handler CallbackHandler) bool {
	var __retVal bool
	__hook := C.uintptr_t(hook)
	__type_ := C.uint8_t(type_)
	__handler := plugify.GetFunctionPointerForDelegate(handler)
	__retVal = bool(C.AddCallback(__hook, __type_, __handler))
	return __retVal
}

// AddCallback2 
//  @brief Adds a callback to existing hook
//
//  @param hook: Hook pointer
//  @param type_: Whether the hook was in post mode (after processing) or pre mode (before processing).
//  @param handler: Callback function which trigger by hook.
//  @param priority: Priority of callback among others
//
//  @return Returns true on success, false otherwise
func AddCallback2(hook HookHandle, type_ CallbackType, handler CallbackHandler, priority int32) bool {
	var __retVal bool
	__hook := C.uintptr_t(hook)
	__type_ := C.uint8_t(type_)
	__handler := plugify.GetFunctionPointerForDelegate(handler)
	__priority := C.int32_t(priority)
	__retVal = bool(C.AddCallback2(__hook, __type_, __handler, __priority))
	return __retVal
}

// RemoveCallback 
//  @brief Removes a callback from existing hook
//
//  @param hook: Hook pointer
//  @param type_: Whether the hook was in post mode (after processing) or pre mode (before processing).
//  @param handler: Callback function which trigger by hook.
//
//  @return Returns true on success, false otherwise
func RemoveCallback(hook HookHandle, type_ CallbackType, handler CallbackHandler) bool {
	var __retVal bool
	__hook := C.uintptr_t(hook)
	__type_ := C.uint8_t(type_)
	__handler := plugify.GetFunctionPointerForDelegate(handler)
	__retVal = bool(C.RemoveCallback(__hook, __type_, __handler))
	return __retVal
}

// IsCallbackRegistered 
//  @brief Checks that a callback exists on existing hook
//
//  @param hook: Hook pointer
//  @param type_: Whether the hook was in post mode (after processing) or pre mode (before processing).
//  @param handler: Callback function which trigger by hook.
//
//  @return Returns true on success, false otherwise
func IsCallbackRegistered(hook HookHandle, type_ CallbackType, handler CallbackHandler) bool {
	var __retVal bool
	__hook := C.uintptr_t(hook)
	__type_ := C.uint8_t(type_)
	__handler := plugify.GetFunctionPointerForDelegate(handler)
	__retVal = bool(C.IsCallbackRegistered(__hook, __type_, __handler))
	return __retVal
}

// AreCallbacksRegistered 
//  @brief Checks that a any callback exists on existing hook
//
//  @param hook: Hook pointer
//
//  @return Returns true on success, false otherwise
func AreCallbacksRegistered(hook HookHandle) bool {
	var __retVal bool
	__hook := C.uintptr_t(hook)
	__retVal = bool(C.AreCallbacksRegistered(__hook))
	return __retVal
}

