package gorf24

/*
  #cgo LDFLAGS: -lrf24_c -lrf24
  #include <RF24_c.h>
  #include <stdio.h>
*/
import "C"

func gobool(b C.cbool) bool {
	if b == C.cbool(0) {
		return false
	}
	return true
}

func cbool(b bool) C.cbool {
	if b {
		return C.cbool(1)
	}
	return C.cbool(0)
}
