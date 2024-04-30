package xerror

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/
// 参数绑定相关错误

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002

// TOKEN_EXPIRE_ERROR TOKEN相关错误码
const TOKEN_EXPIRE_ERROR uint32 = 100050
const TOKEN_GENERATE_ERROR uint32 = 100051
const TOKEN_FORMAT_ERROR uint32 = 100052
const TOKEN_ERROR uint32 = 100051

// DB_ERROR DB相关错误
const DB_ERROR uint32 = 100100
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100101

// PARAM_BIND_ERROR 参数绑定失败
const PARAM_BIND_ERROR uint32 = 100150
const PARAM_VALIDATE_ERROR uint32 = 100150

//用户模块
