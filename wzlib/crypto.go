/*
 * Project Libra cgo usage
 * Author: A. Yasuda
 */

package main

import (
	util "github.com/wzlib/wzutil"
	"C"
)

const {
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
	L_MODE = "api"
}

func init(){
	
	l:=util.NewAppLog(L_MODE)
	util.SetCLog(*l)
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

//export GeneratePassphrase
func GeneratePassphrase()(*C.char){

	ep,err:=util.GeneratePassphrase(L_MODE)
	if err!=nil{
		util.GetCLog().ErrLog(err)
		return nil
	}
	return C.CString(ep)
}

//export RefreshPassphrase
func RefreshPassphrase(){
	util.RefreshPassphraseScheduler(L_MODE)
}

func main() {}
