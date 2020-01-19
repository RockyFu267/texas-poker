package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//PerflopAction 翻牌前操作
func (p *PlayPlayer) PerflopAction(PotChip int64) (potChip int64, error error) {
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
	fmt.Println(p[0].ActionChip)
	fmt.Println(p[1].ActionChip)
	fmt.Println("----------------------------------")
	p[0].ActionChip = 100
	p[1].ActionChip = 200
	for p[0].ActionChip != p[1].ActionChip {
		fmt.Println("zenmebana ")
		time.Sleep(1 * time.Second)
		p[0].PlayPerflopAction(PotChip, p[1])
		p[1].PlayPerflopAction(PotChip, p[0])
	}
	//debug0119------------------------
	fmt.Println("场上选手剩余筹码：", p[0].Name, ":", p[0].Chip, "\t", p[1].Name, ":", p[1].Chip, "\t", "底池:", PotChip)
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
func (p *PlayPlayData) PlayPerflopAction(PotChip int64, PlayPlayer PlayPlayData) (potChip int64, error error) {
	if p.BBCheckStatus == 0 && p.SiteNumber == 1 {
		fmt.Println("你可以进行以下操作\n call\n raise\n fold\n allin\n check")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		switch input.Text() {
		case "call":
			p.Call(PotChip, PlayPlayer)
		case "raise":
			p.Raise(PotChip, PlayPlayer)
		case "fold":
			p.Fold(PotChip, PlayPlayer)
		case "allin":
			p.Allin(PotChip, PlayPlayer)
		case "check":
			p.Check(PotChip, PlayPlayer)
		default:
			fmt.Println("无效的指令，请重新输入：")
			p.PlayPerflopAction(PotChip, PlayPlayer)
		}
		return PotChip, nil
	}
	fmt.Println("你可以进行以下操作\n call\n raise\n fold\n allin")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	switch input.Text() {
	case "call":
		p.Call(PotChip, PlayPlayer)
	case "raise":
		p.Raise(PotChip, PlayPlayer)
	case "fold":
		p.Fold(PotChip, PlayPlayer)
	case "allin":
		p.Allin(PotChip, PlayPlayer)
	default:
		fmt.Println("无效的指令，请重新输入：")
		p.PlayPerflopAction(PotChip, PlayPlayer)
	}
	return PotChip, nil

}

//Call 跟注
func (p *PlayPlayData) Call(PotChip int64, PlayPlayer PlayPlayData) (potChip int64, error error) {
	return PotChip, nil
}

//Raise 加注
func (p *PlayPlayData) Raise(PotChip int64, PlayPlayer PlayPlayData) (potChip int64, error error) {
	return PotChip, nil
}

//Allin 全下 无条件可以看五张 最后开牌
func (p *PlayPlayData) Allin(PotChip int64, PlayPlayer PlayPlayData) (potChip int64, error error) {
	return PotChip, nil
}

//Check 过牌
func (p *PlayPlayData) Check(PotChip int64, PlayPlayer PlayPlayData) (potChip int64, error error) {
	//wait 其他玩家行动
	//next step

	return PotChip, nil
}

//Fold 弃牌
func (p *PlayPlayData) Fold(PotChip int64, PlayPlayer PlayPlayData) (potChip int64, error error) {
	return PotChip, nil
}
