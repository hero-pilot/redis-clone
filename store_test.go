package main

import (
	"reflect"
	"testing"
)

func TestKeys_returns_sorted_keys(t *testing.T) {
	store := NewStore()
	store.Set("alpha", "beta")
	store.Set("pashm", "shishe")
	store.Set("erfan", "dousti")

	got := store.Keys()
	expected := []string{"alpha", "erfan", "pashm"}

	if !reflect.DeepEqual(got, expected){
		t.Errorf("The keys are different!!, got: %v , wanted: %v", got, expected)
	}
}