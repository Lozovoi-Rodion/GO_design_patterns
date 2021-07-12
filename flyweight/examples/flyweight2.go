package flyweight

import (
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{fullName}
}

var AllNames []string

type User2 struct {
	Names []uint8
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.Names {
		parts = append(parts, AllNames[id])
	}
	return strings.Join(parts, " ")
}

func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i := range AllNames {
			if AllNames[i] == s {
				return uint8(i)
			}
		}
		AllNames = append(AllNames, s)
		return uint8(len(AllNames) - 1)
	}
	result := User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.Names = append(result.Names, getOrAdd(p))
	}
	return &result
}
