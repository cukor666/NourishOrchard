package fruitctrl

import (
	"fmt"
	"server/common/levellog"
)

func validName(name string) bool {
	if len(name) == 0 {
		levellog.Valid("没有传递name")
		return true
	}
	if len(name) < 2 || len(name) > 20 {
		levellog.Valid("参数name校验失败，name = " + name)
		return false
	}
	levellog.Valid("参数name校验成功")
	return true
}

func validWater(w int) bool {
	if w == 0 {
		levellog.Valid("没有传递water")
		return true
	}
	if w < 0 || w >= 100 {
		levellog.Valid(fmt.Sprintf("参数water校验失败，water = %d", w))
		return false
	}
	levellog.Valid("参数water校验成功")
	return true
}

func validSugar(s int) bool {
	if s == 0 {
		levellog.Valid("没有传递sugar")
		return true
	}
	if s <= 0 || s >= 100 {
		levellog.Valid(fmt.Sprintf("参数sugar校验失败，sugar = %d", s))
		return false
	}
	levellog.Valid("参数sugar校验成功")
	return true
}

func validShelfLife(sl int) bool {
	if sl == 0 {
		levellog.Valid("没有传递shelLife")
		return true
	}
	if sl < 0 || sl > 90 {
		levellog.Valid(fmt.Sprintf("参数shelfLife校验失败，shelfLife = %d", sl))
		return false
	}
	levellog.Valid("参数shelfLife校验成功")
	return true
}

func validOrigin(o string) bool {
	if len(o) == 0 {
		levellog.Valid("没有传递origin")
		return true
	}
	if len(o) < 2 || len(o) > 200 {
		levellog.Valid("参数origin校验失败，origin = " + o)
		return false
	}
	levellog.Valid("参数origin校验成功")
	return true
}

func validSupplierId(sid int) bool {
	if sid == 0 {
		levellog.Valid("没有传递supplierId")
		return true
	}
	if sid < 0 || sid > 20 {
		levellog.Valid(fmt.Sprintf("参数supplierId校验失败，sid = %d", sid))
		return false
	}
	levellog.Valid("参数supplierId校验成功")
	return true
}
