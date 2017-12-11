package repository

import "testing"

func TestFindById(t *testing.T) {
	group := FindById(2)
	if group.Name == "" {
		t.Errorf("%#v", group)
	}
}

func TestFindAll(t *testing.T) {
	groups := FindAll()
	if len(groups) == 0 {
		for _, group := range groups {
			t.Errorf("%#v", group)
		}
	}
}