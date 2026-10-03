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

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore()

	got := store.Keys()
	if len(got) != 0{
		t.Errorf("The slice is supposed to be empty bitch!; got: %v", got)
	}

}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore()
	store.Set("hello", "world")
	v, exists := store.Get("hello")
	if !exists{
		t.Error("It should exist, but it doesn't ?")
	}else if v != "world"{
		t.Error("Wrong value, which is wierd")
	}
	v, exists = store.Get("missing")
	if exists{
		t.Error("It shouldn't exist  !!1")
	}else if v != ""{
		t.Error("It shouldn't have a value at all!!!")
	}

}