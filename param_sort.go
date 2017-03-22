package main

import "sort"

// Hello
// I hope that this file in no way becomes the unadulterated shitshow that it probably will
// Sincerely
// Everyone who ever tried to do something useful in the spur of the moment, ever

func sortParamData(pd map[string]ParamData) []ParamData {
	list := make([]ParamData, len(pd)) // <-
	var i int
	for _, v := range pd {
		list[i] = v
		i++
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Embed && !list[j].Embed // <-
	})
	return list
}
