package ivtsvc

import (
	"server/dao/empdao"
	"server/dao/fruitdao"
	"server/dao/whdao"
	"sync"
)

func hasFruit(id int64, wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()
	has := fruitdao.Has(id)
	c <- has
}

func hasEmp(id int64, wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()
	has := empdao.Has(id)
	c <- has
}

func hasWarehouse(id int64, wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()
	has := whdao.Has(id)
	c <- has
}
