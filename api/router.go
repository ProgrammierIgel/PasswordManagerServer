package api

import (
	"log"

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

func GetRouter(store storage.Store, logger *log.Logger) *httprouter.Router {
	router := httprouter.New()
	// QUERYS
	router.GET("/ping", ping.Handle())
	// COMMANDS
	router.POST("/addNewAccount", addnewaccount.Handle(store, logger))
	router.POST("/deleteAccount", deleteaccount.Handle(store, logger))
	router.POST("/addNewPassword", addnewpassword.Handle(store, logger))
	router.POST("/checkPassword", checkpassword.Handle(store, logger))
	router.POST("/deletePassword", deletepassword.Handle(store, logger))
	router.POST("/system/disableSync", disablesync.Handle(store, logger))
	router.POST("/system/enableSync", enablesync.Handle(store, logger))
	router.POST("/getAllPasswordsOfAccount", getallpasswordsofaccount.Handle(store, logger))
	router.POST("/getPassword", getpassword.Handle(store, logger))
	router.POST("/getUsername", getusername.Handle(store, logger))
	router.POST("/getUrl", geturl.Handle(store, logger))
	router.POST("/changePasswordName", changepasswordname.Handle(store, logger))
	router.POST("/changeUrl", changepasswordname.Handle(store, logger))
	router.POST("/changeSecret", changesecret.Handle(store, logger))
	router.POST("/changeUsername", changeusername.Handle(store, logger))
	router.GET("/system/syncFromFile", syncfromfile.Handle(store, logger))
	router.GET("/system/syncToFile", synctofile.Handle(store, logger))
	router.GET("/system/isSyncDisabled", issyncdisabled.Handle(store, logger))
	router.POST("/system/devalueAllToken", devaluealltoken.Handle(store, logger))
	router.POST("/devalueAllTokenOfAccount", devaluealltokenofaccount.Handle(store, logger))
	router.POST("/createToken", createtoken.Handle(store, logger))
	router.POST("/devalueToken", devaluetoken.Handle(store, logger))
	router.POST("/checkToken", checktoken.Handle(store, logger))
	router.POST("/getNumberOfAllRegisteredTokens", getnumberofallregisteredtokens.Handle(store, logger))
	return router
}
