package model

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes8 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification5 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	RegistrationForm *FormOfSecurity2Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency3Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus2Choice `xml:"PmtSts,omitempty"`

	// Indicates the direction of payment for asset or mortgage backed securities, ie, whether the repaid capital is distributed (payment direction is down) or capitalized (payment direction is up).
	PaymentDirection *PaymentDirection2Choice `xml:"PmtDrctn,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency3Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Indicates the  level of priority to claim on income and assets of the company in case of the payment of dividends and in the event of a bankruptcy, for example, ordinary/common stocks, preferred stocks, subordinated debt, etc.
	PreferenceToIncome *PreferenceToIncome2Choice `xml:"PrefToIncm,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType2Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number2Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *Number2Choice `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security by an ISIN.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification11 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes8) AddPlaceOfListing() *MarketIdentification5 {
	f.PlaceOfListing = new(MarketIdentification5)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes8) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes8) AddRegistrationForm() *FormOfSecurity2Choice {
	f.RegistrationForm = new(FormOfSecurity2Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes8) AddPaymentFrequency() *Frequency3Choice {
	f.PaymentFrequency = new(Frequency3Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes8) AddPaymentStatus() *SecuritiesPaymentStatus2Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus2Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes8) AddPaymentDirection() *PaymentDirection2Choice {
	f.PaymentDirection = new(PaymentDirection2Choice)
	return f.PaymentDirection
}

func (f *FinancialInstrumentAttributes8) AddVariableRateChangeFrequency() *Frequency3Choice {
	f.VariableRateChangeFrequency = new(Frequency3Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes8) AddPreferenceToIncome() *PreferenceToIncome2Choice {
	f.PreferenceToIncome = new(PreferenceToIncome2Choice)
	return f.PreferenceToIncome
}

func (f *FinancialInstrumentAttributes8) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes8) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes8) AddOptionType() *OptionType2Choice {
	f.OptionType = new(OptionType2Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes8) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes8) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes8) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes8) AddCouponAttachedNumber() *Number2Choice {
	f.CouponAttachedNumber = new(Number2Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes8) AddPoolNumber() *Number2Choice {
	f.PoolNumber = new(Number2Choice)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes8) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes8) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes8) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes8) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes8) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes8) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes8) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes8) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes8) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes8) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes8) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification11 {
	newValue := new(SecurityIdentification11)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func paymentInstruct(doc *Document) (paymentInstruct, error) {
	output := paymentInstruct{
		AdrLine:                doc.AdrLine.Max70Text,                                     // "<AdrLine>"
		Agt:                    doc.Agt.BranchAndFinancialInstitutionIdentification5,      // "<Agt>"
		Assgne:                 doc.Assgne.Party35Choice,                                  // "<Assgne>"
		Assgnr:                 doc.Assgnr.Party7Choice,                                   // "<Assgnr>"
		BICFI:                  doc.BICFI.BankIdentificationCode,                          // "<BICFI>"
		BizMsgIdr:              doc.BizMsgIdr.BusinessMessageIdentifier,                   // "<BizMsgIdr>"
		BldgNb:                 doc.BldgNb.Max16Text,                                      // "<BldgNb>"
		// CanonicalizationMethod https://docs.microsoft.com/en-us/dotnet/api/system.security.cryptography.xml.signedinfo.canonicalizationmethod?view=dotnet-plat-ext-6.0
		// CanonicalizationMethod https://docs.oracle.com/javase/8/docs/api/javax/xml/crypto/dsig/CanonicalizationMethod.html
		// CanonicalizationMethod https://docs.oracle.com/en/java/javase/13/docs/api/java.xml.crypto/javax/xml/crypto/dsig/CanonicalizationMethod.html
		// CanonicalizationMethod https://www.di-mgt.com.au/xmldsig-c14n.html
		CanonicalizationMethod: doc.CanonicalizationMethod.CanonicalizationMethod,         // "<CanonicalizationMethod>"
		CdtrAgt:                doc.CdtrAgt.BranchAndFinancialInstitutionIdentification4,  // "<CdtrAgt>"
		CdtTrfTxInf:            doc.CdtTrfTxInf.CreditTransferTransactionInformation11,    // "<CdtTrfTxInf>"
		ChrgBr:                 doc.ChrgBr.ChargeBearerType1Code,                          // "<ChrgBr>"
		ChrgsInf:               doc.ChrgsInf.ChargesInformation5,                          // "<ChrgsInf>"
		Conf:                   doc.Conf.ExternalInvestigationExecutionConfirmation1Co,    // "<Conf>"
		CreDt:                  doc.CreDt.ISONormalisedDateTime,                           // "<CreDt>"
		CreDtTm:                doc.CreDtTm.ISODateTime,                                   // "<CreDtTm>"
		Cretr:                  doc.Cretr.Party35Choice,                                   // "<Cretr>"
		Ctry:                   doc.Ctry.CountryCode,                                      // "<Ctry>"
		CxlRsnInf:              doc.CxlRsnInf.CancellationReasonInformation3,              // "<CxlRsnInf>"
		Dbtr:                   doc.Dbtr.PartyIdentification32,                            // "<Dbtr>"
		DbtrAgt:                doc.DbtrAgt.BranchAndFinancialInstitutionIdentification4,  // "<DbtrAgt>"
		// EndToEndId https://wiki.xmldation.com/Support/ISO20022/General_Rules/EndToEndId
		// EndToEndId https://www.jam-software.com/sepa-transfer/end-to-end-id.shtml
		// EndToEndId https://answers.sap.com/questions/12267089/element-endtoendid-not-filled-in-xml-payment-file.html
		// EndToEndId https://answers.sap.com/questions/10275743/dmee-%E2%80%93-endtoendid-with-paymantorder.html
		EndToEndId:             doc.EndToEndId.EndToEndId,                                 // "<EndToEndId>"
		Envlp:                  doc.Envlp.SupplementaryDataEnvelope1,                      // "<Envlp>"
		// FIId https://www.iso.org/iso-22000-food-safety-management.html
		// FIId https://www.qyriel.com/FullCatalogue/ISO_HEAD/out/ProtocolReport/xsd_head/head.001.001.01.xsd.html
		// FIId Financial Institution Identification
		// FIId AppHdr/Fr [Choice]
		FIId:                   doc.FIId.FinancialInstitutionIdentification,               // "<FIId>"
		// FinInstnId EPC limits the usage of Debtor Agent (DbtrAgt) and Creditor Agent CdtrAgt to allow only BIC and nothing else.
		// FinInstnId https://wiki.xmldation.com/Support/EPC/FinInstnId
		// FinInstnId https://wiki.xmldation.com/Support/RBS/CT_Rules/Global_Rules/CdtTrfTxInf%2F%2FCdtrAgt%2F%2FFinInstnId%2F%2FPstlAdr
		// FinInstnId CdtTrfTxInf/CdtrAgt/FinInstnId/PstlAdr Following fields from CreditorAgent / FinancialInstitutionIdentification / PostalAddress / Department '<CdtrAgt><FinInstnId><PstlAdr><Dept>'
		FinInstnId:             doc.FinInstnId.FinancialInstitutionIdentification7,        // "<FinInstnId>"
		FIToFICstmrCdtTrf:      doc.FIToFICstmrCdtTrf.FIToFICustomerCreditTransferV02,     // "<FIToFICstmrCdtTrf>"
		FIToFIPmtCxlReq:        doc.FIToFIPmtCxlReq.FIToFIPaymentCancellationRequestV01,   // "<FIToFIPmtCxlReq>"
		FIToFIPmtStsRpt:        doc.FIToFIPmtStsRpt.FIToFIPaymentStatusReportV03,          // "<FIToFIPmtStsRpt>"
		Fr:                     doc.Fr.Party9Choice,                                       // "<Fr>"
		GrpHdr:                 doc.GrpHdr.GroupHeader33,                                  // "<GrpHdr>"
		Id:                     doc.Id.Max35Text,                                          // "<Id>"
		InstdAgt:               doc.InstdAgt.BranchAndFinancialInstitutionIdentification4, // "<InstdAgt>"
		InstdAmt:               doc.InstdAmt.ActiveOrHistoricCurrencyAndAmount,            // "<InstdAmt>"
		// InstgAgt https://www.swift.com/swift-resource/248686/download
		// InstgAgt https://community.oracle.com/tech/developers/discussion/4327286/ora-00904-error-outer-join-19c
		// InstgAgt https://www.nacha.org/content/iso-20022-ach-mapping-guide
		// InstgAgt https://www.iso20022.org/sites/default/files/documents/D7/ISO20022_RTPG_pacs00800106_July_2017_v1_1.pdf
		InstgAgt:               doc.InstgAgt.BranchAndFinancialInstitutionIdentification4, // "<InstgAgt>"
		// InstrId https://wiki.xmldation.com/Support/ISO20022/General_Rules/InstrId
		// InstrId https://www.mathworks.com/help/instrument/instrid.html
		// InstrId https://wiki.xmldation.com/Support/Sampo/InstrId
		// InstrId https://docs.oracle.com/cd/E16582_01/doc.91/e15104/fields_sepa_pay_file_appx.htm#EOAEL01692
		InstrId:                doc.InstrId.InstructionIdentification,                     // "<InstrId>"
		// IntrBkSttlmAmt https://www.ecb.europa.eu/paym/groups/shared/docs/75299-tips-_cg_2017-09-28_presentation_udfs.pdf
		// IntrBkSttlmAmt https://wiki.xmldation.com/General_Information/ISO_20022/Difference_between_InstdAmt_and_EqvtAmt
		// IntrBkSttlmAmt https://www.iotafinance.com/en/SWIFT-ISO15022-Message-type-MT202-COV.html
		// IntrBkSttlmAmt https://www.bnymellon.com/content/dam/bnymellon/documents/pdf/iso-20022/Module%201_September%202020_Demystifying%20ISO20022.pdf
		IntrBkSttlmAmt:         doc.IntrBkSttlmAmt.ActiveOrHistoricCurrencyAndAmount,      // "<IntrBkSttlmAmt>"
		// IntrBkSttlmDt https://www.citibank.com/tts/sa/flippingbook/2021/ISO-20022-Citi-Mini-Series-and-Reference-Guide-Part-2/10/
		// IntrBkSttlmDt https://www.citibank.com/tts/sa/flippingbook/2021/ISO-20022-Citi-Mini-Series-and-Reference-Guide-Part-2/26/
		// IntrBkSttlmDt https://www.paymentstandards.ch/dam/mapping-rules_pacs008_esr.pdf
		// IntrBkSttlmDt https://www.payments.ca/sites/default/files/part_a_of_5_fitofi_customer_credit_transfers.pdf
		IntrBkSttlmDt:          doc.IntrBkSttlmDt.InterbankSettlementDate,                 // "<IntrBkSttlmDt>"
		Issr:                   doc.Issr.Issuer,                                           // "<Issr>"
		Justfn:                 doc.Justfn.CaseForwardingNotification3Code,                // "<Justfn>"
		KeyInfo:                doc.KeyInfo.KeyInfo,                                       // "<KeyInfo>"
		Mod:                    doc.Mod.RequestedModification7,                            // "<Mod>"
		MsgDefIdr:              doc.MsgDefIdr.MessageDefinitionIdentifier,                 // "<MsgDefIdr>"
		MsgId:                  doc.MsgId.MessageIdentification,                           // "<MsgId>"
		MssngOrIncrrctInf:      doc.MssngOrIncrrctInf.MissingOrIncorrectInformation3,      // "<MssngOrIncrrctInf>"
		// NbOfTxs https://wiki.xmldation.com/Support/RBS/DD_Rules/Global_Rules/NbOfTxs
		// NbOfTxs https://support.oracle.com/knowledge/Oracle%20E-Business%20Suite/1571592_1.html
		// NbOfTxs https://docs.oracle.com/cd/E16582_01/doc.91/e15104/fields_sepa_pay_file_appx.htm#EOAEL01692
		// NbOfTxs https://wiki.xmldation.com/Support/ISO20022/General_Rules/NbOfTxs
		NbOfTxs:                doc.NbOfTxs.Max15NumericText,                              // "<NbOfTxs>"
		NtfctnOfCaseAssgnmt:    doc.NtfctnOfCaseAssgnmt.NotificationOfCaseAssignmentV05,   // "<NtfctnOfCaseAssgnmt>"
		OrgnlCreDtTm:           doc.OrgnlCreDtTm.ISODateTime,                              // "<OrgnlCreDtTm>"
		// OrgnlEndToEndId https://wiki.xmldation.com/Support/ISO20022/General_Rules/EndToEndId
		// OrgnlEndToEndId https://paymentcomponents.atlassian.net/wiki/spaces/AH/pages/479428560/Sample+SEPA+messages+for+Testing
		// OrgnlEndToEndId https://answers.sap.com/questions/10275743/dmee-%E2%80%93-endtoendid-with-paymantorder.html
		// OrgnlEndToEndId https://blogs.sap.com/2021/07/30/pain.002-payment-rejections-processing-via-rfebka00/
		// OrgnlEndToEndId https://docs.crbcos.com/unicorncrb/docs/unicorn-output-files
		OrgnlEndToEndId:        doc.OrgnlEndToEndId.OriginalEndToEndIdentification,        // "<OrgnlEndToEndId>"
		// OrgnlGrpInf https://www.payments.ca/sites/default/files/part_c_of_5_payment_return.pdf
		// OrgnlGrpInf https://wiki.xmldation.com/Support/Nordea/CancellationRequest/Cancellation_Request_%2f%2f_CancellationReason2Code
		// OrgnlGrpInf https://www.iso20022.org/sites/default/files/documents/D7/Pacs004%20Real%20Time%20Payment%20Sep2018_v0.1.pdf
		// OrgnlGrpInf https://www.nacha.org/content/iso-20022-ach-mapping-guide
		// OrgnlGrpInf https://www.iso20022.org/sites/default/files/documents/D7/ISO20022_RTPG_pacs00200108_July_2017_v1_1.pdf
		OrgnlGrpInf:            doc.OrgnlGrpInf.OriginalGroupInformation3,                 // "<OrgnlGrpInf>"
		OrgnlGrpInfAndCxl:      doc.OrgnlGrpInfAndCxl.OriginalGroupInformation23,          // "<OrgnlGrpInfAndCxl>"
		OrgnlGrpInfAndSts:      doc.OrgnlGrpInfAndSts.OriginalGroupInformation20,          // "<OrgnlGrpInfAndSts>"
		OrgnlInstdAmt:          doc.OrgnlInstdAmt.OriginalInstructedAmount,                // "<OrgnlInstdAmt>"
		// OrgnlInstrId https://www.iso20022.org/sites/default/files/documents/D7/Pacs004%20Real%20Time%20Payment%20Sep2018_v0.1.pdf
		// OrgnlInstrId https://paymentcomponents.atlassian.net/wiki/spaces/AH/pages/479428560/Sample+SEPA+messages+for+Testing
		// OrgnlInstrId https://stackoverflow.com/questions/65199828/parsing-xml-in-c-sharp-with-xsd-file
		// OrgnlInstrId https://github.com/FasterXML/jackson-dataformat-xml/issues/217
		OrgnlInstrId:           doc.OrgnlInstrId.OriginalInstructionIdentification,        // "<OrgnlInstrId>"
		OrgnlIntrBkSttlmAmt:    doc.OrgnlIntrBkSttlmAmt.ActiveOrHistoricCurrencyAndAmount, // "<OrgnlIntrBkSttlmAmt>"
		OrgnlMsgId:             doc.OrgnlMsgId.OriginalMessageIdentification,              // "<OrgnlMsgId>"
		OrgnlMsgNmId:           doc.OrgnlMsgNmId.OriginalMessageNameIdentification,        // "<OrgnlMsgNmId>"
		OrgnlTxId:              doc.OrgnlTxId.OriginalTransactionIdentification,           // "<OrgnlTxId>"
		OrgnlTxRef:             doc.OrgnlTxRef.OriginalTransactionReference13,             // "<OrgnlTxRef>"
		Orgtr:                  doc.Orgtr.PartyIdentification32,                           // "<Orgtr>"
		PlcAndNm:               doc.PlcAndNm.PlcAndNm,                                     // "<PlcAndNm>"
		PmtTpInf:               doc.PmtTpInf.PaymentTypeInformation21,                     // "<PmtTpInf>"
		PstCd:                  doc.PstCd.Max16Text,                                       // "<PstCd>"
		PstlAdr:                doc.PstlAdr.PostalAddress6,                                // "<PstlAdr>"
		ReqToModfyPmt:          doc.ReqToModfyPmt.RequestToModifyPaymentV05,               // "<ReqToModfyPmt>"
		RsltnOfInvstgtn:        doc.RsltnOfInvstgtn.ResolutionOfInvestigationV08,          // "<RsltnOfInvstgtn>"
		RtrdInstdAmt:           doc.RtrdInstdAmt.ActiveOrHistoricCurrencyAndAmount,        // "<RtrdInstdAmt>"
		RtrdIntrBkSttlmAmt:     doc.RtrdIntrBkSttlmAmt.ActiveCurrencyAndAmount,            // "<RtrdIntrBkSttlmAmt>"
		RtrId:                  doc.RtrId.Max35Text,                                       // "<RtrId>"
		RtrRsnInf:              doc.RtrRsnInf.ReturnReasonInformation9,                    // "<RtrRsnInf>"
		Signature:              doc.Signature.Signature,                                   // "<Signature>"
		SignatureMethod:        doc.SignatureMethod.SignatureMethod,                       // "<SignatureMethod>"
		SplmtryData:            doc.SplmtryData.SupplementaryData1,                        // "<SplmtryData>"
		StrtNm:                 doc.StrtNm.Max70Text,                                      // "<StrtNm>"
		SttlmAcct:              doc.SttlmAcct.CashAccount16,                               // "<SttlmAcct>"
		SttlmInf:               doc.SttlmInf.SettlementInformation13,                      // "<SttlmInf>"
		SttlmMtd:               doc.SttlmMtd.SettlementMethod1Code,                        // "<SttlmMtd>"
		SvcLvl:                 doc.SvcLvl.ServiceLevel8Choice,                            // "<SvcLvl>"
		To:                     doc.To.Party9Choice,                                       // "<To>"
		TwnNm:                  doc.TwnNm.Max35Text,                                       // "<TwnNm>"
		TxId:                   doc.TxId.TransactionIdentification,                        // "<TxId>"
		TxInfAndSts:            doc.TxInfAndSts.PaymentTransactionInformation26,           // "<TxInfAndSts>"
		TxSts:                  doc.TxSts.TransactionIndividualStatus3Code,                // "<TxSts>"
		UblToApply:             doc.UblToApply.UnableToApplyV07,                           // "<UblToApply>"
		UltmtCdtr:              doc.UltmtCdtr.PartyIdentification32,                       // "<UltmtCdtr>"
		Undrlyg:                doc.Undrlyg.UnderlyingTransaction4Choice,                  // "<Undrlyg>"
		X509Data:               doc.X509Data.KeyInfoX509Data,                              // "<X509Data>"
		XchgRate:               doc.XchgRate.BaseOneRate,                                  // "<XchgRate>"
	}

func (f *FinancialInstrumentAttributes8) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}
