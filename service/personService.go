package service

import (
	"context"
	"errors"
	"fmt"
	"handsOnGO/database"
	"handsOnGO/dto"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func GetPersonDetailsById(ctx context.Context, ID string) (*dto.Person, error) {
	response, err := database.Db.Collection(CollectionName).Doc(ID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var person dto.Person
	if err := response.DataTo(&person); err != nil {
		return nil, err
	}

	return &person, nil

}

func UpdatePersonDetailsById(ctx context.Context, person dto.Person, ID string) (*map[string]interface{}, error) {
	// Reference to the document by its ID
	docRef := database.Db.Collection(CollectionName).Doc(ID)

	// Get the snapshot of the document
	snapshot, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}

	// Check if the document exists
	if !snapshot.Exists() {
		return nil, status.Errorf(codes.NotFound, "Document %s not found", ID)
	}

	// Retrieve the existing person data from Firestore
	var existingPerson dto.Person
	if err := snapshot.DataTo(&existingPerson); err != nil {
		return nil, err
	}

	// Check if email already exists (only if it's being updated)
	if person.Email != "" && person.Email != existingPerson.Email {
		queryByEmail := database.Db.Collection(CollectionName).Where("Email", "==", person.Email).Limit(1)
		emailDocs, err := queryByEmail.Documents(ctx).GetAll()
		if err != nil {
			return nil, err
		}
		if len(emailDocs) > 0 {
			return nil, errors.New("email already exists")
		}
	}

	// Ensure the ID remains the same
	if person.ID != "" && person.ID != ID {
		return nil, errors.New("cannot change ID")
	}

	// Update only the non-zero and non-empty fields
	updateFields := make(map[string]interface{})
	updateFields["ID"] = ID // Assuming id is the field in Firestore
	if person.Name != "" {
		updateFields["Name"] = person.Name // Assuming Name is the field in Firestore
	} else {
		updateFields["Name"] = existingPerson.Name
	}
	if person.Email != "" {
		updateFields["Email"] = person.Email // Assuming Email is the field in Firestore
	} else {
		updateFields["Email"] = existingPerson.Email // Assuming Email is the field in Firestore
	}
	if person.Age != 0 {
		updateFields["Age"] = person.Age // Assuming Age is the field in Firestore
	} else {
		updateFields["Age"] = existingPerson.Age // Assuming Age is the field in Firestore
	}

	// Print the updated fields for verification
	fmt.Println("Updated Fields:")
	for key, value := range updateFields {
		fmt.Printf("%s: %v\n", key, value)
	}

	// Apply the update only if there are fields to update
	if len(updateFields) > 0 {
		_, err := docRef.Set(ctx, updateFields, firestore.MergeAll)
		if err != nil {
			return nil, err
		}
	} else {
		// If no fields to update, return an error
		return nil, status.Errorf(codes.InvalidArgument, "No fields to update")
	}

	// Return the updated fields
	return &updateFields, nil
}
