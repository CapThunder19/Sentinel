use axum::{
    extract::State,
    http::StatusCode,
};

use crate::state::AppState;

pub async fn login_proxy(
    State(state): State<AppState>,
) -> Result<String, StatusCode> {
    let url = format!("{}/", state.config.auth_service_url);

    let response = state
        .http_client
        .get(url)
        .send()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    let body = response
        .text()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    Ok(body)
}