package bot

import "data-miner/utils"

var NewLakeCreators []int64

func AddNewLakeCreator(userId int64) {
	NewLakeCreators = append(NewLakeCreators, userId)
}

func CheckNewLakeCreator(userId int64) bool {
	for i := range NewLakeCreators {
		if NewLakeCreators[i] == userId {
			return true
		}
	}
	return false
}

func RemoveNewLakeCreator(userId int64) {
	for i := range NewLakeCreators {
		if NewLakeCreators[i] == userId {
			updatedCreators := utils.DeleteSliceElement(NewLakeCreators, i)
			NewLakeCreators = updatedCreators
		}
	}
}
