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

	util.Init("agent",true)
}

//export ValidateToken
func ValidateToken(om *C.char) (*C.char){

	_om:=C.GoString(om)
	uid,err:=util.ValidateTokenExp(_om)
	if err!=nil{
		util.ErrLog(err)
		return nil
	}
	return C.CString(uid)
}

//export GeneratePassphrase
func GeneratePassphrase()(*C.char){
	/*
gateway:	
server:
client:
gw:client:
gw:server:
x:gateway:
x:server:
x:client:
watcher:
daemon:
api:*/

	ep,err:=util.GeneratePassphrase("api")
	if err!=nil{
		util.ErrLog(err)
		return nil
	}
	return C.CString(ep)
}

//export RefereshPassphrase
func RefreshPassphrase(){
	util.RefreshPassphraseScheduler("api")
}

func main() {}
