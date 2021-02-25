package tools

import "container/list"

func ListToArray(list *list.List) []string {
	actArray := make([]string, 0)
	for i := list.Front(); i != nil; i = i.Next() {
		actArray = append(actArray, i.Value.(string))
	}
	return actArray
}
func In(array []string, ele string) bool {
	for _, tmp := range array {
		if tmp == ele {
			return true
		}
	}
	return false
}
