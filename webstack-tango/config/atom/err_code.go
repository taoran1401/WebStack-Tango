package atom

const (
	SUCCESS                      = 200  //成功
	ERROR_CODE_NO_LOGIN          = 1000 //未登录
	ERROR_CODE_PARAM             = 1001 //参数错误
	ERROR_CODE_NO_TOKEN          = 1002 //token为空
	ERROR_CODE_JSON              = 1003 //json处理失败
	ERROR_CODE_EXCEPTION         = 1004 //处理异常失败
	ERROR_CODE_PERMISSION_DENIED = 1006 //权限拒绝
	ERROR_CODE_SMS               = 1010 //短信失败
	ERROR_CODE_UPLOAD_FAILE      = 1011 //文件上传失败
	ERROR_CODE_DATA_NOT_FIND     = 6000 //数据不存在
)
