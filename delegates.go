package polyhook

import "github.com/untrustedmodders/go-plugify"

var _ = plugify.ApiVersion

// Generated from polyhook

// CallbackHandler - Callback function
type CallbackHandler func(hook HookHandle, params ParametersHandle, count int32, ret ReturnHandle, type_ CallbackType) ResultType


