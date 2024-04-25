// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Savings/ent/job"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Job is the model entity for the Job schema.
type Job struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Date holds the value of the "date" field.
	Date string `json:"date,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Batch holds the value of the "batch" field.
	Batch int `json:"batch,omitempty"`
	// LastProcessedID holds the value of the "last_processed_id" field.
	LastProcessedID uint64 `json:"last_processed_id,omitempty"`
	// TotalProcessed holds the value of the "total_processed" field.
	TotalProcessed uint `json:"total_processed,omitempty"`
	// Data holds the value of the "data" field.
	Data         map[string]interface{} `json:"data,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Job) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case job.FieldData:
			values[i] = new([]byte)
		case job.FieldID, job.FieldBatch, job.FieldLastProcessedID, job.FieldTotalProcessed:
			values[i] = new(sql.NullInt64)
		case job.FieldName, job.FieldDate, job.FieldStatus:
			values[i] = new(sql.NullString)
		case job.FieldCreatedAt, job.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Job fields.
func (j *Job) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case job.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			j.ID = uint64(value.Int64)
		case job.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				j.CreatedAt = value.Time
			}
		case job.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				j.UpdatedAt = value.Time
			}
		case job.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				j.Name = value.String
			}
		case job.FieldDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				j.Date = value.String
			}
		case job.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				j.Status = value.String
			}
		case job.FieldBatch:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field batch", values[i])
			} else if value.Valid {
				j.Batch = int(value.Int64)
			}
		case job.FieldLastProcessedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_processed_id", values[i])
			} else if value.Valid {
				j.LastProcessedID = uint64(value.Int64)
			}
		case job.FieldTotalProcessed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_processed", values[i])
			} else if value.Valid {
				j.TotalProcessed = uint(value.Int64)
			}
		case job.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &j.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %w", err)
				}
			}
		default:
			j.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Job.
// This includes values selected through modifiers, order, etc.
func (j *Job) Value(name string) (ent.Value, error) {
	return j.selectValues.Get(name)
}

// Update returns a builder for updating this Job.
// Note that you need to call Job.Unwrap() before calling this method if this Job
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Job) Update() *JobUpdateOne {
	return NewJobClient(j.config).UpdateOne(j)
}

// Unwrap unwraps the Job entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Job) Unwrap() *Job {
	_tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Job is not a transactional entity")
	}
	j.config.driver = _tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Job) String() string {
	var builder strings.Builder
	builder.WriteString("Job(")
	builder.WriteString(fmt.Sprintf("id=%v, ", j.ID))
	builder.WriteString("created_at=")
	builder.WriteString(j.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(j.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(j.Name)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(j.Date)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(j.Status)
	builder.WriteString(", ")
	builder.WriteString("batch=")
	builder.WriteString(fmt.Sprintf("%v", j.Batch))
	builder.WriteString(", ")
	builder.WriteString("last_processed_id=")
	builder.WriteString(fmt.Sprintf("%v", j.LastProcessedID))
	builder.WriteString(", ")
	builder.WriteString("total_processed=")
	builder.WriteString(fmt.Sprintf("%v", j.TotalProcessed))
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", j.Data))
	builder.WriteByte(')')
	return builder.String()
}

// Jobs is a parsable slice of Job.
type Jobs []*Job
