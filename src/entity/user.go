package entity

type User struct {
	id         int64
	mask       UserMask
	points     int64
	trustLevel int64
}

func (u *User) GetId() int64 {
	return u.id
}

func (u *User) SetId(id int64) {
	u.id = id
}

func (u *User) GetMask() UserMask {
	return u.mask
}

func (u *User) SetMask(mask UserMask) {
	u.mask = mask
}

func (u *User) IncreasePoints(points int64) {
	u.points += points
}

func (u *User) GetPoints() int64 {
	return u.points
}

func (u *User) SetTrustLevel(level int64) {
	u.trustLevel = level
}

func (u *User) GetTrustLevel() int64 {
	return u.trustLevel
}
