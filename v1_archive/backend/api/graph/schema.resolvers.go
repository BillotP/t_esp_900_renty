package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"context"
	"time"

	"github.com/BillotP/gorenty/v2/models"
)

func (r *candidatureResolver) Appointment(ctx context.Context, obj *models.Candidature) (*string, error) {
	var resp string
	if obj.Appointment != nil && !obj.Appointment.IsZero() {
		resp = obj.Appointment.Format(time.RFC3339)
		return &resp, nil
	}
	return nil, nil
}

func (r *offerorResolver) ID(ctx context.Context, obj *models.Offeror) (string, error) {
	return obj.DocumentMeta.Key, nil
}

func (r *rentOfferResolver) ID(ctx context.Context, obj *models.RentOffer) (string, error) {
	return obj.DocumentMeta.Key, nil
}

func (r *rentOfferResolver) CreatedAt(ctx context.Context, obj *models.RentOffer) (string, error) {
	return obj.Base.CreatedAt.Format(time.RFC3339), nil
}

func (r *rentOfferResolver) ExpiredAt(ctx context.Context, obj *models.RentOffer) (*string, error) {
	var resp string
	if !obj.ExpiredAt.IsZero() {
		resp = obj.ExpiredAt.Format(time.RFC3339)
		return &resp, nil
	}
	return nil, nil
}

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.DocumentMeta.Key, nil
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.Base.CreatedAt.Format(time.RFC3339), nil
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	return obj.Base.UpdatedAt.Format(time.RFC3339), nil
}

// Candidature returns generated.CandidatureResolver implementation.
func (r *Resolver) Candidature() generated.CandidatureResolver { return &candidatureResolver{r} }

// Offeror returns generated.OfferorResolver implementation.
func (r *Resolver) Offeror() generated.OfferorResolver { return &offerorResolver{r} }

// RentOffer returns generated.RentOfferResolver implementation.
func (r *Resolver) RentOffer() generated.RentOfferResolver { return &rentOfferResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type candidatureResolver struct{ *Resolver }
type offerorResolver struct{ *Resolver }
type rentOfferResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
