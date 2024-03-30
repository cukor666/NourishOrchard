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

func Config(w string) {
	log.Println("config层：", w)
}

func Valid(w string) {
	log.Println("valid层：", w)
}

func Router(w string) {
	log.Println("router层：", w)
}
