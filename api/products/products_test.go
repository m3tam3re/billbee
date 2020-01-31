package products

import (
	"testing"
)

func TestGetOne(t *testing.T) {
	p := Product{}
	err := p.GetOne("weltkarte-graphit-kork-l_100x70-1-S")
	if err != nil {
		t.Error(err)
	}
	t.Log(p)
}
