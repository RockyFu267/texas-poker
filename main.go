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
	gameNumber = gameNumber + 1
	fmt.Println("------------------第", gameNumber, "局--------------------------------")
	//洗牌-每一局开始必须操作
	new52 := ShuffleCard()
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
	// fmt.Println(p[0].SiteNumber)
	// fmt.Println(p[1].SiteNumber)
	//----
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

	//作弊
	// if p[1].Name == "AI" {
	//  var Card2 = [2]Card{
	//      {Suit: "黑桃", Rank: 14},
	//      {Suit: "黑桃", Rank: 13},
	//  }
	//  p[1].HandsCard = Card2
	//  p[1].CardInfo.Value7[0] = Card{Suit: "黑桃", Rank: 14}
	//  p[1].CardInfo.Value7[1] = Card{Suit: "黑桃", Rank: 13}
	// }
	// if p[0].Name == "AI" {
	//  var Card2 = [2]Card{
	//      {Suit: "黑桃", Rank: 14},
	//      {Suit: "黑桃", Rank: 13},
	//  }
	//  p[0].HandsCard = Card2
	//  p[0].CardInfo.Value7[0] = Card{Suit: "黑桃", Rank: 14}
	//  p[0].CardInfo.Value7[1] = Card{Suit: "黑桃", Rank: 13}
	// }

	//showhand 比大小
	p.ShowHandSort(d)
	// //假设每次AI赢
	// if p[1].Name == "AI" {
	//  p[1].Chip = p[1].Chip + d
	// } else {
	//  p[0].Chip = p[0].Chip + d
	// }
	//判断是否有结果了
	if p[0].Chip == 0 {
		fmt.Println("winner is:", p[1].Name)
		return nil
	}
	if p[1].Chip == 0 {
		fmt.Println("winner is:", p[0].Name)
		return nil
	}

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
	PotChip = 0
	time.Sleep(1 * time.Second)
	StartOneGame(p, PotChip, gameNumber)

	return nil
}

//ShowHandSort 开牌
func (p *PlayPlayer) ShowHandSort(PotChip int64) (error error) {
	//前置位玩家的最大牌值 CardRank要修改 7选5也要改
	CardRank0 := Judge5From7(p[0].CardInfo.Value7, p[0].Name)
	CardRank1 := Judge5From7(p[1].CardInfo.Value7, p[1].Name)
	SortCardMaxList0 := SortCardMaxList(CardRank0)
	SortCardMaxList1 := SortCardMaxList(CardRank1)
	p[0].CardInfo = SortCardMaxList0
	p[1].CardInfo = SortCardMaxList1

	//放进一个切片比大小
	var MaxCardRankSlice []CardRank
	MaxCardRankSlice = append(MaxCardRankSlice, SortCardMaxList0)
	MaxCardRankSlice = append(MaxCardRankSlice, SortCardMaxList1)
	//调用最后的比大小
	MaxSSSList, MaxSSSGetMoney := SortCardMaxListEND(MaxCardRankSlice)
	//debug
	// fmt.Println(CardRank0)
	// fmt.Println(CardRank1)
	fmt.Println("双方最大的牌是:")
	fmt.Println(SortCardMaxList0.PlayName, SortCardMaxList0.Value[0].CardTranslate(), SortCardMaxList0.Value[1].CardTranslate(), SortCardMaxList0.Value[2].CardTranslate(), SortCardMaxList0.Value[3].CardTranslate(), SortCardMaxList0.Value[4].CardTranslate())
	fmt.Println(SortCardMaxList1.PlayName, SortCardMaxList1.Value[0].CardTranslate(), SortCardMaxList1.Value[1].CardTranslate(), SortCardMaxList1.Value[2].CardTranslate(), SortCardMaxList1.Value[3].CardTranslate(), SortCardMaxList1.Value[4].CardTranslate())
	// fmt.Println(MaxSSSList)
	// fmt.Println(MaxSSSGetMoney)
	// fmt.Println(len(MaxSSSGetMoney))
	//--

	//如果大于1 就会出现分钱局面
	if len(MaxSSSGetMoney) > 2 {
		//目前只有两个人所以
		fmt.Println("两人分钱")
		p[0].Chip = PotChip/2 + p[0].Chip
		p[1].Chip = PotChip/2 + p[1].Chip
	} else {
		if MaxSSSList.PlayName == p[0].Name {
			p[0].Chip = PotChip + p[0].Chip
			fmt.Println("玩家：", p[0].Name, "获得胜利")
		} else {
			p[1].Chip = PotChip + p[1].Chip
			fmt.Println("玩家：", p[1].Name, "获得胜利")
		}
	}
	return nil
}

//SortCardMaxList 比较牌型的最大值
func SortCardMaxList(CardRankSlice []CardRank) (CardMaxList CardRank) {
	CardMaxList = CardRankSlice[0]
	for i := 0; i < len(CardRankSlice); i++ {
		if CardMaxList.Grade < CardRankSlice[i].Grade {
			CardMaxList = CardRankSlice[i]
		} else {
			if CardMaxList.Grade == CardRankSlice[i].Grade {
				for j := 0; j < 5; j++ {
					if CardMaxList.Value[j].Rank < CardRankSlice[i].Value[j].Rank {
						CardMaxList = CardRankSlice[i]
						break
					} else {
						if CardMaxList.Value[j].Rank > CardRankSlice[i].Value[j].Rank {
							break
						}
					}
				}
			}
		}
	}
	return CardMaxList
}

//SortCardMaxListEND 两位玩家比较牌型的最大值
func SortCardMaxListEND(CardRankSlice []CardRank) (CardMaxList CardRank, GetMoneyList []CardRank) {
	CardMaxList = CardRankSlice[0]
	for i := 0; i < len(CardRankSlice); i++ {
		if CardMaxList.Grade < CardRankSlice[i].Grade {
			CardMaxList = CardRankSlice[i]
		} else {
			if CardMaxList.Grade == CardRankSlice[i].Grade {
				for j := 0; j < 5; j++ {
					if CardMaxList.Value[j].Rank < CardRankSlice[i].Value[j].Rank {
						CardMaxList = CardRankSlice[i]
						break
					} else {
						if j == 4 {
							if CardMaxList.Value[j].Rank == CardRankSlice[i].Value[j].Rank {
								GetMoneyList = append(GetMoneyList, CardMaxList)
								CardMaxList = CardRankSlice[i]
								GetMoneyList = append(GetMoneyList, CardMaxList)
								break
							}
						}
					}
				}
			}
		}
	}
	return CardMaxList, GetMoneyList
}

//GetCard 获取自己的牌
func (p *PlayPlayData) GetCard(New52CardList [52]Card) (error error) {
	var number int64
	number = 2
	var hunCard [2]Card
	//选手获取牌
	hunCard[0] = New52CardList[p.SiteNumber]
	hunCard[1] = New52CardList[p.SiteNumber+number]
	p.HandsCard = hunCard
	//选手5张公共牌提前注入
	p.CardInfo.Value7[0] = New52CardList[p.SiteNumber]
	p.CardInfo.Value7[1] = New52CardList[p.SiteNumber+number]
	p.CardInfo.Value7[2] = New52CardList[2*number]
	p.CardInfo.Value7[3] = New52CardList[2*number+1]
	p.CardInfo.Value7[4] = New52CardList[2*number+2]
	p.CardInfo.Value7[5] = New52CardList[2*number+3]
	p.CardInfo.Value7[6] = New52CardList[2*number+4]
	return nil
}
