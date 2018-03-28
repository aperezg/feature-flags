package identity

type Identity string

// String parsing the Identity type to string
func (i Identity) String() string {
	return string(i)
}
