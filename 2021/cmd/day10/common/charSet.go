package common

type CharSet map[string]string

func (c CharSet) Has(s string) bool {
	return c[s] != ""
}
