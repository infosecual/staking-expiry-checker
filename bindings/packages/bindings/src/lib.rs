mod querier;
mod query;
mod types;

pub use querier::BabylonQuerier;
pub use query::{
    BabylonQuery, BtcBaseHeaderResponse, BtcHeaderQueryResponse, BtcTipResponse,
    CurrentEpochResponse, LatestFinalizedEpochInfoResponse,
};
pub use types::{BtcBlockHeader, BtcBlockHeaderInfo, FinalizedEpochInfo};

// This export is added to all contracts that import this package, signifying that they require
// "babylon" support on the chain they run on.
#[no_mangle]
extern "C" fn requires_babylon() {}
