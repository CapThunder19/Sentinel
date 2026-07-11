use axum::{
    Router,
    routing::{any, get},
};

use crate::{handlers::health, state::AppState};

use crate::handlers::auth;

pub fn create_router(state: AppState) -> Router {
   Router::new()
    .route("/health", get(health::health))
    .route("/auth/{*path}", any(auth::proxy))
    .with_state(state)
}
