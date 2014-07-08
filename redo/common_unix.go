// +build !darwin

// 7 july 2014

package ui

import (
	"unsafe"
)

// #include "gtk_unix.h"
// /* because cgo doesn't like g_signal_connect() */
// void gSignalConnect(gpointer obj, gchar *sig, GCallback callback, gpointer data)
// {
// 	g_signal_connect(obj, sig, callback, data);
// }
import "C"

func fromgstr(s *C.gchar) string {
	return C.GoString((*C.char)(unsafe.Pointer(s)))
}

func togstr(s string) *C.gchar {
	return (*C.gchar)(unsafe.Pointer(C.CString(s)))
}

func freegstr(s *C.gchar) {
	C.free(unsafe.Pointer(s))
}

func g_signal_connect(object C.gpointer, name string, to C.GCallback, data C.gpointer) {
	cname := togstr(name)
	defer freegstr(cname)
	C.gSignalConnect(object, cname, to, data)
}
