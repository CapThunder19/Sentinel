use axum::{
    Router,
    routing::{get, post},
};

use crate::{handlers::health, state::AppState};

use crate::handlers::auth;

pub fn create_router(state: AppState) -> Router {
    Router::new()
        .route("/health", get(health::health))
        .route("/auth/login", post(auth::login_proxy))
        .route("/auth/register", post(auth::register_proxy))
        .with_state(state)
}
