package models

type Update struct {
	Type   string `json:"type"`
	Object struct {
		Message struct {
			Date                  int    `json:"date"`
			FromID                int    `json:"from_id"`
			ID                    int    `json:"id"`
			Out                   int    `json:"out"`
			PeerID                int    `json:"peer_id"`
			Text                  string `json:"text"`
			ConversationMessageID int    `json:"conversation_message_id"`
			Payload               string `json:"payload"`
		} `json:"message"`
	} `json:"object"`
}

type Keyboard struct {
	OneTime bool       `json:"one_time"`
	Inline  bool       `json:"inline"`
	Buttons [][]Button `json:"buttons"`
}

type Button struct {
	Action struct {
		Type    string `json:"type"`
		Payload string `json:"payload,omitempty"`
		Label   string `json:"label"`
	} `json:"action"`
	Color string `json:"color,omitempty"`
}

type PFC struct {
	Id                 int
	User_id            int
	Calories           int
	Protein            int
	Fats               int
	Carbohydrates      int
	Sex                string
	Age                int
	WaitingForAge      bool
	WaitingForHeight   bool
	WaitingForWeight   bool
	WaitingForSex      bool
	WaitingForActivity bool
	Activity           float64
	Weight             int
	Height             int
}
