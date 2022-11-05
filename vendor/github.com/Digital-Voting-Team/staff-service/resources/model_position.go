/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Position struct {
	Key
	Attributes PositionAttributes `json:"attributes"`
}
type PositionResponse struct {
	Data     Position `json:"data"`
	Included Included `json:"included"`
}

type PositionListResponse struct {
	Data     []Position `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustPosition - returns Position from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPosition(key Key) *Position {
	var position Position
	if c.tryFindEntry(key, &position) {
		return &position
	}
	return nil
}
