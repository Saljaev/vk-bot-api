package keyboards

import (
	"go/model"
)

model

func keyYesRej() model.Keyboard {
	key := keyboard{
		OneTime: true,
		Buttons: [][]button{
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"want\"}",
						Label:   "Да",
					},
					Color: "positive",
				},
			},
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"reject\"}",
						Label:   "Нет",
					},
					Color: "negative",
				},
			},
		},
	}
	return key
}

func keySex() keyboard {
	key := keyboard{
		Inline: true,
		Buttons: [][]button{
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"male\"}",
						Label:   "Мужской",
					},
				},
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"female\"}",
						Label:   "Женский",
					},
				},
			},
		},
	}
	return key
}

func keyActivity() keyboard {
	key := keyboard{
		Inline: true,
		Buttons: [][]button{
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"easy\"}",
						Label:   "Малая активность",
					},
				},
			},
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"medium\"}",
						Label:   "Слабая активность 1-3 раза в неделю",
					},
				},
			},
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"hard\"}",
						Label:   "Средняя активность 3-5 раз в неделю",
					},
				},
			},
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"extreme\"}",
						Label:   "Высокая активность 5-7 раз в неделю",
					},
				},
			},
		},
	}
	return key
}

func keyGoat() keyboard {
	key := keyboard{
		Inline: true,
		Buttons: [][]button{
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"lost\"}",
						Label:   "Похудеть",
					},
				},
			},
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"stay\"}",
						Label:   "Сохранить вес",
					},
				},
			},
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"gain\"}",
						Label:   "Набрать мышечную массу",
					},
				},
			},
		},
	}
	return key
}

func keyCheck() keyboard {
	key := keyboard{
		Inline: true,
		Buttons: [][]button{
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"check\"}",
						Label:   "Посчитать",
					},
					Color: "positive",
				},
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"repeat\"}",
						Label:   "Ввести заново",
					},
					Color: "negative",
				},
			},
		},
	}
	return key
}

func keyCalculate() keyboard {
	key := keyboard{
		OneTime: true,
		Buttons: [][]button{
			[]button{
				button{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"calculate\"}",
						Label:   "Рассчитать",
					},
					Color: "positive",
				},
			},
		},
	}
	return key
}
