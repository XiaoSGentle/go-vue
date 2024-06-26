package xcaptcha

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"time"
	"xcore/common/xerror"

	"github.com/eko/gocache/lib/v4/cache"
	bigcacheStore "github.com/eko/gocache/store/bigcache/v4"
	"github.com/wenlng/go-captcha/captcha"
	"golang.org/x/image/font"
	"os"
	"sort"
)

type CaptchaFontDots struct {
	X     int `json:"x"`
	Y     int `json:"y"`
	Index int `json:"index"`
}

type VerifyCaptchaCodeParam struct {
	Dots []CaptchaFontDots `json:"dots"`
	Key  string            `json:"key"`
}

type Captcha struct {
	captcha     *captcha.Captcha
	catchManger *cache.Cache[[]byte]
	ctx         context.Context
}

type CaptchaCodeResp struct {
	ImageBase64 string `json:"image_base64"`
	ThumbBase64 string `json:"thumb_base64"`
	CaptchaKey  string `json:"captcha_key"`
}

func NewCaptchaManager() *Captcha {
	var ctx = context.Background()
	bigCacheClient, _ := bigcache.New(ctx, bigcache.DefaultConfig(5*time.Minute))
	bigCacheStore := bigcacheStore.NewBigcache(bigCacheClient)
	cacheManager := cache.New[[]byte](bigCacheStore)
	capt := captcha.GetCaptcha()
	path, _ := os.Getwd()
	capt.SetFont([]string{
		path + `/public/resources/captcha/fonts/font_3.ttf`,
	})
	// ====================================================
	// Method: SetBackground(color []string);
	// Desc: 设置验证码背景图，自动仅读取一次并加载到内存中缓存，如需重置可清除缓存
	// ====================================================
	capt.SetBackground([]string{
		path + `/public/resources/captcha/images/1.jpg`,
		path + `/public/resources/captcha/images/2.jpg`,
		path + `/public/resources/captcha/images/3.jpg`,
		path + `/public/resources/captcha/images/4.jpg`,
	})

	// ====================================================
	// Method: SetImageSize(size Size);
	// Desc: 设置验证码主图的尺寸
	// ====================================================
	capt.SetImageSize(captcha.Size{Width: 300, Height: 240})

	// ====================================================
	// Method: SetImageQuality(val int);
	// Desc: 设置验证码主图清晰度，压缩级别范围 QualityCompressLevel1 - 5，QualityCompressNone无压缩，默认为最低压缩级别
	// ====================================================
	capt.SetImageQuality(captcha.QualityCompressNone)

	// ====================================================
	// Method: SetFontHinting(val font.Hinting);
	// Desc: 设置字体Hinting值 (HintingNone,HintingVertical,HintingFull)
	// ====================================================
	capt.SetFontHinting(font.HintingFull)

	// ====================================================
	// Method: SetTextRangLen(val xcaptcha.RangeVal);
	// Desc: 设置验证码文本显示的总数随机范围
	// ====================================================
	capt.SetTextRangLen(captcha.RangeVal{Min: 6, Max: 7})

	// ====================================================
	// Method: SetRangFontSize(val xcaptcha.RangeVal);
	// Desc: 设置验证码文本的随机大小
	// ====================================================
	capt.SetRangFontSize(captcha.RangeVal{Min: 32, Max: 42})

	// ====================================================
	// Method: SetTextRangFontColors(colors []string);
	// Desc: 设置验证码文本的随机十六进制颜色
	// ====================================================
	capt.SetTextRangFontColors([]string{
		"#1d3f84",
		"#3a6a1e",
	})

	// ====================================================
	// Method: SetImageFontAlpha(val float64);
	// Desc:设置验证码字体的透明度
	// ====================================================
	capt.SetImageFontAlpha(0.5)

	// ====================================================
	// Method: SetTextShadow(val bool);
	// Desc:设置字体阴影
	// ====================================================
	capt.SetTextShadow(true)

	// ====================================================
	// Method: SetTextShadowColor(val string);
	// Desc:设置字体阴影颜色
	// ====================================================
	capt.SetTextShadowColor("#101010")

	// ====================================================
	// Method: SetTextShadowPoint(val xcaptcha.Point);
	// Desc:设置字体阴影偏移位置
	// ====================================================
	capt.SetTextShadowPoint(captcha.Point{X: 1, Y: 1})

	// ====================================================
	// Method: SetTextRangAnglePos(pos []xcaptcha.RangeVal);
	// Desc:设置验证码文本的旋转角度
	// ====================================================
	capt.SetTextRangAnglePos([]captcha.RangeVal{
		{1, 15},
		{15, 30},
		{30, 45},
		{315, 330},
		{330, 345},
		{345, 359},
	})

	// ====================================================
	// Method: SetImageFontDistort(val int);
	// Desc:设置验证码字体的扭曲程度
	// ====================================================
	capt.SetImageFontDistort(captcha.DistortLevel2)
	return &Captcha{captcha: capt, catchManger: cacheManager, ctx: ctx}
}

