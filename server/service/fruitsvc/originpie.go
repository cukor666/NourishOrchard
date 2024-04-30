package fruitsvc

import (
	"log"
	"server/common/levellog"
	"server/dao/fruitdao"
	"sync"
)

func OriginPie() (result map[string]int64, err error) {
	result = make(map[string]int64)
	wg := new(sync.WaitGroup)
	wg.Add(5)
	go originCount("广西", result, wg)
	go originCount("广东", result, wg)
	go originCount("海南", result, wg)
	go originCount("福建", result, wg)
	go originCount("四川", result, wg)
	cnt, err := fruitdao.OriginCount("")
	if err != nil {
		levellog.Service("获取全部水果原产地数量失败")
	}
	wg.Wait()
	for _, v := range result {
		cnt -= v
	}
	result["其他"] = cnt
	log.Println("result: ", result)
	return
}

func originCount(key string, m map[string]int64, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
	m[key] = 0
	m[key], err = fruitdao.OriginCount(key)
	if err != nil {
		levellog.Service("获取水果原产地数量失败")
		return
	}
}
