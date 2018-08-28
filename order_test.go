package sdk

import (
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestOrdering(t *testing.T) {
	s := &PluginSpec{}
	if err := yaml.Unmarshal([]byte(SpecChatBot), s); err != nil {
		t.Fatal(err)
	}
	if err := PostProcessSpec(s); err != nil {
		t.Fatal(err)
	}
	if len(s.Actions) == 0 {
		t.Fatal("parsing actions failed")
	}
	uniqueOrders := make(map[int]struct{})
	for _, a := range s.Actions {
		if a.Order == 0 {
			t.Fatal("expected not to see 0s in order, but saw 0s")
		}
		if _, ok := uniqueOrders[a.Order]; ok {
			t.Fatal("expected unqiue orders, but found a dupe")
		}
		uniqueOrders[a.Order] = struct{}{}
	}
	if len(s.Triggers) == 0 {
		t.Fatal("parsing triggers failed")
	}
	uniqueOrders = make(map[int]struct{})
	for _, tr := range s.Triggers {
		if tr.Order == 0 {
			t.Fatal("expected not to see 0s in order, but saw 0s")
		}
		if _, ok := uniqueOrders[tr.Order]; ok {
			t.Fatal("expected unqiue orders, but found a dupe")
		}
		uniqueOrders[tr.Order] = struct{}{}
	}
	for _, rt := range s.RawTypes {
		uniqueOrders = make(map[int]struct{})
		for _, ty := range rt {
			if ty.Order == 0 {
				t.Fatal("expected not to see 0s in order, but saw 0s")
			}
			if _, ok := uniqueOrders[ty.Order]; ok {
				t.Fatal("expected unqiue orders, but found a dupe")
			}
			uniqueOrders[ty.Order] = struct{}{}
		}
	}
	for _, rt := range s.Types {
		uniqueOrders = make(map[int]struct{})
		for _, ty := range rt.Fields {
			if ty.Order == 0 {
				t.Fatal("expected not to see 0s in order, but saw 0s")
			}
			if _, ok := uniqueOrders[ty.Order]; ok {
				t.Fatal("expected unqiue orders, but found a dupe")
			}
			uniqueOrders[ty.Order] = struct{}{}
		}
	}
}
