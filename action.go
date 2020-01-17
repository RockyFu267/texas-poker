package main

import "fmt"

//PerflopAction 翻牌前操作
func (p *PlayPlayer) PerflopAction(PotChip int64) (potChip int64, error error) {
	//--debug
	fmt.Println("小盲是：", p[0].Name)
	fmt.Println("大盲是：", p[1].Name)
	//--debug
	//目前只写机器CALL
	p[0].Chip = p[0].Chip - p[0].BBSBChip
	p[1].Chip = p[1].Chip - p[1].BBSBChip
	PotChip = p[0].BBSBChip + p[1].BBSBChip
	//justcall
	p[0].Chip = p[0].Chip - p[0].BBSBChip
	PotChip = PotChip + p[0].BBSBChip
	fmt.Println("场上选手剩余筹码：", p[0].Name, ":", p[0].Chip, "\t", p[1].Name, ":", p[1].Chip, "\t", "底池:", PotChip)
	return PotChip, nil
}

//FlopAction 翻牌后的一轮操作
func (p *PlayPlayer) FlopAction(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//TurnAction 翻牌后的一轮操作
func (p *PlayPlayer) TurnAction(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//RiverAction 翻牌后的一轮操作
func (p *PlayPlayer) RiverAction(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//Call 跟注
func (p *PlayPlayer) Call(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//Raise 加注
func (p *PlayPlayer) Raise(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//Allin 全下 无条件可以看五张 最后开牌
func (p *PlayPlayer) Allin(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}

//fold 弃牌
func (p *PlayPlayer) Fold(PotChip int64) (potChip int64, error error) {
	return PotChip, nil
}
