package polyhook

/*
#include "lookup.h"
#cgo noescape FindDetour
#cgo noescape FindVirtual
#cgo noescape FindVirtual2
#cgo noescape GetVirtualIndex
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

// Generated from polyhook (group: lookup)

// FindDetour 
//  @brief Attempts to find existing detour hook
//
//  @param pFunc: Function address
//
//  @return Returns hook pointer
func FindDetour(pFunc uintptr) HookHandle {
	var __retVal HookHandle
	__pFunc := C.uintptr_t(pFunc)
	__retVal = uintptr(C.FindDetour(__pFunc))
	return __retVal
}

// FindVirtual 
//  @brief Attempts to find existing virtual hook
//
//  @param pClass: Object pointer
//  @param index: Virtual table index
//
//  @return Returns hook pointer
func FindVirtual(pClass uintptr, index int32) uintptr {
	var __retVal uintptr
	__pClass := C.uintptr_t(pClass)
	__index := C.int32_t(index)
	__retVal = uintptr(C.FindVirtual(__pClass, __index))
	return __retVal
}

// FindVirtual2 
//  @brief Attempts to find existing detour hook by pointer
//
//  @param pClass: Object pointer
//  @param pFunc: Function address
//
//  @return Returns hook pointer
func FindVirtual2(pClass uintptr, pFunc uintptr) HookHandle {
	var __retVal HookHandle
	__pClass := C.uintptr_t(pClass)
	__pFunc := C.uintptr_t(pFunc)
	__retVal = uintptr(C.FindVirtual2(__pClass, __pFunc))
	return __retVal
}

// GetVirtualIndex 
//  @brief Attempts to find virtual table index of virtual function
//
//  @param pFunc: Function address
//
//  @return Virtual table index
func GetVirtualIndex(pFunc uintptr) int32 {
	var __retVal int32
	__pFunc := C.uintptr_t(pFunc)
	__retVal = int32(C.GetVirtualIndex(__pFunc))
	return __retVal
}

