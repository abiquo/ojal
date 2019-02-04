package abiquo

import "github.com/abiquo/ojal/core"

func collectionToList(resource core.Resource, rel string, add func(core.Resource)) {
	collection := resource.Rel(rel).Collection(nil)
	collection.List().Map(add)
	return
}
