use cosmwasm_schema::{cw_serde, QueryResponses};
use cosmwasm_std::{Binary, QueryRequest};

use babylon_bindings::BabylonQuery;

#[cw_serde]
pub struct InstantiateMsg {}

#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {
    #[returns(OwnerResponse)]
    Owner {},
    /// Queries the blockchain and returns the result untouched
    #[returns(ChainResponse)]
    Chain { request: QueryRequest<BabylonQuery> },
}

// We define a custom struct for each query response

#[cw_serde]
pub struct OwnerResponse {
    pub owner: String,
}

#[cw_serde]
pub struct ChainResponse {
    pub data: Binary,
}
