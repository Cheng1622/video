package controller

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"

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
		weblib.NewWebBase(t.W, t.R).WebErr("error")
	}

	// 协程处理
	go func() {
		U, err := url.Parse(t.R.FormValue("url"))

		if err != nil {
			t.ss.Set("vStatus", 0)
			return
		}
		t.ss.Set("vStatus", 2)
		// 可以改为websock长连接 便于完成后主动推送，以及服务进度推送

		//  此处 sql记录提交记录

		vidile, err := http.Get(t.R.FormValue("url"))
		resUrl1 := strings.Split(U.Path, ".")
		resPath := "./Up/" + t.ss.GetId()[:len(t.ss.GetId())-1] + "." + resUrl1[1]
		// download success

		// 需要分块读取，待修改
		vidile1, _ := io.ReadAll(vidile.Body)
		defer vidile.Body.Close()
		ioutil.WriteFile("./Up/"+t.ss.GetId(), vidile1, 0755)
		// whrite success
		t.ss.Set("vStatus", 3)
		if err != nil {
			t.ss.Set("vStatus", 0)
			return
		}

		// ffmpeg 处理
		// resPath:= vidile.Header.Get()
		log.Print(U)
		cmd := exec.Command(
			"/usr/bin/ffmpeg", "-i", "./Up/"+t.ss.GetId(), "-ss", t.R.FormValue("begin"), "-to", t.R.FormValue("end"), "-strict", "-2", "-qscale", "0", "-intra", resPath)
		err = cmd.Run()

		if err != nil {
			t.ss.Set("vStatus", 0)
			return
		}

		// 处理完成后赋值
		resUrl2 := strings.Split(resPath, "/")

		t.ss.Set("vurl", "/api/Index/Down/"+resUrl2[2])

		t.ss.Set("vStatus", 9)
	}()
}

// 获取进度
func (t *Index) GetStatus() {
	if t.ss.Get("vStatus") == nil {
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
func (t *Index) Down() {
	fn1 := strings.Split(t.R.URL.Path, "/")
	fn := "./Up/" + fn1[len(fn1)-1]
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic("")
	}
	t.W.Header().Add("content-type", "video/mp4")
	t.W.WriteHeader(200)
	t.W.Write(data)

}
