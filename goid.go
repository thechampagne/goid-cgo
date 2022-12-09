package main

/*
#include <stdlib.h>
typedef struct {
   char* buffer;
   int is_err;
   char* err_msg;
} goid_uuid_t;
*/
import "C"
import (
	"unsafe"
	"github.com/jakehl/goid"
)

//export goid_get_uuid_from_string
func goid_get_uuid_from_string(str_uuid *C.char) C.goid_uuid_t {
	var self C.goid_uuid_t
	res, err := goid.GetUUIDFromString(C.GoString(str_uuid))
	if err != nil {
		self.err_msg = C.CString(string(res[:]))
		self.is_err = 1
		return self
	}
	self.buffer = C.CString(string(res[:]))
	self.is_err = 0
	return self
}

//export goid_new_v4_uuid
func goid_new_v4_uuid() C.goid_uuid_t {
	var self C.goid_uuid_t
	res := goid.NewV4UUID()
	self.buffer = C.CString(string(res[:]))
	return self
	
}

//export goid_uuid_equals
func goid_uuid_equals(uuid *C.goid_uuid_t, comp *C.goid_uuid_t) C.int {
	var uuid_1 goid.UUID
	var uuid_2 goid.UUID
	copy(uuid_1[:], C.GoString(uuid.buffer))
	copy(uuid_2[:], C.GoString(comp.buffer))
	is_equals := uuid_1.Equals(&uuid_2)
	if is_equals {
		return 1
	} else {
		return 0
	}
}

//export goid_uuid_get_version
func goid_uuid_get_version(uuid *C.goid_uuid_t) *C.char {
	var uuid_1 goid.UUID
	copy(uuid_1[:], C.GoString(uuid.buffer))
	res := uuid_1.GetVersion()
	return C.CString(res)
}

//export goid_uuid_clean
func goid_uuid_clean(uuid *C.goid_uuid_t) {
	if uuid != nil {
		if uuid.buffer != nil { C.free(unsafe.Pointer(uuid.buffer)) }
		if uuid.err_msg != nil { C.free(unsafe.Pointer(uuid.err_msg)) }
	}
}

func main() {}
