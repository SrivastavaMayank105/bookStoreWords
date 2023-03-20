package mapper

type BookRequestParams struct {
	BookText string `json:"booktext" validate:"required_notNull"`
}


