package Scan

import (
	"LinkDigger/OutPut"
	"LinkDigger/help"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

var httpreg = regexp.MustCompile(`<(a|link).*href=["'](.+?)["']`)

var pathList []string

/*
现在有一个关于通信间的问题。那就是我拿到的url应该如何去重与传递。
1. 直接放入数组，一次解决
2. 通过管道传入控制模块 // 这有一些问题 如管道在哪里生产，如何关闭。
3. more？

直接调用太不节能了，容易导致重复啊，看来还是要做个对比表
可以用map  快速查键、不能重复，省去去重 这样的话可以考虑用 map[string]int string是url int是深度

url的结构体
深度
类型：js、jpg、mp4

地下的两个finder高度重复，有无办法继续分割减少重复度？
 */

func Control(tg string,deep bool) {
	wg := &sync.WaitGroup{}

	fmt.Println("start")
	webUrl,err := url.Parse(tg)
	if err != nil {
		log.Fatal(err)
	}




	if deep {
		goPathFinder(webUrl,wg)
	}else {
		pathFinder(webUrl,wg)
	}

	wg.Wait()

	if deep {
		OutPut.TxtOut(webUrl.Host,pathList)
	}else {
		for _,v := range pathList {
			fmt.Println(v)
		}
	}

}


func pathFinder(tg *url.URL,wg *sync.WaitGroup){

	wg.Add(1)
	defer wg.Done()

	// 请求网页

	resp,err := http.Get(tg.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {

		log.Fatal(err)
	}
	// 调用正则，取href
	allUrl := httpreg.FindAllSubmatch(body,-1)


	for _,tmpList := range allUrl {
		tmpUrl,err := url.Parse(string(tmpList[2]))
		if err != nil {
			log.Fatal()
		}
		//判断一下href对象是否为目标域名下，防止跑到其他网站,这里用的host，有点小问题，因为如果用的旁站资源那么就会无法导入。
		if tmpUrl.Host == tg.Host && !help.IndexOf(tmpUrl.String(),pathList){
			pathList = append(pathList,tmpUrl.String())
			// 递归调用，跑完全网站
			// pathFinder(tmpUrl,wg)
		}
	}


}

func goPathFinder(tg *url.URL,wg *sync.WaitGroup){

	wg.Add(1)
	defer wg.Done()

	// 请求网页

	resp,err := http.Get(tg.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {

		log.Fatal(err)
	}
	// 调用正则，取href

	allUrl := httpreg.FindAllSubmatch(body,-1)


	for _,tmpList := range allUrl {
		tmpUrl,err := url.Parse(string(tmpList[2]))
		if err != nil {
			log.Fatal()
		}
		//判断一下href对象是否为目标域名下，防止跑到其他网站,这里用的host，有点小问题，因为如果用的旁站资源那么就会无法导入。
		if tmpUrl.Host == tg.Host && !help.IndexOf(tmpUrl.String(),pathList){
			pathList = append(pathList,tmpUrl.String())
			go goPathFinder(tmpUrl,wg)
		}
	}



}
