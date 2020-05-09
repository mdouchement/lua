package lua

import "github.com/milochristiansen/lua/luautil"

// Function is a Lua or native function with its upvalues.
type Function = function

// CompileAsFunc compiles the given script (line should be set to 1) as Function.
func CompileAsFunc(script, name string, line int) (*Function, error) {
	proto, err := compSource(script, name, line)
	if err != nil {
		return nil, err
	}

	l := NewState()
	return l.asFunc(proto, l.global), nil
}

// AsFunc takes a Block and adds execution environment.
func AsFunc(proto *Block) *Function {
	l := NewState()
	return l.asFunc(proto, l.global)
}

// LoadFunc adds to the stack the given Function with the given env.
// This method is NOT Threadsafe, you cannot share the same Function accross several Lua's VM.
// You can use sync.Pool to reuse instances.
func (l *State) LoadFunc(f *Function, env int) {
	envv := l.global
	if env != 0 {
		var ok bool
		envv, ok = l.get(env).(*table)
		if !ok {
			luautil.Raise("Value used as environment is not a table.", luautil.ErrTypGenRuntime)
		}
	}

	// Top level functions must have their first upvalue as _ENV
	if len(f.up) > 0 {
		if f.up[0].name != "_ENV" && f.up[0].name != "" {
			luautil.Raise("Top level function without _ENV or _ENV in improper position.", luautil.ErrTypGenRuntime)
		}

		f.up[0].val = envv
	}

	l.stack.Push(f)
}

// SetFunc sets to the stack the given Function with the given env ate the given index.
// Use at your own risk.
// This method is NOT Threadsafe, you cannot share the same Function accross several Lua's VM.
// You can use sync.Pool to reuse instances.
func (l *State) SetFunc(i int, f *Function, env int) {
	envv := l.global
	if env != 0 {
		var ok bool
		envv, ok = l.get(env).(*table)
		if !ok {
			luautil.Raise("Value used as environment is not a table.", luautil.ErrTypGenRuntime)
		}
	}

	// Top level functions must have their first upvalue as _ENV
	if len(f.up) > 0 {
		if f.up[0].name != "_ENV" && f.up[0].name != "" {
			luautil.Raise("Top level function without _ENV or _ENV in improper position.", luautil.ErrTypGenRuntime)
		}

		f.up[0].val = envv
	}

	l.stack.Set(i, f)
}
