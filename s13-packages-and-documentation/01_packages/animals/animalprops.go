package animals

// Speaker ... you comment
type Speaker interface {
	Speaks() string
}

// Perform ... your comments
func Perform(s Speaker) string { return s.Speaks() }

// Lion ... you comment
type Lion struct{}

// Speaks ... your comments
func (s Lion) Speaks() string { return "roars" }

// Cat ... you comment
type Cat struct{}

// Speaks ... your comments
func (s Cat) Speaks() string { return "meaws" }

// Human ... you comment
type Human struct{}

// Speaks ... your comments
func (s Human) Speaks() string { return "talks" }
