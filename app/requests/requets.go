package requests

// import (
// 	"github.com/gin-gonic/gin"
// 	"gohub/pkg/response"
// )

// func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {

//     // 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
//     if err := c.ShouldBind(obj); err != nil {
//         response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
//         return false
//     }

//     // 2. 表单验证
//     errs := handler(obj, c)

//     // 3. 判断验证是否通过
//     if len(errs) > 0 {
//         response.ValidationError(c, errs)
//         return false
//     }

//     return true
// }
