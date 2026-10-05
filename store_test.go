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

func TestRename_MovesTheValue(t *testing.T) {
	set := NewStore()
	set.Set("wow", "pashm")
	value, _ := set.Get("wow")
	set.Rename("wow", "yo")
	
	if v, exists := set.Get("yo"); !exists{
		t.Error("It hasn't renamed the key!!")
	}else if _, exists = set.Get("wow"); exists{
		t.Error("The old key still exists dumbass !!!")
	}else if v != value {
		t.Errorf("The value should remain the same! , expected: %v, got:%v", value, v)
	}
}

func TestPop_ReturnsAndRemoves(t *testing.T) {
	store := NewStore()
	store.Set("wow", "wow forever")

	value , exist := store.Get("wow")
	if !exist{
		t.Error("What the hell!")
	}
	if v, exists := store.Pop("wow") ; !exists{
		t.Error("Again, what the hell?")
	}else if v != value{
		t.Error("Now this is some funny shit going on")
	}

	if _, exist= store.Pop("wow") ; exist{
		t.Error("This should not be in the store bitch")
	}
	
}

func TestRename_MissingKeyCreatesNothing(t *testing.T) {

	store := NewStore()
	store.Set("wow", "boo")
	store.Rename("no", "yes")
	if _ ,exists := store.Get("no"); exists{
		t.Error("It shouldn't exist !!!")
	}
	got := store.Keys()
	if len(got) != 1{
		t.Error("Something fishy is going on here")
	}
	if got[0] != "wow"{
		t.Error("This is wrong buddy")
	}
}