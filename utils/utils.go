package utils

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"strings"
	"time"
)

func PrintNow(){
	fmt.Println(time.Now())
}

//Music163TopList 需要下载chrome及对应版本的ChromeDrive TopListID为网易云网页榜单id
func Music163TopList(TopListID int)[]string {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	//service, err := selenium.NewChromeDriverService("./chromedriver", 9515, opts...)
	//if err != nil {
	//log.Printf("Error starting the ChromeDriver server: %v", err)
	//}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d", 9515))
	if err != nil {
		panic(err)
	}
	err=webDriver.Get("https://music.163.com/#/discover/toplist?id=19723756")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)}
	webDriver.SwitchFrame("g_iframe")

	if err!=nil{fmt.Printf("Failed to find element 1: %s\n", err)}

	webE11,err:=webDriver.FindElement(selenium.ByTagName,"tbody")
	if err != nil {
		fmt.Printf("Failed html12: %s\n", err)}
	webE12,err:=webE11.FindElements(selenium.ByTagName,"tr")
	if err != nil {
		fmt.Printf("Failed html13: %s\n", err)}
	fmt.Println(len(webE12))
	id:=make([]string,len(webE12))
	for i:=0;i<len(webE12);i++{
		webE13,err:=webE12[i].FindElements(selenium.ByTagName,"td")
		if err!=nil{fmt.Printf("Failed to find element 14: %s\n", err)}
		webE14,err:=webE13[1].FindElement(selenium.ByClassName,"f-cb")
		if err!=nil{fmt.Printf("Failed to find element 15: %s\n", err)}
		webE15,err:=webE14.FindElement(selenium.ByClassName,"tt")
		if err!=nil{fmt.Printf("Failed to find element 16: %s\n", err)}
		webE155,err:=webE15.FindElement(selenium.ByTagName,"a")
		if err!=nil{fmt.Printf("Failed to find element 17: %s\n", err)}
		id[i],err=webE155.GetAttribute("href")
		id[i]=strings.Split(id[i],"=")[1]
		id[i]="http://music.163.com/song/media/outer/url?id="+id[i]+".mp3"
	}

	//for _,v:=range id{
	//	fmt.Println(v)}
return id
}
func Music163HotComment(id int)[]string{
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	//service, err := selenium.NewChromeDriverService("./chromedriver", 9515, opts...)
	//if err != nil {
	//log.Printf("Error starting the ChromeDriver server: %v", err)
	//}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d", 9515))
	if err != nil {
		panic(err)
	}
	url:=fmt.Sprintf("https://music.163.com/#/song?id=%v",id)
	err=webDriver.Get(url)
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)}
	webDriver.SwitchFrame("g_iframe")
	if err!=nil{fmt.Printf("Failed to find element 1: %s\n", err)}
	webE11,err:=webDriver.FindElements(selenium.ByClassName,"itm")
	if err != nil {
		fmt.Printf("Failed html12: %s\n", err)}
	comments:=make([]string,len(webE11))
	for i:=0;i<len(webE11);i++{
		webE15,err:=webE11[i].FindElement(selenium.ByClassName,"cnt")
		if err!=nil{fmt.Printf("Failed to find element 16: %s\n", err)}

		comments[i],err=webE15.Text()
		if err!=nil{fmt.Printf("Failed to find element 18: %s\n", err)}

	}



return comments


}




