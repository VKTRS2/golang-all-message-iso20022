package model

// Parameters applied to the settlement of a security transfer.
type Transfer33 struct {

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference7 `xml:"CtrPtyRef,omitempty"`

	// Identifies the business process in which the actors are involved. This is important to trigger the right business process, according to the market business model, which may require matching instructions in a CSD environment (double leg process) or not (single leg process).
	BusinessFlowType *BusinessFlowType1Code `xml:"BizFlowTp,omitempty"`

	// Date and time at which the transfer was executed.
	EffectiveTransferDate *DateAndDateTimeChoice `xml:"FctvTrfDt,omitempty"`

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Date and time at which a transaction is completed and cleared, that is, securities are delivered.
	TradeDate *DateAndDateTimeChoice `xml:"TradDt,omitempty"`

	// Identifies in which date the investor signed the transfer order form.
	TransferOrderDateForm *ISODate `xml:"TrfOrdrDtForm,omitempty"`

	// Specifies the reason for the assets transfer.
	TransferReason *TransferReason1 `xml:"TrfRsn,omitempty"`

	// Identifies whether or not saving plan or withdrawal or switch plan are included in the holdings.
	HoldingsPlanType []*HoldingsPlanType1Code `xml:"HldgsPlanTp,omitempty"`

	// Information related to the financial instrument received.
	FinancialInstrumentDetails *FinancialInstrument49 `xml:"FinInstrmDtls"`

	// Total quantity of securities settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Information about the units to be transferred.
	UnitsDetails []*Unit6 `xml:"UnitsDtls,omitempty"`

	// Weighted average price of the units in the account before the transfer was executed.
	AveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"AvrgPric,omitempty"`

	// Weighted average price of the units in the account after the transfer was executed.
	NewAveragePrice *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"NewAvrgPric,omitempty"`

	// Trade date of the average weighted data of units in the account before the transfer was executed.
	AverageDate *ISODate `xml:"AvrgDt,omitempty"`

	// Trade date of the average weighted data of units in the account after the transfer was executed.
	NewAverageDate *ISODate `xml:"NewAvrgDt,omitempty"`

	// Identifies the currency to be used to transfer the holdings.
	TransferCurrency *ActiveOrHistoricCurrencyCode `xml:"TrfCcy,omitempty"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount125 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount125 `xml:"DlvrgAgtDtls,omitempty"`

	// Specifies how the payment of charges, taxes and commissions as a result of the transfer is covered, that is, whether by cash or the redemption of units.
	TransferExpensesPaymentType *ChargePaymentMethod1Choice `xml:"TrfExpnssPmtTp,omitempty"`
}

func (t *Transfer33) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *Transfer33) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer33) AddClientReference() *AdditionalReference7 {
	t.ClientReference = new(AdditionalReference7)
	return t.ClientReference
}

func (t *Transfer33) AddCounterpartyReference() *AdditionalReference7 {
	t.CounterpartyReference = new(AdditionalReference7)
	return t.CounterpartyReference
}

func (t *Transfer33) SetBusinessFlowType(value string) {
	t.BusinessFlowType = (*BusinessFlowType1Code)(&value)
}

func (t *Transfer33) AddEffectiveTransferDate() *DateAndDateTimeChoice {
	t.EffectiveTransferDate = new(DateAndDateTimeChoice)
	return t.EffectiveTransferDate
}

func (t *Transfer33) SetRequestedSettlementDate(value string) {
	t.RequestedSettlementDate = (*ISODate)(&value)
}

func (t *Transfer33) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	t.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return t.EffectiveSettlementDate
}

func (t *Transfer33) AddTradeDate() *DateAndDateTimeChoice {
	t.TradeDate = new(DateAndDateTimeChoice)
	return t.TradeDate
}

func (t *Transfer33) SetTransferOrderDateForm(value string) {
	t.TransferOrderDateForm = (*ISODate)(&value)
}

func (t *Transfer33) AddTransferReason() *TransferReason1 {
	t.TransferReason = new(TransferReason1)
	return t.TransferReason
}

func (t *Transfer33) AddHoldingsPlanType(value string) {
	t.HoldingsPlanType = append(t.HoldingsPlanType, (*HoldingsPlanType1Code)(&value))
}

func (t *Transfer33) AddFinancialInstrumentDetails() *FinancialInstrument49 {
	t.FinancialInstrumentDetails = new(FinancialInstrument49)
	return t.FinancialInstrumentDetails
}

func (t *Transfer33) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer33) AddUnitsDetails() *Unit6 {
	newValue := new(Unit6)
	t.UnitsDetails = append(t.UnitsDetails, newValue)
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

func (t *Transfer33) SetAveragePrice(value, currency string) {
	t.AveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer33) SetNewAveragePrice(value, currency string) {
	t.NewAveragePrice = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *Transfer33) SetAverageDate(value string) {
	t.AverageDate = (*ISODate)(&value)
}

func (t *Transfer33) SetNewAverageDate(value string) {
	t.NewAverageDate = (*ISODate)(&value)
}

func (t *Transfer33) SetTransferCurrency(value string) {
	t.TransferCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (t *Transfer33) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer33) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}

func (t *Transfer33) AddReceivingAgentDetails() *PartyIdentificationAndAccount125 {
	t.ReceivingAgentDetails = new(PartyIdentificationAndAccount125)
	return t.ReceivingAgentDetails
}

func (t *Transfer33) AddDeliveringAgentDetails() *PartyIdentificationAndAccount125 {
	t.DeliveringAgentDetails = new(PartyIdentificationAndAccount125)
	return t.DeliveringAgentDetails
}

func (t *Transfer33) AddTransferExpensesPaymentType() *ChargePaymentMethod1Choice {
	t.TransferExpensesPaymentType = new(ChargePaymentMethod1Choice)
	return t.TransferExpensesPaymentType
}
