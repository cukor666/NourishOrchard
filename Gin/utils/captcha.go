package utils

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 配置Session
func SessionConfig() sessions.Store {
	sessionMaxAge := 3600
	sessionSecret := "cukor.cn"
	store := cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, // session的生命周期，单位：秒
		Path:   "/",
	})
	return store
}

// 中间件，处理session
func Session(keyPairs string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(keyPairs, store)
}

// 生成图片

/*
w: 使用http写操作
r: 使用http读操作
id: 从session中获取的验证码的ID
ext: 过期时间
lang: 语言
download: 是否下载
width: 验证码图片的宽度
height: 验证码图片的高度
*/
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	// 响应头
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	// 图片格式
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

// 生成验证码
func Captcha(c *gin.Context, length ...int) {
	l := captcha.DefaultLen // 验证码的默认长度，例如：9412  ---> 4
	w, h := 107, 36         // 验证码图片的默认宽度和高度
	if len(length) == 1 {
		l = length[0] // 如果传了一个参数，就替换掉默认的验证码长度
	}
	if len(length) == 2 {
		w = length[1] // 如果传了两个参数，就讲默认的宽度替换掉
	}
	if len(length) == 3 {
		h = length[2] // 如果传了三个参数，就讲默认的高度替换掉
	}
	captchaId := captcha.NewLen(l)
	session := sessions.Default(c)
	session.Set("captcha", captchaId)
	captchaId2 := session.Get("captcha") // 获取到验证码ID
	log.Printf("captchaId: %s", captchaId2)
	_ = session.Save()                                                   // 保存更新后的session
	_ = Serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h) // 生成图片
}

// 验证，验证码
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)      // 获取到上下文的session
	captchaId := session.Get("captcha") // 获取到验证码ID
	// 这里出现了bug，无法从session中获取captcha，但是上面明明已经设置了，而且上面可以获取得到
	if captchaId != nil {
		log.Printf("有captchaId")
		session.Delete("captcha")                             // 如果有验证码，那就将session中的captcha去除
		_ = session.Save()                                    // 保存当前session状态
		return captcha.VerifyString(captchaId.(string), code) // 比较系统生成的验证码和用户的验证码是否相同
	}
	log.Printf("没有captchaId")
	return false // 如果没有从session中获取验证码ID就直接返回false
}
