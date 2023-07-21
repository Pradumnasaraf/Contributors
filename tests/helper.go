package tests

import (
	"bytes"
	"encoding/json"

	"github.com/Pradumnasaraf/Contributors/database"
	"github.com/Pradumnasaraf/Contributors/graph/model"
)

var db = database.NewMongoDB()

func JSONEqual(a, b map[string]interface{}) bool {
	aJSON, err := json.Marshal(a)
	if err != nil {
		return false
	}

	bJSON, err := json.Marshal(b)
	if err != nil {
		return false
	}

	return bytes.Equal(aJSON, bJSON)
}

func DeleteAddedContributor(id string) {

	err := db.DeleteByID(id)
	if err != nil {
		return
	}

}

func DeleteAddedContribution(userID, ContributorID string) {

	err := db.DeleteContributionByID(userID, ContributorID)
	if err != nil {
		return
	}

}

func AddDeletedContributor() {

	var contributor = &model.Contributor{
		UserID:         "U7",
		Name:           "user7",
		Email:          "test@test.com",
		GithubUsername: "7",
		Contributions:  []*model.Contribution{},
	}

	err := db.Add(contributor)
	if err != nil {
		return
	}

}

func AddDeletedContribution() {

	var contribution = &model.Contribution{
		ContributionID: "C1",
		Type:           "code",
		ProjectName:    "1",
	}

	err := db.AddContributionByID("U8", contribution)
	if err != nil {
		return
	}
}

func UpdateModifiedContributor() {

	var contributor = &model.Contributor{
		UserID:         "U9",
		Name:           "user9",
		Email:          "test@test.com",
		GithubUsername: "9",
		Contributions:  []*model.Contribution{},
	}

	err := db.UpdateByID(contributor)
	if err != nil {
		return
	}

}
