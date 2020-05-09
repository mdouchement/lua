package lua

import "github.com/milochristiansen/lua/luautil"

// A Block holds compiled statements.
type Block = funcProto

// CompileSource compiles the given script (line should be set to 1).
func CompileSource(script, name string, line int) (*Block, error) {
	return compSource(script, name, line)
}

// LoadBlock loads the Block of compiled statements.
// It Pushes the Block on the stack.
// This method is Threadsafe, you can share the same Block accross several Lua's VM.
func (l *State) LoadBlock(proto *Block, env int) error {
	envv := l.global
	if env != 0 {
		var ok bool
		envv, ok = l.get(env).(*table)
		if !ok {
			return luautil.Error{Msg: "Value used as environment is not a table.", Type: luautil.ErrTypGenRuntime}
		}
	}

	l.stack.Push(l.asFunc(proto, envv))
	return nil
}
