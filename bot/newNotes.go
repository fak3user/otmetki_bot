package bot

type UpdatingRecord struct {
	LakeId   string
	RecordId string
	UserId   int64
}

var RecordUpdaters []UpdatingRecord

func AddNewNoteCreator(userId int64, lakeId string, recordId string) {
	RecordUpdaters = append(RecordUpdaters, UpdatingRecord{
		LakeId:   lakeId,
		RecordId: recordId,
		UserId:   userId,
	})
}

func CheckNewNoteCreator(userId int64) *UpdatingRecord {
	var record UpdatingRecord
	for i := range RecordUpdaters {
		if RecordUpdaters[i].UserId == userId {
			record = RecordUpdaters[i]
		}
	}
	return &record
}

func RemoveNewNoteCreator(userId int64) {
	var updatedCreators []UpdatingRecord
	for i := range RecordUpdaters {
		if RecordUpdaters[i].UserId != userId {
			updatedCreators = append(updatedCreators, RecordUpdaters[i])
		}
	}
	RecordUpdaters = updatedCreators
}
