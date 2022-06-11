package controller

import (
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cjie9759/goWeb/ext/session"
	"github.com/cjie9759/goWeb/ext/weblib"
)

type Index struct {
	BaseApp
	ss session.Session
}

func (t *Index) Init() {

	s := Sm.BeginSession(t.W, t.R)
	Sm.Update(t.W, t.R)
	t.ss = s
	// 此处可以加入登陆判断 ，识别用户身份
}
func (t *Index) Up() {
	// 验证数据
	t.ss.Set("vStatus", 1)
	t.ss.Set("vurl", "")
	if t.R.FormValue("url") != "" {
		weblib.NewWebBase(t.W, t.R).WebSucess("sucess")
	} else {
		weblib.NewWebBase(t.W, t.R).WebErr("sucess")
	}

	// 协程处理
	go func() {
		vidile, err := http.Get(t.R.FormValue("url"))
		// download success
		t.ss.Set("vStatus", 2)
		if err != nil {
			t.ss.Set("vStatus", 0)
			return
		}
		// 可以改为websock长连接 便于完成后主动推送，以及服务进度推送

		//  此处 sql记录提交记录

		// 需要分块读取，待修改
		vidile1, _ := io.ReadAll(vidile.Body)
		defer vidile.Body.Close()
		ioutil.WriteFile("./public/"+t.ss.GetId(), vidile1, 0755)
		// whrite success
		t.ss.Set("vStatus", 3)
		if err != nil {
			t.ss.Set("vStatus", 0)
			return
		}

		// ffmpeg 处理
		// 处理完成后赋值
		t.ss.Set("vurl", "")

		t.ss.Set("vStatus", 9)
	}()
}

// 获取进度
func (t *Index) GetStatus() {
	if t.ss.Get("vStatus") != nil {
		weblib.NewWebBase(t.W, t.R).WebErr("0")
		return
	}
	weblib.NewWebBase(t.W, t.R).WebSucess(strconv.Itoa(t.ss.Get("vStatus").(int)))
}

func (t *Index) GetUrl() {
	if t.ss.Get("vStatus").(int) != 9 {
		weblib.NewWebBase(t.W, t.R).WebErr("0")
		return
	}
	weblib.NewWebBase(t.W, t.R).WebSucess(t.ss.Get("vurl").(string))
}
