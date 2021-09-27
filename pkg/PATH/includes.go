package PATH

func Includes(s string) (bool, error) {
	l, err := List()
	if err != nil {
		return false, err
	}
	for _, e := range l {
		if e == s {
			return true, nil
		}
	}
	return false, nil
}