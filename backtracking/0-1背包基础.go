package backtracking

import (
	"math"
)

/*
背包总的承载重量是Wkg，现在我们有n个物品，每个物品的重量不等，并且不可分割(因为不可分割，所以称为0-1背包问题)。
我们现在期望选择几件物品装到背包中。在不超过背包所能装载重量的前提下，如何让
背包中物品的总重量最大？
 */

/*
现对第一个
 */
var maxWeight = math.MinInt64

// i表示考察到哪个物品了
// cw表示当前已装进背包的物品质量和
// w背包重量
// items表示每个物品的重量
func Bag(i, cw, w int, items []int) {
	n := len(items)
	if cw == w || i == n { // cw==w表示装满了
		if cw > maxWeight {
			maxWeight = cw
		}
		println(maxWeight)
		return
	}

	// 先对第一个物品进行处理，选择装进去或不装进去
	Bag(i+1, cw, w, items) // 不装进去
	if cw + items[i] <= w { // 已经超过可以背包承受的重量的时候，就不要再装了
		Bag(i+1, cw + items[i], w, items) // 装进去
	}
}
