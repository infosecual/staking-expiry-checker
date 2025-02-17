use crate::msg::{ChainResponse, InstantiateMsg, OwnerResponse, QueryMsg};
use crate::state::{State, CONFIG};
use babylon_bindings::BabylonQuery;
use cosmwasm_std::{
    entry_point, to_json_binary, to_json_vec, ContractResult, Deps, DepsMut, Env, MessageInfo,
    QueryRequest, QueryResponse, Response, StdError, StdResult, SystemResult,
};

#[entry_point]
pub fn instantiate(
    deps: DepsMut<BabylonQuery>,
    _env: Env,
    info: MessageInfo,
    _msg: InstantiateMsg,
) -> StdResult<Response> {
    let state = State { owner: info.sender };
    CONFIG.save(deps.storage, &state)?;
    Ok(Response::default())
}

#[entry_point]
pub fn query(deps: Deps<BabylonQuery>, _env: Env, msg: QueryMsg) -> StdResult<QueryResponse> {
    match msg {
        QueryMsg::Owner {} => to_json_binary(&query_owner(deps)?),
        QueryMsg::Chain { request } => to_json_binary(&query_chain(deps, &request)?),
    }
}

fn query_owner(deps: Deps<BabylonQuery>) -> StdResult<OwnerResponse> {
    let state = CONFIG.load(deps.storage)?;
    let resp = OwnerResponse {
        owner: state.owner.into(),
    };
    Ok(resp)
}

fn query_chain(
    deps: Deps<BabylonQuery>,
    request: &QueryRequest<BabylonQuery>,
) -> StdResult<ChainResponse> {
    let raw = to_json_vec(request).map_err(|serialize_err| {
        StdError::generic_err(format!("Serializing QueryRequest: {}", serialize_err))
    })?;
    match deps.querier.raw_query(&raw) {
        SystemResult::Err(system_err) => Err(StdError::generic_err(format!(
            "Querier system error: {}",
            system_err
        ))),
        SystemResult::Ok(ContractResult::Err(contract_err)) => Err(StdError::generic_err(format!(
            "Querier contract error: {}",
            contract_err
        ))),
        SystemResult::Ok(ContractResult::Ok(value)) => Ok(ChainResponse { data: value }),
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::testing::{
        mock_env, mock_info, MockApi, MockQuerier, MockStorage, MOCK_CONTRACT_ADDR,
    };
    use cosmwasm_std::{coins, from_json, AllBalanceResponse, BankQuery, Coin};
    use cosmwasm_std::{OwnedDeps, SystemError};
    use std::marker::PhantomData;

    pub fn mock_dependencies(
        contract_balance: &[Coin],
    ) -> OwnedDeps<MockStorage, MockApi, MockQuerier<BabylonQuery>, BabylonQuery> {
        let custom_querier: MockQuerier<BabylonQuery> =
            MockQuerier::new(&[(MOCK_CONTRACT_ADDR, contract_balance)]).with_custom_handler(|_| {
                SystemResult::Err(SystemError::InvalidRequest {
                    error: "not implemented".to_string(),
                    request: Default::default(),
                })
            });
        OwnedDeps {
            storage: MockStorage::default(),
            api: MockApi::default(),
            querier: custom_querier,
            custom_query_type: PhantomData,
        }
    }

    #[test]
    fn proper_instantialization() {
        let mut deps = mock_dependencies(&[]);

        let msg = InstantiateMsg {};
        let info = mock_info("creator", &coins(1000, "earth"));

        // we can just call .unwrap() to assert this was a success
        let res = instantiate(deps.as_mut(), mock_env(), info, msg).unwrap();
        assert_eq!(0, res.messages.len());

        // it worked, let's query the state
        let value = query_owner(deps.as_ref()).unwrap();
        assert_eq!("creator", value.owner.as_str());
    }

    #[test]
    fn chain_query_works() {
        let deps = mock_dependencies(&coins(123, "ucosm"));

        // with bank query
        let msg = QueryMsg::Chain {
            request: BankQuery::AllBalances {
                address: MOCK_CONTRACT_ADDR.to_string(),
            }
            .into(),
        };
        let response = query(deps.as_ref(), mock_env(), msg).unwrap();
        let outer: ChainResponse = from_json(response).unwrap();
        let inner: AllBalanceResponse = from_json(outer.data).unwrap();
        assert_eq!(inner.amount, coins(123, "ucosm"));

        // TODO? or better in multitest?
        // // with custom query
        // let msg = QueryMsg::Chain {
        //     request: BabylonQuery::Ping {}.into(),
        // };
        // let response = query(deps.as_ref(), mock_env(), msg).unwrap();
        // let outer: ChainResponse = from_json(&response).unwrap();
        // let inner: SpecialResponse = from_json(&outer.data).unwrap();
        // assert_eq!(inner.msg, "pong");
    }

    // this mocks out what happens after reflect_subcall
}
