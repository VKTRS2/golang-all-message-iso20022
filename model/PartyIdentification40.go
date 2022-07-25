package model

// Set of elements used to identify a person or an organisation.
type PartyIdentification40 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *PersonIdentification5 `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification40) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification40) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification40) AddIdentification() *PersonIdentification5 {
	p.Identification = new(PersonIdentification5)
	return p.Identification
}

func (p *PartyIdentification40) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification40) AddContactDetails() *ContactDetails2 {
	p.ContactDetails = new(ContactDetails2)
	return p.ContactDetails
}
