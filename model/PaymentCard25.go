package model

// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
type PaymentCard25 struct {

	// Type of card, for example, credit card.
	Type *CardType1Code `xml:"Tp"`

	// Number embossed on a card that links the card to the account owner and account servicer.
	Number *Max35Text `xml:"Nb"`

	// Party entitled by a card issuer to use a card.
	HolderName *Max35Text `xml:"HldrNm"`

	// Year and month the card is available for use.
	StartDate *ISOYearMonth `xml:"StartDt,omitempty"`

	// Year and month the card expires.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`

	// Party that issues a payment card, as expressed by a numeric identification of the card issuer according to ISO/IEC 7812-1.
	CardIssuerName *Max35Text `xml:"CardIssrNm,omitempty"`

	// Party that issues a payment card, as expressed by a numeric identification of the card issuer according to ISO/IEC 7812-1.
	CardIssuerIdentification *PartyIdentification113 `xml:"CardIssrId,omitempty"`

	// Security code written on, or in, the card.
	SecurityCode *Max35Text `xml:"SctyCd,omitempty"`

	// Number distinguishing two or more payment cards with the same account number.
	SequenceNumber *Max3Text `xml:"SeqNb,omitempty"`
}

func (p *PaymentCard25) SetType(value string) {
	p.Type = (*CardType1Code)(&value)
}

func (p *PaymentCard25) SetNumber(value string) {
	p.Number = (*Max35Text)(&value)
}

func (p *PaymentCard25) SetHolderName(value string) {
	p.HolderName = (*Max35Text)(&value)
}

func (p *PaymentCard25) SetStartDate(value string) {
	p.StartDate = (*ISOYearMonth)(&value)
}

func (p *PaymentCard25) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}

func (p *PaymentCard25) SetCardIssuerName(value string) {
	p.CardIssuerName = (*Max35Text)(&value)
}

func (p *PaymentCard25) AddCardIssuerIdentification() *PartyIdentification113 {
	p.CardIssuerIdentification = new(PartyIdentification113)
	return p.CardIssuerIdentification
}

func (p *PaymentCard25) SetSecurityCode(value string) {
	p.SecurityCode = (*Max35Text)(&value)
}

func (p *PaymentCard25) SetSequenceNumber(value string) {
	p.SequenceNumber = (*Max3Text)(&value)
}
