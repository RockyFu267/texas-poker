package main

//Judge5From7 7选五的21种牌型的牌力
func Judge5From7(CardArry7_5 [7]Card, PlayName string) (CardRankSlice []CardRank) {
	var CardRanktest CardRank
	var ii, j, k, l, m int64
	var iii, jjj int64
	var MaxArr [5]Card
	var temp Card
	for ii = 0; ii <= 2; ii++ {
		for j = ii + 1; j <= 3; j++ {
			for k = j + 1; k <= 4; k++ {
				for l = k + 1; l <= 5; l++ {
					for m = l + 1; m <= 6; m++ {
						// fmt.Println(CardArry7_5[ii], CardArry7_5[j], CardArry7_5[k], CardArry7_5[l], CardArry7_5[m])
						MaxArr[0] = CardArry7_5[ii]
						MaxArr[1] = CardArry7_5[j]
						MaxArr[2] = CardArry7_5[k]
						MaxArr[3] = CardArry7_5[l]
						MaxArr[4] = CardArry7_5[m]
						for iii = 0; iii < 4; iii++ {
							for jjj = 0; jjj < 4; jjj++ {
								if MaxArr[jjj].Rank < MaxArr[jjj+1].Rank {
									temp = MaxArr[jjj+1]
									MaxArr[jjj+1] = MaxArr[jjj]
									MaxArr[jjj] = temp
								}
							}
						}
						// --debug//fmt.Println(MaxArr)
						//判断是否是同花
						if MaxArr[0].Suit == MaxArr[1].Suit &&
							MaxArr[0].Suit == MaxArr[2].Suit &&
							MaxArr[0].Suit == MaxArr[3].Suit &&
							MaxArr[0].Suit == MaxArr[4].Suit {
							//判断是否是同花顺
							if (MaxArr[0].Rank-MaxArr[1].Rank == 1 &&
								MaxArr[1].Rank-MaxArr[2].Rank == 1 &&
								MaxArr[2].Rank-MaxArr[3].Rank == 1 &&
								MaxArr[3].Rank-MaxArr[4].Rank == 1) ||
								(MaxArr[0].Rank-MaxArr[1].Rank == 9 &&
									MaxArr[1].Rank-MaxArr[2].Rank == 1 &&
									MaxArr[2].Rank-MaxArr[3].Rank == 1 &&
									MaxArr[3].Rank-MaxArr[4].Rank == 1) {
								//如果是同花顺
								CardRanktest.Grade = 9
								CardRanktest.Value = MaxArr
								CardRanktest.PlayName = PlayName
								CardRankSlice = append(CardRankSlice, CardRanktest)
								// --debug//fmt.Println("straight flush")
							} else {
								//如果是同花
								CardRanktest.Grade = 6
								CardRanktest.Value = MaxArr
								CardRanktest.PlayName = PlayName
								CardRankSlice = append(CardRankSlice, CardRanktest)
								// --debug//fmt.Println("flush")
							}
						} else {
							//判断是否是金刚
							if (MaxArr[0].Rank == MaxArr[3].Rank) ||
								(MaxArr[1].Rank == MaxArr[4].Rank) {
								//如果是金刚
								//判断重组金刚数组的顺序 保证前四个为相等的值，最后的值为单值
								if MaxArr[1].Rank == MaxArr[4].Rank {
									CardRanktest.Grade = 8
									CardRanktest.Value[4] = MaxArr[0]
									CardRanktest.Value[0] = MaxArr[1]
									CardRanktest.Value[1] = MaxArr[2]
									CardRanktest.Value[2] = MaxArr[3]
									CardRanktest.Value[3] = MaxArr[4]
									CardRanktest.PlayName = PlayName
									CardRankSlice = append(CardRankSlice, CardRanktest)
								} else {
									CardRanktest.Grade = 8
									CardRanktest.Value = MaxArr
									CardRanktest.PlayName = PlayName
									CardRankSlice = append(CardRankSlice, CardRanktest)
									// --debug//fmt.Println("KINGKONG")
								}
							} else {
								//判断是否是葫芦
								if (MaxArr[0].Rank == MaxArr[2].Rank && MaxArr[3].Rank == MaxArr[4].Rank) ||
									(MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[2].Rank == MaxArr[4].Rank) {
									//如果是葫芦
									//判断三条数组的顺序 保证前三个值相等，后两个值为相等得对子
									if MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[2].Rank == MaxArr[4].Rank {
										CardRanktest.Grade = 7
										CardRanktest.Value[0] = MaxArr[2]
										CardRanktest.Value[1] = MaxArr[3]
										CardRanktest.Value[2] = MaxArr[4]
										CardRanktest.Value[3] = MaxArr[0]
										CardRanktest.Value[4] = MaxArr[1]
										CardRanktest.PlayName = PlayName
										CardRankSlice = append(CardRankSlice, CardRanktest)
									} else {
										CardRanktest.Grade = 7
										CardRanktest.Value = MaxArr
										CardRanktest.PlayName = PlayName
										CardRankSlice = append(CardRankSlice, CardRanktest)
										// --debug//fmt.Println("full house")
									}
								} else {
									//判断是否是顺子
									if (MaxArr[0].Rank-MaxArr[1].Rank == 1 &&
										MaxArr[1].Rank-MaxArr[2].Rank == 1 &&
										MaxArr[2].Rank-MaxArr[3].Rank == 1 &&
										MaxArr[3].Rank-MaxArr[4].Rank == 1) ||
										(MaxArr[0].Rank-MaxArr[1].Rank == 9 &&
											MaxArr[1].Rank-MaxArr[2].Rank == 1 &&
											MaxArr[2].Rank-MaxArr[3].Rank == 1 &&
											MaxArr[3].Rank-MaxArr[4].Rank == 1) {
										//如果是顺子
										CardRanktest.Grade = 5
										CardRanktest.Value = MaxArr
										CardRanktest.PlayName = PlayName
										CardRankSlice = append(CardRankSlice, CardRanktest)
										// --debug//fmt.Println("straight")
									} else {
										//判断是否是三条
										if (MaxArr[0].Rank == MaxArr[2].Rank) ||
											(MaxArr[2].Rank == MaxArr[4].Rank) ||
											(MaxArr[1].Rank == MaxArr[3].Rank) {
											//如果是三条
											//判断三条数组的顺序 保证前三个值相等，后两个值为大小顺序的单值
											if MaxArr[2].Rank == MaxArr[4].Rank {
												CardRanktest.Grade = 4
												CardRanktest.Value[0] = MaxArr[2]
												CardRanktest.Value[1] = MaxArr[3]
												CardRanktest.Value[2] = MaxArr[4]
												CardRanktest.Value[3] = MaxArr[0]
												CardRanktest.Value[4] = MaxArr[1]
												CardRanktest.PlayName = PlayName
												CardRankSlice = append(CardRankSlice, CardRanktest)
											}
											if MaxArr[1].Rank == MaxArr[3].Rank {
												CardRanktest.Grade = 4
												CardRanktest.Value[0] = MaxArr[1]
												CardRanktest.Value[1] = MaxArr[2]
												CardRanktest.Value[2] = MaxArr[3]
												CardRanktest.Value[3] = MaxArr[0]
												CardRanktest.Value[4] = MaxArr[4]
												CardRanktest.PlayName = PlayName
												CardRankSlice = append(CardRankSlice, CardRanktest)
											}
											if MaxArr[0].Rank == MaxArr[2].Rank {
												CardRanktest.Grade = 4
												CardRanktest.Value = MaxArr
												CardRanktest.PlayName = PlayName
												CardRankSlice = append(CardRankSlice, CardRanktest)
												// --debug//fmt.Println("three of a kind")
											}
										} else {
											//判断是否是两对
											if (MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[2].Rank == MaxArr[3].Rank) ||
												(MaxArr[1].Rank == MaxArr[2].Rank && MaxArr[3].Rank == MaxArr[4].Rank) ||
												(MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[3].Rank == MaxArr[4].Rank) {
												//如果是两对
												//判断两对数组的顺序 保证前四个值为从大到小的对子，最后一个值为单值
												if MaxArr[1].Rank == MaxArr[2].Rank && MaxArr[3].Rank == MaxArr[4].Rank {
													CardRanktest.Grade = 3
													CardRanktest.Value[0] = MaxArr[1]
													CardRanktest.Value[1] = MaxArr[2]
													CardRanktest.Value[2] = MaxArr[3]
													CardRanktest.Value[3] = MaxArr[4]
													CardRanktest.Value[4] = MaxArr[0]
													CardRanktest.PlayName = PlayName
													CardRankSlice = append(CardRankSlice, CardRanktest)
												}
												if MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[3].Rank == MaxArr[4].Rank {
													CardRanktest.Grade = 3
													CardRanktest.Value[0] = MaxArr[0]
													CardRanktest.Value[1] = MaxArr[1]
													CardRanktest.Value[2] = MaxArr[3]
													CardRanktest.Value[3] = MaxArr[4]
													CardRanktest.Value[4] = MaxArr[2]
													CardRanktest.PlayName = PlayName
													CardRankSlice = append(CardRankSlice, CardRanktest)
												}
												if MaxArr[0].Rank == MaxArr[1].Rank && MaxArr[2].Rank == MaxArr[3].Rank {
													CardRanktest.Grade = 3
													CardRanktest.Value = MaxArr
													CardRanktest.PlayName = PlayName
													CardRankSlice = append(CardRankSlice, CardRanktest)
													// --debug//fmt.Println("two pair")
												}
											} else {
												//判断是否是对子
												if MaxArr[0].Rank == MaxArr[1].Rank || MaxArr[1].Rank == MaxArr[2].Rank || MaxArr[2].Rank == MaxArr[3].Rank || MaxArr[3].Rank == MaxArr[4].Rank {
													//如果是对子
													//判断对子数组的顺序 保证前2个值为对子，最后三个值为从大到小单值
													if MaxArr[1].Rank == MaxArr[2].Rank {
														CardRanktest.Grade = 2
														CardRanktest.Value[0] = MaxArr[1]
														CardRanktest.Value[1] = MaxArr[2]
														CardRanktest.Value[2] = MaxArr[0]
														CardRanktest.Value[3] = MaxArr[3]
														CardRanktest.Value[4] = MaxArr[4]
														CardRanktest.PlayName = PlayName
														CardRankSlice = append(CardRankSlice, CardRanktest)
													}
													if MaxArr[2].Rank == MaxArr[3].Rank {
														CardRanktest.Grade = 2
														CardRanktest.Value[0] = MaxArr[2]
														CardRanktest.Value[1] = MaxArr[3]
														CardRanktest.Value[2] = MaxArr[0]
														CardRanktest.Value[3] = MaxArr[1]
														CardRanktest.Value[4] = MaxArr[4]
														CardRanktest.PlayName = PlayName
														CardRankSlice = append(CardRankSlice, CardRanktest)
													}
													if MaxArr[3].Rank == MaxArr[4].Rank {
														CardRanktest.Grade = 2
														CardRanktest.Value[0] = MaxArr[3]
														CardRanktest.Value[1] = MaxArr[4]
														CardRanktest.Value[2] = MaxArr[0]
														CardRanktest.Value[3] = MaxArr[1]
														CardRanktest.Value[4] = MaxArr[2]
														CardRanktest.PlayName = PlayName
														CardRankSlice = append(CardRankSlice, CardRanktest)
													}
													if MaxArr[0].Rank == MaxArr[1].Rank {
														CardRanktest.Grade = 2
														CardRanktest.Value = MaxArr
														CardRanktest.PlayName = PlayName
														CardRankSlice = append(CardRankSlice, CardRanktest)
														// --debug//fmt.Println("one pair")
													}
												} else {
													//为高张
													CardRanktest.Grade = 1
													CardRanktest.Value = MaxArr
													CardRanktest.PlayName = PlayName
													CardRankSlice = append(CardRankSlice, CardRanktest)
													// --debug//fmt.Println("high card")
												}
											}
										}

									}
								}
							}
						}
					}
				}
			}
		}
	}
	return CardRankSlice
}
