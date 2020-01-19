package main

import (
	"math/rand"
	"time"
)

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

//ShuffleCard 洗牌
func ShuffleCard() (New52CardList [52]Card) {
	//初始化52张牌
	var Card52 = [52]Card{
		{Suit: "黑桃", Rank: 14},
		{Suit: "黑桃", Rank: 2},
		{Suit: "黑桃", Rank: 3},
		{Suit: "黑桃", Rank: 4},
		{Suit: "黑桃", Rank: 5},
		{Suit: "黑桃", Rank: 6},
		{Suit: "黑桃", Rank: 7},
		{Suit: "黑桃", Rank: 8},
		{Suit: "黑桃", Rank: 9},
		{Suit: "黑桃", Rank: 10},
		{Suit: "黑桃", Rank: 11},
		{Suit: "黑桃", Rank: 12},
		{Suit: "黑桃", Rank: 13},
		{Suit: "红桃", Rank: 14},
		{Suit: "红桃", Rank: 2},
		{Suit: "红桃", Rank: 3},
		{Suit: "红桃", Rank: 4},
		{Suit: "红桃", Rank: 5},
		{Suit: "红桃", Rank: 6},
		{Suit: "红桃", Rank: 7},
		{Suit: "红桃", Rank: 8},
		{Suit: "红桃", Rank: 9},
		{Suit: "红桃", Rank: 10},
		{Suit: "红桃", Rank: 11},
		{Suit: "红桃", Rank: 12},
		{Suit: "红桃", Rank: 13},
		{Suit: "梅花", Rank: 14},
		{Suit: "梅花", Rank: 2},
		{Suit: "梅花", Rank: 3},
		{Suit: "梅花", Rank: 4},
		{Suit: "梅花", Rank: 5},
		{Suit: "梅花", Rank: 6},
		{Suit: "梅花", Rank: 7},
		{Suit: "梅花", Rank: 8},
		{Suit: "梅花", Rank: 9},
		{Suit: "梅花", Rank: 10},
		{Suit: "梅花", Rank: 11},
		{Suit: "梅花", Rank: 12},
		{Suit: "梅花", Rank: 13},
		{Suit: "方片", Rank: 14},
		{Suit: "方片", Rank: 2},
		{Suit: "方片", Rank: 3},
		{Suit: "方片", Rank: 4},
		{Suit: "方片", Rank: 5},
		{Suit: "方片", Rank: 6},
		{Suit: "方片", Rank: 7},
		{Suit: "方片", Rank: 8},
		{Suit: "方片", Rank: 9},
		{Suit: "方片", Rank: 10},
		{Suit: "方片", Rank: 11},
		{Suit: "方片", Rank: 12},
		{Suit: "方片", Rank: 13},
	}
	//洗牌
	var new52 [52]Card
	b := 0
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(Card52)) {
		val := Card52[i]
		// fmt.Println(val)
		// fmt.Println(i)
		new52[b] = val
		b = b + 1
	}
	//fmt.Println(new52)
	return new52
}
