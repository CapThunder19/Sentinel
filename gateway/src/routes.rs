use axum::{
    routing::get,
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
        .route("/auth/test", get(auth::login_proxy))
        .with_state(state)
}