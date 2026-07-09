use axum::{
    extract::State,
    http::StatusCode,
    Json,
};

use crate::{
    models::{LoginRequest, LoginResponse},
    state::AppState,
};

pub async fn login_proxy(
    State(state): State<AppState>,
    Json(payload): Json<LoginRequest>,
) -> Result<Json<LoginResponse>, StatusCode> {

    let url = format!("{}/login", state.config.auth_service_url);

    let response = state
        .http_client
        .post(url)
        .json(&payload)
        .send()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    if !response.status().is_success() {
        return Err(StatusCode::UNAUTHORIZED);
    }

    let login_response = response
        .json::<LoginResponse>()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    Ok(Json(login_response))
}