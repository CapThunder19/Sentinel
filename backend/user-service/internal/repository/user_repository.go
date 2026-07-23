type UserRepository interface {
    GetByID(ctx context.Context, id string) (*models.User, error)
    GetByEmail(ctx context.Context, email string) (*models.User, error)
}