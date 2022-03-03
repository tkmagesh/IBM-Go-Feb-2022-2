package models

import "fmt"

func (p Product) WhoAmI() string {
	return fmt.Sprintf("I am a %q", p.Name)
}
