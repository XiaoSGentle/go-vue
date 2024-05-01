package xcore

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"xcore/common/xcaptcha"
	"xcore/common/xconfig"
	"xcore/common/xgorm"
	"xcore/common/xlogger"
	"xcore/common/xvalidate"
	"xcore/core/xdig"
	"xcore/core/xvariable"
	"xcore/dao/model"
	"xcore/dao/query"
)

type GinCore struct {
	regFunctions []interface{}
	routerGroup  []*GroupBase
	router       *gin.Engine
}

type RouterGroup struct {
}

func NewGinCore() *GinCore {
	CheckRequiredFolds([]string{
		"./public", "./config/",
	})
	// 初始化全局变量
	InitializeGlobalVariables()
	var router *gin.Engine
	// 默认跨域访问

	// 非调试模式（生产模式） 日志写到日志文件
	if xvariable.GlobalYmlConfig.GetBool("AppDebug") == false {
		//1.gin自行记录接口访问日志，不需要nginx，如果开启以下3行，那么请屏蔽第 34 行代码
		//gin.DisableConsoleColor()
		//f, _ := platform.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		//gin.DefaultWriter = io.MultiWriter(f)
		//【生产模式】
		// 根据 gin 官方的说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
		// 如果部署到生产环境，请使用以下模式：
		// 1.生产模式(release) 和开发模式的变化主要是禁用 gin 记录接口访问日志，
		// 2.go服务就必须使用nginx作为前置代理服务，这样也方便实现负载均衡
		// 3.如果程序发生 panic 等异常使用自定义的 panic 恢复中间件拦截、记录到日志
		router = ReleaseRouter()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		pprof.Register(router)
	}

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))
	// 定义结构体返回
	core := &GinCore{
		regFunctions: make([]interface{}, 0),
		routerGroup:  make([]*GroupBase, 0),
		router:       router,
	}
	return core
}

func (receiver *GinCore) RegisterRouterGroups(router []*GroupBase) {
	receiver.routerGroup = append(receiver.routerGroup, router...)
}

func (receiver *GinCore) RegisterRouterGroup(router *GroupBase) {
	receiver.routerGroup = append(receiver.routerGroup, router)
}

func (receiver *GinCore) RegisterRegFunctions(functions []interface{}) {
	receiver.regFunctions = append(receiver.regFunctions, functions...)
}

func (receiver *GinCore) RegisterRegFunction(function interface{}) {
	receiver.regFunctions = append(receiver.regFunctions, function)
}

func (receiver *GinCore) RegisterGinRouter(relativePath string, handlers ...gin.HandlerFunc) {
	receiver.router.Group(relativePath, handlers...)
}

func (receiver *GinCore) Run(port string) {

	for _, regFunction := range receiver.regFunctions {
		err := xdig.ProvideForDI(regFunction)
		if err != nil {
			xvariable.Logger.ErrorContext(context.Background(), "注册DIG出错:"+err.Error())
		}
	}
	mainRouter := receiver.router.Group("/api")
	for _, groupBase := range receiver.routerGroup {
		RegisterGroup(mainRouter, groupBase)
	}

	if xvariable.GlobalYmlConfig.GetBool("AppDebug") == true {
		InitSqlInfo(receiver.router.Routes())
	}

	err := receiver.router.Run(port)
	if err != nil {
		xvariable.Logger.ErrorContext(context.Background(), "启动失败！"+err.Error())
	}
}

func InitSqlInfo(routers []gin.RouteInfo) {
	apiQuery := query.Use(xvariable.GormDB).SysAPI

	// 清除所欲路由
	_, _ = apiQuery.Where(apiQuery.APICode.IsNotNull()).Delete()

	for _, router := range routers {
		if !strings.HasPrefix(router.Path, "/debug/") {
			_ = apiQuery.Create(&model.SysAPI{
				APICode:       fmt.Sprintf("%s::%s", router.Method, router.Path),
				Version:       0,
				SoftDeleteTag: 0,
				UpdateTime:    time.Now(),
				UpdateUID:     0,
				CreateUID:     0,
				CreateBy:      "",
				CreateTime:    time.Now(),
				UpdateBy:      "",
			})

		}

	}

}

func InitializeGlobalVariables() {
	// 初始化系统
	xvariable.Logger = xlogger.NewSlog(xvariable.BasePath + "/logs")
	// 系统配置文件初始化 并启动文件监听
	xvariable.GlobalYmlConfig = xconfig.CreateYamlFactory("config")
	xvariable.GlobalYmlConfig.ConfigFileChangeListen()
	// GORM配置文件初始化 并启动文件监听
	xvariable.GormYmlConfig = xconfig.CreateYamlFactory("gorm")
	xvariable.GormYmlConfig.ConfigFileChangeListen()
	// 初始化全局字段翻译器
	xvalidate.InitTransValidator(xvariable.GlobalYmlConfig.GetString("HttpServer.ValidateLang"))
	// 全局GORM链接
	xvariable.GormDB = xgorm.GetMysqlConnection()
	// 全局验证码
	xvariable.Captcha = xcaptcha.InitCaptcha()
}

func CheckRequiredFolds(paths []string) {
	if len(paths) == 0 {
		return
	}
	for _, path := range paths {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			fmt.Printf("文件夹 %s 不存在,请检查环境配置\n", path)
		}
	}
}

// ReleaseRouter 根据 gin 路由包官方的建议，gin 路由引擎如果在生产模式使用，官方建议设置为 release 模式
// 官方原版提示说明：[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
// 这里我们将按照官方指导进行生产模式精细化处理
func ReleaseRouter() *gin.Engine {
	// 切换到生产模式禁用 gin 输出接口访问日志，经过并发测试验证，可以提升5%的性能
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	engine := gin.New()
	// 载入gin的中间件，关键是第二个中间件，我们对它进行了自定义重写，将可能的 panic 异常等，统一使用 slog 接管，保证全局日志打印统一
	engine.Use(gin.Logger(), CustomRecovery())
	return engine
}

// CustomRecovery 自定义错误(panic等)拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() gin.HandlerFunc {
	DefaultErrorWriter := &PanicExceptionRecord{}
	return gin.RecoveryWithWriter(DefaultErrorWriter, func(c *gin.Context, err interface{}) {
		// 这里针对发生的panic等异常进行统一响应即可
		// 这里的 err 数据类型为 ：runtime.boundsError  ，需要转为普通数据类型才可以输出
		xvariable.Logger.ErrorContext(c, fmt.Sprintf("%s", err))
	})
}

// PanicExceptionRecord  panic等异常记录
type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	errStr := string(b)
	err = xerrors.New(errStr)
	xvariable.Logger.ErrorContext(context.Background(), "系统出错！")
	return len(errStr), err
}
