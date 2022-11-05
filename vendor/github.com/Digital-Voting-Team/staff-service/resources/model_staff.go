/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Staff struct {
	Key
	Attributes    StaffAttributes    `json:"attributes"`
	Relationships StaffRelationships `json:"relationships"`
}
type StaffResponse struct {
	Data     Staff    `json:"data"`
	Included Included `json:"included"`
}

type StaffListResponse struct {
	Data     []Staff  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustStaff - returns Staff from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustStaff(key Key) *Staff {
	var staff Staff
	if c.tryFindEntry(key, &staff) {
		return &staff
	}
	return nil
}
