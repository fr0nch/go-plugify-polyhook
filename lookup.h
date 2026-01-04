#pragma once

#include "shared.h"

extern uintptr_t (*__polyhook_FindDetour)(uintptr_t);

static uintptr_t FindDetour(uintptr_t pFunc) {
	return __polyhook_FindDetour(pFunc);
}

extern uintptr_t (*__polyhook_FindVirtual)(uintptr_t, int32_t);

static uintptr_t FindVirtual(uintptr_t pClass, int32_t index) {
	return __polyhook_FindVirtual(pClass, index);
}

extern uintptr_t (*__polyhook_FindVirtual2)(uintptr_t, uintptr_t);

static uintptr_t FindVirtual2(uintptr_t pClass, uintptr_t pFunc) {
	return __polyhook_FindVirtual2(pClass, pFunc);
}

extern int32_t (*__polyhook_GetVirtualIndex)(uintptr_t);

static int32_t GetVirtualIndex(uintptr_t pFunc) {
	return __polyhook_GetVirtualIndex(pFunc);
}

