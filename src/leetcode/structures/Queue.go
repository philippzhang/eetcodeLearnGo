package structures

type Queue struct {
	nums []int
}

//构造函数
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

//Push
func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

//Pop
func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

//Len
func (q *Queue) Len() int {
	return len(q.nums)
}

//IsEmpty
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
