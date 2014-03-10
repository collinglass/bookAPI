package schema

type Epub struct {
	Metadata Metadata
	Data     Data
}

type Metadata struct {
	Title       []string
	Language    []string
	Identifier  []string
	Creator     []string
	Subject     []string
	Description []string
	Publisher   []string
	Contributor []string
	Date        []string
	EpubType    []string
	Format      []string
	Source      []string
	Relation    []string
	Coverage    []string
	Rights      []string
	Meta        []string
}

type Data struct {
	Chapter []Chapter
}

type Chapter struct {
	Title string
	Text  []string
}

type Section struct {
	Title string
	Text  []string
}
