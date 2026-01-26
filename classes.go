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

//  @brief RAII wrapper for Callback pointer from hook operations.
//
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

// AddCallback 
//  @brief Adds a callback to existing hook
//
//  @param hook: Hook pointer
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

// AddCallback2 
//  @brief Adds a callback to existing hook
//
//  @param hook: Hook pointer
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

// RemoveCallback 
//  @brief Removes a callback from existing hook
//
//  @param hook: Hook pointer
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

// IsCallbackRegistered 
//  @brief Checks that a callback exists on existing hook
//
//  @param hook: Hook pointer
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

// AreCallbacksRegistered 
//  @brief Checks that a any callback exists on existing hook
//
//  @param hook: Hook pointer
//
//  @return Returns true on success, false otherwise
func (w *Callback) AreCallbacksRegistered() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, CallbackErrEmptyHandle
	}
	return AreCallbacksRegistered(w.handle), nil
}

// GetFunctionAddr 
//  @brief Get generated function address
//
//  @param hook: Hook pointer
//
//  @return Returns jit generated function pointer
func (w *Callback) GetFunctionAddr() (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, CallbackErrEmptyHandle
	}
	return GetFunctionAddr(w.handle), nil
}

// GetOriginalAddr 
//  @brief Get original function address
//
//  @param hook: Hook pointer
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

//  @brief RAII wrapper for Callback::Parameters pointer used in callback handlers.
//
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

// GetBool 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetBool(index uint64) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentBool(w.handle, index), nil
}

// GetInt8 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetInt8(index uint64) (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentInt8(w.handle, index), nil
}

// GetUInt8 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetUInt8(index uint64) (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentUInt8(w.handle, index), nil
}

// GetInt16 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetInt16(index uint64) (int16, error) {
	if w.handle == 0 {
		var zero int16
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentInt16(w.handle, index), nil
}

// GetUInt16 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetUInt16(index uint64) (uint16, error) {
	if w.handle == 0 {
		var zero uint16
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentUInt16(w.handle, index), nil
}

// GetInt32 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetInt32(index uint64) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentInt32(w.handle, index), nil
}

// GetUInt32 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetUInt32(index uint64) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentUInt32(w.handle, index), nil
}

// GetInt64 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetInt64(index uint64) (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentInt64(w.handle, index), nil
}

// GetUInt64 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetUInt64(index uint64) (uint64, error) {
	if w.handle == 0 {
		var zero uint64
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentUInt64(w.handle, index), nil
}

// GetFloat 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetFloat(index uint64) (float32, error) {
	if w.handle == 0 {
		var zero float32
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentFloat(w.handle, index), nil
}

// GetDouble 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetDouble(index uint64) (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentDouble(w.handle, index), nil
}

// GetPointer 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetPointer(index uint64) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentPointer(w.handle, index), nil
}

// GetString 
//  @brief Get argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to get
func (w *Parameters) GetString(index uint64) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ParametersErrEmptyHandle
	}
	return GetArgumentString(w.handle, index), nil
}

// SetBool 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetBool(index uint64, value bool) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentBool(w.handle, index, value)
	return nil
}

// SetInt8 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetInt8(index uint64, value int8) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentInt8(w.handle, index, value)
	return nil
}

// SetUInt8 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetUInt8(index uint64, value uint8) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentUInt8(w.handle, index, value)
	return nil
}

// SetInt16 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetInt16(index uint64, value int16) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentInt16(w.handle, index, value)
	return nil
}

// SetUInt16 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetUInt16(index uint64, value uint16) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentUInt16(w.handle, index, value)
	return nil
}

// SetInt32 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetInt32(index uint64, value int32) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentInt32(w.handle, index, value)
	return nil
}

// SetUInt32 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetUInt32(index uint64, value uint32) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentUInt32(w.handle, index, value)
	return nil
}

// SetInt64 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetInt64(index uint64, value int64) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentInt64(w.handle, index, value)
	return nil
}

// SetUInt64 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetUInt64(index uint64, value uint64) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentUInt64(w.handle, index, value)
	return nil
}

// SetFloat 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetFloat(index uint64, value float32) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentFloat(w.handle, index, value)
	return nil
}

