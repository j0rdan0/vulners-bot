package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type VulnersResponse struct {
	Result string `json:"result"`
	Data   struct {
		Result []struct {
			Source struct {
				Lastseen    string `json:"lastseen"`
				Description string `json:"description"`

				Title string `json:"title"`

				Href string `json:"href"`
			} `json:"_source"`
		} `json:"result"`
	} `json:"data"`
}

var subscription = map[string]string{
	"securityNews":   "HGCUAIG7TNIUHV1GFMVQTSCG219N105MHBDUQZZJ4GU3WFWPSSP4IBJZV0YGXAJ&apiKey=QBHKFCICBU5Q4VVTODZXJ9Z9FOELNW5CBE8TFO04YLMNNMG7RRU7FQKNBVAOMRLZ",
	"exploitUpdates": "S16UEBTEIR7WHYJYPUCW1LSTIUNGL5QO60Y68UKODLZ38Z78Z984CIGI9H24PRZB&apiKey=QBHKFCICBU5Q4VVTODZXJ9Z9FOELNW5CBE8TFO04YLMNNMG7RRU7FQKNBVAOMRLZ",
	"linuxVulners":   "OVAXFIQFYBQU21JJC007O7SM43EN4HS4PGNQAW70LKUKCDRA9ZCLQLKJD7OSRUZP&apiKey=QBHKFCICBU5Q4VVTODZXJ9Z9FOELNW5CBE8TFO04YLMNNMG7RRU7FQKNBVAOMRLZ",
	"windowsVulners": "D8KLC97XPTFTBCJPU0JPTO880KD68HHUO4VV7CBDTKRIZO597MN45GIE03KRH4MV&apiKey=QBHKFCICBU5Q4VVTODZXJ9Z9FOELNW5CBE8TFO04YLMNNMG7RRU7FQKNBVAOMRLZ",
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Security News", "securityNews"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Exploit Updates", "exploitUpdates"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Linux Vulners", "linuxVulners"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Windows Vulners", "windowsVulners")))
