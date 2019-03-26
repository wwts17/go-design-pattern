package main

import "fmt"

type HotelBoutique struct{}

type HotelBoutiqueProxy struct {
	subject *HotelBoutique
}

func (s *HotelBoutique) Book() {
	fmt.Println("Booking done on external site")
}

func (p *HotelBoutiqueProxy) Book() {
	if p.subject == nil {
		p.subject = new(HotelBoutique)
	}
	fmt.Println("Proxy Delegating Booking call")

	p.subject.Book()
}

func main() {
	h := new(HotelBoutiqueProxy)
	h.Book()
}
