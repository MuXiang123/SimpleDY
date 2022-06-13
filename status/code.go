package status

const (
	Success                          = iota // 成功
	RequestParamError                       // 请求参数错误
	UnknownError                            // 未知错误
	UsernameHasExistedError                 // 用户名已存在
	GenerateTokenError                      // 生成token出错
	TokenIsNULL                             // Token为空
	TokenParseError                         // 通过token获取id出错
	TokenIsExpired                          // Token过期
	UserNotExistOrPasswordWrongError        // 用户名不存在或密码错误
	LoadFileError                           // 加载文件出错
	SaveUploadedFileError                   // 保存文件出错
	AttentionExistsError                    // 已经关注错误
	AttentionNullError                      // 未关注错误
	InabilityToFocusOnYourself              // 无法关注自己
)

//
