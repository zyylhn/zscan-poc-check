package cmd

import (
	"embed"
	"flag"
	"fmt"
	"github.com/zyylhn/zscan_poc_check/lib"
	"net/http"
	"strings"
)

//go:embed pocs
var Pocs embed.FS
var Target string
var Pocpath string
var PocName string
var Thread int
var Listpoc bool



func Execute()  {
	flag.StringVar(&Target,"t","","set target")
	flag.StringVar(&Pocpath,"path","","set external poc path")
	flag.StringVar(&PocName,"name","","set poc name,defaule use all")
	flag.IntVar(&Thread,"thread",10,"set Thread")
	flag.BoolVar(&Listpoc,"list",false,"List built in poc")
	flag.Parse()
	//lib.Inithttp()  //初始化httpclient
	req, err := http.NewRequest("GET", Target, nil)
	if err != nil {
		return
	}
	if Listpoc{
		lib.ListBuiltinPoc(Pocs,PocName)
		return
	}
	req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	if Target==""{
		fmt.Println("必须指定目标")
		return
	}else {
		if Pocpath==""{
			lib.CheckBuiltinPoc(req, Pocs, Thread, PocName)
		}else {
			if strings.HasSuffix(Pocpath, ".yml") || strings.HasSuffix(Pocpath, ".yaml"){
				lib.CheckSinglePoc(req,Pocpath)
			}else {
				lib.CheckExternalPoc(req,Pocpath,10,PocName)
			}
		}
	}
}
