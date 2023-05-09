package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/database"
	"go/keyboards"
	"go/models"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	longPollServer string
	longPollKey    string
	ts             string
)

func GetLongPollServer() error {
	params := url.Values{}
	params.Set("group_id", os.Getenv("groupID"))
	params.Set("access_token", os.Getenv("accessToken"))
	params.Set("v", os.Getenv("apiVersion"))

	resp, err := http.Get("https://api.vk.com/method/groups.getLongPollServer?" + params.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data struct {
		Response struct {
			Server string `json:"server"`
			Key    string `json:"key"`
			Ts     string `json:"ts"`
		} `json:"response"`
		Error struct {
			ErrorCode int    `json:"error_code"`
			ErrorMsg  string `json:"error_msg"`
		} `json:"error"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	if data.Error.ErrorCode != 0 {
		return fmt.Errorf("getLongPollServer error: %d %s", data.Error.ErrorCode, data.Error.ErrorMsg)
	}

	longPollServer = data.Response.Server
	longPollKey = data.Response.Key
	ts = data.Response.Ts

	return nil
}

func LongPollHandler() {
	db, err := database.CreateDB()
	if err != nil {
		panic(err)
	}
	pfc := models.PFC{}
	for {
		updates, err := GetLongPollUpdates()
		if err != nil {
			log.Println(err)
			continue
		}
		for _, update := range updates {
			if update.Type == "message_new" {
				check, err := database.CheckIfPFCExists(db, update.Object.Message.FromID)
				if err != nil {
					panic(err)
				}
				if check == false {
					pfc = CreateZeroStruct(update.Object.Message.FromID, false)
					err := database.InsertPFC(db, pfc)
					if err != nil {
						if !strings.Contains(err.Error(), "already exists") {
							panic(err)
						}
					}
				} else {
					pfc, err = database.GetPFCByUserID(db, update.Object.Message.FromID)
					if err != nil {
						panic(err)
					}
				}
				if (update.Object.Message.Text == "Рассчитать" || update.Object.Message.Text == "рассчитать") && update.Object.Message.Payload != "{\"button\":\"calculate\"}" {
					SendMessage(update.Object.Message.FromID, "Введите ваш возраст", models.Keyboard{})
					pfc = CreateZeroStruct(update.Object.Message.FromID, false)
					pfc.WaitingForAge = true
				} else if update.Object.Message.Payload == "" && !pfc.WaitingForAge && !pfc.WaitingForHeight && !pfc.WaitingForWeight && !pfc.WaitingForSex && !pfc.WaitingForActivity {
					SendMessage(update.Object.Message.FromID, "Я умею считать БЖУ, если хотите посчитать напишите Рассчитать или нажмите кнопку снизу", keyboards.KeyCalculate())
				} else {
					if pfc.WaitingForAge && update.Object.Message.Payload == "" {
						age, err := strconv.Atoi(update.Object.Message.Text)
						if err != nil || age < 1 || age > 150 {
							SendMessage(update.Object.Message.FromID, "Пожалуйста, введите возраст корректно", models.Keyboard{})
						} else {
							SendMessage(update.Object.Message.FromID, "Введите ваш рост в см", models.Keyboard{})
							pfc.Age = age
							pfc.WaitingForAge = false
							pfc.WaitingForHeight = true
						}

					} else if pfc.WaitingForHeight && update.Object.Message.Payload == "" {
						height, err := strconv.Atoi(update.Object.Message.Text)
						if err != nil || height < 10 || height > 300 {
							SendMessage(update.Object.Message.FromID, "Пожалуйста, введите рост корректно", models.Keyboard{})
						} else {
							SendMessage(update.Object.Message.FromID, "Введите ваш вес в кг", models.Keyboard{})
							pfc.Height = height
							pfc.WaitingForHeight = false
							pfc.WaitingForWeight = true
						}
					} else if pfc.WaitingForWeight && update.Object.Message.Payload == "" {
						weight, err := strconv.Atoi(update.Object.Message.Text)
						if err != nil || weight < 0 || weight > 300 {
							SendMessage(update.Object.Message.FromID, "Пожалуйста, введите вес корректно", models.Keyboard{})
						} else {
							SendMessage(update.Object.Message.FromID, "Укажите ваш пол", keyboards.KeySex())
							pfc.Weight = weight
							pfc.WaitingForWeight = false
							pfc.WaitingForSex = true
						}
					} else if pfc.WaitingForSex && update.Object.Message.Payload == "" {
						if update.Object.Message.Payload == "" {
							SendMessage(update.Object.Message.FromID, "Пожалуйста, укажите пол", models.Keyboard{})
						} else {
							pfc.WaitingForSex = false
							pfc.WaitingForActivity = true
						}
					} else if pfc.WaitingForActivity && update.Object.Message.Payload == "" {
						if update.Object.Message.Payload == "" {
							SendMessage(update.Object.Message.FromID, "Пожалуйста, выберите вашу активность", models.Keyboard{})
						} else {
							pfc.WaitingForActivity = false
						}
					}
				}
				switch update.Object.Message.Payload {
				case "{\"button\":\"lost\"}":
					pfc.Protein = int(float64(pfc.Calories) * 0.25 / 4)
					pfc.Fats = int(float64(pfc.Calories) * 0.25 / 9)
					pfc.Carbohydrates = int(float64(pfc.Calories) * 0.4 / 4)
					SendMessage(update.Object.Message.FromID, fmt.Sprintf("Для похудения вам следует съедать\nБелка: %d грамм\nЖиров: %d грамм\nУглеводов: %d грамм", pfc.Protein, pfc.Fats, pfc.Carbohydrates), keyboards.KeyRepeat())

				case "{\"button\":\"stay\"}":
					pfc.Protein = int(float64(pfc.Calories) * 0.3 / 4)
					pfc.Fats = int(float64(pfc.Calories) * 0.3 / 9)
					pfc.Carbohydrates = int(float64(pfc.Calories) * 0.4 / 4)
					SendMessage(update.Object.Message.FromID, fmt.Sprintf("Для сохранения вам следует съедать\nБелка: %d грамм\nЖиров: %d грамм\nУглеводов: %d грамм", pfc.Protein, pfc.Fats, pfc.Carbohydrates), keyboards.KeyRepeat())

				case "{\"button\":\"gain\"}":
					pfc.Protein = int(float64(pfc.Calories) * 0.35 / 4)
					pfc.Fats = int(float64(pfc.Calories) * 0.3 / 9)
					pfc.Carbohydrates = int(float64(pfc.Calories) * 0.55 / 4)
					SendMessage(update.Object.Message.FromID, fmt.Sprintf("Для набора вам следует съедать\nБелка: %d грамм\nЖиров: %d грамм\nУглеводов: %d грамм", pfc.Protein, pfc.Fats, pfc.Carbohydrates), keyboards.KeyRepeat())

				case "{\"button\":\"check\"}":
					var calories float64
					if pfc.Sex == "мужской" {
						calories = (10*float64(pfc.Weight) + 6.25*float64(pfc.Height) - 5*float64(pfc.Age) + 5) * pfc.Activity
					} else {
						calories = (10*float64(pfc.Weight) + 6.25*float64(pfc.Height) - 5*float64(pfc.Age) - 161) * pfc.Activity
					}
					pfc.Calories = int(calories)
					SendMessage(update.Object.Message.FromID, fmt.Sprintf("Ваша суточная норма калорий: %d ккал", pfc.Calories), models.Keyboard{})
					SendMessage(update.Object.Message.FromID, "Хотите ли вы узнать БЖУ для таких каллорий?", keyboards.KeyYesRej())

				case "{\"button\":\"want\"}":
					SendMessage(update.Object.Message.FromID, "Отлично, выберите желаемую цель", keyboards.KeyGoal())

				case "{\"button\":\"reject\"}":
					SendMessage(update.Object.Message.FromID, "Был рад помочь)", models.Keyboard{})

				case "{\"button\":\"repeat\"}":
					SendMessage(update.Object.Message.FromID, "Введите ваш возраст", models.Keyboard{})
					pfc = CreateZeroStruct(update.Object.Message.FromID, true)

				case "{\"button\":\"easy\"}":
					if pfc.Activity == 0 {
						pfc.Activity = 1.2
						SendMessage(update.Object.Message.FromID, fmt.Sprintf("Вы ввели \nВозраст: %d\nПол: %s\nВес: %d\nРост: %d\nАктивность: малоподвижность", pfc.Age, pfc.Sex, pfc.Weight, pfc.Height), keyboards.KeyCheck())
					} else {
						SendMessage(update.Object.Message.FromID, "Вы уже указали активность", models.Keyboard{})
					}

				case "{\"button\":\"medium\"}":
					if pfc.Activity == 0 {
						pfc.Activity = 1.375
						SendMessage(update.Object.Message.FromID, fmt.Sprintf("Вы ввели \nВозраст: %d\nПол: %s\nВес: %d\nРост: %d\nАктивность: слабая", pfc.Age, pfc.Sex, pfc.Weight, pfc.Height), keyboards.KeyCheck())
					} else {
						SendMessage(update.Object.Message.FromID, "Вы уже указали активность", models.Keyboard{})
					}

				case "{\"button\":\"hard\"}":
					if pfc.Activity == 0 {
						pfc.Activity = 1.55
						SendMessage(update.Object.Message.FromID, fmt.Sprintf("Вы ввели \nВозраст: %d\nПол: %s\nВес: %d\nРост: %d\nАктивность: средняя", pfc.Age, pfc.Sex, pfc.Weight, pfc.Height), keyboards.KeyCheck())
					} else {
						SendMessage(update.Object.Message.FromID, "Вы уже указали активность", models.Keyboard{})
					}

				case "{\"button\":\"extreme\"}":
					if pfc.Activity == 0 {
						pfc.Activity = 1.7
						SendMessage(update.Object.Message.FromID, fmt.Sprintf("Вы ввели \nВозраст: %d\nПол: %s\nВес: %d\nРост: %d\nАктивность: высокая", pfc.Age, pfc.Sex, pfc.Weight, pfc.Height), keyboards.KeyCheck())
					} else {
						SendMessage(update.Object.Message.FromID, "Вы уже указали активность", models.Keyboard{})
					}

				case "{\"button\":\"male\"}":
					if pfc.Sex == "" {
						SendMessage(update.Object.Message.FromID, "Укажите вашу активность", keyboards.KeyActivity())
						pfc.Sex = "мужской"
					} else {
						SendMessage(update.Object.Message.FromID, "Вы уже выбрали пол", models.Keyboard{})
					}
					pfc.WaitingForSex = false
				case "{\"button\":\"female\"}":
					if pfc.Sex == "" {
						SendMessage(update.Object.Message.FromID, "Укажите вашу активность", keyboards.KeyActivity())
						pfc.Sex = "женский"
					} else {
						SendMessage(update.Object.Message.FromID, "Вы уже выбрали пол", models.Keyboard{})
					}
					pfc.WaitingForSex = false
				case "{\"button\":\"calculate\"}":
					pfc = CreateZeroStruct(update.Object.Message.FromID, false)
					if err != nil {
						if !strings.Contains(err.Error(), "already exists") {
							panic(err)
						}
					}
					SendMessage(update.Object.Message.FromID, "Введите ваш возраст", models.Keyboard{})
					pfc.WaitingForAge = true

				case "{\"command\":\"start\"}":
					pfc = CreateZeroStruct(update.Object.Message.FromID, false)
					if err != nil {
						if !strings.Contains(err.Error(), "already exists") {
							panic(err)
						}
					}
					SendMessage(update.Object.Message.FromID, "Введите ваш возраст", models.Keyboard{})
					pfc.WaitingForAge = true
				default:
					break
				}

				err = database.UpdatePFC(db, pfc)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func CreateZeroStruct(userId int, waitForAge bool) models.PFC {
	user := models.PFC{
		User_id:            userId,
		Calories:           0,
		Height:             0,
		Weight:             0,
		Protein:            0,
		Fats:               0,
		Carbohydrates:      0,
		Sex:                "",
		Age:                0,
		WaitingForHeight:   false,
		WaitingForWeight:   false,
		WaitingForAge:      waitForAge,
		Activity:           0,
		WaitingForSex:      false,
		WaitingForActivity: false,
	}
	return user
}

func GetLongPollUpdates() ([]models.Update, error) {
	params := url.Values{}
	params.Set("act", "a_check")
	params.Set("key", longPollKey)
	params.Set("ts", ts)
	params.Set("wait", "25")
	params.Set("mode", "2")
	params.Set("version", os.Getenv("apiVersion"))
	resp, err := http.Get(longPollServer + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Updates []models.Update `json:"updates"`
		TS      string          `json:"ts"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	ts = data.TS

	return data.Updates, nil
}
func SendMessage(userID int, message string, key models.Keyboard) error {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(userID))
	params.Set("message", message)
	params.Set("access_token", os.Getenv("accessToken"))
	params.Set("v", os.Getenv("apiVersion"))
	rand.Seed(time.Now().UnixNano())
	params.Set("random_id", strconv.Itoa(rand.Intn(1000000)))
	if key.Buttons != nil {
		keyboardJSON, err := json.Marshal(key)
		if err != nil {
			return err
		}
		params.Set("keyboard", string(keyboardJSON))
	}

	apiURL := "https://api.vk.com/method/messages.send?" + params.Encode()
	body := []byte("123")
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return nil
}
