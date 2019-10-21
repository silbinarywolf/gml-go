package monotime

// Now returns the current time more precisely for Web and Windows targets
func Now() int64 {
	return now()
}
