package sort

func BitmapSort(items []int, min, max int) {
	lens := (max-min) / 8 + 1
	s := make([]int, lens)

	for i := range items {
		s[items[i]/8] |= (1 << uint(items[i]%8))
	}

	var k int
	for i := range s {
		for j := 0; j < 8; j++ {
			if s[i] & (1 << uint(j)) > 0 {
				items[k] = i * 8 + j
				k++
			}
		}
	}
}





/*
* 10亿int整型数，以及一台可用内存为1GB的机器，时间复杂度要求O(n)，统计只出现一次的数？
*/

//分析：一个Int型占4个字节，10亿个整数全部存入内存要40亿字节，即4GB，显然不可能将所有数据读入内存；
//方法：1位图法 2分治法，归并排序
//1.位图法 int整型数占4字节（Byte）,也就是32位(bit)。那么把所有int整型数字表示出来需要2^32bit的空间，
//换算成字节单位也就是2^32/8 = 2^29 Byte，大约等于512MB。
