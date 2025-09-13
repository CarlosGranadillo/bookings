package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap      map[string]string
	IntMap         map[string]int
	FloatMap       map[string]float32
	OtherTypesMap  map[string]interface{}
	CSRFToken      string
	FlashMessage   string
	WarningMessage string
	ErrorMessage   string
}
