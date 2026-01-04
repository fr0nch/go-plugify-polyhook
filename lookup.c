#include "shared.h"

PLUGIFY_EXPORT uintptr_t (*__polyhook_FindDetour)(uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_FindVirtual)(uintptr_t, int32_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_FindVirtual2)(uintptr_t, uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__polyhook_GetVirtualIndex)(uintptr_t) = NULL;


