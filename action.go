package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//PerflopAction 翻牌前操作
func (p *PlayPlayer) PerflopAction(PotChip int64) (potChip int64) {
	fmt.Println("场上选手剩余筹码：", p[0].Name, ":", p[0].Chip, "\t", p[1].Name, ":", p[1].Chip, "\t", "底池:", PotChip)
	//--debug
	fmt.Println("小盲是：", p[0].Name)
	fmt.Println("大盲是：", p[1].Name)

	//判断翻前操作次数
	var number1 int64
	number1 = 0
	for p[0].ActionChip != p[1].ActionChip {
		number1 = number1 + 1
		fmt.Println("开始对抗")
		time.Sleep(1 * time.Second)
		PotChip = p[0].PlayPerflopAction(number1, PotChip, p[1])
		fmt.Println(PotChip, "debug0121")
		if number1 != 1 && p[0].ActionChip == p[1].ActionChip {
			//debug0121
			fmt.Println(PotChip)
			//debug0121
			//归零
			p[0].ActionChip = 0
			p[1].ActionChip = 0
			p[0].ActionTurnNumber = 1
			p[1].ActionTurnNumber = 1
			//展示下一轮的牌
			//debug0121---
			fmt.Println(p[0].Name, p[0].Chip)
			fmt.Println(p[1].Name, p[1].Chip)
			//debug0121---
			PotChip = p.FlopAction(PotChip)
			return PotChip
		}
		fmt.Println(PotChip, "debug0121")
		PotChip = p[1].PlayPerflopAction(number1, PotChip, p[0])
	}
	//debug0121
	fmt.Println(PotChip)
	//debug0121
	//归零
	p[0].ActionChip = 0
	p[1].ActionChip = 0
	p[0].ActionTurnNumber = 1
	p[1].ActionTurnNumber = 1
	//展示下一轮的牌
	fmt.Println("FLOP:", p[0].CardInfo.Value7[2].CardTranslate(), p[0].CardInfo.Value7[3].CardTranslate(), p[0].CardInfo.Value7[4].CardTranslate())
	//debug0121---
	fmt.Println(p[0].Name, p[0].Chip)
	fmt.Println(p[1].Name, p[1].Chip)
	fmt.Println("AAAAAAAAAAAAAAAAAAA")
	//debug0121---
	PotChip = p.FlopAction(PotChip)
	return PotChip
}

//FlopAction 翻牌后的一轮操作
func (p *PlayPlayer) FlopAction(PotChip int64) (potChip int64) {
	var number1 int64
	number1 = 0
	PotChip = p[1].PlayflopAction(number1, PotChip, p[0])
	//无意义就位消费掉number1？
	if number1 != 0 && p[0].ActionChip == p[1].ActionChip {
		return PotChip
	}
	PotChip = p[0].PlayflopAction(number1, PotChip, p[1])
	for p[0].ActionChip != p[1].ActionChip {
		PotChip = p[1].PlayflopAction(number1, PotChip, p[0])
		if p[0].ActionChip == p[1].ActionChip {
			//debug0121
			fmt.Println(PotChip)
			//debug0121
			//归零
			p[0].ActionChip = 0
			p[1].ActionChip = 0
			p[0].ActionTurnNumber = 2
			p[1].ActionTurnNumber = 2
			//展示下一轮的牌
			fmt.Println("TURN:", p[0].CardInfo.Value7[5].CardTranslate())
			//debug0121---
			fmt.Println(p[0].Name, p[0].Chip)
			fmt.Println(p[1].Name, p[1].Chip)
			//debug0121---
			PotChip = p.TurnAction(PotChip)
			return PotChip
		}
		PotChip = p[0].PlayflopAction(number1, PotChip, p[1])
	}
	//debug0121
	fmt.Println(PotChip)
	//debug0121
	//归零
	p[0].ActionChip = 0
	p[1].ActionChip = 0
	p[0].ActionTurnNumber = 2
	p[1].ActionTurnNumber = 2
	//展示下一轮的牌
	fmt.Println("TURN:", p[0].CardInfo.Value7[5].CardTranslate())
	//debug0121---
	fmt.Println(p[0].Name, p[0].Chip)
	fmt.Println(p[1].Name, p[1].Chip)
	//debug0121---
	PotChip = p.TurnAction(PotChip)
	return PotChip
}

