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


