type UserService interface {
    GetUserByID(ctx context.Context, id string) (*models.User, error)
    GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}