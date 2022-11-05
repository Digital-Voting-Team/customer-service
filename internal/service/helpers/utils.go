package helpers

import "net/http"

func GetIdsForGivenUser(r *http.Request, userId int64) (int64, int64, int64, error) {
	resultCustomer, err := CustomersQ(r).FilterByUserID(userId).Get()
	if err != nil {
		return 0, 0, 0, err
	}
	resultPerson, err := PersonsQ(r).FilterByID(resultCustomer.PersonID).Get()
	if err != nil {
		return 0, 0, 0, err
	}
	return resultCustomer.ID, resultPerson.ID, resultPerson.AddressID, nil
}
