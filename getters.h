#pragma once

#include "shared.h"

extern uintptr_t (*__polyhook_GetFunctionAddr)(uintptr_t);

static uintptr_t GetFunctionAddr(uintptr_t hook) {
	return __polyhook_GetFunctionAddr(hook);
}

extern uintptr_t (*__polyhook_GetOriginalAddr)(uintptr_t);

static uintptr_t GetOriginalAddr(uintptr_t hook) {
	return __polyhook_GetOriginalAddr(hook);
}

extern bool (*__polyhook_GetArgumentBool)(uintptr_t, uint64_t);

static bool GetArgumentBool(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentBool(params, index);
}

extern int8_t (*__polyhook_GetArgumentInt8)(uintptr_t, uint64_t);

static int8_t GetArgumentInt8(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentInt8(params, index);
}

extern uint8_t (*__polyhook_GetArgumentUInt8)(uintptr_t, uint64_t);

static uint8_t GetArgumentUInt8(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentUInt8(params, index);
}

extern int16_t (*__polyhook_GetArgumentInt16)(uintptr_t, uint64_t);

static int16_t GetArgumentInt16(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentInt16(params, index);
}

extern uint16_t (*__polyhook_GetArgumentUInt16)(uintptr_t, uint64_t);

static uint16_t GetArgumentUInt16(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentUInt16(params, index);
}

extern int32_t (*__polyhook_GetArgumentInt32)(uintptr_t, uint64_t);

static int32_t GetArgumentInt32(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentInt32(params, index);
}

extern uint32_t (*__polyhook_GetArgumentUInt32)(uintptr_t, uint64_t);

static uint32_t GetArgumentUInt32(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentUInt32(params, index);
}

extern int64_t (*__polyhook_GetArgumentInt64)(uintptr_t, uint64_t);

static int64_t GetArgumentInt64(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentInt64(params, index);
}

extern uint64_t (*__polyhook_GetArgumentUInt64)(uintptr_t, uint64_t);

static uint64_t GetArgumentUInt64(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentUInt64(params, index);
}

extern float (*__polyhook_GetArgumentFloat)(uintptr_t, uint64_t);

static float GetArgumentFloat(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentFloat(params, index);
}

extern double (*__polyhook_GetArgumentDouble)(uintptr_t, uint64_t);

static double GetArgumentDouble(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentDouble(params, index);
}

extern uintptr_t (*__polyhook_GetArgumentPointer)(uintptr_t, uint64_t);

static uintptr_t GetArgumentPointer(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentPointer(params, index);
}

extern String (*__polyhook_GetArgumentString)(uintptr_t, uint64_t);

static String GetArgumentString(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgumentString(params, index);
}

extern Variant (*__polyhook_GetArgument)(uintptr_t, uint64_t);

static Variant GetArgument(uintptr_t params, uint64_t index) {
	return __polyhook_GetArgument(params, index);
}

extern bool (*__polyhook_GetReturnBool)(uintptr_t);

static bool GetReturnBool(uintptr_t ret) {
	return __polyhook_GetReturnBool(ret);
}

extern int8_t (*__polyhook_GetReturnInt8)(uintptr_t);

static int8_t GetReturnInt8(uintptr_t ret) {
	return __polyhook_GetReturnInt8(ret);
}

extern uint8_t (*__polyhook_GetReturnUInt8)(uintptr_t);

static uint8_t GetReturnUInt8(uintptr_t ret) {
	return __polyhook_GetReturnUInt8(ret);
}

extern int16_t (*__polyhook_GetReturnInt16)(uintptr_t);

static int16_t GetReturnInt16(uintptr_t ret) {
	return __polyhook_GetReturnInt16(ret);
}

extern uint16_t (*__polyhook_GetReturnUInt16)(uintptr_t);

static uint16_t GetReturnUInt16(uintptr_t ret) {
	return __polyhook_GetReturnUInt16(ret);
}

extern int32_t (*__polyhook_GetReturnInt32)(uintptr_t);

static int32_t GetReturnInt32(uintptr_t ret) {
	return __polyhook_GetReturnInt32(ret);
}

