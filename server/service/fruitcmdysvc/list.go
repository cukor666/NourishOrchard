package fruitcmdysvc

import (
	"encoding/json"
	"errors"
	"log"
	"server/common/levellog"
	"server/dao/fruitcmdydao"
	"server/dto/fruitdto"
	"server/models"
	"strings"
)

func List(flag int) ([]models.FruitCommodity, int64, error) {
	// 把所有的水果都先查询出来
	var (
		result []models.FruitCommodity
		r      []fruitdto.FruitCommodity
		total  int64
		err    error
	)
	r, total, err = fruitcmdydao.List()
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	// 根据flag来过滤水果
	switch flag {
	case 0:
		// 查询全部
		bytes, err := json.Marshal(r)
		if err != nil {
			levellog.Service("序列化失败")
			return nil, 0, err
		}
		err = json.Unmarshal(bytes, &result)
		return result, total, err
	case 1:
		// 查询南方水果
		result, total = filterSouthFruits(r)
	case 2:
		// 查询北方水果
		result, total = filterNorthFruits(r)
	default:
		err = errors.New("参数错误")
		return nil, 0, err
	}
	return result, total, nil
}

// 过滤南方水果
func filterSouthFruits(fruits []fruitdto.FruitCommodity) ([]models.FruitCommodity, int64) {
	/*
		上海 江苏 浙江 安徽 福建
		广东 广西 海南 重庆 四川 贵州 云南 西藏
		台湾 香港 澳门
	*/
	cities := []string{
		"上海", "江苏", "浙江", "安徽", "福建", "江西", "湖北", "湖南",
		"广东", "广西", "海南", "重庆", "四川", "贵州", "云南", "西藏",
		"台湾", "香港", "澳门",
	}
	result := make([]models.FruitCommodity, 0)
	total := int64(0)
	for _, fruit := range fruits {
		origin := fruit.Origin
		// 如果origin包含cities中的元素，则添加到结果中
		for _, city := range cities {
			if strings.Contains(origin, city) {
				var f models.FruitCommodity
				bytes, err := json.Marshal(fruit)
				if err != nil {
					levellog.Service("序列化失败")
					return nil, 0
				}
				err = json.Unmarshal(bytes, &f)
				if err != nil {
					levellog.Service("反序列化失败")
					return nil, 0
				}
				result = append(result, f)
				total++
				log.Println(f.ID)
			}
		}
	}

	return result, total
}

// 过滤北方水果
func filterNorthFruits(fruits []fruitdto.FruitCommodity) ([]models.FruitCommodity, int64) {
	/*
		北京 天津 河北 山西 内蒙 辽宁 吉林 黑龙江
		河南 山东 陕西 甘肃 宁夏 青海 新疆
	*/
	cities := []string{
		"北京", "天津", "河北", "山西", "内蒙", "辽宁", "吉林", "黑龙江",
		"河南", "山东", "陕西", "甘肃", "宁夏", "青海", "新疆",
	}
	result := make([]models.FruitCommodity, 0)
	total := int64(0)
	for _, fruit := range fruits {
		origin := fruit.Origin
		// 如果origin包含cities中的元素，则添加到结果中
		for _, city := range cities {
			if strings.Contains(origin, city) {
				var f models.FruitCommodity
				bytes, err := json.Marshal(fruit)
				if err != nil {
					levellog.Service("序列化失败")
					return nil, 0
				}
				err = json.Unmarshal(bytes, &f)
				if err != nil {
					levellog.Service("反序列化失败")
					return nil, 0
				}
				result = append(result, f)
				total++
				log.Println(f.ID)
			}
		}
	}

	return result, total
}

/*

// 北方水果
9
12
15
17
22
28
33
35


// 南方水果
1
2
3
4
5
6
8
10
11
18
19
20
23
24
25
26
27
29
30
31
32
34
36
38
40
41
42
43
*/
