#include "shared.h"

PLUGIFY_EXPORT uintptr_t (*__polyhook_HookDetour)(uintptr_t, uint8_t, Vector*, int32_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_HookDetour2)(uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_HookVirtualTable)(uintptr_t, int32_t, uint8_t, Vector*, int32_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_HookVirtualTable2)(uintptr_t, uintptr_t, uint8_t, Vector*, int32_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_HookVirtualFunc)(uintptr_t, int32_t, uint8_t, Vector*, int32_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_HookVirtualFunc2)(uintptr_t, uintptr_t, uint8_t, Vector*, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_UnhookDetour)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_UnhookVirtualTable)(uintptr_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_UnhookVirtualTable2)(uintptr_t, uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_UnhookVirtualFunc)(uintptr_t, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_UnhookVirtualFunc2)(uintptr_t, uintptr_t) = NULL;


PLUGIFY_EXPORT void (*__polyhook_UnhookAll)() = NULL;


PLUGIFY_EXPORT void (*__polyhook_UnhookAllVirtual)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_AddCallback)(uintptr_t, uint8_t, void*) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_AddCallback2)(uintptr_t, uint8_t, void*, int32_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_RemoveCallback)(uintptr_t, uint8_t, void*) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_IsCallbackRegistered)(uintptr_t, uint8_t, void*) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_AreCallbacksRegistered)(uintptr_t) = NULL;