// SetDouble 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
//  @param value: Value to set
func (w *Parameters) SetDouble(index uint64, value float64) error {
	if w.handle == 0 {
		return ParametersErrEmptyHandle
	}
	SetArgumentDouble(w.handle, index, value)
	return nil
}

// SetPointer 
//  @brief Set argument value
//
//  @param params: Pointer to params structure
//  @param index: Index to set
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

//  @brief RAII wrapper for Callback::Return pointer used in callback handlers.
//
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

// GetBool 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetBool() (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnBool(w.handle), nil
}

// GetInt8 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetInt8() (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnInt8(w.handle), nil
}

// GetUInt8 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetUInt8() (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnUInt8(w.handle), nil
}

// GetInt16 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetInt16() (int16, error) {
	if w.handle == 0 {
		var zero int16
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnInt16(w.handle), nil
}

// GetUInt16 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetUInt16() (uint16, error) {
	if w.handle == 0 {
		var zero uint16
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnUInt16(w.handle), nil
}

// GetInt32 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetInt32() (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnInt32(w.handle), nil
}

// GetUInt32 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetUInt32() (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnUInt32(w.handle), nil
}

// GetInt64 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetInt64() (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnInt64(w.handle), nil
}

// GetUInt64 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetUInt64() (uint64, error) {
	if w.handle == 0 {
		var zero uint64
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnUInt64(w.handle), nil
}

// GetFloat 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetFloat() (float32, error) {
	if w.handle == 0 {
		var zero float32
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnFloat(w.handle), nil
}

// GetDouble 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetDouble() (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnDouble(w.handle), nil
}

// GetPointer 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetPointer() (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnPointer(w.handle), nil
}

// GetString 
//  @brief Get return value
//
//  @param ret: Pointer to return structure
func (w *Return) GetString() (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, ReturnErrEmptyHandle
	}
	return GetReturnString(w.handle), nil
}

// SetBool 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetBool(value bool) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnBool(w.handle, value)
	return nil
}

// SetInt8 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetInt8(value int8) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnInt8(w.handle, value)
	return nil
}

// SetUInt8 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetUInt8(value uint8) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnUInt8(w.handle, value)
	return nil
}

// SetInt16 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetInt16(value int16) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnInt16(w.handle, value)
	return nil
}

// SetUInt16 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetUInt16(value uint16) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnUInt16(w.handle, value)
	return nil
}

// SetInt32 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetInt32(value int32) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnInt32(w.handle, value)
	return nil
}

// SetUInt32 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetUInt32(value uint32) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnUInt32(w.handle, value)
	return nil
}

// SetInt64 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetInt64(value int64) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnInt64(w.handle, value)
	return nil
}

// SetUInt64 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetUInt64(value uint64) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnUInt64(w.handle, value)
	return nil
}

// SetFloat 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetFloat(value float32) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnFloat(w.handle, value)
	return nil
}

// SetDouble 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetDouble(value float64) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnDouble(w.handle, value)
	return nil
}

// SetPointer 
//  @brief Set return value
//
//  @param ret: Pointer to return structure
//  @param value: Value to set
func (w *Return) SetPointer(value uintptr) error {
	if w.handle == 0 {
		return ReturnErrEmptyHandle
	}
	SetReturnPointer(w.handle, value)
	return nil
}

var (
	RegistersErrEmptyHandle = errors.New("Registers: empty handle")
)

//  @brief RAII wrapper for Callback::Registers pointer used in callback handlers.
//
type Registers struct {
	handle    uintptr
}

// NewRegisters creates a Registers from a handle
func NewRegisters(handle uintptr) *Registers {
	return &Registers{
		handle:    handle,
	}
}

// Get returns the underlying handle
func (w *Registers) Get() uintptr {
	return w.handle
}

// Release releases ownership and returns the handle
func (w *Registers) Release() uintptr {
	handle := w.handle
	w.handle = 0
	return handle
}

// Reset destroys and resets the handle
func (w *Registers) Reset() {
	w.handle = 0
}

// IsValid returns true if handle is not nil
func (w *Registers) IsValid() bool {
	return w.handle != 0
}

// GetBool 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetBool(reg RegisterType) (bool, error) {
	if w.handle == 0 {
		var zero bool
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterBool(w.handle, reg), nil
}

