//! Generated by Snowball 2.0.0 - https://snowballstem.org/

package german

import (
	snowballRuntime "github.com/blevesearch/snowballstem"
)

var A_0 = []*snowballRuntime.Among{
	{Str: "", A: -1, B: 5, F: nil},
	{Str: "U", A: 0, B: 2, F: nil},
	{Str: "Y", A: 0, B: 1, F: nil},
	{Str: "\u00E4", A: 0, B: 3, F: nil},
	{Str: "\u00F6", A: 0, B: 4, F: nil},
	{Str: "\u00FC", A: 0, B: 2, F: nil},
}

var A_1 = []*snowballRuntime.Among{
	{Str: "e", A: -1, B: 2, F: nil},
	{Str: "em", A: -1, B: 1, F: nil},
	{Str: "en", A: -1, B: 2, F: nil},
	{Str: "ern", A: -1, B: 1, F: nil},
	{Str: "er", A: -1, B: 1, F: nil},
	{Str: "s", A: -1, B: 3, F: nil},
	{Str: "es", A: 5, B: 2, F: nil},
}

var A_2 = []*snowballRuntime.Among{
	{Str: "en", A: -1, B: 1, F: nil},
	{Str: "er", A: -1, B: 1, F: nil},
	{Str: "st", A: -1, B: 2, F: nil},
	{Str: "est", A: 2, B: 1, F: nil},
}

var A_3 = []*snowballRuntime.Among{
	{Str: "ig", A: -1, B: 1, F: nil},
	{Str: "lich", A: -1, B: 1, F: nil},
}

var A_4 = []*snowballRuntime.Among{
	{Str: "end", A: -1, B: 1, F: nil},
	{Str: "ig", A: -1, B: 2, F: nil},
	{Str: "ung", A: -1, B: 1, F: nil},
	{Str: "lich", A: -1, B: 3, F: nil},
	{Str: "isch", A: -1, B: 2, F: nil},
	{Str: "ik", A: -1, B: 2, F: nil},
	{Str: "heit", A: -1, B: 3, F: nil},
	{Str: "keit", A: -1, B: 4, F: nil},
}

var G_v = []byte{17, 65, 16, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 0, 32, 8}

var G_s_ending = []byte{117, 30, 5}

var G_st_ending = []byte{117, 30, 4}

type Context struct {
	i_x  int
	i_p2 int
	i_p1 int
}

func r_prelude(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	var v_1 = env.Cursor
replab0:
	for {
		var v_2 = env.Cursor
	lab1:
		for range [2]struct{}{} {
		lab2:
			for {
				var v_3 = env.Cursor
			lab3:
				for {
					env.Bra = env.Cursor
					if !env.EqS("\u00DF") {
						break lab3
					}
					env.Ket = env.Cursor
					if !env.SliceFrom("ss") {
						return false
					}
					break lab2
				}
				env.Cursor = v_3
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
				break lab2
			}
			continue replab0
		}
		env.Cursor = v_2
		break replab0
	}
	env.Cursor = v_1
replab4:
	for {
		var v_4 = env.Cursor
	lab5:
		for range [2]struct{}{} {
		golab6:
			for {
				var v_5 = env.Cursor
			lab7:
				for {
					if !env.InGrouping(G_v, 97, 252) {
						break lab7
					}
					env.Bra = env.Cursor
				lab8:
					for {
						var v_6 = env.Cursor
					lab9:
						for {
							if !env.EqS("u") {
								break lab9
							}
							env.Ket = env.Cursor
							if !env.InGrouping(G_v, 97, 252) {
								break lab9
							}
							if !env.SliceFrom("U") {
								return false
							}
							break lab8
						}
						env.Cursor = v_6
						if !env.EqS("y") {
							break lab7
						}
						env.Ket = env.Cursor
						if !env.InGrouping(G_v, 97, 252) {
							break lab7
						}
						if !env.SliceFrom("Y") {
							return false
						}
						break lab8
					}
					env.Cursor = v_5
					break golab6
				}
				env.Cursor = v_5
				if env.Cursor >= env.Limit {
					break lab5
				}
				env.NextChar()
			}
			continue replab4
		}
		env.Cursor = v_4
		break replab4
	}
	return true
}

