package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16 //机器id
)

func getMachineID() (uint16, error) {
	//返回全局定义的机器ID
	return sonyMachineID, nil
}

//需要传入当前的机器ID
func Init(machineID uint16) (err error) {
	sonyMachineID = machineID
	//初始化一个开始的时间
	timer, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Sep 17,2022 at 5:39pm (MST)")
	setting := sonyflake.Settings{
		StartTime: timer,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(setting) //根据配置生成snoyflake节点
	return
}

//返回生成id值
func GetID() (id uint64, err error) {
	//拿到sonyflake节点生成id
	if sonyFlake == nil {
		err = fmt.Errorf("sonyflake not init")
		return
	}
	//生成下一个ID,时间溢出后产生错误
	id, err = sonyFlake.NextID()
	return
}
