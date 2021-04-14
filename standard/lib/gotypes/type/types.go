package types

import (
	"fmt"
)

type Address = string

type Home struct {
	Size    float32
	Address Address
}

func (h *Home) SetHomeDetails(size float32, addr Address) {
	h.Address = addr
	h.Size = size
}

func (h *Home) GetHomeDetails() string {
	details := fmt.Sprintf("Home Address: %s, Home Size: %f", h.Address, h.Size)
	return details
}
