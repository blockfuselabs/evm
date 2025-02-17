package types

import ("testing")

func TestStack(t *testing.T){
	

	stack := NewStack()
	//....Test Push....//
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if len(stack.Data) != 3 {
		t.Errorf("expected stack size 3, got %d", len(stack.Data))
	}

	//.... Test Pop....//
	val := stack.Pop()
	if val != 30 {
		t.Errorf("expected 30, got %d", val)
	}

	val = stack.Pop()
	if val != 20 {
		t.Errorf("expected 20, got %d", val)
	}

	val = stack.Pop()
	if val != 10 {
		t.Errorf("expected 10, got %d", val)
	}

	//....Test Stack Underflow (should panic)....//
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic on empty stack pop, but got none")
		}
	}()
	stack.Pop()
}

