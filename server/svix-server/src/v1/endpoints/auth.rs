use aide::{
    axum::{routing::post_with, ApiRouter},
    transform::TransformOperation,
};
use axum::{
    extract::{Path, State},
    Json,
};
use schemars::JsonSchema;
use serde::{Deserialize, Serialize};
use svix_server_derive::aide_annotate;
use validator::Validate;

use crate::{
    core::{permissions, security::generate_app_token, types::FeatureFlagSet},
    db::models::application,
    error::{HttpError, Result},
    v1::{
        endpoints::{
            application::{create_app_from_app_in, ApplicationIn},
            message::validate_create_app_uid,
        },
        utils::{api_not_implemented, openapi_tag, ApplicationPath, ValidatedJson},
    },
    AppState,
};

fn login_url_example() -> &'static str {
    "https://app.svix.com/login#key=eyJhcHBJZCI6ICJhcHBfMXRSdFl"
}

fn token_example() -> &'static str {
    "appsk_kV3ts5tKPNJN4Dl25cMTfUNdmabxbX0O"
}

#[derive(Deserialize, Serialize, JsonSchema)]
pub struct DashboardAccessOut {
    #[schemars(url, example = "login_url_example", length(min = 1, max = 65_536))]
    pub url: String,
    #[schemars(example = "token_example")]
    pub token: String,
}

fn feature_flag_set_example() -> FeatureFlagSet {
    FeatureFlagSet::new()
}

#[derive(Deserialize, Serialize, Validate, JsonSchema)]
#[serde(rename_all = "camelCase")]
pub struct AppPortalAccessIn {
    /// The set of feature flags the created token will have access to.
    #[serde(default, skip_serializing_if = "FeatureFlagSet::is_empty")]
    #[schemars(example = "feature_flag_set_example")]
    pub feature_flags: FeatureFlagSet,
    /// Optionally creates a new application alongside the message.
    ///
    /// If the application id or uid that is used in the path already exists,
    /// this argument is ignored.
    #[validate]
    #[serde(skip_serializing_if = "Option::is_none")]
    pub application: Option<ApplicationIn>,
}

#[derive(Serialize, JsonSchema)]
pub struct AppPortalAccessOut {
    #[serde(flatten)]
    common_: DashboardAccessOut,
}

impl From<DashboardAccessOut> for AppPortalAccessOut {
    fn from(common_: DashboardAccessOut) -> Self {
        Self { common_ }
    }
}

/// Use this function to get magic links (and authentication codes) for connecting your users to the Consumer Application Portal.
#[aide_annotate(op_id = "v1.authentication.app-portal-access")]
async fn app_portal_access(
    State(AppState { ref db, cfg, .. }): State<AppState>,
    Path(ApplicationPath { app_id }): Path<ApplicationPath>,
    permissions::Organization { org_id }: permissions::Organization,
    ValidatedJson(data): ValidatedJson<AppPortalAccessIn>,
) -> Result<Json<AppPortalAccessOut>> {
    let app_from_path_app_id =
        application::Entity::secure_find_by_id_or_uid(org_id.clone(), app_id.to_owned())
            .one(db)
            .await?;

    let app = match (&data.application, app_from_path_app_id) {
        (None, None) => {
            return Err(
                HttpError::not_found(None, Some("Application not found".to_string())).into(),
            );
        }

        (_, Some(app_from_path_param)) => app_from_path_param,
        (Some(app_from_body), None) => {
            validate_create_app_uid(&app_id, app_from_body)?;
            let (app, _metadata) =
                create_app_from_app_in(db, app_from_body.to_owned(), org_id).await?;

            app
        }
    };

    let token = generate_app_token(
        &cfg.jwt_signing_config,
        app.org_id,
        app.id.clone(),
        data.feature_flags,
    )?;

    Ok(Json(AppPortalAccessOut::from(DashboardAccessOut {
        url: "https://docs.svix.com/app-portal/oss".to_owned(),
        token,
    })))
}

/// DEPRECATED: Please use `app-portal-access` instead.
///
/// Use this function to get magic links (and authentication codes) for connecting your users to the Consumer Application Portal.
#[aide_annotate(op_id = "v1.authentication.dashboard-access")]
async fn dashboard_access(
    state: State<AppState>,
    path: Path<ApplicationPath>,
    permissions: permissions::Organization,
) -> Result<Json<DashboardAccessOut>> {
    app_portal_access(
        state,
        path,
        permissions,
        ValidatedJson(AppPortalAccessIn {
            feature_flags: FeatureFlagSet::default(),
            application: None,
        }),
    )
    .await
    .map(|Json(AppPortalAccessOut { common_: out })| Json(out))
}

const LOGOUT_DESCRIPTION: &str = r#"
Logout an app token.

Trying to log out other tokens will fail.
"#;

fn logout_operation(op: TransformOperation<'_>) -> TransformOperation<'_> {
    op.id("logout_api_v1_auth_logout__post")
        .summary("Logout")
        .description(LOGOUT_DESCRIPTION)
}

pub fn router() -> ApiRouter<AppState> {
    let tag = openapi_tag("Authentication");
    ApiRouter::new()
        .api_route_with(
            "/auth/dashboard-access/:app_id",
            post_with(dashboard_access, dashboard_access_operation),
            &tag,
        )
        .api_route_with(
            "/auth/logout",
            post_with(api_not_implemented, logout_operation),
            &tag,
        )
        .api_route_with(
            "/auth/app-portal-access/:app_id",
            post_with(app_portal_access, app_portal_access_operation),
            tag,
        )
}
