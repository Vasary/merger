package entity

type UserList struct {
	Value string
	List []int64
}

func (u UserList) GetValue() string  {
	return u.Value
}

func (u UserList) GetIdList() []int64 {
	return u.List
}