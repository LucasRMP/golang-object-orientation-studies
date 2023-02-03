package holder

type Holder struct {
	Name       string
	SSN        string
	Profession string
}

func NewHolder(name string, ssn string, profession string) *Holder {
	return &Holder{name, ssn, profession}
}
