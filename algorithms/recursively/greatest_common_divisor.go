/*
* 1、该算法为CPU密集型算法，分别使用了递归思想、迭代思想实现最大公约数求值
* 2、使用递归会不断增加函数调用栈，有栈溢出风险，并且函数调用的开销比较大，系统要为每次函数调用分配存储空间，并将调用点压栈予以记录，函数调用后，还要释放空间，弹栈恢复断点。所以说函数调用不仅浪费空间，还浪费时间；
* 3、使用循环迭代，则消除了系统进行函数调用带来的额外开销；
*/

package gcd

func gcdRecur(n1, n2 int) int {
	if n2 == 0 {
		return n1
	}

	r := n1 % n2

	return gcdRecur(n2, r)
}

func gcdIteration(n1, n2 int) int {
	for n2 != 0 {
		r := n1 % n2
		n1 = n2
		n2 = r
	}

	return n1
}



