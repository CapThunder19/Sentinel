use axum::{
    body::{Body, to_bytes},
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
    let (parts, body) = request.into_parts();

    let method = parts.method.clone();

    let headers = parts.headers.clone();

    let body = to_bytes(body, usize::MAX)
        .await
        .map_err(|_| StatusCode::BAD_REQUEST)?;

    let url = format!("{}/{}", state.config.auth_service_url, path,);

    println!("Forwarding {} {}", method, url);

    let mut builder = state.http_client.request(method, url);

    for (name, value) in &headers {
        builder = builder.header(name, value);
    }

    let response = builder
        .body(body)
        .send()
        .await
        .map_err(|_| StatusCode::BAD_GATEWAY)?;

    let status = StatusCode::from_u16(response.status().as_u16()).unwrap();

    let body = response.text().await.map_err(|_| StatusCode::BAD_GATEWAY)?;

    Ok((status, body))
}
