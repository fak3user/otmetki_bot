package db

// func AddRecord(userId int64, minerId ) (bool, error) {
// 	collection := GetCollection("app", "users")
// 	filter := bson.M{"tg_id": fromUser.ID}
// 	var existingUser types.User

// 	err := collection.FindOne(context.TODO(), filter).Decode(&existingUser)
// 	if err == nil {
// 		return false, nil
// 	}

// 	user := types.User{
// 		TgID: fromUser.ID,
// 		Name: fromUser.FirstName + " " + fromUser.LastName,
// 	}
// 	// Insert user document into MongoDB
// 	_, err = collection.InsertOne(context.TODO(), user)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }
