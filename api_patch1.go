package lua

// GetGlobalRaw returns the the value for the given k variable name.
func (l *State) GetGlobalRaw(k interface{}) interface{} {
	return l.global.GetRaw(k)
}

// SetGlobalRaw adds k,v as global.
//
// It's better performance alternative to:
// l.Push("my greate message")
// l.SetGlobal("message")
func (l *State) SetGlobalRaw(k, v interface{}) {
	l.global.SetRaw(k, v)
}

// GetStack returns a raw value from the stack givent its index.
// Use at your own risk.
func (l *State) GetStack(i int) interface{} {
	return l.stack.Get(i)
}

// SetStack sets the given value at the given index.
// Use at your own risk.
func (l *State) SetStack(i int, v interface{}) {
	l.stack.Set(i, v)
}
