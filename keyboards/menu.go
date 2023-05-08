package keyboards

import "go/models"

func KeyYesRej() models.Keyboard {
	key := models.Keyboard{
		OneTime: true,
		Buttons: [][]models.Button{
			{
				models.Button{
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
			{
				models.Button{
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

func KeySex() models.Keyboard {
	key := models.Keyboard{
		Inline: true,
		Buttons: [][]models.Button{
			{
				models.Button{
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
				models.Button{
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

func KeyActivity() models.Keyboard {
	key := models.Keyboard{
		Inline: true,
		Buttons: [][]models.Button{
			{
				models.Button{
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
			{
				models.Button{
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
			{
				models.Button{
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
			{
				models.Button{
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

func KeyGoal() models.Keyboard {
	key := models.Keyboard{
		Inline: true,
		Buttons: [][]models.Button{
			{
				models.Button{
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
			{
				models.Button{
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
			{
				models.Button{
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

func KeyCheck() models.Keyboard {
	key := models.Keyboard{
		Inline: true,
		Buttons: [][]models.Button{
			{
				models.Button{
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
				models.Button{
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

func KeyCalculate() models.Keyboard {
	key := models.Keyboard{
		OneTime: true,
		Buttons: [][]models.Button{
			{
				{
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

func KeyRepeat() models.Keyboard {
	key := models.Keyboard{
		OneTime: true,
		Buttons: [][]models.Button{
			{
				{
					Action: struct {
						Type    string `json:"type"`
						Payload string `json:"payload,omitempty"`
						Label   string `json:"label"`
					}{
						Type:    "text",
						Payload: "{\"button\":\"repeat\"}",
						Label:   "Рассчитать ещё раз",
					},
					Color: "positive",
				},
			},
		},
	}
	return key
}
