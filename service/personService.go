package service

import (
	"context"
	"errors"
	"fmt"
	"handsOnGO/database"
	"handsOnGO/dto"
	"cloud.google.com/go/firestore"
)

var CollectionName = "sohel"

func CreatePerson(ctx context.Context, person dto.Person) (*dto.Person, *firestore.WriteResult, error) {
	fmt.Println(person.Email, "person.Email")

	// Check if email already exists
	queryByEmail := database.Db.Collection(CollectionName).Where("Email", "==", person.Email).Limit(1)
	emailDocs, err := queryByEmail.Documents(ctx).GetAll()
	if err != nil {
		return nil, nil, err
	}
	if len(emailDocs) > 0 {
		return nil, nil, errors.New("email already exists")
	}

	// Check if ID already exists
	queryByID := database.Db.Collection(CollectionName).Where("ID", "==", person.ID).Limit(1)
	idDocs, err := queryByID.Documents(ctx).GetAll()
	if err != nil {
		return nil, nil, err
	}
	if len(idDocs) > 0 {
		return nil, nil, errors.New("ID already exists")
	}

	response, err := database.Db.Collection(CollectionName).Doc(person.ID).Set(ctx, person)
	if err != nil {
		return nil, nil, err
	}
	return &person, response, nil
}

func GetAllPersons(ctx context.Context) ([]dto.Person, error) {
	var persons []dto.Person
	documents, err := database.Db.Collection(CollectionName).Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get documents: %v", err)
	}

	// Iterate over the documents
	for _, doc := range documents {
		// Unmarshal the document data into a Person struct
		var person dto.Person
		if err := doc.DataTo(&person); err != nil {
			return nil, fmt.Errorf("failed to unmarshal document: %v", err)
		}
		persons = append(persons, person)
	}
	return persons, nil
}
