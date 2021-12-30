# 单链表：前驱节点使用场景
* 删除某个节点
* 尾部添加节点

# 单链表：快慢指针中跳出循环的的判断
> fast != nil && fast.Nex != nil

* 偶数个节点中，如果希望fast结束遍历时，slow指向后半段的第一个node，则fast, slow = head, head;
* 偶数个节点中，如果希望fast结束遍历时，slow指向前半段的最后一个node，则slow, fast = head, head.Next;
* 判断链表是否循环时，则slow, fast = head, head.Next;
