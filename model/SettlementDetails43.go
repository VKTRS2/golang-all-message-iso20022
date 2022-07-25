package model

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails43 struct {

	// Indicates whether the trade is to be settled as principal or netted off against another trade.
	SettlementTransactionType *SettlementTransactionType1Choice `xml:"SttlmTxTp"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric3Choice `xml:"Prty,omitempty"`

	// Specifies if the Electronic Trade Confirmation (ETC) service provider is to generate or not a settlement instruction.
	SettlementInstructionGeneration *SettlementInstructionGeneration1Choice `xml:"SttlmInstrGnrtn,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition11Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership3Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade3Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility3Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem3Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType9Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction3Choice `xml:"FxStgInstr,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide3Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility3Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration6Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType11Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction3Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS3Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity3Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod3Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty3Choice `xml:"TaxCpcty,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification38 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking3Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing5Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee3Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed3Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails43) AddSettlementTransactionType() *SettlementTransactionType1Choice {
	s.SettlementTransactionType = new(SettlementTransactionType1Choice)
	return s.SettlementTransactionType
}

func (s *SettlementDetails43) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddPriority() *PriorityNumeric3Choice {
	s.Priority = new(PriorityNumeric3Choice)
	return s.Priority
}

func (s *SettlementDetails43) AddSettlementInstructionGeneration() *SettlementInstructionGeneration1Choice {
	s.SettlementInstructionGeneration = new(SettlementInstructionGeneration1Choice)
	return s.SettlementInstructionGeneration
}

func (s *SettlementDetails43) AddSettlementTransactionCondition() *SettlementTransactionCondition11Choice {
	newValue := new(SettlementTransactionCondition11Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
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

func (s *SettlementDetails43) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddBeneficialOwnership() *BeneficialOwnership3Choice {
	s.BeneficialOwnership = new(BeneficialOwnership3Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails43) AddBlockTrade() *BlockTrade3Choice {
	s.BlockTrade = new(BlockTrade3Choice)
	return s.BlockTrade
}

func (s *SettlementDetails43) AddCCPEligibility() *CentralCounterPartyEligibility3Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility3Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails43) AddCashClearingSystem() *CashSettlementSystem3Choice {
	s.CashClearingSystem = new(CashSettlementSystem3Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails43) AddExposureType() *ExposureType9Choice {
	s.ExposureType = new(ExposureType9Choice)
	return s.ExposureType
}

func (s *SettlementDetails43) AddFXStandingInstruction() *FXStandingInstruction3Choice {
	s.FXStandingInstruction = new(FXStandingInstruction3Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails43) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SettlementDetails43) AddMarketClientSide() *MarketClientSide3Choice {
	s.MarketClientSide = new(MarketClientSide3Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails43) AddNettingEligibility() *NettingEligibility3Choice {
	s.NettingEligibility = new(NettingEligibility3Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails43) AddRegistration() *Registration6Choice {
	s.Registration = new(Registration6Choice)
	return s.Registration
}

func (s *SettlementDetails43) AddRepurchaseType() *RepurchaseType11Choice {
	s.RepurchaseType = new(RepurchaseType11Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails43) AddLegalRestrictions() *Restriction3Choice {
	s.LegalRestrictions = new(Restriction3Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails43) AddSecuritiesRTGS() *SecuritiesRTGS3Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS3Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails43) AddSettlingCapacity() *SettlingCapacity3Choice {
	s.SettlingCapacity = new(SettlingCapacity3Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails43) AddSettlementSystemMethod() *SettlementSystemMethod3Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod3Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails43) AddTaxCapacity() *TaxCapacityParty3Choice {
	s.TaxCapacity = new(TaxCapacityParty3Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails43) SetStampDutyIndicator(value string) {
	s.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddStampDutyTaxBasis() *GenericIdentification38 {
	s.StampDutyTaxBasis = new(GenericIdentification38)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails43) AddTracking() *Tracking3Choice {
	s.Tracking = new(Tracking3Choice)
	return s.Tracking
}

func (s *SettlementDetails43) AddAutomaticBorrowing() *AutomaticBorrowing5Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing5Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails43) AddLetterOfGuarantee() *LetterOfGuarantee3Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee3Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails43) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddModificationCancellationAllowed() *ModificationCancellationAllowed3Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed3Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails43) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}
