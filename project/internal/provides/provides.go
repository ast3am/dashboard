package provides

type provides []string

func NewSmsProvidesList() *provides {
	return &provides{
		"Topolo", "Rond", "Kildy",
	}
}

func NewVoiceProvidesList() *provides {
	return &provides{
		"TransparentCalls", "E-Voice", "JustPhone",
	}
}

func NewMailProvidesList() *provides {
	return &provides{
		"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Protonmail", "Yandex", "Mail.ru",
	}
}

func (pa provides) CheckProvide(in string) bool {
	for _, val := range pa {
		if val == in {
			return true
		}
	}
	return false
}
