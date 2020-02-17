package model

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Root struct {
	XMLName xml.Name `xml:"root"`
	Medias  []Media  `xml:"media"`
}

type Media struct {
	XMLName           xml.Name     `xml:"media"`
	Id                string       `xml:"id"`
	NavisionTitleCode string       `xml:"navisionTitleCode"`
	ExtTitleCode      string       `xml:"extTitleCode"`
	Title             string       `xml:"title"`
	Format            string       `xml:"format"`
	Modified          TimeWithoutZ `xml:"modified"`
}

func (media Media) String() string {
	return fmt.Sprintf("Id: %s NavisionTitleCode: %s, ExtTitleCode: %s Title: %s Format: %s Modified: %s", media.Id, media.NavisionTitleCode, media.ExtTitleCode, media.Title, media.Format, media.Modified)
}

type TimeWithoutZ struct {
	time.Time
}

func (c *TimeWithoutZ) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, _ := time.Parse("2006-01-02T03:04:05", v)
	*c = TimeWithoutZ{parse}
	return nil
}

type Specification struct {
	Url          string `required:"true"`
	SiteCode     string `required:"true"`
	OperatorCode string `required:"true"`
	Country      string `required:"true"`
	Language     string `required:"true"`
}
