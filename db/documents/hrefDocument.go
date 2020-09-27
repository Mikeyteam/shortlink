package documents

type HrefDocument struct {
	Id        string `bson:"_id,omitempty"`
	LongHref  string
	ShortHref string
}
