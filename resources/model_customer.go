/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Customer struct {
	Key
	Attributes    CustomerAttributes    `json:"attributes"`
	Relationships CustomerRelationships `json:"relationships"`
}
type CustomerResponse struct {
	Data     Customer `json:"data"`
	Included Included `json:"included"`
}

type CustomerListResponse struct {
	Data     []Customer `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustCustomer - returns Customer from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCustomer(key Key) *Customer {
	var customer Customer
	if c.tryFindEntry(key, &customer) {
		return &customer
	}
	return nil
}
