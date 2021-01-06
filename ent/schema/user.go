package schema

import (
	"github.com/facebook/ent"
	"github.com/fabiustech/workspace/pkg/ent/privacy"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
    return privacy.Policy{
        Mutation: privacy.MutationPolicy{
            // Deny if not set otherwise. 
            privacy.AlwaysDenyRule(),
        },
        Query: privacy.QueryPolicy{
            // Allow any viewer to read anything.
            privacy.AlwaysAllowRule(),
        },
    }
}

