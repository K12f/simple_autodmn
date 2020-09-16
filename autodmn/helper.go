package autodmn

func SliceKeyExist(v int, sl []interface{}) bool {
	for k, _ := range sl {
		if k == v {
			return true
		}
	}
	return false
}
