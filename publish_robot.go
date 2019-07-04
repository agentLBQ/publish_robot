package main

import (
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.Create("publish.log")

	logger := log.New(file, "", log.Llongfile)
	logger.Println("winter is coming")

	times := 0
	publishOnce := false

	for {
		if time.Now().Hour() > 7 {
			times += 1
			logger.Println("开始发帖，第" + strconv.Itoa(times) + "次,时间：" + strconv.Itoa(time.Now().Hour()))
			httpHandle(logger, publishOnce)
			publishOnce = false
			logger.Println("休眠开始,时间：" + strconv.Itoa(time.Now().Hour()))
			time.Sleep(time.Duration(60+rand.Intn(30)) * time.Minute)
			logger.Println("休眠结束,时间：" + strconv.Itoa(time.Now().Hour()))
		} else {
			publishOnce = true
			logger.Println("不到时间点，暂不发帖，休眠5分钟，时间：" + strconv.Itoa(time.Now().Hour()))
			time.Sleep(time.Duration(5) * time.Minute)
		}
	}
}

func httpHandle(logs *log.Logger, publishOnce bool) {
	client := &http.Client{}
	//登陆，获取cookie
	cookie1, cookie2, cookie3, cookie4, cookie5 := Login(client, logs)

	//兼职实习信息，顶帖
	time.Sleep(time.Duration(10+rand.Intn(10)) * time.Second)
	ParttimeJob(cookie1, cookie2, cookie3, cookie4, cookie5, client, logs)

	//跳槽就业，社招
	time.Sleep(time.Duration(10+rand.Intn(10)) * time.Second)
	Jump(cookie1, cookie2, cookie3, cookie4, cookie5, client, logs)

	//已发过，不再发帖
	if !publishOnce {
		return
	}

	//招聘信息专版，社招
	time.Sleep(time.Duration(110+rand.Intn(10)) * time.Second)
	JobInfo(cookie1, cookie2, cookie3, cookie4, cookie5, client, logs)
}

func JobInfo(cookie1 *http.Cookie, cookie2 *http.Cookie, cookie3 *http.Cookie, cookie4 *http.Cookie, cookie5 *http.Cookie, client *http.Client, logs *log.Logger) {
	JobInfo, _ := http.NewRequest("POST", "https://bbs.byr.cn/article/JobInfo/ajax_post.json", strings.NewReader("subject="))
	JobInfo.Header.Add("x-requested-with", "XMLHttpRequest")
	JobInfo.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	JobInfo.AddCookie(cookie1)
	JobInfo.AddCookie(cookie2)
	JobInfo.AddCookie(cookie3)
	JobInfo.AddCookie(cookie4)
	JobInfo.AddCookie(cookie5)
	respJobInfo, _ := client.Do(JobInfo)
	defer respJobInfo.Body.Close()
	b3, _ := ioutil.ReadAll(respJobInfo.Body)
	str3 := ConvertToString(string(b3), "gbk", "utf-8")
	if strings.Contains(string(str3), "成功") {
		logs.Println("发布成功【招聘信息专版】【社招】")
	} else {
		logs.Println("发布失败【招聘信息专版】【社招】，返回结果：%v", string(str3))
	}
}

func Jump(cookie1 *http.Cookie, cookie2 *http.Cookie, cookie3 *http.Cookie, cookie4 *http.Cookie, cookie5 *http.Cookie, client *http.Client, logs *log.Logger) {
	reqJump, _ := http.NewRequest("POST", "https://bbs.byr.cn/article/Jump/ajax_post.json", strings.NewReader("content="))
	reqJump.Header.Add("x-requested-with", "XMLHttpRequest")
	reqJump.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqJump.AddCookie(cookie1)
	reqJump.AddCookie(cookie2)
	reqJump.AddCookie(cookie3)
	reqJump.AddCookie(cookie4)
	reqJump.AddCookie(cookie5)
	respJump, _ := client.Do(reqJump)
	defer respJump.Body.Close()
	b4, _ := ioutil.ReadAll(respJump.Body)
	str4 := ConvertToString(string(b4), "gbk", "utf-8")
	if strings.Contains(string(str4), "成功") {
		logs.Println("发布成功【跳槽就业】【社招】")
	} else {
		logs.Println("发布失败【跳槽就业】【社招】，返回结果：%v", string(str4))
	}
}

