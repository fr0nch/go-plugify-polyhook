#pragma once

#include "shared.h"

extern uintptr_t (*__polyhook_HookDetour)(uintptr_t, uint8_t, Vector*, int32_t);

static uintptr_t HookDetour(uintptr_t pFunc, uint8_t returnType, Vector* arguments, int32_t varIndex) {
	return __polyhook_HookDetour(pFunc, returnType, arguments, varIndex);
}

extern uintptr_t (*__polyhook_HookDetour2)(uintptr_t);

static uintptr_t HookDetour2(uintptr_t pFunc) {
	return __polyhook_HookDetour2(pFunc);
}

extern uintptr_t (*__polyhook_HookVirtualTable)(uintptr_t, int32_t, uint8_t, Vector*, int32_t);

static uintptr_t HookVirtualTable(uintptr_t pClass, int32_t index, uint8_t returnType, Vector* arguments, int32_t varIndex) {
	return __polyhook_HookVirtualTable(pClass, index, returnType, arguments, varIndex);
}

extern uintptr_t (*__polyhook_HookVirtualTable2)(uintptr_t, uintptr_t, uint8_t, Vector*, int32_t);

static uintptr_t HookVirtualTable2(uintptr_t pClass, uintptr_t pFunc, uint8_t returnType, Vector* arguments, int32_t varIndex) {
	return __polyhook_HookVirtualTable2(pClass, pFunc, returnType, arguments, varIndex);
}

extern uintptr_t (*__polyhook_HookVirtualFunc)(uintptr_t, int32_t, uint8_t, Vector*, int32_t);

static uintptr_t HookVirtualFunc(uintptr_t pClass, int32_t index, uint8_t returnType, Vector* arguments, int32_t varIndex) {
	return __polyhook_HookVirtualFunc(pClass, index, returnType, arguments, varIndex);
}

extern uintptr_t (*__polyhook_HookVirtualFunc2)(uintptr_t, uintptr_t, uint8_t, Vector*, int32_t);

static uintptr_t HookVirtualFunc2(uintptr_t pClass, uintptr_t pFunc, uint8_t returnType, Vector* arguments, int32_t varIndex) {
	return __polyhook_HookVirtualFunc2(pClass, pFunc, returnType, arguments, varIndex);
}

extern bool (*__polyhook_UnhookDetour)(uintptr_t);

static bool UnhookDetour(uintptr_t pFunc) {
	return __polyhook_UnhookDetour(pFunc);
}

extern bool (*__polyhook_UnhookVirtualTable)(uintptr_t, int32_t);

static bool UnhookVirtualTable(uintptr_t pClass, int32_t index) {
	return __polyhook_UnhookVirtualTable(pClass, index);
}

extern bool (*__polyhook_UnhookVirtualTable2)(uintptr_t, uintptr_t);

static bool UnhookVirtualTable2(uintptr_t pClass, uintptr_t pFunc) {
	return __polyhook_UnhookVirtualTable2(pClass, pFunc);
}

extern bool (*__polyhook_UnhookVirtualFunc)(uintptr_t, int32_t);

static bool UnhookVirtualFunc(uintptr_t pClass, int32_t index) {
	return __polyhook_UnhookVirtualFunc(pClass, index);
}

extern bool (*__polyhook_UnhookVirtualFunc2)(uintptr_t, uintptr_t);

static bool UnhookVirtualFunc2(uintptr_t pClass, uintptr_t pFunc) {
	return __polyhook_UnhookVirtualFunc2(pClass, pFunc);
}

extern void (*__polyhook_UnhookAll)();

static void UnhookAll() {
	__polyhook_UnhookAll();
}

extern void (*__polyhook_UnhookAllVirtual)(uintptr_t);

static void UnhookAllVirtual(uintptr_t pClass) {
	__polyhook_UnhookAllVirtual(pClass);
}

extern bool (*__polyhook_AddCallback)(uintptr_t, uint8_t, void*);

static bool AddCallback(uintptr_t hook, uint8_t type_, void* handler) {
	return __polyhook_AddCallback(hook, type_, handler);
}

extern bool (*__polyhook_AddCallback2)(uintptr_t, uint8_t, void*, int32_t);

static bool AddCallback2(uintptr_t hook, uint8_t type_, void* handler, int32_t priority) {
	return __polyhook_AddCallback2(hook, type_, handler, priority);
}

extern bool (*__polyhook_RemoveCallback)(uintptr_t, uint8_t, void*);

static bool RemoveCallback(uintptr_t hook, uint8_t type_, void* handler) {
	return __polyhook_RemoveCallback(hook, type_, handler);
}

extern bool (*__polyhook_IsCallbackRegistered)(uintptr_t, uint8_t, void*);

static bool IsCallbackRegistered(uintptr_t hook, uint8_t type_, void* handler) {
	return __polyhook_IsCallbackRegistered(hook, type_, handler);
}

extern bool (*__polyhook_AreCallbacksRegistered)(uintptr_t);

static bool AreCallbacksRegistered(uintptr_t hook) {
	return __polyhook_AreCallbacksRegistered(hook);
}

