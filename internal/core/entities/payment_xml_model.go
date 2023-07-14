package entities

import (
	"encoding/xml"
)

const (
	SUCCESS_STATUS_CODE string = "0"
)

type Return struct {
	Text string `xml:",chardata"`
}

type BpVerifyRequestResponse struct {
	XMLName xml.Name `xml:"bpVerifyRequestResponse"`
	Return  Return   `xml:"return"`
}

type BpPayRequestResponse struct {
	XMLName xml.Name `xml:"bpPayRequestResponse"`
	Return  Return   `xml:"return"`
}

type VerifyBody struct {
	XMLName                 xml.Name                `xml:"Body"`
	BpVerifyRequestResponse BpVerifyRequestResponse `xml:"bpVerifyRequestResponse"`
}

type RequestBody struct {
	XMLName              xml.Name             `xml:"Body"`
	BpPayRequestResponse BpPayRequestResponse `xml:"bpPayRequestResponse"`
}

type EnvelopeRequest struct {
	XMLName xml.Name    `xml:"Envelope"`
	Body    RequestBody `xml:"Body"`
}

type EnvelopeVerify struct {
	XMLName xml.Name   `xml:"Envelope"`
	Body    VerifyBody `xml:"Body"`
}

var RequestXMLBody string = `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
	<soapenv:Header/>
	<soapenv:Body>
		<web:bpPayRequest>
			<web:terminalId>%s</web:terminalId>
			<web:userName>%s</web:userName>
			<web:userPassword>%s</web:userPassword>
			<web:orderId>%v</web:orderId>
			<web:amount>%s</web:amount>
			<web:localDate>%s</web:localDate>
			<web:localTime>%s</web:localTime>
			<web:additionalData>%s</web:additionalData>
			<web:callBackUrl>%s</web:callBackUrl>
			<web:payerId>%s</web:payerId>
		</web:bpPayRequest>
	</soapenv:Body>
</soapenv:Envelope>`

var VerifyXMLBody string = `<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://interfaces.core.sw.bps.com/">
	<soapenv:Header/>
	<soapenv:Body>
		<web:bpVerifyRequest>
			<web:terminalId>%s</web:terminalId>
			<web:userName>%s</web:userName>
			<web:userPassword>%s</web:userPassword>
			<web:orderId>%s</web:orderId>
			<web:saleOrderId>%s</web:saleOrderId>
			<web:saleReferenceId>%s</web:saleReferenceId>
		</web:bpVerifyRequest>
	</soapenv:Body>
</soapenv:Envelope>`
