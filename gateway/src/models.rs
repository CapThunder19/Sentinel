use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct LoginRequest {
    pub email: String,
    pub password: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct LoginResponse {
    #[serde(rename = "AccessToken")]
    pub access_token: String,

    #[serde(rename = "TokenType")]
    pub token_type: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct RegisterRequest {
    pub username: String,
    pub email: String,
    pub password: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct UserResponse {
    #[serde(rename = "ID")]
    pub id: String,

    #[serde(rename = "Username")]
    pub username: String,

    #[serde(rename = "Email")]
    pub email: String,

    #[serde(rename = "Role")]
    pub role: String,

    #[serde(rename = "IsVerified")]
    pub is_verified: bool,
}
