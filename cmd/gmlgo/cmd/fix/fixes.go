package fix

type fix struct {
	name string
	date string // date that fix was introduced, in YYYY-MM-DD format
	f    func(*File) bool
	desc string
	//disabled bool // whether this fix should be disabled by default
}

var fixes []fix

func register(f fix) {
	fixes = append(fixes, f)
}
