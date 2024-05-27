// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::coreaccounts::*;

#[derive(Clone)]
pub struct ImmutableFoundryCreateNewResults {
    pub proxy: Proxy,
}

impl ImmutableFoundryCreateNewResults {
    // serial number of the newly created foundry
    pub fn foundry_sn(&self) -> ScImmutableUint32 {
        ScImmutableUint32::new(self.proxy.root(RESULT_FOUNDRY_SN))
    }
}

#[derive(Clone)]
pub struct MutableFoundryCreateNewResults {
    pub proxy: Proxy,
}

impl MutableFoundryCreateNewResults {
    pub fn new() -> MutableFoundryCreateNewResults {
        MutableFoundryCreateNewResults {
            proxy: results_proxy(),
        }
    }

    // serial number of the newly created foundry
    pub fn foundry_sn(&self) -> ScMutableUint32 {
        ScMutableUint32::new(self.proxy.root(RESULT_FOUNDRY_SN))
    }
}

#[derive(Clone)]
pub struct MapUint32ToImmutableBool {
    pub(crate) proxy: Proxy,
}

impl MapUint32ToImmutableBool {
    pub fn get_bool(&self, key: u32) -> ScImmutableBool {
        ScImmutableBool::new(self.proxy.key(&uint32_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableAccountFoundriesResults {
    pub proxy: Proxy,
}

impl ImmutableAccountFoundriesResults {
    // foundry serial numbers owned by the given account
    pub fn foundries(&self) -> MapUint32ToImmutableBool {
        MapUint32ToImmutableBool { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MapUint32ToMutableBool {
    pub(crate) proxy: Proxy,
}

impl MapUint32ToMutableBool {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_bool(&self, key: u32) -> ScMutableBool {
        ScMutableBool::new(self.proxy.key(&uint32_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableAccountFoundriesResults {
    pub proxy: Proxy,
}

impl MutableAccountFoundriesResults {
    pub fn new() -> MutableAccountFoundriesResults {
        MutableAccountFoundriesResults {
            proxy: results_proxy(),
        }
    }

    // foundry serial numbers owned by the given account
    pub fn foundries(&self) -> MapUint32ToMutableBool {
        MapUint32ToMutableBool { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct ImmutableAccountNFTAmountResults {
    pub proxy: Proxy,
}

impl ImmutableAccountNFTAmountResults {
    // amount of NFTs owned by the account
    pub fn amount(&self) -> ScImmutableUint32 {
        ScImmutableUint32::new(self.proxy.root(RESULT_AMOUNT))
    }
}

#[derive(Clone)]
pub struct MutableAccountNFTAmountResults {
    pub proxy: Proxy,
}

impl MutableAccountNFTAmountResults {
    pub fn new() -> MutableAccountNFTAmountResults {
        MutableAccountNFTAmountResults {
            proxy: results_proxy(),
        }
    }

    // amount of NFTs owned by the account
    pub fn amount(&self) -> ScMutableUint32 {
        ScMutableUint32::new(self.proxy.root(RESULT_AMOUNT))
    }
}

#[derive(Clone)]
pub struct ImmutableAccountNFTAmountInCollectionResults {
    pub proxy: Proxy,
}

impl ImmutableAccountNFTAmountInCollectionResults {
    // amount of NFTs in collection owned by the account
    pub fn amount(&self) -> ScImmutableUint32 {
        ScImmutableUint32::new(self.proxy.root(RESULT_AMOUNT))
    }
}

#[derive(Clone)]
pub struct MutableAccountNFTAmountInCollectionResults {
    pub proxy: Proxy,
}

impl MutableAccountNFTAmountInCollectionResults {
    pub fn new() -> MutableAccountNFTAmountInCollectionResults {
        MutableAccountNFTAmountInCollectionResults {
            proxy: results_proxy(),
        }
    }

    // amount of NFTs in collection owned by the account
    pub fn amount(&self) -> ScMutableUint32 {
        ScMutableUint32::new(self.proxy.root(RESULT_AMOUNT))
    }
}

#[derive(Clone)]
pub struct ArrayOfImmutableNftID {
    pub(crate) proxy: Proxy,
}

impl ArrayOfImmutableNftID {
    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_nft_id(&self, index: u32) -> ScImmutableNftID {
        ScImmutableNftID::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct ImmutableAccountNFTsResults {
    pub proxy: Proxy,
}

impl ImmutableAccountNFTsResults {
    // NFT IDs owned by the account
    pub fn nft_i_ds(&self) -> ArrayOfImmutableNftID {
        ArrayOfImmutableNftID { proxy: self.proxy.root(RESULT_NFT_I_DS) }
    }
}

#[derive(Clone)]
pub struct ArrayOfMutableNftID {
    pub(crate) proxy: Proxy,
}

impl ArrayOfMutableNftID {
    pub fn append_nft_id(&self) -> ScMutableNftID {
        ScMutableNftID::new(self.proxy.append())
    }

    pub fn clear(&self) {
        self.proxy.clear_array();
    }

    pub fn length(&self) -> u32 {
        self.proxy.length()
    }

    pub fn get_nft_id(&self, index: u32) -> ScMutableNftID {
        ScMutableNftID::new(self.proxy.index(index))
    }
}

#[derive(Clone)]
pub struct MutableAccountNFTsResults {
    pub proxy: Proxy,
}

impl MutableAccountNFTsResults {
    pub fn new() -> MutableAccountNFTsResults {
        MutableAccountNFTsResults {
            proxy: results_proxy(),
        }
    }

    // NFT IDs owned by the account
    pub fn nft_i_ds(&self) -> ArrayOfMutableNftID {
        ArrayOfMutableNftID { proxy: self.proxy.root(RESULT_NFT_I_DS) }
    }
}

#[derive(Clone)]
pub struct ImmutableAccountNFTsInCollectionResults {
    pub proxy: Proxy,
}

impl ImmutableAccountNFTsInCollectionResults {
    // NFT IDs in collection owned by the account
    pub fn nft_i_ds(&self) -> ArrayOfImmutableNftID {
        ArrayOfImmutableNftID { proxy: self.proxy.root(RESULT_NFT_I_DS) }
    }
}

#[derive(Clone)]
pub struct MutableAccountNFTsInCollectionResults {
    pub proxy: Proxy,
}

impl MutableAccountNFTsInCollectionResults {
    pub fn new() -> MutableAccountNFTsInCollectionResults {
        MutableAccountNFTsInCollectionResults {
            proxy: results_proxy(),
        }
    }

    // NFT IDs in collection owned by the account
    pub fn nft_i_ds(&self) -> ArrayOfMutableNftID {
        ArrayOfMutableNftID { proxy: self.proxy.root(RESULT_NFT_I_DS) }
    }
}

#[derive(Clone)]
pub struct MapTokenIDToImmutableBigInt {
    pub(crate) proxy: Proxy,
}

impl MapTokenIDToImmutableBigInt {
    pub fn get_big_int(&self, key: &ScTokenID) -> ScImmutableBigInt {
        ScImmutableBigInt::new(self.proxy.key(&token_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceResults {
    pub proxy: Proxy,
}

impl ImmutableBalanceResults {
    // balance per token ID, zero length indicates base token
    pub fn balances(&self) -> MapTokenIDToImmutableBigInt {
        MapTokenIDToImmutableBigInt { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MapTokenIDToMutableBigInt {
    pub(crate) proxy: Proxy,
}

impl MapTokenIDToMutableBigInt {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_big_int(&self, key: &ScTokenID) -> ScMutableBigInt {
        ScMutableBigInt::new(self.proxy.key(&token_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableBalanceResults {
    pub proxy: Proxy,
}

impl MutableBalanceResults {
    pub fn new() -> MutableBalanceResults {
        MutableBalanceResults {
            proxy: results_proxy(),
        }
    }

    // balance per token ID, zero length indicates base token
    pub fn balances(&self) -> MapTokenIDToMutableBigInt {
        MapTokenIDToMutableBigInt { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceBaseTokenResults {
    pub proxy: Proxy,
}

impl ImmutableBalanceBaseTokenResults {
    // amount of base tokens in the account
    pub fn balance(&self) -> ScImmutableUint64 {
        ScImmutableUint64::new(self.proxy.root(RESULT_BALANCE))
    }
}

#[derive(Clone)]
pub struct MutableBalanceBaseTokenResults {
    pub proxy: Proxy,
}

impl MutableBalanceBaseTokenResults {
    pub fn new() -> MutableBalanceBaseTokenResults {
        MutableBalanceBaseTokenResults {
            proxy: results_proxy(),
        }
    }

    // amount of base tokens in the account
    pub fn balance(&self) -> ScMutableUint64 {
        ScMutableUint64::new(self.proxy.root(RESULT_BALANCE))
    }
}

#[derive(Clone)]
pub struct ImmutableBalanceNativeTokenResults {
    pub proxy: Proxy,
}

impl ImmutableBalanceNativeTokenResults {
    // amount of native tokens in the account
    pub fn tokens(&self) -> ScImmutableBigInt {
        ScImmutableBigInt::new(self.proxy.root(RESULT_TOKENS))
    }
}

#[derive(Clone)]
pub struct MutableBalanceNativeTokenResults {
    pub proxy: Proxy,
}

impl MutableBalanceNativeTokenResults {
    pub fn new() -> MutableBalanceNativeTokenResults {
        MutableBalanceNativeTokenResults {
            proxy: results_proxy(),
        }
    }

    // amount of native tokens in the account
    pub fn tokens(&self) -> ScMutableBigInt {
        ScMutableBigInt::new(self.proxy.root(RESULT_TOKENS))
    }
}

#[derive(Clone)]
pub struct ImmutableGetAccountNonceResults {
    pub proxy: Proxy,
}

impl ImmutableGetAccountNonceResults {
    // account nonce
    pub fn account_nonce(&self) -> ScImmutableUint64 {
        ScImmutableUint64::new(self.proxy.root(RESULT_ACCOUNT_NONCE))
    }
}

#[derive(Clone)]
pub struct MutableGetAccountNonceResults {
    pub proxy: Proxy,
}

impl MutableGetAccountNonceResults {
    pub fn new() -> MutableGetAccountNonceResults {
        MutableGetAccountNonceResults {
            proxy: results_proxy(),
        }
    }

    // account nonce
    pub fn account_nonce(&self) -> ScMutableUint64 {
        ScMutableUint64::new(self.proxy.root(RESULT_ACCOUNT_NONCE))
    }
}

#[derive(Clone)]
pub struct MapTokenIDToImmutableBool {
    pub(crate) proxy: Proxy,
}

impl MapTokenIDToImmutableBool {
    pub fn get_bool(&self, key: &ScTokenID) -> ScImmutableBool {
        ScImmutableBool::new(self.proxy.key(&token_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableGetNativeTokenIDRegistryResults {
    pub proxy: Proxy,
}

impl ImmutableGetNativeTokenIDRegistryResults {
    // token IDs
    pub fn mapping(&self) -> MapTokenIDToImmutableBool {
        MapTokenIDToImmutableBool { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MapTokenIDToMutableBool {
    pub(crate) proxy: Proxy,
}

impl MapTokenIDToMutableBool {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_bool(&self, key: &ScTokenID) -> ScMutableBool {
        ScMutableBool::new(self.proxy.key(&token_id_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableGetNativeTokenIDRegistryResults {
    pub proxy: Proxy,
}

impl MutableGetNativeTokenIDRegistryResults {
    pub fn new() -> MutableGetNativeTokenIDRegistryResults {
        MutableGetNativeTokenIDRegistryResults {
            proxy: results_proxy(),
        }
    }

    // token IDs
    pub fn mapping(&self) -> MapTokenIDToMutableBool {
        MapTokenIDToMutableBool { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct ImmutableNativeTokenResults {
    pub proxy: Proxy,
}

impl ImmutableNativeTokenResults {
    // serialized foundry output
    pub fn foundry_output_bin(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(RESULT_FOUNDRY_OUTPUT_BIN))
    }
}

#[derive(Clone)]
pub struct MutableNativeTokenResults {
    pub proxy: Proxy,
}

impl MutableNativeTokenResults {
    pub fn new() -> MutableNativeTokenResults {
        MutableNativeTokenResults {
            proxy: results_proxy(),
        }
    }

    // serialized foundry output
    pub fn foundry_output_bin(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(RESULT_FOUNDRY_OUTPUT_BIN))
    }
}

#[derive(Clone)]
pub struct ImmutableNftDataResults {
    pub proxy: Proxy,
}

impl ImmutableNftDataResults {
    // serialized NFT data
    pub fn nft_data(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(RESULT_NFT_DATA))
    }
}

#[derive(Clone)]
pub struct MutableNftDataResults {
    pub proxy: Proxy,
}

impl MutableNftDataResults {
    pub fn new() -> MutableNftDataResults {
        MutableNftDataResults {
            proxy: results_proxy(),
        }
    }

    // serialized NFT data
    pub fn nft_data(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(RESULT_NFT_DATA))
    }
}

#[derive(Clone)]
pub struct ImmutableTotalAssetsResults {
    pub proxy: Proxy,
}

impl ImmutableTotalAssetsResults {
    // balance per token ID, zero length indicates base token
    pub fn assets(&self) -> MapTokenIDToImmutableBigInt {
        MapTokenIDToImmutableBigInt { proxy: self.proxy.clone() }
    }
}

#[derive(Clone)]
pub struct MutableTotalAssetsResults {
    pub proxy: Proxy,
}

impl MutableTotalAssetsResults {
    pub fn new() -> MutableTotalAssetsResults {
        MutableTotalAssetsResults {
            proxy: results_proxy(),
        }
    }

    // balance per token ID, zero length indicates base token
    pub fn assets(&self) -> MapTokenIDToMutableBigInt {
        MapTokenIDToMutableBigInt { proxy: self.proxy.clone() }
    }
}
