// @APIVersion 1.0.0
// @Title 致简生活 API
// @Description 致简生活是一款微信小工具，您可以通过这个查看天气和快递等生活服务
// @Contact shiboven@foxmail.com
// @TermsOfServiceUrl https://tolife.yuhanle.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @Schemes http, https
package routers

import (
	"zjlife/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/weather",
			beego.NSInclude(
				&controllers.WeatherController{},
			),
		),
		beego.NSNamespace("/express",
			beego.NSInclude(
				&controllers.ExpressController{},
			),
		),
		beego.NSNamespace("/toilet",
			beego.NSInclude(
				&controllers.ToiletController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{})
}
