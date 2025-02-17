use crate::types::{BtcBlockHeaderInfo, FinalizedEpochInfo};
use cosmwasm_schema::{cw_serde, QueryResponses};
use cosmwasm_std::CustomQuery;

#[cw_serde]
#[derive(QueryResponses)]
pub enum BabylonQuery {
    #[returns(CurrentEpochResponse)]
    Epoch {},

    #[returns(LatestFinalizedEpochInfoResponse)]
    LatestFinalizedEpochInfo {},

    #[returns(BtcBaseHeaderResponse)]
    BtcBaseHeader {},

    #[returns(BtcTipResponse)]
    BtcTip {},

    #[returns(BtcHeaderQueryResponse)]
    BtcHeaderByHeight { height: u64 },

    #[returns(BtcHeaderQueryResponse)]
    BtcHeaderByHash { hash: String },
}

#[cw_serde]
pub struct CurrentEpochResponse {
    pub epoch: u64,
}
#[cw_serde]
pub struct LatestFinalizedEpochInfoResponse {
    pub epoch_info: FinalizedEpochInfo,
}

#[cw_serde]
pub struct BtcTipResponse {
    pub header_info: BtcBlockHeaderInfo,
}

#[cw_serde]
pub struct BtcBaseHeaderResponse {
    pub header_info: BtcBlockHeaderInfo,
}
#[cw_serde]
pub struct BtcHeaderQueryResponse {
    pub header_info: Option<BtcBlockHeaderInfo>,
}

impl CustomQuery for BabylonQuery {}
