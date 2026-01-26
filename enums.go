package polyhook

// Generated from polyhook

// DataType - Enum representing different data types in the system.
type DataType = uint8

const (
	// Void - Represents no data type (void).
	DataType_Void DataType = 0
	// Bool - Represents a boolean data type (true or false).
	DataType_Bool DataType = 1
	// Int8 - Represents an 8-bit signed integer.
	DataType_Int8 DataType = 2
	// UInt8 - Represents an 8-bit unsigned integer.
	DataType_UInt8 DataType = 3
	// Int16 - Represents a 16-bit signed integer.
	DataType_Int16 DataType = 4
	// UInt16 - Represents a 16-bit unsigned integer.
	DataType_UInt16 DataType = 5
	// Int32 - Represents a 32-bit signed integer.
	DataType_Int32 DataType = 6
	// UInt32 - Represents a 32-bit unsigned integer.
	DataType_UInt32 DataType = 7
	// Int64 - Represents a 64-bit signed integer.
	DataType_Int64 DataType = 8
	// UInt64 - Represents a 64-bit unsigned integer.
	DataType_UInt64 DataType = 9
	// Float - Represents a 32-bit floating point number.
	DataType_Float DataType = 10
	// Double - Represents a 64-bit double precision floating point number.
	DataType_Double DataType = 11
	// Pointer - Represents a pointer to any type of data.
	DataType_Pointer DataType = 12
	// String - Represents a string data type.
	DataType_String DataType = 13
)

// CallbackType - Enum representing the type of callback.
type CallbackType = uint8

const (
	// Pre - Callback will be executed before the original function
	CallbackType_Pre CallbackType = 0
	// Post - Callback will be executed after the original function
	CallbackType_Post CallbackType = 1
)

// ResultType - Enum representing the possible results of an operation.
type ResultType = int32

const (
	// Ignored - Handler didn't take any action.
	ResultType_Ignored ResultType = 0
	// Handled - We did something, but real function should still be called.
	ResultType_Handled ResultType = 1
	// Override - Call real function, but use my return value.
	ResultType_Override ResultType = 2
	// Supercede - Skip real function; use my return value.
	ResultType_Supercede ResultType = 3
)

// RegisterType - Enum representing register storage offsets based on RegisterType layout (x64)
type RegisterType = uint64

const (
	// XMM0 - XMM register 0.
	RegisterType_XMM0 RegisterType = 0
	// XMM1 - XMM register 1.
	RegisterType_XMM1 RegisterType = 2
	// XMM2 - XMM register 2.
	RegisterType_XMM2 RegisterType = 4
	// XMM3 - XMM register 3.
	RegisterType_XMM3 RegisterType = 6
	// XMM4 - XMM register 4.
	RegisterType_XMM4 RegisterType = 8
	// XMM5 - XMM register 5.
	RegisterType_XMM5 RegisterType = 10
	// XMM6 - XMM register 6.
	RegisterType_XMM6 RegisterType = 12
	// XMM7 - XMM register 7.
	RegisterType_XMM7 RegisterType = 14
	// XMM8 - XMM register 8.
	RegisterType_XMM8 RegisterType = 16
	// XMM9 - XMM register 9.
	RegisterType_XMM9 RegisterType = 18
	// XMM10 - XMM register 10.
	RegisterType_XMM10 RegisterType = 20
	// XMM11 - XMM register 11.
	RegisterType_XMM11 RegisterType = 22
	// XMM12 - XMM register 12.
	RegisterType_XMM12 RegisterType = 24
	// XMM13 - XMM register 13.
	RegisterType_XMM13 RegisterType = 26
	// XMM14 - XMM register 14.
	RegisterType_XMM14 RegisterType = 28
	// XMM15 - XMM register 15.
	RegisterType_XMM15 RegisterType = 30
	// R15 - General-purpose register R15.
	RegisterType_R15 RegisterType = 32
	// R14 - General-purpose register R14.
	RegisterType_R14 RegisterType = 33
	// R13 - General-purpose register R13.
	RegisterType_R13 RegisterType = 34
	// R12 - General-purpose register R12.
	RegisterType_R12 RegisterType = 35
	// R11 - General-purpose register R11.
	RegisterType_R11 RegisterType = 36
	// R10 - General-purpose register R10.
	RegisterType_R10 RegisterType = 37
	// R9 - General-purpose register R9.
	RegisterType_R9 RegisterType = 38
	// R8 - General-purpose register R8.
	RegisterType_R8 RegisterType = 39
	// RDI - Destination Index register.
	RegisterType_RDI RegisterType = 40
	// RSI - Source Index register.
	RegisterType_RSI RegisterType = 41
	// RBP - Base Pointer register.
	RegisterType_RBP RegisterType = 42
	// RBX - Base register.
	RegisterType_RBX RegisterType = 43
	// RDX - Data register.
	RegisterType_RDX RegisterType = 44
	// RCX - Counter register.
	RegisterType_RCX RegisterType = 45
	// RAX - Accumulator register.
	RegisterType_RAX RegisterType = 46
	// RFLAGS - CPU flags register.
	RegisterType_RFLAGS RegisterType = 47
	// COUNT - Number of register entries.
	RegisterType_COUNT RegisterType = 48
)


