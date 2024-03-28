package flags

type Doctor struct {
	Common
}


// ff:
func (b *Doctor) Default() *Doctor {
	return &Doctor{}
}
