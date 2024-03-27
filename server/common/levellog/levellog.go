package levellog

import "log"

func Controller(w string) {
	log.Println("controller层：", w)
}

func Service(w string) {
	log.Println("service层：", w)
}

func Dao(w string) {
	log.Println("dao层：", w)
}
