package maybe

// Int is just a shortcut for Option[int].
type Int = Option[int]

// NoneInt is a shortcut for None[int].
func NoneInt() Int { return None[int]() }
