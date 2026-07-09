use axum::{
    routing::{get,post},
    Router,
};

use crate::{
    handlers::health,
    state::AppState,
};

use crate::handlers::{
    auth,
};

pub fn create_router(state: AppState) -> Router {
    Router::new()
        .route("/health", get(health::health))
        .route("/auth/login", post(auth::login_proxy))
        .with_state(state)
}