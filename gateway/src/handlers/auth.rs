use axum::{
    body::{to_bytes, Body},
    extract::{Path, State},
    http::{Request, StatusCode},
    response::IntoResponse,
};

use crate::state::AppState;

pub async fn proxy(
    Path(path): Path<String>,
    State(state): State<AppState>,
    request: Request<Body>,
) -> Result<impl IntoResponse, StatusCode> {

    let method = request.method().clone();

    let body = to_bytes(request.into_body(), usize::MAX)
        .await
        .map_err(|_| StatusCode::BAD_REQUEST)?;

    let url = format!(
        "{}/{}",
        state.config.auth_service_url,
        path,
    );

    println!("Forwarding {} {}", method, url);

    let response = state
        .http_client
        .request(method, url)
        .body(body)
        .send()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    let status = StatusCode::from_u16(response.status().as_u16())
        .unwrap();

    let body = response
        .text()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    Ok((status, body))
}