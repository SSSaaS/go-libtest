package tests

import (
)

func (t Tests) XTestCreateCombine() {
	t.XChain([]string{"1", "4", "Hello World!"}, "Combine", "Create", "Hello World!")
	t.XChain([]string{"1", "4", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}, "Combine", "Create", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	t.Success()
}
