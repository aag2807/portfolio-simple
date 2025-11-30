package portfolio

type Contact struct {
	Email    string `json:"email"`
	LinkedIn string `json:"linkedin"`
	GitHub   string `json:"github"`
}

type Personal struct {
	Name           string  `json:"name"`
	Title          string  `json:"title"`
	Subtitle       string  `json:"subtitle"`
	Location       string  `json:"location"`
	RelocationNote string  `json:"relocationNote"`
	Contact        Contact `json:"contact"`
}

type About struct {
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
	CurrentFocus string   `json:"currentFocus"`
}

type PolyglotLanguage struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Years    int    `json:"years"`
}

type Polyglot struct {
	Tagline     string             `json:"tagline"`
	Description string             `json:"description"`
	Languages   []PolyglotLanguage `json:"languages"`
}

type Skills struct {
	Backend   []string `json:"backend"`
	Frontend  []string `json:"frontend"`
	Database  []string `json:"database"`
	Tools     []string `json:"tools"`
	Exploring []string `json:"exploring"`
}

type Experience struct {
	Company      string   `json:"company"`
	Title        string   `json:"title"`
	Period       string   `json:"period"`
	Location     string   `json:"location"`
	Type         string   `json:"type"`
	Achievements []string `json:"achievements"`
	Technologies []string `json:"technologies"`
}

type Project struct {
	Name         string   `json:"name"`
	Tagline      string   `json:"tagline"`
	Description  string   `json:"description"`
	Status       string   `json:"status"`
	Technologies []string `json:"technologies"`
	Highlights   []string `json:"highlights"`
}

type Certification struct {
	Name         string   `json:"name"`
	Issuer       string   `json:"issuer"`
	Date         string   `json:"date"`
	CredentialID string   `json:"credentialId"`
	Skills       []string `json:"skills"`
}

type Education struct {
	Type        string `json:"type"`
	Field       string `json:"field"`
	Period      string `json:"period"`
	Description string `json:"description"`
}

type Language struct {
	Language    string `json:"language"`
	Proficiency string `json:"proficiency"`
}

type Portfolio struct {
	Personal       Personal        `json:"personal"`
	About          About           `json:"about"`
	Polyglot       Polyglot        `json:"polyglot"`
	Skills         Skills          `json:"skills"`
	Experience     []Experience    `json:"experience"`
	Projects       []Project       `json:"projects"`
	Certifications []Certification `json:"certifications"`
	Education      Education       `json:"education"`
	Languages      []Language      `json:"languages"`
	Interests      []string        `json:"interests"`
}

