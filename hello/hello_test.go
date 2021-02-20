package main

import "testing"

func TestHello(t *testing.T){
	assertMessage := func(t *testing.T, got, want string){
		t.Helper()
		if got != want{
			t.Errorf("got '%q'  want '%q'", got, want)
		}

	}

	t.Run("say hello in english", func(t *testing.T){
	got := Hello("ewan", "english")
	want := "hello  ewan"
	assertMessage(t, got, want)
	})


	t.Run("if empty say beijing 0", func(t *testing.T){
		got := Hello("", "spanish")
		want := "hala  world"
		assertMessage(t, got, want)
		
	})
}