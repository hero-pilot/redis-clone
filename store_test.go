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
	v, error := store.Get("hello")
	if error != nil{
		t.Error("It should exist, but it doesn't ?")
	}else if v != "world"{
		t.Error("Wrong value, which is wierd")
	}
	v, error = store.Get("missing")
	if error ==nil{
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
	
	if v, error := set.Get("yo"); error !=nil{
		t.Error("It hasn't renamed the key!!")
	}else if _, error = set.Get("wow"); error ==nil{
		t.Error("The old key still exists dumbass !!!")
	}else if v != value {
		t.Errorf("The value should remain the same! , expected: %v, got:%v", value, v)
	}
}

func TestPop_ReturnsAndRemoves(t *testing.T) {
	store := NewStore()
	store.Set("wow", "wow forever")

	value , error := store.Get("wow")
	if error !=nil{
		t.Error("What the hell!")
	}
	if v, error := store.Pop("wow") ; error !=nil{
		t.Error("Again, what the hell?")
	}else if v != value{
		t.Error("Now this is some funny shit going on")
	}

	if _, error= store.Pop("wow") ; error ==nil{
		t.Error("This should not be in the store bitch")
	}
	
}

func TestRename_MissingKeyCreatesNothing(t *testing.T) {

	store := NewStore()
	store.Set("wow", "boo")
	store.Rename("no", "yes")
	if _ ,error := store.Get("no"); error ==nil{
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