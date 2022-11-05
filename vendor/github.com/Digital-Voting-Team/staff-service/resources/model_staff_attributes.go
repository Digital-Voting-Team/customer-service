/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "time"

type StaffAttributes struct {
	EmploymentDate time.Time     `json:"employment_date"`
	Salary         float32       `json:"salary"`
	Status         *WorkerStatus `json:"status,omitempty"`
}
