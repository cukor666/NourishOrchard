package ivtsvc

import (
	"fmt"
	"server/common/levellog"
	"server/models"
	"sync"
)

// has 检查必要的是否都存在
func has(inventory models.Inventory) bool {
	var wg sync.WaitGroup
	gCnt := 3
	wg.Add(gCnt)
	okChan := make(chan bool, gCnt)

	// 拿出水果id查看是否有该水果，有则继续执行，否则结束
	go hasFruit(inventory.CommodityId, &wg, okChan)
	// 拿出员工id查看是否有该员工，有则继续执行，否则结束
	go hasEmp(inventory.EmployeeId, &wg, okChan)
	// 拿出仓库id查看是否有该仓库，有则继续执行，否则结束
	go hasWarehouse(inventory.WarehouseId, &wg, okChan)

	wg.Wait()
	close(okChan) // 记得关闭，否则会死锁，关闭后仍然可以继续从中读取，但是不能再次写入
	for ok := range okChan {
		if !ok {
			w := fmt.Sprintf("数据不一致，请检查，%s表、%s表和%s表中是否存在req=%v对应id数据",
				(&models.Fruit{}).TableName(),
				(&models.Employee{}).TableName(),
				(&models.Warehouse{}).TableName(), inventory)
			levellog.Service(w)
			return false
		}
	}
	return true
}
