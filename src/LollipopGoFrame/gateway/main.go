package gateway

import (
	"FenDZ/glog-master"
	"xcode/xutil"
	"fmt"
)

var FPS  = 20

func main()  {
	glog.Info("Fream Server data!!!")
	app:=xutil.GApp
	app.SetFps(FPS).SetPidFile("LollipopGo.json").SetVersion(1,0,1)
	if !app.Init(){
		return
	}
	setConsoleTitle()
	app.Run(Twst)
	app.Distroy()
}

func Twst(i int64)  {
   fmt.Println("---- 帧同步服务器 -----",i)
}
