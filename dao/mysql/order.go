package mysql

import (
	"armor_plate/model"
	"armor_plate/pkg/baiduMap"
	priqueue "armor_plate/pkg/pri_queue"
	"errors"
	"gorm.io/gorm"
)

//根据订单id查询
func GetOrderByID(id uint64) (order *model.Order, err error) {
	err = db.Model(&model.Order{}).Where("order_id = ?", id).First(&order).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return order, errors.New("不存在此订单")
	}
	return order, nil
}

// CreateOrder 生成订单
func CreateOrder(mes *model.Order) (err error) {
	//通过产品ID获取库存数量
	amount, err := GetProductAmount(mes.ProductID)
	if err != nil {
		return
	}
	//如果订单数量大于库存数量的3倍，则取消该笔订单
	if mes.Amount > amount*3 {
		return errors.New("订单数量过大，工人加班也做不了这么多，换人吧")
	}
	price, err := GetPriceById(mes.ProductID)
	if err != nil {
		return
	}
	//记录订单金额
	money := float64(mes.Amount) * price
	mes.Money = money
	err = db.Model(&model.Order{}).Create(&mes).Error
	return
}

// DeliveryOrder 交付订单业务
func DeliveryOrder(mes *model.Order) (err error) {
	amount, err := GetProductAmount(mes.ProductID)
	if err != nil {
		return
	}
	if mes.Amount > amount {
		return errors.New("仓库数量不足，请及时补货")
	}
	price, err := GetPriceById(mes.ProductID)
	if err != nil {
		return
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		//订单的支付状态是已支付
		if mes.State != 0 {
			lng, lat := baiduMap.Get_JWData_By_Ctiy(mes.Destination) //获取订单目的地址的经纬度
			//根据货物id获取仓库列表
			//ckList, err := GetPathListByID(mes.ProductID)
			//if err != nil {
			//	return err
			//}
			type distance = struct {
				DepotID  uint64  `json:"depot_id"`
				Distance float64 `json:"distance"` //距离
			}
			var dis []distance
			tx.Raw(`
				SELECT d.depot_id AS depot_id,
					(6371 * acos(cos(radians(?)) * cos(radians(d.latitude)) * cos(radians(d.longitude) - radians(?)) + sin(radians(?)) * sin(radians(d.latitude)))) AS distance
				FROM ap_depot d,ap_address a where d.state != 0 and d.depot_id = a.depot_id and a.product_id = ? ORDER BY distance LIMIT 0,5;`, lng, lat, lng, mes.ProductID).Scan(&dis)
			//fmt.Println(dis)
			//将获取到距离最近的5条仓库记录保存到列表里
			
			//err = exec.Model(&distance{}).Find(&dis).Error
			//if err != nil {
			//	return err
			//}
			
			//生成优先级队列
			priorityQueue := priqueue.PriorityQueue{}
			for i, v := range dis {
				priorityQueue = append(priorityQueue, &priqueue.Queue{
					Index:    i,
					DepotID:  v.DepotID,
					Distance: v.Distance,
					Priority: GetPriByID(v.DepotID),
					Amount:   GetAmuByID(mes.ProductID, v.DepotID),
				})
			}
			priqueue.GenQueue(&priorityQueue)
			//按照比较器排序后推出优先级最高的
			queue := priorityQueue.Pop().(*priqueue.Queue)
			//fmt.Println(queue)
			//建立出库单模型
			outOrder := []*model.OutgoingOrder{}
			num := mes.Amount //参数记录一下订单需求数量
			//fmt.Println(num)
			for num > 0 {
				var count int
				if num-queue.Amount > 0 {
					count = queue.Amount
				} else {
					count = num
				}
				outOrder = append(outOrder, &model.OutgoingOrder{
					OrderID:      mes.OrderID,
					EmployeeID:   mes.EmployeeID,
					DepotID:      queue.DepotID,
					ProductID:    mes.ProductID,
					CustomerName: mes.CustomerName,
					Amount:       count,
					Money:        price * float64(count),
					Destination:  mes.Destination,
					Path:         GetPathByID(queue.DepotID),
				})
				//修改存储信息表的存货量
				tx.Model(&model.Address{}).Where("product_id = ? and depot_id = ?", mes.ProductID, queue.DepotID).Update("count", queue.Amount-count)
				//生成一个出库单，订单数量也减去出库数量
				//fmt.Println("num=, queue=", num, queue.Amount)
				num = num - queue.Amount
				//当第一个仓库的数量不足时，再推出第二高优先级仓库发货
				if num > 0 {
					queue = priorityQueue.Pop().(*priqueue.Queue)
				}
			}
			//fmt.Println(outOrder)
			err = CreateOutgoingOrder(outOrder)
			if err != nil {
				return err
			}
		} else {
			return errors.New("都可以，但是要先给钱")
		}
		//提交事务
		return nil
	})
	//交付订单后将订单的状态改为1
	db.Table("ap_order").Where("order_id = ?", mes.OrderID).Update("state", 1)
	return
}

// CreateOutgoingOrder 生成出库单
func CreateOutgoingOrder(orders []*model.OutgoingOrder) (err error) {
	for _, v := range orders {
		err = db.Table("ap_outgoing_order").Create(&v).Error
		if err != nil {
			return
		}
	}
	return nil
}
