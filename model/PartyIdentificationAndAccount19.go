package model

// Party and account details.
type PartyIdentificationAndAccount19 struct {

	// Identification of the party.
	Identification *PartyIdentification10Choice `xml:"Id,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount19) AddIdentification() *PartyIdentification10Choice {
	p.Identification = new(PartyIdentification10Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount19) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount19) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentificationAndAccount19) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount19) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount19) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}