// GetInt8 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetInt8(reg RegisterType) (int8, error) {
	if w.handle == 0 {
		var zero int8
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterInt8(w.handle, reg), nil
}

// GetUInt8 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetUInt8(reg RegisterType) (uint8, error) {
	if w.handle == 0 {
		var zero uint8
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterUInt8(w.handle, reg), nil
}

// GetInt16 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetInt16(reg RegisterType) (int16, error) {
	if w.handle == 0 {
		var zero int16
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterInt16(w.handle, reg), nil
}

// GetUInt16 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetUInt16(reg RegisterType) (uint16, error) {
	if w.handle == 0 {
		var zero uint16
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterUInt16(w.handle, reg), nil
}

// GetInt32 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetInt32(reg RegisterType) (int32, error) {
	if w.handle == 0 {
		var zero int32
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterInt32(w.handle, reg), nil
}

// GetUInt32 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetUInt32(reg RegisterType) (uint32, error) {
	if w.handle == 0 {
		var zero uint32
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterUInt32(w.handle, reg), nil
}

// GetInt64 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetInt64(reg RegisterType) (int64, error) {
	if w.handle == 0 {
		var zero int64
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterInt64(w.handle, reg), nil
}

// GetUInt64 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetUInt64(reg RegisterType) (uint64, error) {
	if w.handle == 0 {
		var zero uint64
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterUInt64(w.handle, reg), nil
}

// GetFloat 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetFloat(reg RegisterType) (float32, error) {
	if w.handle == 0 {
		var zero float32
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterFloat(w.handle, reg), nil
}

// GetDouble 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetDouble(reg RegisterType) (float64, error) {
	if w.handle == 0 {
		var zero float64
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterDouble(w.handle, reg), nil
}

// GetPointer 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetPointer(reg RegisterType) (uintptr, error) {
	if w.handle == 0 {
		var zero uintptr
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterPointer(w.handle, reg), nil
}

// GetString 
//  @brief Get register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to get
func (w *Registers) GetString(reg RegisterType) (string, error) {
	if w.handle == 0 {
		var zero string
		return zero, RegistersErrEmptyHandle
	}
	return GetRegisterString(w.handle, reg), nil
}

// SetBool 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetBool(reg RegisterType, value bool) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterBool(w.handle, reg, value)
	return nil
}

// SetInt8 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetInt8(reg RegisterType, value int8) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterInt8(w.handle, reg, value)
	return nil
}

// SetUInt8 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetUInt8(reg RegisterType, value uint8) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterUInt8(w.handle, reg, value)
	return nil
}

// SetInt16 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetInt16(reg RegisterType, value int16) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterInt16(w.handle, reg, value)
	return nil
}

// SetUInt16 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetUInt16(reg RegisterType, value uint16) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterUInt16(w.handle, reg, value)
	return nil
}

// SetInt32 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetInt32(reg RegisterType, value int32) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterInt32(w.handle, reg, value)
	return nil
}

// SetUInt32 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetUInt32(reg RegisterType, value uint32) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterUInt32(w.handle, reg, value)
	return nil
}

// SetInt64 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetInt64(reg RegisterType, value int64) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterInt64(w.handle, reg, value)
	return nil
}

// SetUInt64 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetUInt64(reg RegisterType, value uint64) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterUInt64(w.handle, reg, value)
	return nil
}

// SetFloat 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetFloat(reg RegisterType, value float32) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterFloat(w.handle, reg, value)
	return nil
}

// SetDouble 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetDouble(reg RegisterType, value float64) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterDouble(w.handle, reg, value)
	return nil
}

// SetPointer 
//  @brief Set register value
//
//  @param registers: Pointer to registers structure
//  @param reg: Register to set
//  @param value: Value to set
func (w *Registers) SetPointer(reg RegisterType, value uintptr) error {
	if w.handle == 0 {
		return RegistersErrEmptyHandle
	}
	SetRegisterPointer(w.handle, reg, value)
	return nil
}

//  @brief Global hooking functions for PolyHook plugin.
//
type PolyHook struct {
}

// HookDetour 
//  @brief Sets a detour hook
//
//  @param pFunc: Function address
//  @param returnType: Return type
//  @param arguments: Arguments type array
//  @param varIndex: Index of a first variadic argument or -1
//
//  @return Returns hook pointer
func (w *PolyHook) HookDetour(pFunc uintptr, returnType DataType, arguments []DataType, varIndex int32) *Callback {
	return NewCallback(HookDetour(pFunc, returnType, arguments, varIndex))
}

