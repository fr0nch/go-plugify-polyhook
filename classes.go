package polyhook

/*
#include "classes.h"
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

// Generated from polyhook (group: classes)

var (
	CallbackErrEmptyHandle = errors.New("Callback: empty handle")
)

// Callback - RAII wrapper for Callback pointer from hook operations.
type Callback struct {
	handle    uintptr
}

// NewCallback creates a Callback from a handle
func NewCallback(handle uintptr) *Callback {
	return &Callback{
		handle:    handle,
	}
}

// Get returns the underlying handle
func (w *Callback) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *Callback) Release() uintptr {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *Callback) Reset() {
	w.handle = 0
}

// IsValid returns true if handle is not nil
func (w *Callback) IsValid() bool {
	return w.handle != 0
}

// AddCallback - Adds a callback to existing hook
//  @param type_: Whether the hook was in post mode (after processing) or pre mode (before processing).
//  @param handler: Callback function which trigger by hook.
//
//  @return Returns true on success, false otherwise
func (w *Callback) AddCallback(type_ CallbackType, handler CallbackHandler) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CallbackErrEmptyHandle
	}
	return AddCallback(w.handle, type_, handler), nil
}

// AddCallback2 - Adds a callback to existing hook
//  @param type_: Whether the hook was in post mode (after processing) or pre mode (before processing).
//  @param handler: Callback function which trigger by hook.
//  @param priority: Priority of callback among others
//
//  @return Returns true on success, false otherwise
func (w *Callback) AddCallback2(type_ CallbackType, handler CallbackHandler, priority int32) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CallbackErrEmptyHandle
	}
	return AddCallback2(w.handle, type_, handler, priority), nil
}

// RemoveCallback - Removes a callback from existing hook
//  @param type_: Whether the hook was in post mode (after processing) or pre mode (before processing).
//  @param handler: Callback function which trigger by hook.
//
//  @return Returns true on success, false otherwise
func (w *Callback) RemoveCallback(type_ CallbackType, handler CallbackHandler) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CallbackErrEmptyHandle
	}
	return RemoveCallback(w.handle, type_, handler), nil
}

// IsCallbackRegistered - Checks that a callback exists on existing hook
//  @param type_: Whether the hook was in post mode (after processing) or pre mode (before processing).
//  @param handler: Callback function which trigger by hook.
//
//  @return Returns true on success, false otherwise
func (w *Callback) IsCallbackRegistered(type_ CallbackType, handler CallbackHandler) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CallbackErrEmptyHandle
	}
	return IsCallbackRegistered(w.handle, type_, handler), nil
}

// AreCallbacksRegistered - Checks that a any callback exists on existing hook
//
//  @return Returns true on success, false otherwise
func (w *Callback) AreCallbacksRegistered() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CallbackErrEmptyHandle
	}
	return AreCallbacksRegistered(w.handle), nil
}

// GetFunctionAddr - Get generated function address
//
//  @return Returns jit generated function pointer
func (w *Callback) GetFunctionAddr() (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, CallbackErrEmptyHandle
	}
	return GetFunctionAddr(w.handle), nil
}

// GetOriginalAddr - Get original function address
//
//  @return Returns original function pointer
func (w *Callback) GetOriginalAddr() (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, CallbackErrEmptyHandle
	}
	return GetOriginalAddr(w.handle), nil
}

var (
	ParametersErrEmptyHandle = errors.New("Parameters: empty handle")
)

// Parameters - RAII wrapper for Callback::Parameters pointer used in callback handlers.
type Parameters struct {
	handle    uintptr
}

// NewParameters creates a Parameters from a handle
func NewParameters(handle uintptr) *Parameters {
	return &Parameters{
		handle:    handle,
	}
}

// Get returns the underlying handle
func (w *Parameters) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *Parameters) Release() uintptr {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *Parameters) Reset() {
	w.handle = 0
}

// IsValid returns true if handle is not nil
func (w *Parameters) IsValid() bool {
	return w.handle != 0
}

// GetBool - Get argument value
//  @param index: Value to set
func (w *Parameters) GetBool(index uint64) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentBool(w.handle, index), nil
}

// GetInt8 - Get argument value
//  @param index: Value to set
func (w *Parameters) GetInt8(index uint64) (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentInt8(w.handle, index), nil
}

// GetUInt8 - Get argument value
//  @param index: Value to set
func (w *Parameters) GetUInt8(index uint64) (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentUInt8(w.handle, index), nil
}

// GetInt16 - Get argument value
//  @param index: Value to set
func (w *Parameters) GetInt16(index uint64) (int16, error) {
	if w.handle == 0 {
		var zero int16
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentInt16(w.handle, index), nil
}

// GetUInt16 - Get argument value
//  @param index: Value to set
func (w *Parameters) GetUInt16(index uint64) (uint16, error) {
	if w.handle == 0 {
		var zero uint16
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentUInt16(w.handle, index), nil
}

// GetInt32 - Get argument value
//  @param index: Value to set
func (w *Parameters) GetInt32(index uint64) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentInt32(w.handle, index), nil
}

// GetUInt32 - Get argument value
//  @param index: Value to set
func (w *Parameters) GetUInt32(index uint64) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentUInt32(w.handle, index), nil
}

// GetInt64 - Get argument value
//  @param index: Value to set
func (w *Parameters) GetInt64(index uint64) (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentInt64(w.handle, index), nil
}

// GetUInt64 - Get argument value
//  @param index: Value to set
func (w *Parameters) GetUInt64(index uint64) (uint64, error) {
	if w.handle == 0 {
		var zero uint64
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentUInt64(w.handle, index), nil
}

// GetFloat - Get argument value
//  @param index: Value to set
func (w *Parameters) GetFloat(index uint64) (float32, error) {
	if w.handle == 0 {
		var zero float32
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentFloat(w.handle, index), nil
}

// GetDouble - Get argument value
//  @param index: Value to set
func (w *Parameters) GetDouble(index uint64) (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentDouble(w.handle, index), nil
}

// GetPointer - Get argument value
//  @param index: Value to set
func (w *Parameters) GetPointer(index uint64) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentPointer(w.handle, index), nil
}

// GetString - Get argument value
//  @param index: Value to set
func (w *Parameters) GetString(index uint64) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentString(w.handle, index), nil
}

// SetBool - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetBool(index uint64, value bool) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentBool(w.handle, index, value)
	return nil
}

// SetInt8 - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetInt8(index uint64, value int8) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentInt8(w.handle, index, value)
	return nil
}

// SetUInt8 - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetUInt8(index uint64, value uint8) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentUInt8(w.handle, index, value)
	return nil
}

// SetInt16 - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetInt16(index uint64, value int16) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentInt16(w.handle, index, value)
	return nil
}

// SetUInt16 - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetUInt16(index uint64, value uint16) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentUInt16(w.handle, index, value)
	return nil
}

// SetInt32 - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetInt32(index uint64, value int32) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentInt32(w.handle, index, value)
	return nil
}

// SetUInt32 - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetUInt32(index uint64, value uint32) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentUInt32(w.handle, index, value)
	return nil
}

// SetInt64 - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetInt64(index uint64, value int64) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentInt64(w.handle, index, value)
	return nil
}

// SetUInt64 - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetUInt64(index uint64, value uint64) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentUInt64(w.handle, index, value)
	return nil
}

// SetFloat - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetFloat(index uint64, value float32) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentFloat(w.handle, index, value)
	return nil
}

// SetDouble - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetDouble(index uint64, value float64) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentDouble(w.handle, index, value)
	return nil
}

// SetPointer - Set argument value
//  @param index: Value to set
//  @param value: Value to set
func (w *Parameters) SetPointer(index uint64, value uintptr) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentPointer(w.handle, index, value)
	return nil
}

var (
	ReturnErrEmptyHandle = errors.New("Return: empty handle")
)

// Return - RAII wrapper for Callback::Return pointer used in callback handlers.
type Return struct {
	handle    uintptr
}

// NewReturn creates a Return from a handle
func NewReturn(handle uintptr) *Return {
	return &Return{
		handle:    handle,
	}
}

// Get returns the underlying handle
func (w *Return) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *Return) Release() uintptr {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *Return) Reset() {
	w.handle = 0
}

// IsValid returns true if handle is not nil
func (w *Return) IsValid() bool {
	return w.handle != 0
}

// GetBool - Get return value
func (w *Return) GetBool() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnBool(w.handle), nil
}

// GetInt8 - Get return value
func (w *Return) GetInt8() (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnInt8(w.handle), nil
}

// GetUInt8 - Get return value
func (w *Return) GetUInt8() (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnUInt8(w.handle), nil
}

// GetInt16 - Get return value
func (w *Return) GetInt16() (int16, error) {
	if w.handle == 0 {
		var zero int16
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnInt16(w.handle), nil
}

// GetUInt16 - Get return value
func (w *Return) GetUInt16() (uint16, error) {
	if w.handle == 0 {
		var zero uint16
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnUInt16(w.handle), nil
}

// GetInt32 - Get return value
func (w *Return) GetInt32() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnInt32(w.handle), nil
}

// GetUInt32 - Get return value
func (w *Return) GetUInt32() (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnUInt32(w.handle), nil
}

// GetInt64 - Get return value
func (w *Return) GetInt64() (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnInt64(w.handle), nil
}

// GetUInt64 - Get return value
func (w *Return) GetUInt64() (uint64, error) {
	if w.handle == 0 {
		var zero uint64
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnUInt64(w.handle), nil
}

// GetFloat - Get return value
func (w *Return) GetFloat() (float32, error) {
	if w.handle == 0 {
		var zero float32
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnFloat(w.handle), nil
}

// GetDouble - Get return value
func (w *Return) GetDouble() (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnDouble(w.handle), nil
}

// GetPointer - Get return value
func (w *Return) GetPointer() (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnPointer(w.handle), nil
}

// GetString - Get return value
func (w *Return) GetString() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnString(w.handle), nil
}

// SetBool - Set return value
//  @param value: Value to set
func (w *Return) SetBool(value bool) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnBool(w.handle, value)
	return nil
}

// SetInt8 - Set return value
//  @param value: Value to set
func (w *Return) SetInt8(value int8) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnInt8(w.handle, value)
	return nil
}

// SetUInt8 - Set return value
//  @param value: Value to set
func (w *Return) SetUInt8(value uint8) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnUInt8(w.handle, value)
	return nil
}

// SetInt16 - Set return value
//  @param value: Value to set
func (w *Return) SetInt16(value int16) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnInt16(w.handle, value)
	return nil
}

// SetUInt16 - Set return value
//  @param value: Value to set
func (w *Return) SetUInt16(value uint16) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnUInt16(w.handle, value)
	return nil
}

// SetInt32 - Set return value
//  @param value: Value to set
func (w *Return) SetInt32(value int32) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnInt32(w.handle, value)
	return nil
}

// SetUInt32 - Set return value
//  @param value: Value to set
func (w *Return) SetUInt32(value uint32) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnUInt32(w.handle, value)
	return nil
}

// SetInt64 - Set return value
//  @param value: Value to set
func (w *Return) SetInt64(value int64) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnInt64(w.handle, value)
	return nil
}

// SetUInt64 - Set return value
//  @param value: Value to set
func (w *Return) SetUInt64(value uint64) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnUInt64(w.handle, value)
	return nil
}

// SetFloat - Set return value
//  @param value: Value to set
func (w *Return) SetFloat(value float32) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnFloat(w.handle, value)
	return nil
}

// SetDouble - Set return value
//  @param value: Value to set
func (w *Return) SetDouble(value float64) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnDouble(w.handle, value)
	return nil
}

// SetPointer - Set return value
//  @param value: Value to set
func (w *Return) SetPointer(value uintptr) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnPointer(w.handle, value)
	return nil
}

// PolyHook - Global hooking functions for PolyHook plugin.
type PolyHook struct {
}

// HookDetour - Sets a detour hook
//  @param pFunc: Function address
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func (w *PolyHook) HookDetour(pFunc uintptr, returnType DataType, arguments []DataType, varIndex int32) *Callback {
	return NewCallback(HookDetour(pFunc, returnType, arguments, varIndex))
}

// UnhookDetour - Removes a detour hook
//  @param pFunc: Function address
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookDetour(pFunc uintptr) bool {
	return UnhookDetour(pFunc)
}

// HookVirtualTable - Sets a virtual table hook
//  @param pClass: Object pointer
//  @param index: Vtable offset
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func (w *PolyHook) HookVirtualTable(pClass uintptr, index int32, returnType DataType, arguments []DataType, varIndex int32) *Callback {
	return NewCallback(HookVirtualTable(pClass, index, returnType, arguments, varIndex))
}

// HookVirtualTableByFunc - Sets a virtual table hook by pointer
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func (w *PolyHook) HookVirtualTableByFunc(pClass uintptr, pFunc uintptr, returnType DataType, arguments []DataType, varIndex int32) *Callback {
	return NewCallback(HookVirtualTable2(pClass, pFunc, returnType, arguments, varIndex))
}

// UnhookVirtualTable - Removes a virtual hook table
//  @param pClass: Object pointer
//  @param index: Value to set
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookVirtualTable(pClass uintptr, index int32) bool {
	return UnhookVirtualTable(pClass, index)
}

// UnhookVirtualTableByFunc - Removes a virtual table hook by pointer
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookVirtualTableByFunc(pClass uintptr, pFunc uintptr) bool {
	return UnhookVirtualTable2(pClass, pFunc)
}

// HookVirtualFunc - Sets a virtual function hook
//  @param pClass: Object pointer
//  @param index: Vtable offset
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func (w *PolyHook) HookVirtualFunc(pClass uintptr, index int32, returnType DataType, arguments []DataType, varIndex int32) *Callback {
	return NewCallback(HookVirtualFunc(pClass, index, returnType, arguments, varIndex))
}

// HookVirtualFuncByFunc - Sets a virtual function hook by pointer
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func (w *PolyHook) HookVirtualFuncByFunc(pClass uintptr, pFunc uintptr, returnType DataType, arguments []DataType, varIndex int32) *Callback {
	return NewCallback(HookVirtualFunc2(pClass, pFunc, returnType, arguments, varIndex))
}

// UnhookVirtualFunc - Removes a virtual function table
//  @param pClass: Object pointer
//  @param index: Value to set
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookVirtualFunc(pClass uintptr, index int32) bool {
	return UnhookVirtualFunc(pClass, index)
}

// UnhookVirtualFuncByFunc - Removes a virtual function hook by pointer
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookVirtualFuncByFunc(pClass uintptr, pFunc uintptr) bool {
	return UnhookVirtualFunc2(pClass, pFunc)
}

// FindDetour - Attempts to find existing detour hook
//  @param pFunc: Function address
//
//  @return Returns hook pointer
func (w *PolyHook) FindDetour(pFunc uintptr) *Callback {
	return NewCallback(FindDetour(pFunc))
}

// FindVirtual - Attempts to find existing virtual hook
//  @param pClass: Object pointer
//  @param index: Value to set
//
//  @return Returns hook pointer
func (w *PolyHook) FindVirtual(pClass uintptr, index int32) *Callback {
	return NewCallback(FindVirtual(pClass, index))
}

// FindVirtualByFunc - Attempts to find existing detour hook by pointer
//  @param pClass: Object pointer
//  @param pFunc: Function address
//
//  @return Returns hook pointer
func (w *PolyHook) FindVirtualByFunc(pClass uintptr, pFunc uintptr) *Callback {
	return NewCallback(FindVirtual2(pClass, pFunc))
}

// GetVirtualIndex - Attempts to find virtual table index of virtual function
//  @param pFunc: Function address
//
//  @return Virtual table index
func (w *PolyHook) GetVirtualIndex(pFunc uintptr) int32 {
	return GetVirtualIndex(pFunc)
}

// UnhookAll - Removes all previously set hooks
func (w *PolyHook) UnhookAll() {
	UnhookAll()
}

// UnhookAllVirtual - Removes all previously set hooks on the object
//  @param pClass: Object pointer
func (w *PolyHook) UnhookAllVirtual(pClass uintptr) {
	UnhookAllVirtual(pClass)
}