func r_mark_regions(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	context.i_p1 = env.Limit
	context.i_p2 = env.Limit
	var v_1 = env.Cursor
	{
		var c = env.ByteIndexForHop((3))
		if int32(0) > c || c > int32(env.Limit) {
			return false
		}
		env.Cursor = int(c)
	}
	context.i_x = env.Cursor
	env.Cursor = v_1
golab0:
	for {
	lab1:
		for {
			if !env.InGrouping(G_v, 97, 252) {
				break lab1
			}
			break golab0
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
golab2:
	for {
	lab3:
		for {
			if !env.OutGrouping(G_v, 97, 252) {
				break lab3
			}
			break golab2
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
	context.i_p1 = env.Cursor
lab4:
	for {
		if !(context.i_p1 < context.i_x) {
			break lab4
		}
		context.i_p1 = context.i_x
		break lab4
	}
golab5:
	for {
	lab6:
		for {
			if !env.InGrouping(G_v, 97, 252) {
				break lab6
			}
			break golab5
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
golab7:
	for {
	lab8:
		for {
			if !env.OutGrouping(G_v, 97, 252) {
				break lab8
			}
			break golab7
		}
		if env.Cursor >= env.Limit {
			return false
		}
		env.NextChar()
	}
	context.i_p2 = env.Cursor
	return true
}

func r_postlude(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	var among_var int32
replab0:
	for {
		var v_1 = env.Cursor
	lab1:
		for range [2]struct{}{} {
			env.Bra = env.Cursor
			among_var = env.FindAmong(A_0, context)
			if among_var == 0 {
				break lab1
			}
			env.Ket = env.Cursor
			if among_var == 1 {
				if !env.SliceFrom("y") {
					return false
				}
			} else if among_var == 2 {
				if !env.SliceFrom("u") {
					return false
				}
			} else if among_var == 3 {
				if !env.SliceFrom("a") {
					return false
				}
			} else if among_var == 4 {
				if !env.SliceFrom("o") {
					return false
				}
			} else if among_var == 5 {
				if env.Cursor >= env.Limit {
					break lab1
				}
				env.NextChar()
			}
			continue replab0
		}
		env.Cursor = v_1
		break replab0
	}
	return true
}

func r_R1(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	if !(context.i_p1 <= env.Cursor) {
		return false
	}
	return true
}

func r_R2(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	if !(context.i_p2 <= env.Cursor) {
		return false
	}
	return true
}

func r_standard_suffix(env *snowballRuntime.Env, ctx interface{}) bool {
	context := ctx.(*Context)
	_ = context
	var among_var int32
	var v_1 = env.Limit - env.Cursor
lab0:
	for {
		env.Ket = env.Cursor
		among_var = env.FindAmongB(A_1, context)
		if among_var == 0 {
			break lab0
		}
		env.Bra = env.Cursor
		if !r_R1(env, context) {
			break lab0
		}
		if among_var == 1 {
			if !env.SliceDel() {
				return false
			}
		} else if among_var == 2 {
			if !env.SliceDel() {
				return false
			}
			var v_2 = env.Limit - env.Cursor
		lab1:
			for {
				env.Ket = env.Cursor
				if !env.EqSB("s") {
					env.Cursor = env.Limit - v_2
					break lab1
				}
				env.Bra = env.Cursor
				if !env.EqSB("nis") {
					env.Cursor = env.Limit - v_2
					break lab1
				}
				if !env.SliceDel() {
					return false
				}
				break lab1
			}
		} else if among_var == 3 {
			if !env.InGroupingB(G_s_ending, 98, 116) {
				break lab0
			}
			if !env.SliceDel() {
				return false
			}
		}
		break lab0
	}
	env.Cursor = env.Limit - v_1
	var v_3 = env.Limit - env.Cursor
lab2:
	for {
		env.Ket = env.Cursor
		among_var = env.FindAmongB(A_2, context)
		if among_var == 0 {
			break lab2
		}
		env.Bra = env.Cursor
		if !r_R1(env, context) {
			break lab2
		}
		if among_var == 1 {
			if !env.SliceDel() {
				return false
			}
		} else if among_var == 2 {
			if !env.InGroupingB(G_st_ending, 98, 116) {
				break lab2
			}
			{
				var c = env.ByteIndexForHop(-(3))
				if int32(env.LimitBackward) > c || c > int32(env.Limit) {
					break lab2
				}
				env.Cursor = int(c)
			}
			if !env.SliceDel() {
				return false
			}
		}
		break lab2
	}
	env.Cursor = env.Limit - v_3
	var v_4 = env.Limit - env.Cursor
lab3:
	for {
		env.Ket = env.Cursor
		among_var = env.FindAmongB(A_4, context)
		if among_var == 0 {
			break lab3
		}
		env.Bra = env.Cursor
		if !r_R2(env, context) {
			break lab3
		}
		if among_var == 1 {
			if !env.SliceDel() {
				return false
			}
			var v_5 = env.Limit - env.Cursor
		lab4:
			for {
				env.Ket = env.Cursor
				if !env.EqSB("ig") {
					env.Cursor = env.Limit - v_5
					break lab4
				}
				env.Bra = env.Cursor
				var v_6 = env.Limit - env.Cursor
			lab5:
				for {
					if !env.EqSB("e") {
						break lab5
					}
					env.Cursor = env.Limit - v_5
					break lab4
				}
				env.Cursor = env.Limit - v_6
				if !r_R2(env, context) {
					env.Cursor = env.Limit - v_5
					break lab4
				}
				if !env.SliceDel() {
					return false
				}
				break lab4
			}
		} else if among_var == 2 {
			var v_7 = env.Limit - env.Cursor
		lab6:
			for {
				if !env.EqSB("e") {
					break lab6
				}
				break lab3
			}
			env.Cursor = env.Limit - v_7
			if !env.SliceDel() {
				return false
			}
		} else if among_var == 3 {
			if !env.SliceDel() {
				return false
			}
			var v_8 = env.Limit - env.Cursor
		lab7:
			for {
				env.Ket = env.Cursor
			lab8:
				for {
					var v_9 = env.Limit - env.Cursor
				lab9:
					for {
						if !env.EqSB("er") {
							break lab9
						}
						break lab8
					}
					env.Cursor = env.Limit - v_9
					if !env.EqSB("en") {
						env.Cursor = env.Limit - v_8
						break lab7
					}
					break lab8
				}
				env.Bra = env.Cursor
				if !r_R1(env, context) {
					env.Cursor = env.Limit - v_8
					break lab7
				}
				if !env.SliceDel() {
					return false
				}
				break lab7
			}
		} else if among_var == 4 {
			if !env.SliceDel() {
				return false
			}
			var v_10 = env.Limit - env.Cursor
		lab10:
			for {
				env.Ket = env.Cursor
				if env.FindAmongB(A_3, context) == 0 {
					env.Cursor = env.Limit - v_10
					break lab10
				}
				env.Bra = env.Cursor
				if !r_R2(env, context) {
					env.Cursor = env.Limit - v_10
					break lab10
				}
				if !env.SliceDel() {
					return false
				}
				break lab10
			}
		}
		break lab3
	}
	env.Cursor = env.Limit - v_4
	return true
}

func Stem(env *snowballRuntime.Env) bool {
	var context = &Context{
		i_x:  0,
		i_p2: 0,
		i_p1: 0,
	}
	_ = context
	var v_1 = env.Cursor
	r_prelude(env, context)
	env.Cursor = v_1
	var v_2 = env.Cursor
	r_mark_regions(env, context)
	env.Cursor = v_2
	env.LimitBackward = env.Cursor
	env.Cursor = env.Limit
	r_standard_suffix(env, context)
	env.Cursor = env.LimitBackward
	var v_4 = env.Cursor
	r_postlude(env, context)
	env.Cursor = v_4
	return true
}
