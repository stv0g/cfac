package warnung

import "time"

type ResponseMapData []MapDataWarning

type ResponseAppCovid struct {
	Key         string              `json:"key"`
	Level       Level               `json:"level"`
	GeneralInfo string              `json:"generalInfo"`
	Rules       []AppCovidRule      `json:"rules"`
	Regulations AppCovidRegulations `json:"regulations"`
	Common      []struct {
		ID      string `json:"id"`
		Caption string `json:"caption"`
		Text    string `json:"text"`
	} `json:"common"`
}

type ResponseDashboard []Overview

type Overview struct {
	ID      string `json:"id"`
	Payload struct {
		Version int    `json:"version"`
		Type    string `json:"type"`
		ID      string `json:"id"`
		Hash    string `json:"hash"`
		Data    struct {
			Headline  string    `json:"headline"`
			Provider  string    `json:"provider"`
			Severity  string    `json:"severity"`
			MsgType   string    `json:"msgType"`
			TransKeys TransKeys `json:"transKeys"`
			Area      Area      `json:"area"`
		} `json:"data"`
	} `json:"payload"`
	I18NTitle I18NTitle `json:"i18nTitle"`
	Sent      time.Time `json:"sent"`
}

type AppCovidRule struct {
	ID      string `json:"id"`
	Caption string `json:"caption"`
	Text    string `json:"text"`
	Source  string `json:"source"`
	Icon    Icon   `json:"icon"`
}

type AppCovidRegulations struct {
	ValidFromUntil string             `json:"validFromUntil"`
	Sections       map[string]Section `json:"sections"`
}

type MapDataWarning struct {
	ID        string    `json:"id"`
	Version   int       `json:"version"`
	StartDate time.Time `json:"startDate"`
	Severity  string    `json:"severity"`
	Type      string    `json:"type"`
	I18NTitle I18NTitle `json:"i18nTitle"`
	TransKeys TransKeys `json:"transKeys"`
}

type Section struct {
	Caption string `json:"caption"`
	URL     string `json:"url"`
	Icon    Icon   `json:"icon"`
}

type Icon struct {
	Src  string `json:"src"`
	Hash string `json:"hash"`
}

type Level struct {
	Headline        string      `json:"headline"`
	Range           string      `json:"range"`
	BackgroundColor interface{} `json:"backgroundColor"`
	TextColor       interface{} `json:"textColor"`
}

type TransKeys struct {
	Event string `json:"event"`
}

type I18NTitle map[string]string

type Area struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
