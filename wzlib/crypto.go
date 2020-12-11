/*
 * Project Libra cgo usage
 * Author: A. Yasuda
 */

package main

import (
	util "github.com/wzlib/wzutil"
	"C"
)

func init(){

	//util.Init("agent",true)
}

//export ValidateToken
func ValidateToken(om *C.char) (*C.char){

	_om:=C.GoString(om)
	uid,err:=util.ValidateTokenExp(_om)
	if err!=nil{
		util.GetCLog().ErrLog(err)
		return nil
	}
	return C.CString(uid)
}

func main() {}
