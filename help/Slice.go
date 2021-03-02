package help

func IndexOf(atom string,array []string) bool{

	for _,value := range array {
		if atom == value {
			return true
		}
	}
	return false
}