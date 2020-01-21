package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

//Card 牌的结构体
type Card struct {
	Suit string `json:"suit"`
	Rank int64  `json:"rank"`
}

//CardRank 翻拍后玩家的最大牌行同级数组
type CardRank struct {
	Grade    int64   `json:"grade"`
	Value    [5]Card `json:"value"`
	PlayName string  `json:"playName"`
	Value7   [7]Card `json:"value7"`
}

//PlayPlayer 指针测试用
type PlayPlayer [2]PlayPlayData

//PlayPlayData 玩家的信息、还不知道怎么用
type PlayPlayData struct {
	Name string `json:"name"`
	Chip int64  `json:"chip"`
	//不包含anti的单轮次行动
	ActionChip int64 `json:"actionChip"`
	//包含anti的多轮次下注总和
	PotChip int64 `json:"potChip"`
	//anti
	AntiChip int64 `json:"antiChip"`
	//大盲小盲
	BBSBChip int64 `json:"bbsbChip"`
	//起手两张牌
	HandsCard [2]Card `json:"handsCard"`
	//将来的牌桌
	TableNumber int64 `json:"tableNumber"`
	//将来的赛事编号
	GameNumber int64 `json:"gameNumber"`
	//本桌的位置
	SiteNumber int64 `json:"siteNumber"`
	//开牌比牌用的信息
	CardInfo CardRank `json:"cardInfo"`
	//需要跟注的筹码量
	NeedCallNumber int64 `json:"needcallnumber"`
	//是否可以被跟注后过牌的状态码--在大盲位的时候特殊
	BBCheckStatus int64 `json:"bbcheckstatus"`
	//所在牌局的四个轮次
	ActionTurnNumber int64 `json:"actionturnnumber"`
	//桌面总筹码
	SumChip int64 `json:"sumchip"`
	//Allin 状态
	AllinStatus int64 `json:"allinstatus"`
	//最低级别的加注单位
	RaiseChipLeve int64 `json:"raisechipleve"`
}

func main() {
	// counts := make(map[string]int)
	fmt.Println("请输入你的用户名：")
	Input := bufio.NewScanner(os.Stdin)
	Input.Scan()
	PlayName := Input.Text()
	fmt.Println("你的用户名是：", PlayName)
	// fmt.Println("准备开始比赛，输入ready开始发牌：")
	// input.Scan()
	// if input.Text() != "ready" {
	//  fmt.Println("fuck")
	//  return
	// }

	//初始化AI和玩家的剩余筹码&pot值
	var PotChip int64
	//如果出现分池
	//var PotChipMulti []int64  --还没开始写
	//c初始化加注量
	//var RaiseNumber int64
	//初始化玩家数组
	var play1 PlayPlayData
	var play2 PlayPlayData
	//台桌1号位
	play1.Name = PlayName
	play1.Chip = 1000
	play1.BBSBChip = 0
	play1.AntiChip = 0
	play1.PotChip = 0
	play1.ActionChip = 0

	//台桌2号位
	play2.Name = "AI"
	play2.Chip = 1000
	play2.BBSBChip = 0
	play2.AntiChip = 0
	play2.PotChip = 0
	play2.ActionChip = 0
	//这个应该是个切片 不一定两人 现在是1v1

	//--
	var playPlayer PlayPlayer
	playPlayer[0] = play1
	playPlayer[1] = play2
	// var playPlayer = [2]PlayPlayData{}
	// playPlayer[0] = play1
	// playPlayer[1] = play2
	//--

	//初始化potchip
	PotChip = 0
	//初始化被跟注后过牌的状态码
	playPlayer[0].BBCheckStatus = 0
	playPlayer[1].BBCheckStatus = 0
	playPlayer[0].ActionTurnNumber = 0
	playPlayer[1].ActionTurnNumber = 0
	playPlayer[0].AllinStatus = 0
	playPlayer[1].AllinStatus = 0
	playPlayer[0].RaiseChipLeve = 200
	playPlayer[1].RaiseChipLeve = 200
	//var playPlayer = [2]string{PlayName, "AI"}
	//fmt.Println("开始发牌")
	StartGame(playPlayer, PotChip)
}

//StartGame 牌局开始 直到最后的出现唯一的胜利者 结束牌局--1v1的前提下
func StartGame(p PlayPlayer, PotChip int64) (error error) {
	p[0].SiteNumber = 0
	p[1].SiteNumber = 1
	p[0].BBSBChip = 100
	p[1].BBSBChip = 200
	p[0].ActionChip = 100
	p[1].ActionChip = 200
	p[0].Chip = p[0].Chip - p[0].ActionChip
	p[1].Chip = p[1].Chip - p[1].ActionChip
	fmt.Println(p[0].Chip, p[1].Chip)
	var play PlayPlayer
	play[0] = p[0]
	play[1] = p[1]
	var gameNumber int64
	gameNumber = 0
	StartOneGame(play, PotChip, gameNumber)
	return nil
}

