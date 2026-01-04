#include "shared.h"

PLUGIFY_EXPORT uintptr_t (*__polyhook_GetFunctionAddr)(uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_GetOriginalAddr)(uintptr_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_GetArgumentBool)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT int8_t (*__polyhook_GetArgumentInt8)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT uint8_t (*__polyhook_GetArgumentUInt8)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT int16_t (*__polyhook_GetArgumentInt16)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT uint16_t (*__polyhook_GetArgumentUInt16)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT int32_t (*__polyhook_GetArgumentInt32)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT uint32_t (*__polyhook_GetArgumentUInt32)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT int64_t (*__polyhook_GetArgumentInt64)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT uint64_t (*__polyhook_GetArgumentUInt64)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT float (*__polyhook_GetArgumentFloat)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT double (*__polyhook_GetArgumentDouble)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_GetArgumentPointer)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT String (*__polyhook_GetArgumentString)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT Variant (*__polyhook_GetArgument)(uintptr_t, uint64_t) = NULL;


PLUGIFY_EXPORT bool (*__polyhook_GetReturnBool)(uintptr_t) = NULL;


PLUGIFY_EXPORT int8_t (*__polyhook_GetReturnInt8)(uintptr_t) = NULL;


PLUGIFY_EXPORT uint8_t (*__polyhook_GetReturnUInt8)(uintptr_t) = NULL;


PLUGIFY_EXPORT int16_t (*__polyhook_GetReturnInt16)(uintptr_t) = NULL;


PLUGIFY_EXPORT uint16_t (*__polyhook_GetReturnUInt16)(uintptr_t) = NULL;


PLUGIFY_EXPORT int32_t (*__polyhook_GetReturnInt32)(uintptr_t) = NULL;


PLUGIFY_EXPORT uint32_t (*__polyhook_GetReturnUInt32)(uintptr_t) = NULL;


PLUGIFY_EXPORT int64_t (*__polyhook_GetReturnInt64)(uintptr_t) = NULL;


PLUGIFY_EXPORT uint64_t (*__polyhook_GetReturnUInt64)(uintptr_t) = NULL;


PLUGIFY_EXPORT float (*__polyhook_GetReturnFloat)(uintptr_t) = NULL;


PLUGIFY_EXPORT double (*__polyhook_GetReturnDouble)(uintptr_t) = NULL;


PLUGIFY_EXPORT uintptr_t (*__polyhook_GetReturnPointer)(uintptr_t) = NULL;


PLUGIFY_EXPORT String (*__polyhook_GetReturnString)(uintptr_t) = NULL;


PLUGIFY_EXPORT Variant (*__polyhook_GetReturn)(uintptr_t) = NULL;