func (receiver Captcha) GenerateCaptureCode() (c *CaptchaCodeResp, err error) {

	var captchaCodeResp CaptchaCodeResp
	dotInfos, imageBase64, thumbBase64, captchaKey, _err := receiver.captcha.Generate()
	if _err != nil {
		dotInfos, imageBase64, thumbBase64, captchaKey, _err = receiver.captcha.Generate()
	}
	captchaCodeResp = CaptchaCodeResp{
		ImageBase64: imageBase64,
		ThumbBase64: thumbBase64,
		CaptchaKey:  captchaKey,
	}
	// 存入BigCatch
	err = receiver.setDotCatch(dotInfos, captchaCodeResp.CaptchaKey)
	return &captchaCodeResp, _err
}

// 取出指定Key的dotInfos
func (receiver Captcha) getDotCatch(dotKey string) (dotInfos map[int]captcha.CharDot, err error) {
	dotsInRedisBytes, _ := receiver.catchManger.Get(receiver.ctx, GenerateCatchKey(dotKey))
	// 反序列化[]byte为map[int]xcaptcha.CharDot
	decoder := gob.NewDecoder(bytes.NewReader(dotsInRedisBytes))
	var decodedData map[int]captcha.CharDot
	err = decoder.Decode(&decodedData)
	dotInfos = decodedData
	return
}

// 放进缓存
func (receiver Captcha) setDotCatch(dotInfos map[int]captcha.CharDot, dotsKey string) (err error) {
	// 存入BigCatch
	var bufferByte = new(bytes.Buffer)
	newEncoder := gob.NewEncoder(bufferByte)
	err = newEncoder.Encode(dotInfos)
	err = receiver.catchManger.Set(receiver.ctx, GenerateCatchKey(dotsKey), bufferByte.Bytes())
	return
}

// VerifyCaptcha  中间件使用该函数去验证是否正确
func (receiver Captcha) VerifyCaptcha(g *gin.Context) (result bool, err error) {

	captureCode := g.Request.Header.Get("CaptureCode")
	captureDots := g.Request.Header.Get("CaptureDots")
	if captureCode == "" || captureDots == "" {
		return false, xerror.NewErrCode(xerror.CAPTCHA_KEY_NOT_FOUND_ERROR)
	}
	var dots []CaptchaFontDots
	_err := json.Unmarshal([]byte(captureDots), &dots)
	if _err != nil {
		return false, xerror.NewErrCode(xerror.CAPTCHA_VERIFY_ERROR)
	}
	param := VerifyCaptchaCodeParam{dots, captureCode}
	result = receiver._checkCaptchaDots(param)
	return
}

func (receiver Captcha) _checkCaptchaDots(v VerifyCaptchaCodeParam) bool {
	// get := redis_factory.RedisUtil.Get(v.Key)
	dotsInRedis, _ := receiver.getDotCatch(v.Key)
	chkRet := false
	for i, dotInRedis := range dotsInRedis {
		// 如果长度不一样 失败
		if len(dotsInRedis) != len(v.Dots) {
			break
		}
		sort.Slice(v.Dots, func(i, j int) bool {
			return v.Dots[i].Index < v.Dots[j].Index
		})
		UX := v.Dots[i].X
		UY := v.Dots[i].Y
		// 检测点位置
		// chkRet = xcaptcha.CheckPointDist(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height))

		// 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
		// 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50/
		/**
		@Description: 计算点的位置在扩张区域(原区域+外边距)是否命中
		@param sx 用户点击的x轴
		@param sy 用户点击的y轴
		@param dx 校验文本的x轴
		@param dy 校验文本的y轴
		@param width 校验文本的宽度
		@param height 校验文本的高度
		@param padding 在原有的区域上添加额外边距进行扩张计算区域，不推荐设置padding
		@return bool
		*/
		chkRet = captcha.CheckPointDistWithPadding(int64(UX), int64(UY), int64(dotInRedis.Dx), int64(dotInRedis.Dy), int64(dotInRedis.Width), int64(dotInRedis.Height), 7)
		if !chkRet {
			break
		}
	}
	// 无论验证结果都要删除redis中数据
	_ = receiver.catchManger.Delete(receiver.ctx, GenerateCatchKey(v.Key))
	return chkRet
}

func GenerateCatchKey(code string) string {
	return "Captcha:Code:" + code
}
