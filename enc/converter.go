package enc

type Converter interface {
	// ToEnc converts a type to Enc.
	ToEnc() *Enc
}
