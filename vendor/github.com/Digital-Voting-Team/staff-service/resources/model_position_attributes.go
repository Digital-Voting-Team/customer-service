/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PositionAttributes struct {
	// Guest -> 1; Worker -> 2; Accountant -> 3; Manager -> 4; Admin -> 5.
	AccessLevel AccessLevel `json:"access_level"`
	Name        string      `json:"name"`
}
