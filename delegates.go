package polyhook

import "github.com/untrustedmodders/go-plugify"

var _ = plugify.Plugin.Loaded

// Generated from polyhook

// CallbackHandler - Callback function
type CallbackHandler func(hook uintptr, params uintptr, count int32, ret uintptr, type_ CallbackType) ResultType


