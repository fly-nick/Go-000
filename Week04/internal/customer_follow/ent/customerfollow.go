// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/fly-nick/Go-000/Week04/internal/customer_follow/ent/customerfollow"
	"github.com/facebook/ent/dialect/sql"
)

// CustomerFollow is the model entity for the CustomerFollow schema.
type CustomerFollow struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// StaffId holds the value of the "staffId" field.
	StaffId int64 `json:"staffId,omitempty"`
	// CustomerId holds the value of the "customerId" field.
	CustomerId int64 `json:"customerId,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreateTime holds the value of the "createTime" field.
	CreateTime time.Time `json:"createTime,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted int64 `json:"deleted,omitempty"`
	// DeleteBy holds the value of the "deleteBy" field.
	DeleteBy string `json:"deleteBy,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CustomerFollow) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // staffId
		&sql.NullInt64{},  // customerId
		&sql.NullString{}, // content
		&sql.NullTime{},   // createTime
		&sql.NullInt64{},  // deleted
		&sql.NullString{}, // deleteBy
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CustomerFollow fields.
func (cf *CustomerFollow) assignValues(values ...interface{}) error {
	if m, n := len(values), len(customerfollow.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	cf.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field staffId", values[0])
	} else if value.Valid {
		cf.StaffId = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field customerId", values[1])
	} else if value.Valid {
		cf.CustomerId = value.Int64
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field content", values[2])
	} else if value.Valid {
		cf.Content = value.String
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createTime", values[3])
	} else if value.Valid {
		cf.CreateTime = value.Time
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field deleted", values[4])
	} else if value.Valid {
		cf.Deleted = value.Int64
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field deleteBy", values[5])
	} else if value.Valid {
		cf.DeleteBy = value.String
	}
	return nil
}

// Update returns a builder for updating this CustomerFollow.
// Note that, you need to call CustomerFollow.Unwrap() before calling this method, if this CustomerFollow
// was returned from a transaction, and the transaction was committed or rolled back.
func (cf *CustomerFollow) Update() *CustomerFollowUpdateOne {
	return (&CustomerFollowClient{config: cf.config}).UpdateOne(cf)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (cf *CustomerFollow) Unwrap() *CustomerFollow {
	tx, ok := cf.config.driver.(*txDriver)
	if !ok {
		panic("ent: CustomerFollow is not a transactional entity")
	}
	cf.config.driver = tx.drv
	return cf
}

// String implements the fmt.Stringer.
func (cf *CustomerFollow) String() string {
	var builder strings.Builder
	builder.WriteString("CustomerFollow(")
	builder.WriteString(fmt.Sprintf("id=%v", cf.ID))
	builder.WriteString(", staffId=")
	builder.WriteString(fmt.Sprintf("%v", cf.StaffId))
	builder.WriteString(", customerId=")
	builder.WriteString(fmt.Sprintf("%v", cf.CustomerId))
	builder.WriteString(", content=")
	builder.WriteString(cf.Content)
	builder.WriteString(", createTime=")
	builder.WriteString(cf.CreateTime.Format(time.ANSIC))
	builder.WriteString(", deleted=")
	builder.WriteString(fmt.Sprintf("%v", cf.Deleted))
	builder.WriteString(", deleteBy=")
	builder.WriteString(cf.DeleteBy)
	builder.WriteByte(')')
	return builder.String()
}

// CustomerFollows is a parsable slice of CustomerFollow.
type CustomerFollows []*CustomerFollow

func (cf CustomerFollows) config(cfg config) {
	for _i := range cf {
		cf[_i].config = cfg
	}
}
