package models

type Book struct {
	NumFound      int  `json:"numFound"`
	Start         int  `json:"start"`
	NumFoundExact bool `json:"numFoundExact"`
	Docs          []struct {
		Key                   string   `json:"key"`
		Type                  string   `json:"type"`
		Seed                  []string `json:"seed"`
		Title                 string   `json:"title"`
		TitleSuggest          string   `json:"title_suggest"`
		EditionCount          int      `json:"edition_count"`
		EditionKey            []string `json:"edition_key,omitempty"`
		PublishDate           []string `json:"publish_date,omitempty"`
		PublishYear           []int    `json:"publish_year,omitempty"`
		FirstPublishYear      int      `json:"first_publish_year,omitempty"`
		NumberOfPagesMedian   int      `json:"number_of_pages_median,omitempty"`
		Lccn                  []string `json:"lccn,omitempty"`
		PublishPlace          []string `json:"publish_place,omitempty"`
		Oclc                  []string `json:"oclc,omitempty"`
		Contributor           []string `json:"contributor,omitempty"`
		Lcc                   []string `json:"lcc,omitempty"`
		Ddc                   []string `json:"ddc,omitempty"`
		Isbn                  []string `json:"isbn,omitempty"`
		LastModifiedI         int      `json:"last_modified_i"`
		EbookCountI           int      `json:"ebook_count_i"`
		EbookAccess           string   `json:"ebook_access"`
		HasFulltext           bool     `json:"has_fulltext"`
		PublicScanB           bool     `json:"public_scan_b"`
		Ia                    []string `json:"ia,omitempty"`
		IaCollectionS         string   `json:"ia_collection_s,omitempty"`
		LendingEditionS       string   `json:"lending_edition_s,omitempty"`
		LendingIdentifierS    string   `json:"lending_identifier_s,omitempty"`
		PrintdisabledS        string   `json:"printdisabled_s,omitempty"`
		CoverEditionKey       string   `json:"cover_edition_key,omitempty"`
		CoverI                int      `json:"cover_i,omitempty"`
		Publisher             []string `json:"publisher,omitempty"`
		Language              []string `json:"language,omitempty"`
		AuthorKey             []string `json:"author_key"`
		AuthorName            []string `json:"author_name"`
		AuthorAlternativeName []string `json:"author_alternative_name,omitempty"`
		Person                []string `json:"person,omitempty"`
		Place                 []string `json:"place,omitempty"`
		Subject               []string `json:"subject,omitempty"`
		Time                  []string `json:"time,omitempty"`
		IDGoodreads           []string `json:"id_goodreads,omitempty"`
		IDLibrarything        []string `json:"id_librarything,omitempty"`
		IDOverdrive           []string `json:"id_overdrive,omitempty"`
		IaLoadedID            []string `json:"ia_loaded_id,omitempty"`
		IaBoxID               []string `json:"ia_box_id,omitempty"`
		PublisherFacet        []string `json:"publisher_facet,omitempty"`
		PersonKey             []string `json:"person_key,omitempty"`
		PlaceKey              []string `json:"place_key,omitempty"`
		TimeFacet             []string `json:"time_facet,omitempty"`
		PersonFacet           []string `json:"person_facet,omitempty"`
		SubjectFacet          []string `json:"subject_facet,omitempty"`
		Version               int64    `json:"_version_"`
		PlaceFacet            []string `json:"place_facet,omitempty"`
		LccSort               string   `json:"lcc_sort,omitempty"`
		AuthorFacet           []string `json:"author_facet"`
		SubjectKey            []string `json:"subject_key,omitempty"`
		DdcSort               string   `json:"ddc_sort,omitempty"`
		TimeKey               []string `json:"time_key,omitempty"`
		FirstSentence         []string `json:"first_sentence,omitempty"`
		IDAmazon              []string `json:"id_amazon,omitempty"`
		Subtitle              string   `json:"subtitle,omitempty"`
	} `json:"docs"`
	NumFound0 int         `json:"num_found"`
	Q         string      `json:"q"`
	Offset    interface{} `json:"offset"`
}

type Key struct {
	Description struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"description"`
	Title         string   `json:"title"`
	Covers        []int    `json:"covers"`
	SubjectPlaces []string `json:"subject_places"`
	Subjects      []string `json:"subjects"`
	SubjectPeople []string `json:"subject_people"`
	Key           string   `json:"key"`
	Authors       []struct {
		Type struct {
			Key string `json:"key"`
		} `json:"type"`
		Author struct {
			Key string `json:"key"`
		} `json:"author"`
	} `json:"authors"`
	SubjectTimes []string `json:"subject_times"`
	Type         struct {
		Key string `json:"key"`
	} `json:"type"`
	LatestRevision int `json:"latest_revision"`
	Revision       int `json:"revision"`
	Created        struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
}