// HookDetour2 
//  @brief Sets a mid hook
//
//  @param pFunc: Function address
//
//  @return Returns hook pointer
func (w *PolyHook) HookDetour2(pFunc uintptr) *Callback {
	return NewCallback(HookDetour2(pFunc))
}

// UnhookDetour 
//  @brief Removes a detour hook
//
//  @param pFunc: Function address
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookDetour(pFunc uintptr) bool {
	return UnhookDetour(pFunc)
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
func (w *PolyHook) HookVirtualTable(pClass uintptr, index int32, returnType DataType, arguments []DataType, varIndex int32) *Callback {
	return NewCallback(HookVirtualTable(pClass, index, returnType, arguments, varIndex))
}

// HookVirtualTableByFunc 
//  @brief Sets a virtual table hook by pointer
//
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

// UnhookVirtualTable 
//  @brief Removes a virtual hook table
//
//  @param pClass: Object pointer
//  @param index: Virtual table index
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookVirtualTable(pClass uintptr, index int32) bool {
	return UnhookVirtualTable(pClass, index)
}

// UnhookVirtualTableByFunc 
//  @brief Removes a virtual table hook by pointer
//
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookVirtualTableByFunc(pClass uintptr, pFunc uintptr) bool {
	return UnhookVirtualTable2(pClass, pFunc)
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
func (w *PolyHook) HookVirtualFunc(pClass uintptr, index int32, returnType DataType, arguments []DataType, varIndex int32) *Callback {
	return NewCallback(HookVirtualFunc(pClass, index, returnType, arguments, varIndex))
}

// HookVirtualFuncByFunc 
//  @brief Sets a virtual function hook by pointer
//
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

// UnhookVirtualFunc 
//  @brief Removes a virtual function table
//
//  @param pClass: Object pointer
//  @param index: Virtual table index
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookVirtualFunc(pClass uintptr, index int32) bool {
	return UnhookVirtualFunc(pClass, index)
}

// UnhookVirtualFuncByFunc 
//  @brief Removes a virtual function hook by pointer
//
//  @param pClass: Object pointer
//  @param pFunc: Vtable member function address
//
//  @return Returns true on success, false otherwise
func (w *PolyHook) UnhookVirtualFuncByFunc(pClass uintptr, pFunc uintptr) bool {
	return UnhookVirtualFunc2(pClass, pFunc)
}

// FindDetour 
//  @brief Attempts to find existing detour hook
//
//  @param pFunc: Function address
//
//  @return Returns hook pointer
func (w *PolyHook) FindDetour(pFunc uintptr) *Callback {
	return NewCallback(FindDetour(pFunc))
}

// FindVirtual 
//  @brief Attempts to find existing virtual hook
//
//  @param pClass: Object pointer
//  @param index: Virtual table index
//
//  @return Returns hook pointer
func (w *PolyHook) FindVirtual(pClass uintptr, index int32) *Callback {
	return NewCallback(FindVirtual(pClass, index))
}

// FindVirtualByFunc 
//  @brief Attempts to find existing detour hook by pointer
//
//  @param pClass: Object pointer
//  @param pFunc: Function address
//
//  @return Returns hook pointer
func (w *PolyHook) FindVirtualByFunc(pClass uintptr, pFunc uintptr) *Callback {
	return NewCallback(FindVirtual2(pClass, pFunc))
}

// GetVirtualIndex 
//  @brief Attempts to find virtual table index of virtual function
//
//  @param pFunc: Function address
//
//  @return Virtual table index
func (w *PolyHook) GetVirtualIndex(pFunc uintptr) int32 {
	return GetVirtualIndex(pFunc)
}

// UnhookAll 
//  @brief Removes all previously set hooks
//
func (w *PolyHook) UnhookAll() {
	UnhookAll()
}

// UnhookAllVirtual 
//  @brief Removes all previously set hooks on the object
//
//  @param pClass: Object pointer
func (w *PolyHook) UnhookAllVirtual(pClass uintptr) {
	UnhookAllVirtual(pClass)
}

