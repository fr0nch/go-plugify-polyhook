#pragma once

#include "shared.h"

extern void (*__polyhook_SetArgumentBool)(uintptr_t, uint64_t, bool);

static void SetArgumentBool(uintptr_t params, uint64_t index, bool value) {
	__polyhook_SetArgumentBool(params, index, value);
}

extern void (*__polyhook_SetArgumentInt8)(uintptr_t, uint64_t, int8_t);

static void SetArgumentInt8(uintptr_t params, uint64_t index, int8_t value) {
	__polyhook_SetArgumentInt8(params, index, value);
}

extern void (*__polyhook_SetArgumentUInt8)(uintptr_t, uint64_t, uint8_t);

static void SetArgumentUInt8(uintptr_t params, uint64_t index, uint8_t value) {
	__polyhook_SetArgumentUInt8(params, index, value);
}

extern void (*__polyhook_SetArgumentInt16)(uintptr_t, uint64_t, int16_t);

static void SetArgumentInt16(uintptr_t params, uint64_t index, int16_t value) {
	__polyhook_SetArgumentInt16(params, index, value);
}

extern void (*__polyhook_SetArgumentUInt16)(uintptr_t, uint64_t, uint16_t);

static void SetArgumentUInt16(uintptr_t params, uint64_t index, uint16_t value) {
	__polyhook_SetArgumentUInt16(params, index, value);
}

extern void (*__polyhook_SetArgumentInt32)(uintptr_t, uint64_t, int32_t);

static void SetArgumentInt32(uintptr_t params, uint64_t index, int32_t value) {
	__polyhook_SetArgumentInt32(params, index, value);
}

extern void (*__polyhook_SetArgumentUInt32)(uintptr_t, uint64_t, uint32_t);

static void SetArgumentUInt32(uintptr_t params, uint64_t index, uint32_t value) {
	__polyhook_SetArgumentUInt32(params, index, value);
}

extern void (*__polyhook_SetArgumentInt64)(uintptr_t, uint64_t, int64_t);

static void SetArgumentInt64(uintptr_t params, uint64_t index, int64_t value) {
	__polyhook_SetArgumentInt64(params, index, value);
}

extern void (*__polyhook_SetArgumentUInt64)(uintptr_t, uint64_t, uint64_t);

static void SetArgumentUInt64(uintptr_t params, uint64_t index, uint64_t value) {
	__polyhook_SetArgumentUInt64(params, index, value);
}

extern void (*__polyhook_SetArgumentFloat)(uintptr_t, uint64_t, float);

static void SetArgumentFloat(uintptr_t params, uint64_t index, float value) {
	__polyhook_SetArgumentFloat(params, index, value);
}

extern void (*__polyhook_SetArgumentDouble)(uintptr_t, uint64_t, double);

static void SetArgumentDouble(uintptr_t params, uint64_t index, double value) {
	__polyhook_SetArgumentDouble(params, index, value);
}

extern void (*__polyhook_SetArgumentPointer)(uintptr_t, uint64_t, uintptr_t);

static void SetArgumentPointer(uintptr_t params, uint64_t index, uintptr_t value) {
	__polyhook_SetArgumentPointer(params, index, value);
}

extern void (*__polyhook_SetArgumentString)(uintptr_t, uintptr_t, uint64_t, String*);

static void SetArgumentString(uintptr_t hook, uintptr_t params, uint64_t index, String* value) {
	__polyhook_SetArgumentString(hook, params, index, value);
}

extern void (*__polyhook_SetArgument)(uintptr_t, uintptr_t, uint64_t, Variant*);

static void SetArgument(uintptr_t hook, uintptr_t params, uint64_t index, Variant* value) {
	__polyhook_SetArgument(hook, params, index, value);
}

extern void (*__polyhook_SetReturnBool)(uintptr_t, bool);

static void SetReturnBool(uintptr_t ret, bool value) {
	__polyhook_SetReturnBool(ret, value);
}

extern void (*__polyhook_SetReturnInt8)(uintptr_t, int8_t);

static void SetReturnInt8(uintptr_t ret, int8_t value) {
	__polyhook_SetReturnInt8(ret, value);
}

extern void (*__polyhook_SetReturnUInt8)(uintptr_t, uint8_t);

static void SetReturnUInt8(uintptr_t ret, uint8_t value) {
	__polyhook_SetReturnUInt8(ret, value);
}

extern void (*__polyhook_SetReturnInt16)(uintptr_t, int16_t);

static void SetReturnInt16(uintptr_t ret, int16_t value) {
	__polyhook_SetReturnInt16(ret, value);
}

extern void (*__polyhook_SetReturnUInt16)(uintptr_t, uint16_t);

static void SetReturnUInt16(uintptr_t ret, uint16_t value) {
	__polyhook_SetReturnUInt16(ret, value);
}

extern void (*__polyhook_SetReturnInt32)(uintptr_t, int32_t);

static void SetReturnInt32(uintptr_t ret, int32_t value) {
	__polyhook_SetReturnInt32(ret, value);
}

extern void (*__polyhook_SetReturnUInt32)(uintptr_t, uint32_t);

static void SetReturnUInt32(uintptr_t ret, uint32_t value) {
	__polyhook_SetReturnUInt32(ret, value);
}

extern void (*__polyhook_SetReturnInt64)(uintptr_t, int64_t);

static void SetReturnInt64(uintptr_t ret, int64_t value) {
	__polyhook_SetReturnInt64(ret, value);
}

extern void (*__polyhook_SetReturnUInt64)(uintptr_t, uint64_t);

static void SetReturnUInt64(uintptr_t ret, uint64_t value) {
	__polyhook_SetReturnUInt64(ret, value);
}

extern void (*__polyhook_SetReturnFloat)(uintptr_t, float);

static void SetReturnFloat(uintptr_t ret, float value) {
	__polyhook_SetReturnFloat(ret, value);
}

extern void (*__polyhook_SetReturnDouble)(uintptr_t, double);

static void SetReturnDouble(uintptr_t ret, double value) {
	__polyhook_SetReturnDouble(ret, value);
}

extern void (*__polyhook_SetReturnPointer)(uintptr_t, uintptr_t);

static void SetReturnPointer(uintptr_t ret, uintptr_t value) {
	__polyhook_SetReturnPointer(ret, value);
}

extern void (*__polyhook_SetReturnString)(uintptr_t, uintptr_t, String*);

static void SetReturnString(uintptr_t hook, uintptr_t ret, String* value) {
	__polyhook_SetReturnString(hook, ret, value);
}

extern void (*__polyhook_SetReturn)(uintptr_t, uintptr_t, Variant*);

static void SetReturn(uintptr_t hook, uintptr_t ret, Variant* value) {
	__polyhook_SetReturn(hook, ret, value);
}

