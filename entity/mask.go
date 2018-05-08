package entity

type UserMask struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Birthday string `json:"birthday"`
}

func (u *UserMask) HasName() bool  {
	return u.Name != ""
}

func (u *UserMask) HashSurname() bool  {
	return u.Surname != ""
}

func (u *UserMask) HashPatronymic() bool  {
	return u.Patronymic != ""
}

func (u *UserMask) HasBirthday() bool  {
	return u.Birthday != ""
}