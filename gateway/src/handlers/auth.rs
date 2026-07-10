use axum::{Json, extract::State, http::StatusCode};

use crate::{
    models::{LoginRequest, LoginResponse, RegisterRequest, UserResponse},
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

pub async fn register_proxy(
    State(state): State<AppState>,
    Json(payload): Json<RegisterRequest>,
) -> Result<(StatusCode, String), StatusCode> {

    let url = format!("{}/register", state.config.auth_service_url);

    let response = state
        .http_client
        .post(url)
        .json(&payload)
        .send()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    let status = StatusCode::from_u16(response.status().as_u16())
        .unwrap_or(StatusCode::BAD_GATEWAY);

    let body = response
        .text()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    Ok((status, body))
}
