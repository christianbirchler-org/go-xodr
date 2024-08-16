package main

import "encoding/xml"

// Schema was generated 2024-08-16 13:10:00 by https://xml-to-go.github.io/ in Ukraine.
type CoreSchema struct {
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
	Element struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		ComplexType struct {
			Text       string `xml:",chardata"`
			Annotation struct {
				Text          string `xml:",chardata"`
				Documentation struct {
					Text string `xml:",chardata"`
				} `xml:"documentation"`
			} `xml:"annotation"`
			Sequence struct {
				Text    string `xml:",chardata"`
				Element []struct {
					Text        string `xml:",chardata"`
					MaxOccurs   string `xml:"maxOccurs,attr"`
					MinOccurs   string `xml:"minOccurs,attr"`
					Name        string `xml:"name,attr"`
					Type        string `xml:"type,attr"`
					Alternative []struct {
						Text string `xml:",chardata"`
						Test string `xml:"test,attr"`
						Type string `xml:"type,attr"`
					} `xml:"alternative"`
				} `xml:"element"`
				Group struct {
					Text      string `xml:",chardata"`
					MaxOccurs string `xml:"maxOccurs,attr"`
					MinOccurs string `xml:"minOccurs,attr"`
					Ref       string `xml:"ref,attr"`
				} `xml:"group"`
			} `xml:"sequence"`
		} `xml:"complexType"`
		Key []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Selector struct {
				Text  string `xml:",chardata"`
				Xpath string `xml:"xpath,attr"`
			} `xml:"selector"`
			Field struct {
				Text  string `xml:",chardata"`
				Xpath string `xml:"xpath,attr"`
			} `xml:"field"`
		} `xml:"key"`
		Keyref []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Refer    string `xml:"refer,attr"`
			Selector struct {
				Text  string `xml:",chardata"`
				Xpath string `xml:"xpath,attr"`
			} `xml:"selector"`
			Field struct {
				Text  string `xml:",chardata"`
				Xpath string `xml:"xpath,attr"`
			} `xml:"field"`
		} `xml:"keyref"`
	} `xml:"element"`
	SimpleType []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		Restriction struct {
			Text        string `xml:",chardata"`
			Base        string `xml:"base,attr"`
			Enumeration []struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"enumeration"`
			MinInclusive struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"minInclusive"`
			MinExclusive struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"minExclusive"`
			MaxInclusive struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"maxInclusive"`
		} `xml:"restriction"`
		Union struct {
			Text        string `xml:",chardata"`
			MemberTypes string `xml:"memberTypes,attr"`
		} `xml:"union"`
	} `xml:"simpleType"`
	ComplexType []struct {
		Text     string `xml:",chardata"`
		Abstract string `xml:"abstract,attr"`
		Name     string `xml:"name,attr"`
		Mixed    string `xml:"mixed,attr"`
		Sequence struct {
			Text    string `xml:",chardata"`
			Element []struct {
				Text      string `xml:",chardata"`
				MaxOccurs string `xml:"maxOccurs,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
			} `xml:"element"`
			Any struct {
				Text            string `xml:",chardata"`
				MaxOccurs       string `xml:"maxOccurs,attr"`
				MinOccurs       string `xml:"minOccurs,attr"`
				ProcessContents string `xml:"processContents,attr"`
			} `xml:"any"`
		} `xml:"sequence"`
		Annotation struct {
			Text          string `xml:",chardata"`
			Documentation struct {
				Text string `xml:",chardata"`
			} `xml:"documentation"`
		} `xml:"annotation"`
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
				} `xml:"sequence"`
				Attribute []struct {
					Text       string `xml:",chardata"`
					Name       string `xml:"name,attr"`
					Type       string `xml:"type,attr"`
					Use        string `xml:"use,attr"`
					Fixed      string `xml:"fixed,attr"`
					Annotation struct {
						Text          string `xml:",chardata"`
						Documentation struct {
							Text string `xml:",chardata"`
						} `xml:"documentation"`
					} `xml:"annotation"`
				} `xml:"attribute"`
			} `xml:"extension"`
		} `xml:"complexContent"`
	} `xml:"complexType"`
	Group struct {
		Text       string `xml:",chardata"`
		Name       string `xml:"name,attr"`
		Annotation struct {
			Text          string `xml:",chardata"`
			Documentation struct {
				Text string `xml:",chardata"`
			} `xml:"documentation"`
		} `xml:"annotation"`
		Sequence struct {
			Text    string `xml:",chardata"`
			Element []struct {
				Text      string `xml:",chardata"`
				MaxOccurs string `xml:"maxOccurs,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
			} `xml:"element"`
		} `xml:"sequence"`
	} `xml:"group"`
}
