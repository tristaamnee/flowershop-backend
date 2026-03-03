package repository

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/tristaamnee/flowershop-be/common/repository"
	"github.com/tristaamnee/flowershop-be/tokens/model"
)

func CreateToken(db *pg.DB, token *model.Token) error {
	return repository.Create(db, token)

}

func GetToken(db *pg.DB, ID uuid.UUID) (*model.Token, error) {
	return repository.GetByCondition[model.Token](db, model.ColTokenID, ID)
}

func ListTokens(db *pg.DB) ([]*model.Token, error) {
	var tokens []*model.Token
	repository.GetWholeTable[model.Token](db, &tokens)

	//create a map to store token activities associated with each token
	activitiesMap := make(map[uuid.UUID][]*model.TokenActivity)

	//Fetch all token activities
	var activities []*model.TokenActivity
	err := db.Model(&activities).Select()
	if err != nil {
		log.Printf("Error getting tokens: %v\n", err)
		return nil, err
	}

	//Organize token activities by token ID
	for _, activity := range activities {
		activitiesMap[activity.TokenID] = append(activitiesMap[activity.TokenID], activity)
	}

	//Populate token activities for each token
	for _, token := range tokens {
		token.Activities = activitiesMap[token.ID]
	}

	return tokens, nil
}

func DeleteToken(db *pg.DB, ID uuid.UUID) error {
	return repository.DeleteByCondition[model.Token](db, model.ColTokenID, ID)
}

// [TOKEN ACTIVITIES] //

func CreateTokenUsage(db *pg.DB, tokenUsage *model.TokenActivity) error {
	return repository.Create(db, tokenUsage)
}

func GetTokenActivity(db *pg.DB, ID uuid.UUID) ([]*model.TokenActivity, error) {
	var tokenActivity []*model.TokenActivity
	repository.GetWholeTable[model.TokenActivity](db, &tokenActivity, model.ColTokenID)
	return tokenActivity, nil
}

func DeleteTokenActivity(db *pg.DB, ID uuid.UUID) error {
	return repository.DeleteByCondition[model.TokenActivity](db, model.ColTokenID, ID)
}
