package radmon

// https://radmon.org/index.php/forum/howtos-and-faqs/864-radmon-org-api

const (
	UrlAPI = "https://radmon.org/radmon.php" // + ?function=showuserpage&user= oder ?function=lastreading
)

var (
	Users []string = []string{
		"Rotter",
		"jokri",
	}
)
