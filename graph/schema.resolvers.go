package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MohammadZeyaAhmad/golang-gql/database"
	"github.com/MohammadZeyaAhmad/golang-gql/graph/generated"
	"github.com/MohammadZeyaAhmad/golang-gql/graph/model"
)

// CreateJobListing is the resolver for the createJobListing field.
func (r *mutationResolver) CreateJobListing(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	return db.CreateJobListing(input), nil
}

// UpdateJobListing is the resolver for the updateJobListing field.
func (r *mutationResolver) UpdateJobListing(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	return db.UpdateJobListing(id, input), nil
}

// DeleteJobListing is the resolver for the deleteJobListing field.
func (r *mutationResolver) DeleteJobListing(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	return db.DeleteJobListing(id), nil
}

// Jobs is the resolver for the jobs field.
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	return db.GetJobs(), nil
}

// Job is the resolver for the job field.
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	return db.GetJob(id), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()