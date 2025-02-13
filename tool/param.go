package tool

import "fmt"

type Param struct {
	Encrypt   string
	Signature string
}

func (p *Param) ToString() string {
	return fmt.Sprintf("Param{encrypt='%s', signature='%s'}", p.Encrypt, p.Signature)
}
