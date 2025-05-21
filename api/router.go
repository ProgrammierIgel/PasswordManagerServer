package api

import (
	"github.com/julienschmidt/httprouter"
	addnewaccount "github.com/programmierigel/pwmanager/api/addNewAccount"
	addnewpassword "github.com/programmierigel/pwmanager/api/addNewPassword"
	"github.com/programmierigel/pwmanager/api/changepasswordname"
	"github.com/programmierigel/pwmanager/api/changesecret"
	changeusername "github.com/programmierigel/pwmanager/api/changeusername"
	"github.com/programmierigel/pwmanager/api/checkpassword"
	"github.com/programmierigel/pwmanager/api/checktoken"
	"github.com/programmierigel/pwmanager/api/createtoken"
	deletepassword "github.com/programmierigel/pwmanager/api/deletePassword"
	"github.com/programmierigel/pwmanager/api/deleteaccount"
	"github.com/programmierigel/pwmanager/api/devaluealltoken"
	"github.com/programmierigel/pwmanager/api/devaluealltokenofaccount"
	"github.com/programmierigel/pwmanager/api/devaluetoken"
	disablesync "github.com/programmierigel/pwmanager/api/disableSync"
	enablesync "github.com/programmierigel/pwmanager/api/enableSync"
	"github.com/programmierigel/pwmanager/api/getallpasswordsofaccount"
	"github.com/programmierigel/pwmanager/api/getnumberofallregisteredtokens"
	"github.com/programmierigel/pwmanager/api/getpassword"
	"github.com/programmierigel/pwmanager/api/geturl"
	"github.com/programmierigel/pwmanager/api/getusername"
	issyncdisabled "github.com/programmierigel/pwmanager/api/isSyncDisabled"
	"github.com/programmierigel/pwmanager/api/ping"
	synctofile "github.com/programmierigel/pwmanager/api/syncToFile"
	"github.com/programmierigel/pwmanager/api/syncfromfile"
	"github.com/programmierigel/pwmanager/storage"
)

func GetRouter(store storage.Store) *httprouter.Router {
	router := httprouter.New()
	// QUERYS
	router.GET("/ping", ping.Handle())
	// COMMANDS
	router.POST("/addNewAccount", addnewaccount.Handle(store))
	router.POST("/deleteAccount", deleteaccount.Handle(store))
	router.POST("/addNewPassword", addnewpassword.Handle(store))
	router.POST("/checkPassword", checkpassword.Handle(store))
	router.POST("/deletePassword", deletepassword.Handle(store))
	router.POST("/system/disableSync", disablesync.Handle(store))
	router.POST("/system/enableSync", enablesync.Handle(store))
	router.POST("/getAllPasswordsOfAccount", getallpasswordsofaccount.Handle(store))
	router.POST("/getPassword", getpassword.Handle(store))
	router.POST("/getUsername", getusername.Handle(store))
	router.POST("/getUrl", geturl.Handle(store))
	router.POST("/changePasswordName", changepasswordname.Handle(store))
	router.POST("/changeUrl", changepasswordname.Handle(store))
	router.POST("/changeSecret", changesecret.Handle(store))
	router.POST("/changeUsername", changeusername.Handle(store))
	router.GET("/system/syncFromFile", syncfromfile.Handle(store))
	router.GET("/system/syncToFile", synctofile.Handle(store))
	router.GET("/system/isSyncDisabled", issyncdisabled.Handle(store))
	router.POST("/system/devalueAllToken", devaluealltoken.Handle(store))
	router.POST("/devalueAllTokenOfAccount", devaluealltokenofaccount.Handle(store))
	router.POST("/createToken", createtoken.Handle(store))
	router.POST("/devalueToken", devaluetoken.Handle(store))
	router.POST("/checkToken", checktoken.Handle(store))
	router.POST("/getNumberOfAllRegisteredTokens", getnumberofallregisteredtokens.Handle(store))
	return router
}
