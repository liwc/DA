package line

import (
	"fmt"
	// "DA/4_queue/queue"
	"strconv"
)

type record struct {
	key int
	next int
}

type queue struct {
	head int
	tail int
}

// BucketSort 桶排序
func BucketSort(a []int, n int) {
	if n <= 1 {
		return
	}

	// 对排序序列进行封装
	rec := make([]record, n)
	for i := 0; i < n; i++ {
		rec[i].key = a[i]
		rec[i].next = i + 1
	}
	rec[n - 1].next = -1

	/*
	创建静态队列，每一个队列相当于一个桶
	对于十进制数，只需10个桶即可
	*/
	que := make([]queue, 10)

	// 获取最大排序位数
	max := a[0]
	for i := 1; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	r := len(strconv.Itoa(max))
	first := 0
	for i := 1; i <= r; i++ {
		distribute(rec, i, first, que)
		collect(rec, que, &first)
	}

	// 重排
	arrange(rec, n, first)

	// 回放
	for i := 0; i < n; i++ {
		a[i] = rec[i].key
	}
}

/*
分配
把n个记录分配到10个桶中
r：从右数按第r位排序(r=1,2...)
first：第一个记录的下标
*/
func distribute(rec []record, r, first int, que []queue) {
	// 先把队列初始化
	for i := 0; i < 10; i++ {
		que[i].head = -1
	}

	d := 1
	for i := 1; i < r; i++ {
		d *= 10
	}

	cur := first
	for cur != -1 {
		k := rec[cur].key / d % 10

		if que[k].head == -1 {
			// 桶空，则当前记录是第一个
			que[k].head = cur
		} else {
			// 桶非空，则当前记录连接到尾部
			rec[que[k].tail].next = cur
		}

		// 修改队列尾部，这一句不可忘了
		que[k].tail = cur

		// 继续给下一个记录分桶
		cur = rec[cur].next
	}
}


/*
收集
把10个桶连接起来
*/
func collect(rec []record, que []queue, first *int) {
	// k是桶号
	k := 0
	// 先找到第一个非空的桶
	for que[k].head == -1 {
		k++
	}

	// 用first记录第一个记录的位置，为下次分配使用
	*first = que[k].head
	tail := que[k].tail
	for k < 9 {
		k++
		for k < 9 && que[k].head == -1 {
			k++
		}

		// 把上一个桶的尾部和当前桶的头部连接起来
		if que[k].head != -1 {
			rec[tail].next = que[k].head
			tail = que[k].tail
		}
	}

	rec[tail].next = -1
}


/*
重排
first：第一个记录的下标
*/
func arrange(rec []record, n, first int) {
	j := first
	for i := 0; i < n - 1; i++ {
		temp := rec[j]
		rec[j] = rec[i]
		rec[i] = temp
		rec[i].next = j
		j = temp.next
		for j <= i {
			j = rec[j].next
		}
	}
}

// BucketTest ...
func BucketTest(array []int) {
	fmt.Println("BucketTest begin")
	// data := make([]int, len(array))
	// copy(data, array)
	data := []int{123, 234, 45, 111, 6, 128}
	fmt.Printf("src: %v\n", data)
	BucketSort(data, len(data))
	fmt.Printf("dst: %v\n", data)
	fmt.Println("BucketTest end")
}