//TurnAction 转牌后的一轮操作
func (p *PlayPlayer) TurnAction(PotChip int64) (potChip int64) {
	var number1 int64
	number1 = 0
	PotChip = p[1].PlayflopAction(number1, PotChip, p[0])
	//无意义就位消费掉number1？
	if number1 != 0 && p[0].ActionChip == p[1].ActionChip {
		return PotChip
	}
	PotChip = p[0].PlayflopAction(number1, PotChip, p[1])
	for p[0].ActionChip != p[1].ActionChip {
		PotChip = p[1].PlayflopAction(number1, PotChip, p[0])
		if p[0].ActionChip == p[1].ActionChip {
			//debug0121
			fmt.Println(PotChip)
			//debug0121
			//归零
			p[0].ActionChip = 0
			p[1].ActionChip = 0
			p[0].ActionTurnNumber = 3
			p[1].ActionTurnNumber = 3
			//展示下一轮的牌
			fmt.Println("RIVER:", p[0].CardInfo.Value7[6].CardTranslate())
			//debug0121---
			fmt.Println(p[0].Name, p[0].Chip)
			fmt.Println(p[1].Name, p[1].Chip)
			//debug0121---
			PotChip = p.RiverAction(PotChip)
			return PotChip
		}
		PotChip = p[0].PlayflopAction(number1, PotChip, p[1])
	}
	//debug0121
	fmt.Println(PotChip)
	//debug0121
	//归零
	p[0].ActionChip = 0
	p[1].ActionChip = 0
	p[0].ActionTurnNumber = 3
	p[1].ActionTurnNumber = 3
	//展示下一轮的牌
	fmt.Println("RIVER:", p[0].CardInfo.Value7[6].CardTranslate())
	//debug0121---
	fmt.Println(p[0].Name, p[0].Chip)
	fmt.Println(p[1].Name, p[1].Chip)
	//debug0121---
	PotChip = p.RiverAction(PotChip)
	return PotChip
}

//RiverAction 河牌后的一轮操作
func (p *PlayPlayer) RiverAction(PotChip int64) (potChip int64) {
	var number1 int64
	number1 = 0
	PotChip = p[1].PlayflopAction(number1, PotChip, p[0])
	if number1 != 0 && p[0].ActionChip == p[1].ActionChip {
		return PotChip
	}
	PotChip = p[0].PlayflopAction(number1, PotChip, p[1])
	for p[0].ActionChip != p[1].ActionChip {
		PotChip = p[1].PlayflopAction(number1, PotChip, p[0])
		if p[0].ActionChip == p[1].ActionChip {
			//debug0121---
			fmt.Println(p[0].Name, p[0].Chip)
			fmt.Println(p[1].Name, p[1].Chip)
			//debug0121---
			return PotChip
		}
		PotChip = p[0].PlayflopAction(number1, PotChip, p[1])
	}
	//debug0121
	fmt.Println(PotChip)
	//debug0121---
	fmt.Println(p[0].Name, p[0].Chip)
	fmt.Println(p[1].Name, p[1].Chip)
	//debug0121---
	//debug0121
	return PotChip
}