extern uint32_t (*__polyhook_GetReturnUInt32)(uintptr_t);

static uint32_t GetReturnUInt32(uintptr_t ret) {
	return __polyhook_GetReturnUInt32(ret);
}

extern int64_t (*__polyhook_GetReturnInt64)(uintptr_t);

static int64_t GetReturnInt64(uintptr_t ret) {
	return __polyhook_GetReturnInt64(ret);
}

extern uint64_t (*__polyhook_GetReturnUInt64)(uintptr_t);

static uint64_t GetReturnUInt64(uintptr_t ret) {
	return __polyhook_GetReturnUInt64(ret);
}

extern float (*__polyhook_GetReturnFloat)(uintptr_t);

static float GetReturnFloat(uintptr_t ret) {
	return __polyhook_GetReturnFloat(ret);
}

extern double (*__polyhook_GetReturnDouble)(uintptr_t);

static double GetReturnDouble(uintptr_t ret) {
	return __polyhook_GetReturnDouble(ret);
}

extern uintptr_t (*__polyhook_GetReturnPointer)(uintptr_t);

static uintptr_t GetReturnPointer(uintptr_t ret) {
	return __polyhook_GetReturnPointer(ret);
}

extern String (*__polyhook_GetReturnString)(uintptr_t);

static String GetReturnString(uintptr_t ret) {
	return __polyhook_GetReturnString(ret);
}

extern Variant (*__polyhook_GetReturn)(uintptr_t);

static Variant GetReturn(uintptr_t ret) {
	return __polyhook_GetReturn(ret);
}

extern bool (*__polyhook_GetRegisterBool)(uintptr_t, uint64_t);

static bool GetRegisterBool(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterBool(registers, reg);
}

extern int8_t (*__polyhook_GetRegisterInt8)(uintptr_t, uint64_t);

static int8_t GetRegisterInt8(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterInt8(registers, reg);
}

extern uint8_t (*__polyhook_GetRegisterUInt8)(uintptr_t, uint64_t);

static uint8_t GetRegisterUInt8(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterUInt8(registers, reg);
}

extern int16_t (*__polyhook_GetRegisterInt16)(uintptr_t, uint64_t);

static int16_t GetRegisterInt16(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterInt16(registers, reg);
}

extern uint16_t (*__polyhook_GetRegisterUInt16)(uintptr_t, uint64_t);

static uint16_t GetRegisterUInt16(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterUInt16(registers, reg);
}

extern int32_t (*__polyhook_GetRegisterInt32)(uintptr_t, uint64_t);

static int32_t GetRegisterInt32(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterInt32(registers, reg);
}

extern uint32_t (*__polyhook_GetRegisterUInt32)(uintptr_t, uint64_t);

static uint32_t GetRegisterUInt32(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterUInt32(registers, reg);
}

extern int64_t (*__polyhook_GetRegisterInt64)(uintptr_t, uint64_t);

static int64_t GetRegisterInt64(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterInt64(registers, reg);
}

extern uint64_t (*__polyhook_GetRegisterUInt64)(uintptr_t, uint64_t);

static uint64_t GetRegisterUInt64(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterUInt64(registers, reg);
}

extern float (*__polyhook_GetRegisterFloat)(uintptr_t, uint64_t);

static float GetRegisterFloat(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterFloat(registers, reg);
}

extern double (*__polyhook_GetRegisterDouble)(uintptr_t, uint64_t);

static double GetRegisterDouble(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterDouble(registers, reg);
}

extern uintptr_t (*__polyhook_GetRegisterPointer)(uintptr_t, uint64_t);

static uintptr_t GetRegisterPointer(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterPointer(registers, reg);
}

extern String (*__polyhook_GetRegisterString)(uintptr_t, uint64_t);

static String GetRegisterString(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegisterString(registers, reg);
}

extern Variant (*__polyhook_GetRegister)(uintptr_t, uint64_t);

static Variant GetRegister(uintptr_t registers, uint64_t reg) {
	return __polyhook_GetRegister(registers, reg);
}

