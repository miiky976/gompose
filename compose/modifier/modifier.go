package modifier

type Modifier struct {
	Space int
	FillMaxWidth bool
	FillMaxHeight bool
}

func NewModifier() Modifier{
	return Modifier{
		Space: 0,
		FillMaxWidth: true,
		FillMaxHeight: false,
	}
}

func (m Modifier) FitMaxHeight() Modifier {
	m.FillMaxHeight = true
	return m
}

func (m Modifier) FitMaxWidth() Modifier {
	m.FillMaxWidth = true
	return m
}
