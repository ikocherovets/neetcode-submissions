type ListNode struct {
    key, val int
    next     *ListNode
}

type MyHashMap struct {
    data []*ListNode
}

func Constructor() MyHashMap {
    data := make([]*ListNode, 1000)
    for i := range data {
        data[i] = &ListNode{key: -1, val: -1}
    }
    return MyHashMap{data: data}
}

func (this *MyHashMap) hash(key int) int {
    return key % len(this.data)
}

func (this *MyHashMap) Put(key int, value int) {
    cur := this.data[this.hash(key)]
    for cur.next != nil {
        if cur.next.key == key {
            cur.next.val = value
            return
        }
        cur = cur.next
    }
    cur.next = &ListNode{key: key, val: value}
}

func (this *MyHashMap) Get(key int) int {
    cur := this.data[this.hash(key)].next
    for cur != nil {
        if cur.key == key {
            return cur.val
        }
        cur = cur.next
    }
    return -1
}

func (this *MyHashMap) Remove(key int) {
    cur := this.data[this.hash(key)]
    for cur.next != nil {
        if cur.next.key == key {
            cur.next = cur.next.next
            return
        }
        cur = cur.next
    }
}