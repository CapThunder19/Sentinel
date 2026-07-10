mod config;
mod handlers;
mod models;
mod routes;
mod state;

use axum::serve;
use std::net::SocketAddr;
use tokio::net::TcpListener;
use tracing::info;

use crate::{config::Config, routes::create_router, state::AppState};

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();

    let config = Config::from_env();

    info!("Starting Sentinel Gateway...");

    let state = AppState {
        config: config.clone(),
        http_client: reqwest::Client::new(),
    };

    let app = create_router(state);

    let addr = SocketAddr::from(([127, 0, 0, 1], config.port));

    info!("Gateway listening on http://{}", addr);

    let listener = TcpListener::bind(addr)
        .await
        .expect("Failed to bind TCP listener");

    serve(listener, app).await.expect("Failed to start server");
}
