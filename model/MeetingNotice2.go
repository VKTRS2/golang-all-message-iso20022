package model

// Information about the shareholders meeting, specifying the participation requirements and the voting procedures. Alternatively, it may indicate where such information may be obtained.
type MeetingNotice2 struct {

	// Identification assigned to a general meeting by the party notifying the meeting. It must be unique for the party notifying the meeting.
	MeetingIdentification *Max35Text `xml:"MtgId,omitempty"`

	// Identification assigned to a meeting by the issuer. It must be unique for the issuer.
	IssuerMeetingIdentification *Max35Text `xml:"IssrMtgId,omitempty"`

	// Specifies the type of security holders meeting.
	Type *MeetingType2Code `xml:"Tp"`

	// Classifies the type of meeting.
	Classification *MeetingTypeClassification1Code `xml:"Clssfctn,omitempty"`

	// This code can be used in case another meeting classifications is required.
	ExtendedClassification *Extended350Code `xml:"XtndedClssfctn,omitempty"`

	// Official meeting announcement date.
	AnnouncementDate *ISODate `xml:"AnncmntDt,omitempty"`

	// Indicates whether physical participation to a meeting is required in order to be allowed to vote.
	AttendanceRequired *YesNoIndicator `xml:"AttndncReqrd"`

	// Indicates how to order the attendance card or to give notice of attendance.
	AttendanceConfirmationInformation *Max350Text `xml:"AttndncConfInf,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of their intention to participate in a meeting. This deadline is set by an intermediary.
	AttendanceConfirmationDeadline *DateFormat2Choice `xml:"AttndncConfDdln,omitempty"`

	// Date and time by which the beneficial owner or agent must notify of their intention to participate in a meeting (STP mode). This deadline is set by an intermediary.
	AttendanceConfirmationSTPDeadline *DateFormat2Choice `xml:"AttndncConfSTPDdln,omitempty"`

	// Date and time by which the attendance to the meeting should be confirmed. This deadline is set by the issuer.
	AttendanceConfirmationMarketDeadline *DateFormat2Choice `xml:"AttndncConfMktDdln,omitempty"`

	// Address to use over the www (HTTP) service where addtional information on the meeting may be found.
	AdditionalDocumentationURLAddress *Max256Text `xml:"AddtlDcmnttnURLAdr,omitempty"`

	// Date and time by which security holders can propose amendments or new resolutions. This deadline is set by an intermediary.
	ResolutionProposalDeadline *DateFormat2Choice `xml:"RsltnPrpslDdln,omitempty"`

	// Date and time by which security holders can propose amendments or new resolutions. This deadline is set by the issuer.
	ResolutionProposalMarketDeadline *DateFormat2Choice `xml:"RsltnPrpslMktDdln,omitempty"`

	// Specifies the minimum stake in share capital or cash value or number of security holders required to table resolutions.
	ResolutionProposalThreshold *Max350Text `xml:"RsltnPrpslThrshld,omitempty"`

	// Specifies the minimum stake in share capital or cash value or number of security holders required to table resolutions. This minimum is expressed as a percentage.
	ResolutionProposalThresholdPercentage *PercentageRate `xml:"RsltnPrpslThrshldPctg,omitempty"`

	// Number of securities admitted to the vote, expressed as an amount and a currency.
	TotalNumberOfSecuritiesOutstanding *CurrencyAndAmount `xml:"TtlNbOfSctiesOutsdng,omitempty"`

	// Number of rights admitted to the vote.
	TotalNumberOfVotingRights *Number `xml:"TtlNbOfVtngRghts,omitempty"`

	// Address where the information on the proxy should be sent.
	ProxyAppointmentNotificationAddress *PostalAddress1 `xml:"PrxyAppntmntNtfctnAdr,omitempty"`

	// Indicates that no proxy is allowed for a meeting.
	ProxyNotAllowed *ProxyNotAllowedCode `xml:"PrxyNotAllwd,omitempty"`

	// Specifies the elements required to assign a proxy.
	Proxy *ProxyAppointmentInformation1 `xml:"Prxy,omitempty"`

	// Contact person at the party organising the meeting, at the issuer or at an intermediary.
	ContactPersonDetails []*MeetingContactPerson1 `xml:"CtctPrsnDtls,omitempty"`

	// Date on which a company publishes the results of its meeting.
	ResultPublicationDate *DateFormat3Choice `xml:"RsltPblctnDt,omitempty"`
}

func (m *MeetingNotice2) SetMeetingIdentification(value string) {
	m.MeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice2) SetIssuerMeetingIdentification(value string) {
	m.IssuerMeetingIdentification = (*Max35Text)(&value)
}

func (m *MeetingNotice2) SetType(value string) {
	m.Type = (*MeetingType2Code)(&value)
}

func (m *MeetingNotice2) SetClassification(value string) {
	m.Classification = (*MeetingTypeClassification1Code)(&value)
}

func (m *MeetingNotice2) SetExtendedClassification(value string) {
	m.ExtendedClassification = (*Extended350Code)(&value)
}

func (m *MeetingNotice2) SetAnnouncementDate(value string) {
	m.AnnouncementDate = (*ISODate)(&value)
}

func (m *MeetingNotice2) SetAttendanceRequired(value string) {
	m.AttendanceRequired = (*YesNoIndicator)(&value)
}

func (m *MeetingNotice2) SetAttendanceConfirmationInformation(value string) {
	m.AttendanceConfirmationInformation = (*Max350Text)(&value)
}

func (m *MeetingNotice2) AddAttendanceConfirmationDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationDeadline
}

func (m *MeetingNotice2) AddAttendanceConfirmationSTPDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationSTPDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationSTPDeadline
}

func (m *MeetingNotice2) AddAttendanceConfirmationMarketDeadline() *DateFormat2Choice {
	m.AttendanceConfirmationMarketDeadline = new(DateFormat2Choice)
	return m.AttendanceConfirmationMarketDeadline
}

func (m *MeetingNotice2) SetAdditionalDocumentationURLAddress(value string) {
	m.AdditionalDocumentationURLAddress = (*Max256Text)(&value)
}

func (m *MeetingNotice2) AddResolutionProposalDeadline() *DateFormat2Choice {
	m.ResolutionProposalDeadline = new(DateFormat2Choice)
	return m.ResolutionProposalDeadline
}

func (m *MeetingNotice2) AddResolutionProposalMarketDeadline() *DateFormat2Choice {
	m.ResolutionProposalMarketDeadline = new(DateFormat2Choice)
	return m.ResolutionProposalMarketDeadline
}

func (m *MeetingNotice2) SetResolutionProposalThreshold(value string) {
	m.ResolutionProposalThreshold = (*Max350Text)(&value)
}

func (m *MeetingNotice2) SetResolutionProposalThresholdPercentage(value string) {
	m.ResolutionProposalThresholdPercentage = (*PercentageRate)(&value)
}

func (m *MeetingNotice2) SetTotalNumberOfSecuritiesOutstanding(value, currency string) {
	m.TotalNumberOfSecuritiesOutstanding = NewCurrencyAndAmount(value, currency)
}

func (m *MeetingNotice2) SetTotalNumberOfVotingRights(value string) {
	m.TotalNumberOfVotingRights = (*Number)(&value)
}

func (m *MeetingNotice2) AddProxyAppointmentNotificationAddress() *PostalAddress1 {
	m.ProxyAppointmentNotificationAddress = new(PostalAddress1)
	return m.ProxyAppointmentNotificationAddress
}

func (m *MeetingNotice2) SetProxyNotAllowed(value string) {
	m.ProxyNotAllowed = (*ProxyNotAllowedCode)(&value)
}

func (m *MeetingNotice2) AddProxy() *ProxyAppointmentInformation1 {
	m.Proxy = new(ProxyAppointmentInformation1)
	return m.Proxy
}

func (m *MeetingNotice2) AddContactPersonDetails() *MeetingContactPerson1 {
	newValue := new(MeetingContactPerson1)
	m.ContactPersonDetails = append(m.ContactPersonDetails, newValue)
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

func (m *MeetingNotice2) AddResultPublicationDate() *DateFormat3Choice {
	m.ResultPublicationDate = new(DateFormat3Choice)
	return m.ResultPublicationDate
}
