package modifier

type Modifier struct {
	Space int
	FillMaxWidth bool
	FillMaxHeight bool
}

func NewEmptyModifier() Modifier{
	return Modifier{
		Space: 0,
		FillMaxWidth: true,
		FillMaxHeight: false,
	}
}

