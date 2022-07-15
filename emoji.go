package httpcord

type Emoji struct {
	ID            Snowflake    `json:"id,omitempty"`
	Name          string       `json:"name,omitempty"`
	Roles         []*Snowflake `json:"roles,omitempty"`
	User          *User        `json:"user,omitempty"`
	RequireColons bool         `json:"require_colons,omitempty"`
	Managed       bool         `json:"managed,omitempty"`
	Animated      bool         `json:"animated,omitempty"`
	Available     bool         `json:"available,omitempty"`
}

func (e *Emoji) Mention() string {
	if e.ID.String() != "" {
		m := "<"

		if e.Animated {
			m += "a"
		}

		m += ":" + e.Name + ":" + e.ID.String() + ">"
		return m
	}

	return e.Name
}

func (e *Emoji) String() string {
	return e.Mention()
}

type Reaction struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emoji `json:"emoji"`
}
