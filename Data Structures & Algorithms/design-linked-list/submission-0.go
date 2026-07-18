type ListNode struct {
    val  int
    next *ListNode
}

type MyLinkedList struct {
    head *ListNode
    size int
}

func Constructor() MyLinkedList {
    return MyLinkedList{head: &ListNode{val: 0}, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
    if index >= this.size {
        return -1
    }
    cur := this.head.next
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
    node := &ListNode{val: val}
    node.next = this.head.next
    this.head.next = node
    this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
    node := &ListNode{val: val}
    cur := this.head
    for cur.next != nil {
        cur = cur.next
    }
    cur.next = node
    this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index > this.size {
        return
    }
    cur := this.head
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    node := &ListNode{val: val}
    node.next = cur.next
    cur.next = node
    this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index >= this.size {
        return
    }
    cur := this.head
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    cur.next = cur.next.next
    this.size--
}