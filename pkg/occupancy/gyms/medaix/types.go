package medaix

type VisitorCounter struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Badge      string     `json:"badge"`
	BadgeClass string     `json:"badgeClass"`
	ShowBadge  string     `json:"showBadge"`
	Value      string     `json:"value"`
	Max        string     `json:"max"`
	A          string     `json:"a"`
	B          string     `json:"b"`
	Active     string     `json:"active"`
	Timestamp  CustomTime `json:"timestamp"`
}
