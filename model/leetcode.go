package model

type TopicTag struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

type Question struct {
	AcRate             float64     `json:"acRate"`
	Difficulty         string      `json:"difficulty"`
	FreqBar            interface{} `json:"freqBar"`
	FrontendQuestionID string      `json:"frontendQuestionId"`
	IsFavor            bool        `json:"isFavor"`
	PaidOnly           bool        `json:"paidOnly"`
	Status             interface{} `json:"status"`
	Title              string      `json:"title"`
	TitleSlug          string      `json:"titleSlug"`
	TopicTags          []TopicTag  `json:"topicTags"`
	HasSolution        bool        `json:"hasSolution"`
	HasVideoSolution   bool        `json:"hasVideoSolution"`
}

type ProblemsetQuestionList struct {
	Total     int        `json:"total"`
	Questions []Question `json:"questions"`
}

type LeetCode struct {
	ProblemsetQuestionList ProblemsetQuestionList `json:"problemsetQuestionList"`
}
