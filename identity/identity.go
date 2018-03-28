package identity

type Identity string

func (i Identity) String() string {
	return string(i)
}
