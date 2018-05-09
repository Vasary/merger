package entity

import "time"

type User struct {
	Id int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Mask UserMask
	points int64
}

func (u *User) GetId() int64  {
	return u.Id
}

func (u *User) GetCreatedAt() time.Time  {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time  {
	return u.UpdatedAt
}

func (u *User) GetMask() UserMask  {
	return u.Mask
}

func (u *User) IncreasePoints(points int64)  {
	u.points += points
}

func (u *User) GetPoints() int64 {
	return u.points
}