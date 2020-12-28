package structures

type Stack struct {
	nums []int
}

//构造函数
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

//Push
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

//Pop
func (s *Stack) Pop() int {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

//Len
func (s *Stack) Len() int {
	return len(s.nums)
}

//IsEmpty
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
