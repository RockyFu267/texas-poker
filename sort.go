package main

import "fmt"

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
		p[0].Chip = p[0].SumChip/2 + p[0].Chip
		p[1].Chip = p[1].SumChip/2 + p[1].Chip
	} else {
		if MaxSSSList.PlayName == p[0].Name {
			p[0].Chip = p[0].SumChip + p[0].Chip
			fmt.Println("玩家：", p[0].Name, "获得胜利")
		} else {
			p[1].Chip = p[1].SumChip + p[1].Chip
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
