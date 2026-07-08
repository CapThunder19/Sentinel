use dotenvy::dotenv;
use std::env;

#[derive(Clone)]
pub struct Config {
    pub port: u16,
    pub auth_service_url: String,
}

impl Config {
    pub fn from_env() -> Self {
        dotenv().ok();

        Self {
            port: env::var("PORT")
                .unwrap_or_else(|_| "3000".to_string())
                .parse()
                .expect("Invalid PORT"),

            auth_service_url: env::var("AUTH_SERVICE_URL")
                .expect("AUTH_SERVICE_URL is required"),
        }
    }
}