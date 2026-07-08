use axum::{
    routing::get,
    Router,
};

use crate::{
    handlers::health,
    state::AppState,
};

pub fn create_router(state: AppState) -> Router {
    Router::new()
        .route("/health", get(health::health))
        .with_state(state)
}