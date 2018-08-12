package holiday

import "encoding/xml"

type GetHolidaysAvailable struct {
	XMLName     xml.Name `xml:"hs:GetHolidaysAvailable"`
	CountryCode string   `xml:"hs:countryCode"`
}