//PlayflopAction 玩家翻牌前的一轮操作
func (p *PlayPlayData) PlayflopAction(number1 int64, PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	//正常对手还没行动 或者 行动为check的情况下
	if PlayPlayer.ActionChip == 0 {
		//筹码不够支付一个大盲
		if p.Chip < p.RaiseChipLeve {
			fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n allin\n check")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			switch input.Text() {
			case "allin":
				PotChip = p.Allin(PotChip, PlayPlayer)
				return PotChip
			case "check":
				PotChip = p.Check(PotChip, PlayPlayer)
				return PotChip
			default:
				fmt.Println("无效的指令，请重新输入：")
				PotChip = p.PlayflopAction(number1, PotChip, PlayPlayer)
				return PotChip
			}
			return PotChip
		}
		//正常对手还没行动 或者 行动为check的情况下
		fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n raise\n allin\n check")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		switch input.Text() {
		case "raise":
			PotChip = p.Raise(number1, PotChip, PlayPlayer)
			return PotChip
		case "allin":
			PotChip = p.Allin(PotChip, PlayPlayer)
			return PotChip
		case "check":
			PotChip = p.Check(PotChip, PlayPlayer)
			return PotChip
		default:
			fmt.Println("无效的指令，请重新输入：")
			PotChip = p.PlayflopAction(number1, PotChip, PlayPlayer)
			return PotChip
		}
		return PotChip
	}
	//对手已经行动过且不是check
	//如果筹码不够了
	if p.Chip < p.RaiseChipLeve {
		fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n fold\n allin")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		switch input.Text() {
		case "fold":
			PotChip = p.Fold(PotChip, PlayPlayer)
			return PotChip
		case "allin":
			PotChip = p.Allin(PotChip, PlayPlayer)
			return PotChip
		default:
			fmt.Println("无效的指令，请重新输入：")
			PotChip = p.PlayflopAction(number1, PotChip, PlayPlayer)
			return PotChip
		}
		return PotChip
	}
	fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n call\n raise\n fold\n allin")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	switch input.Text() {
	case "call":
		PotChip = p.Call(PotChip, PlayPlayer)
		return PotChip
	case "raise":
		PotChip = p.Raise(number1, PotChip, PlayPlayer)
		return PotChip
	case "fold":
		PotChip = p.Fold(PotChip, PlayPlayer)
		return PotChip
	case "allin":
		PotChip = p.Allin(PotChip, PlayPlayer)
		return PotChip
	default:
		fmt.Println("无效的指令，请重新输入：")
		PotChip = p.PlayflopAction(number1, PotChip, PlayPlayer)
		return PotChip
	}
	return PotChip
}

//PlayPerflopAction 玩家翻牌前的一轮操作
func (p *PlayPlayData) PlayPerflopAction(number1 int64, PotChip int64, PlayPlayer PlayPlayData) (potChip int64) {
	//如果是第一轮次操作的大盲行为，他有权check
	// if p.BBCheckStatus == 0 && p.SiteNumber == 1 {
	if number1 == 1 && p.SiteNumber == 1 && p.ActionChip == PlayPlayer.ActionChip {
		//如果筹码不够了
		if p.Chip < p.RaiseChipLeve {
			fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n allin\n check")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			switch input.Text() {
			case "allin":
				PotChip = p.Allin(PotChip, PlayPlayer)
				return PotChip
			case "check":
				PotChip = p.Check(PotChip, PlayPlayer)
				return PotChip
			default:
				fmt.Println("无效的指令，请重新输入：")
				PotChip = p.PlayPerflopAction(number1, PotChip, PlayPlayer)
				return PotChip
			}
			return PotChip
		}
		fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n raise\n allin\n check")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		switch input.Text() {
		case "raise":
			PotChip = p.Raise(number1, PotChip, PlayPlayer)
			return PotChip
		case "allin":
			PotChip = p.Allin(PotChip, PlayPlayer)
			return PotChip
		case "check":
			PotChip = p.Check(PotChip, PlayPlayer)
			return PotChip
		default:
			fmt.Println("无效的指令，请重新输入：")
			PotChip = p.PlayPerflopAction(number1, PotChip, PlayPlayer)
			return PotChip
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
			return PotChip
		case "allin":
			PotChip = p.Allin(PotChip, PlayPlayer)
			return PotChip
		default:
			fmt.Println("无效的指令，请重新输入：")
			PotChip = p.PlayPerflopAction(number1, PotChip, PlayPlayer)
			return PotChip
		}
		return PotChip
	}
	fmt.Println("玩家：", p.Name, "-", "你可以进行以下操作\n call\n raise\n fold\n allin")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	switch input.Text() {
	case "call":
		PotChip = p.Call(PotChip, PlayPlayer)
		return PotChip
	case "raise":
		PotChip = p.Raise(number1, PotChip, PlayPlayer)
		return PotChip
	case "fold":
		PotChip = p.Fold(PotChip, PlayPlayer)
		return PotChip
	case "allin":
		PotChip = p.Allin(PotChip, PlayPlayer)
		return PotChip
	default:
		fmt.Println("无效的指令，请重新输入：")
		PotChip = p.PlayPerflopAction(number1, PotChip, PlayPlayer)
		return PotChip
	}
	return PotChip

}
