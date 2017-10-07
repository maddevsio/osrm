package osrm

type MatchResponse struct {
	Code  string    `json:"code"`
	Matchings []Matching `json:"matchings"`
}

type Matching struct {
	Confidence float64    `json:"confidence"`
	Distance float64    `json:"distance"`
	Duration float64    `json:"duration"`
	Geometry string    `json:"geometry"`
	Legs []Leg `json:"legs"`
}

type Leg struct {
	Distance float64    `json:"distance"`
	Duration float64    `json:"duration"`
	Steps []Step `json:"steps"`
	Summary string    `json:"summary"`
	Weight float64    `json:"weight"`
}

type Step struct {
	Distance float64    `json:"distance"`
	Duration float64    `json:"duration"`
	Geometry string    `json:"geometry"`
	Intersections []Intersection `json:"intersections"`
	Maneuver Maneuver `json:"maneuver"`
	Mode string    `json:"mode"`
	Name string    `json:"name"`
	Weight float64    `json:"weight"`
}

type Intersection struct {
	Bearings []int `json:"bearings"`
	Entry []bool `json:"entry"`
	Location []float64 `json:"location"`
	Type string    `json:"type"`
}

type Maneuver struct {
	BearingAfter int `json:"bearing_after"`
	BearingBefore int `json:"bearing_before"`
	Entry []bool `json:"entry"`
	Location []float64 `json:"location"`
}
