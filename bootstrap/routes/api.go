// Package routes 注册路由
package routes

import (
	"gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，所有的 v1 版本的路由都放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			// 手机验证码
            authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			// Email 验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
			// 手机号注册
            authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
		}
	}
}
