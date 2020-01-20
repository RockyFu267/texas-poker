package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

//PerflopAction 翻牌前操作
func (p *PlayPlayer) PerflopAction(PotChip int64) (potChip int64, error error) {
	fmt.Println("场上选手剩余筹码：", p[0].Name, ":", p[0].Chip, "\t", p[1].Name, ":", p[1].Chip, "\t", "底池:", PotChip)
	//--debug
	fmt.Println("小盲是：", p[0].Name)
	fmt.Println("大盲是：", p[1].Name)
	//--debug
	// //目前只写机器CALL-----
	// p[0].Chip = p[0].Chip - p[0].BBSBChip
	// p[1].Chip = p[1].Chip - p[1].BBSBChip
	// PotChip = p[0].BBSBChip + p[1].BBSBChip
	// //justcall
	// p[0].Chip = p[0].Chip - p[0].BBSBChip
	// PotChip = PotChip + p[0].BBSBChip
	// //目前只写机器CALL-----
	//debug0119------------------------
	//判断翻前操作次数
	var number1 int64
	for p[0].ActionChip != p[1].ActionChip {
		number1 = number1 + 1
		fmt.Println("开始对抗")
		time.Sleep(1 * time.Second)
		PotChip = p[0].PlayPerflopAction(number1, PotChip, p[1])
		if number1 != 1 && p[0].ActionChip == p[1].ActionChip {
			return PotChip, nil
		}
		PotChip = p[1].PlayPerflopAction(number1, PotChip, p[0])
	}
	//debug0119------------------------
	return PotChip, nil
}

//FlopAction 翻牌后的一轮操作
func (p *PlayPlayer) FlopAction(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//TurnAction 转牌后的一轮操作
func (p *PlayPlayer) TurnAction(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//RiverAction 河牌后的一轮操作
func (p *PlayPlayer) RiverAction(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//PlayPerflopAction 玩家翻牌前的一轮操作
func (p *PlayPlayData) PlayPerflopAction(number1 int64, PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	//如果是第一轮次操作的大盲行为，他有权check
	// if p.BBCheckStatus == 0 && p.SiteNumber == 1 {
	if number1 == 1 && p.SiteNumber == 1 && p.ActionChip == PlayPlayer.ActionChip {
		if p.Chip < p.RaiseChipLeve {
			fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n fold\n allin")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			switch input.Text() {
			case "fold":
				PotChip = p.Fold(PotChip, PlayPlayer)
			case "allin":
				PotChip = p.Allin(PotChip, PlayPlayer)
			case "check":
				PotChip = p.Check(PotChip, PlayPlayer)
			default:
				fmt.Println("无效的指令，请重新输入：")
				p.PlayPerflopAction(number1, PotChip, PlayPlayer)
			}
			return PotChip
		}
		fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n raise\n fold\n allin\n check")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		switch input.Text() {
		case "raise":
			PotChip = p.Raise(number1, PotChip, PlayPlayer)
		case "fold":
			PotChip = p.Fold(PotChip, PlayPlayer)
		case "allin":
			PotChip = p.Allin(PotChip, PlayPlayer)
		case "check":
			PotChip = p.Check(PotChip, PlayPlayer)
		default:
			fmt.Println("无效的指令，请重新输入：")
			p.PlayPerflopAction(number1, PotChip, PlayPlayer)
		}
		return PotChip
	}
	//如果筹码不够了
	if p.Chip < p.RaiseChipLeve {
		fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n fold\n allin")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		switch input.Text() {
		case "fold":
			PotChip = p.Fold(PotChip, PlayPlayer)
		case "allin":
			PotChip = p.Allin(PotChip, PlayPlayer)
		default:
			fmt.Println("无效的指令，请重新输入：")
			p.PlayPerflopAction(number1, PotChip, PlayPlayer)
		}
		return PotChip
	}
	fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n call\n raise\n fold\n allin")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	switch input.Text() {
	case "call":
		PotChip = p.Call(PotChip, PlayPlayer)
	case "raise":
		PotChip = p.Raise(number1, PotChip, PlayPlayer)
	case "fold":
		PotChip = p.Fold(PotChip, PlayPlayer)
	case "allin":
		PotChip = p.Allin(PotChip, PlayPlayer)
	default:
		fmt.Println("无效的指令，请重新输入：")
		p.PlayPerflopAction(number1, PotChip, PlayPlayer)
	}
	return PotChip

}

//Call 跟注
func (p *PlayPlayData) Call(PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	//debug0120
	fmt.Println(p.Name, p.Chip)
	fmt.Println("1", PotChip)
	fmt.Println("1", PlayPlayer.ActionChip)
	fmt.Println("1", p.ActionChip)
	//debug0120

	p.Chip = p.Chip + p.ActionChip - PlayPlayer.ActionChip
	p.ActionChip = PlayPlayer.ActionChip

	PotChip = p.ActionChip + PlayPlayer.ActionChip

	//debug0120
	fmt.Println(p.Name, p.Chip)
	fmt.Println("dichi", PotChip)
	//debug0120

	return PotChip
}

//Raise 加注
func (p *PlayPlayData) Raise(number1 int64, PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {

	fmt.Println("请输入要加注的金额：")
	//获取输入
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()
	raiseNumber, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println("请认真输入金额，只有数字！！！")
		p.Raise(number1, PotChip, PlayPlayer)
	}
	//和对位的差价
	differenceNumber := PlayPlayer.ActionChip - p.ActionChip
	//得到最大能加注的值
	canBeRaiseNumber := p.Chip - differenceNumber
	if raiseNumber > canBeRaiseNumber {
		fmt.Println("最多只能加注：", canBeRaiseNumber)
		p.Raise(number1, PotChip, PlayPlayer)
	}
	if raiseNumber < p.RaiseChipLeve {
		fmt.Println("最小加注：", p.RaiseChipLeve)
		p.Raise(number1, PotChip, PlayPlayer)
	}
	p.Chip = p.Chip - raiseNumber - differenceNumber
	if p.Chip == 0 {
		p.AllinStatus = 1
	}
	p.ActionChip = p.ActionChip + differenceNumber + raiseNumber
	PotChip = p.ActionChip + PlayPlayer.ActionChip

	//debug0120
	fmt.Println(p.Name, p.Chip)
	fmt.Println(p.Name, p.ActionChip)
	fmt.Println("dichi", PotChip)
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
	if p.BBCheckStatus == 0 && p.SiteNumber == 1 && p.ActionTurnNumber == 0 {
		p.Chip = p.Chip - p.ActionChip
		PotChip = p.ActionChip + PlayPlayer.ActionChip
		return PotChip
	}
	return PotChip
}

//Fold 弃牌
func (p *PlayPlayData) Fold(PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	return PotChip
}