//StartOneGame 开始一局
func StartOneGame(p PlayPlayer, PotChip int64, gameNumber int64) (error error) {
	//判断是否有结果了
	if p[0].Chip == 0 {
		fmt.Println("winner is:", p[1].Name)
		return nil
	}
	if p[1].Chip == 0 {
		fmt.Println("winner is:", p[0].Name)
		return nil
	}
	gameNumber = gameNumber + 1
	fmt.Println("------------------第", gameNumber, "局--------------------------------")
	//洗牌-每一局开始必须操作
	new52 := ShuffleCard()
	// //debug-指定特殊牌型
	// new52[0].Suit = "红桃"
	// new52[0].Rank = 3
	// new52[1].Suit = "梅花"
	// new52[1].Rank = 10
	// new52[2].Suit = "黑桃"
	// new52[2].Rank = 11
	// new52[3].Suit = "梅花"
	// new52[3].Rank = 2
	// new52[4].Suit = "梅花"
	// new52[4].Rank = 13
	// new52[5].Suit = "黑桃"
	// new52[5].Rank = 14
	// new52[6].Suit = "梅花"
	// new52[6].Rank = 12
	// new52[7].Suit = "梅花"
	// new52[7].Rank = 8
	// new52[8].Suit = "方片"
	// new52[8].Rank = 5
	// //debug指定特殊牌型
	//--debug
	//fmt.Println(new52)
	//--
	p[0].GetCard(new52)
	p[1].GetCard(new52)
	//debug
	fmt.Println(p[0].HandsCard[0].CardTranslate(), p[0].HandsCard[1].CardTranslate())
	fmt.Println(p[1].HandsCard[0].CardTranslate(), p[1].HandsCard[1].CardTranslate())
	//--debug
	//fmt.Println(p[0].Name, p[0].CardInfo.Value7[0].CardTranslate(), p[0].CardInfo.Value7[1].CardTranslate(), p[0].CardInfo.Value7[2].CardTranslate(), p[0].CardInfo.Value7[3].CardTranslate(), p[0].CardInfo.Value7[4].CardTranslate(), p[0].CardInfo.Value7[5].CardTranslate(), p[0].CardInfo.Value7[6].CardTranslate())
	//fmt.Println(p[1].Name, p[1].CardInfo.Value7[0].CardTranslate(), p[1].CardInfo.Value7[1].CardTranslate(), p[1].CardInfo.Value7[2].CardTranslate(), p[1].CardInfo.Value7[3].CardTranslate(), p[1].CardInfo.Value7[4].CardTranslate(), p[1].CardInfo.Value7[5].CardTranslate(), p[1].CardInfo.Value7[6].CardTranslate())
	//--
	//翻前操作
	//目前只写机器CALL

	a, err := p.PerflopAction(PotChip)
	if err != nil {
		return err
	}
	//flop三张
	fmt.Println("Flop:", new52[4].CardTranslate(), new52[5].CardTranslate(), new52[6].CardTranslate())
	//flop操作
	b, err := p.FlopAction(a)
	if err != nil {
		return err
	}
	//trun操作
	c, err := p.TurnAction(b)
	if err != nil {
		return err
	}
	//turn发一张
	fmt.Println("Turn:", new52[7].CardTranslate())
	//river操作
	d, err := p.RiverAction(c)
	if err != nil {
		return err
	}
	fmt.Println("River:", new52[8].CardTranslate())

	//showhand 比大小
	p.ShowHandSort(d)

	//每局结束先换sitenumber 关系到发牌顺序
	var tmpSiteNum int64
	tmpSiteNum = p[0].SiteNumber
	p[0].SiteNumber = p[1].SiteNumber
	p[1].SiteNumber = tmpSiteNum
	//每局结束后初始化 交换大小盲
	var tmpSite PlayPlayData
	tmpSite = p[0]
	p[0] = p[1]
	p[1] = tmpSite

	//每局结束后初始化 交换大小盲
	var tmpBBSBChip int64
	tmpBBSBChip = p[0].BBSBChip
	p[0].BBSBChip = p[1].BBSBChip
	p[1].BBSBChip = tmpBBSBChip
	//PotChip 归零
	p[0].BBCheckStatus = 0
	p[1].BBCheckStatus = 0
	p[0].ActionTurnNumber = 0
	p[1].ActionTurnNumber = 0
	p[0].SumChip = 0
	p[1].SumChip = 0
	p[0].AllinStatus = 0
	p[1].AllinStatus = 0
	//大盲动作不归0 初始化
	p[0].ActionChip = 100
	p[1].ActionChip = 200
	PotChip = 0
	time.Sleep(1 * time.Second)
	StartOneGame(p, PotChip, gameNumber)

	return nil
}
