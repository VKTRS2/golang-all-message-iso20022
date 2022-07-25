# Repo Golang for RTP systems

**Release 0.7.3**

## A secondary git-repository for the RTP system.
Thank you to everyone who participated in the creation of these go elements and projects.

# ISO20022

This repository contains a full set of Golang parse to the ISO-20022 data catalogs [ISO-20022 specifications](https://www.iso20022.org/full_catalogue.page)

## Usage

See `examples/` directory for an example of usage

```go
import (
	"encoding/xml"
	"github.com/yudaprama/iso20022/pacs"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	messages, err := ioutil.ReadFile("./example-message.xml")
	if err != nil {
		log.Fatalf("Unable to read file:  %v", err)
		os.Exit(1)
	}

	var messageParsed pacs.Document00800106
	err = xml.Unmarshal(messages, &messageParsed)
	if err != nil {
		log.Fatalf("Unable to parse file:  %v", err)
		os.Exit(1)
	}
	
	log.Printf("Interbank Settlement Date:  %v", messageParsed.Message.GroupHeader.InterbankSettlementDate)
}

```

## Message Catalogs

Message types covers ISO-20022 messages:

* [ACMT - Account Management](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/acmt)
* [ADMI - Administration](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/admi)
* [AUTH - Authorities](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/auth)
* [CAAA - Acceptor to Acquirer Card Transactions](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/caaa)
* [CAAM - ATM Management](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/caam)
* [CAIN - Acquirer to Issuer Card Transactions](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/cain)
* [CAMT - Cash Management](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/camt)
* [CATM - Terminal Management](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/catm)
* [CATP - ATM Card Transactions](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/catp)
* [COLR - Collateral Management](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/colr)
* [FXTR - Foreign Exchange Trade](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/fxtr)
* [HEAD - Business Application Header](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/head)
* [PACS - Payments Clearing and Settlement](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/pacs)
* [PAIN - Payments Initiation](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/pain)
* [REDA - Reference Data](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/reda)
* [REMT - Payments Remittance Advice](https://github.com/yudaprama/iso20022/tree/master/remt)
* [SECL - Securities Clearing](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/remt)
* [SEEV - Securities Events](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/seev)
* [SEMT - Securities Management](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/semt)
* [SESE - Securities Settlement](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/sese)
* [SETR - Securities Trade](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/setr)
* [SUPL - Supplementary Data](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/supl)
* [TREA - Treasury](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/trea)
* [TSIN - Trade Services Initiation](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/tsin)
* [TSMT - Trade Services Management](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/tsmt)
* [TSRV - Trade Services](https://github.com/VKTRS2/golang-all-message-iso20022/tree/main/tsrv)
