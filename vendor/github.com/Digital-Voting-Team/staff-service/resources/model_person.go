/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Person struct {
	Key
	Attributes    PersonAttributes    `json:"attributes"`
	Relationships PersonRelationships `json:"relationships"`
}
type PersonResponse struct {
	Data     Person   `json:"data"`
	Included Included `json:"included"`
}

type PersonListResponse struct {
	Data     []Person `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustPerson - returns Person from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPerson(key Key) *Person {
	var person Person
	if c.tryFindEntry(key, &person) {
		return &person
	}
	return nil
}
