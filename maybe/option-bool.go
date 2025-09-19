package maybe

// Bool is just a shortcut for Option[bool].
type Bool = Option[bool]

// True is a shorthand for Some(true).
func True() Option[bool] { return Some(true) }

// False is a shorthand for Some(false).
func False() Option[bool] { return Some(false) }

// NoneBool is a shorthand for None[bool]().
func NoneBool() Option[bool] { return None[bool]() }
