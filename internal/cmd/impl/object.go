package main

import "encoding/xml"

// Schema was generated 2024-08-16 14:57:38 by https://xml-to-go.github.io/ in Ukraine.
type ObjectSchema struct {
	XMLName            xml.Name `xml:"schema"`
	Text               string   `xml:",chardata"`
	Xs                 string   `xml:"xs,attr"`
	Xmlns              string   `xml:"xmlns,attr"`
	ElementFormDefault string   `xml:"elementFormDefault,attr"`
	TargetNamespace    string   `xml:"targetNamespace,attr"`
	Include            []struct {
		Text           string `xml:",chardata"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"include"`
	SimpleType []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		Restriction struct {
			Text        string `xml:",chardata"`
			Base        string `xml:"base,attr"`
			Enumeration []struct {
				Text       string `xml:",chardata"`
				Value      string `xml:"value,attr"`
				Annotation struct {
					Text          string `xml:",chardata"`
					Documentation struct {
						Text string `xml:",chardata"`
					} `xml:"documentation"`
				} `xml:"annotation"`
			} `xml:"enumeration"`
		} `xml:"restriction"`
	} `xml:"simpleType"`
	ComplexType []struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name,attr"`
		Annotation struct {
			Text          string `xml:",chardata"`
			Documentation struct {
				Text string `xml:",chardata"`
			} `xml:"documentation"`
		} `xml:"annotation"`
		ComplexContent struct {
			Text      string `xml:",chardata"`
			Extension struct {
				Text     string `xml:",chardata"`
				Base     string `xml:"base,attr"`
				Sequence struct {
					Text    string `xml:",chardata"`
					Element []struct {
						Text      string `xml:",chardata"`
						MaxOccurs string `xml:"maxOccurs,attr"`
						MinOccurs string `xml:"minOccurs,attr"`
						Name      string `xml:"name,attr"`
						Type      string `xml:"type,attr"`
					} `xml:"element"`
					Group struct {
						Text      string `xml:",chardata"`
						MaxOccurs string `xml:"maxOccurs,attr"`
						MinOccurs string `xml:"minOccurs,attr"`
						Ref       string `xml:"ref,attr"`
					} `xml:"group"`
					Choice struct {
						Text      string `xml:",chardata"`
						MaxOccurs string `xml:"maxOccurs,attr"`
						MinOccurs string `xml:"minOccurs,attr"`
						Element   []struct {
							Text      string `xml:",chardata"`
							MaxOccurs string `xml:"maxOccurs,attr"`
							MinOccurs string `xml:"minOccurs,attr"`
							Name      string `xml:"name,attr"`
							Type      string `xml:"type,attr"`
						} `xml:"element"`
					} `xml:"choice"`
				} `xml:"sequence"`
				Attribute []struct {
					Text       string `xml:",chardata"`
					Name       string `xml:"name,attr"`
					Type       string `xml:"type,attr"`
					Use        string `xml:"use,attr"`
					Annotation struct {
						Text          string `xml:",chardata"`
						Documentation struct {
							Text string `xml:",chardata"`
						} `xml:"documentation"`
					} `xml:"annotation"`
				} `xml:"attribute"`
				Assert struct {
					Text string `xml:",chardata"`
					Test string `xml:"test,attr"`
				} `xml:"assert"`
			} `xml:"extension"`
		} `xml:"complexContent"`
		Sequence struct {
			Text    string `xml:",chardata"`
			Element struct {
				Text      string `xml:",chardata"`
				MaxOccurs string `xml:"maxOccurs,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
			} `xml:"element"`
			Group struct {
				Text      string `xml:",chardata"`
				MaxOccurs string `xml:"maxOccurs,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				Ref       string `xml:"ref,attr"`
			} `xml:"group"`
			Choice struct {
				Text      string `xml:",chardata"`
				MaxOccurs string `xml:"maxOccurs,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				Element   []struct {
					Text      string `xml:",chardata"`
					MaxOccurs string `xml:"maxOccurs,attr"`
					MinOccurs string `xml:"minOccurs,attr"`
					Name      string `xml:"name,attr"`
					Type      string `xml:"type,attr"`
				} `xml:"element"`
			} `xml:"choice"`
		} `xml:"sequence"`
		Attribute []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Type       string `xml:"type,attr"`
			Use        string `xml:"use,attr"`
			Annotation struct {
				Text          string `xml:",chardata"`
				Documentation struct {
					Text string `xml:",chardata"`
				} `xml:"documentation"`
			} `xml:"annotation"`
		} `xml:"attribute"`
	} `xml:"complexType"`
}
