package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Call 跟注
func (p *PlayPlayData) Call(PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	//debug0120
	fmt.Println("剩余筹码", p.Name, p.Chip)
	fmt.Println("底池", PotChip)
	fmt.Println("对方行动", PlayPlayer.ActionChip)
	fmt.Println("我的行动", p.Name, p.ActionChip)
	//debug0120
	//差额
	differenceNumber := PlayPlayer.ActionChip - p.ActionChip
	p.Chip = p.Chip + p.ActionChip - PlayPlayer.ActionChip
	//底池就等于我方补齐对方的行动差额
	PotChip = differenceNumber + PotChip

	p.ActionChip = PlayPlayer.ActionChip

	//debug0120
	fmt.Println(p.Name, p.Chip)
	fmt.Println("底池", PotChip)
	//debug0120

	return PotChip
}

//Raise 加注
func (p *PlayPlayData) Raise(number1 int64, PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	var differenceNumber int64
	var canBeRaiseNumber int64
	differenceNumber = 0
	canBeRaiseNumber = 0
	fmt.Println("请输入要加注的金额：")
	//获取输入
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	raiseNumber, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println("请认真输入金额，只有数字！！！")
		p.Raise(number1, PotChip, PlayPlayer)
		return PotChip
	}
	//和对位的差价
	differenceNumber = PlayPlayer.ActionChip - p.ActionChip
	//得到最大能加注的值
	canBeRaiseNumber = p.Chip - differenceNumber
	if raiseNumber > canBeRaiseNumber {
		fmt.Println("最多只能加注：", canBeRaiseNumber)
		p.Raise(number1, PotChip, PlayPlayer)
		return PotChip
	}
	if raiseNumber < p.RaiseChipLeve {
		fmt.Println("最小加注：", p.RaiseChipLeve)
		p.Raise(number1, PotChip, PlayPlayer)
		return PotChip
	}
	//debug0121
	fmt.Println("底池", PotChip)
	fmt.Println("差价", differenceNumber)
	fmt.Println("加注额度", raiseNumber)
	//debug0121
	PotChip = differenceNumber + raiseNumber + PotChip
	p.Chip = p.Chip - raiseNumber - differenceNumber
	if p.Chip == 0 {
		p.AllinStatus = 1
	}
	p.ActionChip = p.ActionChip + differenceNumber + raiseNumber
	//debug0120
	fmt.Println(p.Name, p.Chip)
	fmt.Println(p.Name, p.ActionChip)
	fmt.Println("底池", PotChip)
	//debug0120
	return PotChip
}

//Allin 全下 无条件可以看五张 最后开牌
func (p *PlayPlayData) Allin(PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {

	return PotChip
}

//Check 过牌
func (p *PlayPlayData) Check(PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	//wait 其他玩家行动
	//next step
	// //这里是判断如果是翻前  应该是后期调整 了入口的逻辑 check在这里不用判断是不是翻前
	// if p.SiteNumber == 1 && p.ActionTurnNumber == 0 {
	// 	p.Chip = p.Chip - p.ActionChip
	// 	PotChip = p.ActionChip + PlayPlayer.ActionChip + PotChip
	// 	return PotChip
	// }
	return PotChip
}

//Fold 弃牌
func (p *PlayPlayData) Fold(PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	return PotChip
}