func ParttimeJob(cookie1 *http.Cookie, cookie2 *http.Cookie, cookie3 *http.Cookie, cookie4 *http.Cookie, cookie5 *http.Cookie, client *http.Client, logs *log.Logger) {

	reqParttimeJob, _ := http.NewRequest("POST", "https://bbs.byr.cn/article/ParttimeJob/ajax_post.json", strings.NewReader("content="))
	reqParttimeJob.Header.Add("x-requested-with", "XMLHttpRequest")
	reqParttimeJob.Header.Add("Accept-Encoding", "")
	reqParttimeJob.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqParttimeJob.AddCookie(cookie1)
	reqParttimeJob.AddCookie(cookie2)
	reqParttimeJob.AddCookie(cookie3)
	reqParttimeJob.AddCookie(cookie4)
	reqParttimeJob.AddCookie(cookie5)
	respParttimeJob, _ := client.Do(reqParttimeJob)
	defer respParttimeJob.Body.Close()
	b2, _ := ioutil.ReadAll(respParttimeJob.Body)
	str2 := ConvertToString(string(b2), "gbk", "utf-8")
	if strings.Contains(string(str2), "成功") {
		logs.Println("发布成功【兼职实习信息】")
	} else {
		logs.Println("发布失败【兼职实习信息】，返回结果：%v", string(str2))
	}
}

func Login(client *http.Client, logs *log.Logger) (*http.Cookie, *http.Cookie, *http.Cookie, *http.Cookie, *http.Cookie) {
	reqLogin, _ := http.NewRequest("POST", "https://bbs.byr.cn/user/ajax_login.json", strings.NewReader("id=&passwd="))
	reqLogin.Header.Add("x-requested-with", "XMLHttpRequest")
	reqLogin.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqLogin.Header.Set("Cookie", "name=anny")
	respLogin, _ := client.Do(reqLogin)
	b, _ := ioutil.ReadAll(respLogin.Body)
	str := ConvertToString(string(b), "gbk", "utf-8")
	//登陆
	if strings.Contains(string(str), "成功") {
		logs.Println("登陆成功")
	} else {
		logs.Println("登陆失败，返回结果：%v", string(str))
	}
	var UTMPKEY string
	var UTMPNUM string
	for k, v := range respLogin.Header {
		com01 := strings.Compare(k, "Set-Cookie")
		if com01 == 0 {
			UTMPKEY = v[4]
			UTMPNUM = v[5]
		}
	}
	UTMPKEY = strings.Split(strings.Split(UTMPKEY, "=")[1], ";")[0]
	UTMPNUM = strings.Split(strings.Split(UTMPNUM, "=")[1], ";")[0]
	logs.Println("解析cookie成功，UTMPKEY：%s,UTMPNUM：%s", UTMPKEY, UTMPNUM)
	cookie1 := &http.Cookie{Name: "nforum[UTMPUSERID]", Value: "lbq000lbq", HttpOnly: true}
	cookie2 := &http.Cookie{Name: "nforum[UTMPKEY]", Value: UTMPKEY, HttpOnly: true}
	cookie3 := &http.Cookie{Name: "nforum[UTMPNUM]", Value: UTMPNUM, HttpOnly: true}
	cookie4 := &http.Cookie{Name: "nforum[PASSWORD]", Value: "AHZrliSSzwYQp2oxn56N3g%3D%3D", HttpOnly: true}
	cookie5 := &http.Cookie{Name: "nforum[XWJOKE]", Value: "hoho", HttpOnly: true}
	return cookie1, cookie2, cookie3, cookie4, cookie5
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
