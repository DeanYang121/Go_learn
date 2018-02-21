package router

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go_dev/day_14/Seckill/SecProxy/controller"
)

func init(){
	logs.Debug("enter router init")
	beego.Router("/seckill",&controller.SkillController{},"*:SecKill") //*代表是post还是get  冒号后面是skillController的函数
	beego.Router("/secinfo",&controller.SkillController{},"*:SecInfo")
	
}