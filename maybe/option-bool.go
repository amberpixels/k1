package maybe

// Bool is just a shortcut for Option[bool].
type Bool = Option[bool]

func True() Option[bool]     { return Some(true) }
func False() Option[bool]    { return Some(false) }
func NoneBool() Option[bool] { return None[bool]() }
