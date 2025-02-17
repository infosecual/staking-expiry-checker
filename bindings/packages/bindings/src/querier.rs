use cosmwasm_std::{QuerierWrapper, StdResult, Uint64};

use crate::query::{
    BabylonQuery, BtcBaseHeaderResponse, BtcHeaderQueryResponse, BtcTipResponse,
    CurrentEpochResponse, LatestFinalizedEpochInfoResponse,
};
use crate::types::{BtcBlockHeaderInfo, FinalizedEpochInfo};

pub struct BabylonQuerier<'a> {
    querier: &'a QuerierWrapper<'a, BabylonQuery>,
}

impl<'a> BabylonQuerier<'a> {
    pub fn new(querier: &'a QuerierWrapper<BabylonQuery>) -> Self {
        BabylonQuerier { querier }
    }

    pub fn current_epoch(&self) -> StdResult<Uint64> {
        let request = BabylonQuery::Epoch {}.into();
        let res: CurrentEpochResponse = self.querier.query(&request)?;
        Ok(Uint64::new(res.epoch))
    }

    pub fn latest_finalized_epoch_info(&self) -> StdResult<FinalizedEpochInfo> {
        let request = BabylonQuery::LatestFinalizedEpochInfo {}.into();
        let res: LatestFinalizedEpochInfoResponse = self.querier.query(&request)?;
        Ok(res.epoch_info)
    }

    pub fn btc_tip(&self) -> StdResult<BtcBlockHeaderInfo> {
        let request = BabylonQuery::BtcTip {}.into();
        let res: BtcTipResponse = self.querier.query(&request)?;
        Ok(res.header_info)
    }

    pub fn btc_base_header(&self) -> StdResult<BtcBlockHeaderInfo> {
        let request = BabylonQuery::BtcBaseHeader {}.into();
        let res: BtcBaseHeaderResponse = self.querier.query(&request)?;
        Ok(res.header_info)
    }

    pub fn btc_header_by_height(&self, height: u64) -> StdResult<Option<BtcBlockHeaderInfo>> {
        let request = BabylonQuery::BtcHeaderByHeight { height }.into();
        let res: BtcHeaderQueryResponse = self.querier.query(&request)?;
        Ok(res.header_info)
    }

    pub fn btc_header_by_hash(&self, hash: &str) -> StdResult<Option<BtcBlockHeaderInfo>> {
        let request = BabylonQuery::BtcHeaderByHash {
            hash: hash.to_string(),
        }
        .into();
        let res: BtcHeaderQueryResponse = self.querier.query(&request)?;
        Ok(res.header_info)
    }
}
