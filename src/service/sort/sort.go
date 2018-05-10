package sort

import (
	"merger/src/entity"
	"sort"
)

type customSort []entity.User

func (a customSort) Len() int {
	return len(a)
}

func (a customSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a customSort) Less(i, j int) bool {
	if a[i].GetPoints() > a[j].GetPoints() {
		return true
	}

	if a[i].GetPoints() < a[j].GetPoints() {
		return false
	}

	return a[i].GetTrustLevel() > a[j].GetTrustLevel()
}

func Sort(users []entity.User) []entity.User {

	sort.Sort(customSort(users))

	return users
}
