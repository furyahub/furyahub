import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "stafiprotocol.stafihub.ledger";
export interface MsgAddNewPool {
    creator: string;
    denom: string;
    addr: string;
}
export interface MsgAddNewPoolResponse {
}
export interface MsgRemovePool {
    creator: string;
    denom: string;
    addr: string;
}
export interface MsgRemovePoolResponse {
}
export interface MsgSetEraUnbondLimit {
    creator: string;
    denom: string;
    limit: number;
}
export interface MsgSetEraUnbondLimitResponse {
}
export interface MsgSetInitBond {
    creator: string;
    denom: string;
    pool: string;
    amount: string;
    receiver: string;
}
export interface MsgSetInitBondResponse {
}
export interface MsgSetChainBondingDuration {
    creator: string;
    denom: string;
    era: number;
}
export interface MsgSetChainBondingDurationResponse {
}
export interface MsgSetPoolDetail {
    creator: string;
    denom: string;
    pool: string;
    subAccounts: string[];
    threshold: number;
}
export interface MsgSetPoolDetailResponse {
}
export interface MsgSetLeastBond {
    creator: string;
    denom: string;
    amount: string;
}
export interface MsgSetLeastBondResponse {
}
export interface MsgClearCurrentEraSnapShots {
    creator: string;
    denom: string;
}
export interface MsgClearCurrentEraSnapShotsResponse {
}
export interface MsgSetCommission {
    creator: string;
    commission: string;
}
export interface MsgSetCommissionResponse {
}
export interface MsgSetReceiver {
    creator: string;
    receiver: string;
}
export interface MsgSetReceiverResponse {
}
export interface MsgSetUnbondFee {
    creator: string;
    denom: string;
    value: string;
}
export interface MsgSetUnbondFeeResponse {
}
export interface MsgLiquidityUnbond {
    creator: string;
    pool: string;
    value: string;
    recipient: string;
}
export interface MsgLiquidityUnbondResponse {
}
export interface MsgSetUnbondCommission {
    creator: string;
    commission: string;
}
export interface MsgSetUnbondCommissionResponse {
}
export declare const MsgAddNewPool: {
    encode(message: MsgAddNewPool, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddNewPool;
    fromJSON(object: any): MsgAddNewPool;
    toJSON(message: MsgAddNewPool): unknown;
    fromPartial(object: DeepPartial<MsgAddNewPool>): MsgAddNewPool;
};
export declare const MsgAddNewPoolResponse: {
    encode(_: MsgAddNewPoolResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddNewPoolResponse;
    fromJSON(_: any): MsgAddNewPoolResponse;
    toJSON(_: MsgAddNewPoolResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddNewPoolResponse>): MsgAddNewPoolResponse;
};
export declare const MsgRemovePool: {
    encode(message: MsgRemovePool, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRemovePool;
    fromJSON(object: any): MsgRemovePool;
    toJSON(message: MsgRemovePool): unknown;
    fromPartial(object: DeepPartial<MsgRemovePool>): MsgRemovePool;
};
export declare const MsgRemovePoolResponse: {
    encode(_: MsgRemovePoolResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRemovePoolResponse;
    fromJSON(_: any): MsgRemovePoolResponse;
    toJSON(_: MsgRemovePoolResponse): unknown;
    fromPartial(_: DeepPartial<MsgRemovePoolResponse>): MsgRemovePoolResponse;
};
export declare const MsgSetEraUnbondLimit: {
    encode(message: MsgSetEraUnbondLimit, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetEraUnbondLimit;
    fromJSON(object: any): MsgSetEraUnbondLimit;
    toJSON(message: MsgSetEraUnbondLimit): unknown;
    fromPartial(object: DeepPartial<MsgSetEraUnbondLimit>): MsgSetEraUnbondLimit;
};
export declare const MsgSetEraUnbondLimitResponse: {
    encode(_: MsgSetEraUnbondLimitResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetEraUnbondLimitResponse;
    fromJSON(_: any): MsgSetEraUnbondLimitResponse;
    toJSON(_: MsgSetEraUnbondLimitResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetEraUnbondLimitResponse>): MsgSetEraUnbondLimitResponse;
};
export declare const MsgSetInitBond: {
    encode(message: MsgSetInitBond, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetInitBond;
    fromJSON(object: any): MsgSetInitBond;
    toJSON(message: MsgSetInitBond): unknown;
    fromPartial(object: DeepPartial<MsgSetInitBond>): MsgSetInitBond;
};
export declare const MsgSetInitBondResponse: {
    encode(_: MsgSetInitBondResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetInitBondResponse;
    fromJSON(_: any): MsgSetInitBondResponse;
    toJSON(_: MsgSetInitBondResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetInitBondResponse>): MsgSetInitBondResponse;
};
export declare const MsgSetChainBondingDuration: {
    encode(message: MsgSetChainBondingDuration, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetChainBondingDuration;
    fromJSON(object: any): MsgSetChainBondingDuration;
    toJSON(message: MsgSetChainBondingDuration): unknown;
    fromPartial(object: DeepPartial<MsgSetChainBondingDuration>): MsgSetChainBondingDuration;
};
export declare const MsgSetChainBondingDurationResponse: {
    encode(_: MsgSetChainBondingDurationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetChainBondingDurationResponse;
    fromJSON(_: any): MsgSetChainBondingDurationResponse;
    toJSON(_: MsgSetChainBondingDurationResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetChainBondingDurationResponse>): MsgSetChainBondingDurationResponse;
};
export declare const MsgSetPoolDetail: {
    encode(message: MsgSetPoolDetail, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetPoolDetail;
    fromJSON(object: any): MsgSetPoolDetail;
    toJSON(message: MsgSetPoolDetail): unknown;
    fromPartial(object: DeepPartial<MsgSetPoolDetail>): MsgSetPoolDetail;
};
export declare const MsgSetPoolDetailResponse: {
    encode(_: MsgSetPoolDetailResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetPoolDetailResponse;
    fromJSON(_: any): MsgSetPoolDetailResponse;
    toJSON(_: MsgSetPoolDetailResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetPoolDetailResponse>): MsgSetPoolDetailResponse;
};
export declare const MsgSetLeastBond: {
    encode(message: MsgSetLeastBond, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetLeastBond;
    fromJSON(object: any): MsgSetLeastBond;
    toJSON(message: MsgSetLeastBond): unknown;
    fromPartial(object: DeepPartial<MsgSetLeastBond>): MsgSetLeastBond;
};
export declare const MsgSetLeastBondResponse: {
    encode(_: MsgSetLeastBondResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetLeastBondResponse;
    fromJSON(_: any): MsgSetLeastBondResponse;
    toJSON(_: MsgSetLeastBondResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetLeastBondResponse>): MsgSetLeastBondResponse;
};
export declare const MsgClearCurrentEraSnapShots: {
    encode(message: MsgClearCurrentEraSnapShots, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgClearCurrentEraSnapShots;
    fromJSON(object: any): MsgClearCurrentEraSnapShots;
    toJSON(message: MsgClearCurrentEraSnapShots): unknown;
    fromPartial(object: DeepPartial<MsgClearCurrentEraSnapShots>): MsgClearCurrentEraSnapShots;
};
export declare const MsgClearCurrentEraSnapShotsResponse: {
    encode(_: MsgClearCurrentEraSnapShotsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgClearCurrentEraSnapShotsResponse;
    fromJSON(_: any): MsgClearCurrentEraSnapShotsResponse;
    toJSON(_: MsgClearCurrentEraSnapShotsResponse): unknown;
    fromPartial(_: DeepPartial<MsgClearCurrentEraSnapShotsResponse>): MsgClearCurrentEraSnapShotsResponse;
};
export declare const MsgSetCommission: {
    encode(message: MsgSetCommission, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetCommission;
    fromJSON(object: any): MsgSetCommission;
    toJSON(message: MsgSetCommission): unknown;
    fromPartial(object: DeepPartial<MsgSetCommission>): MsgSetCommission;
};
export declare const MsgSetCommissionResponse: {
    encode(_: MsgSetCommissionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetCommissionResponse;
    fromJSON(_: any): MsgSetCommissionResponse;
    toJSON(_: MsgSetCommissionResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetCommissionResponse>): MsgSetCommissionResponse;
};
export declare const MsgSetReceiver: {
    encode(message: MsgSetReceiver, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetReceiver;
    fromJSON(object: any): MsgSetReceiver;
    toJSON(message: MsgSetReceiver): unknown;
    fromPartial(object: DeepPartial<MsgSetReceiver>): MsgSetReceiver;
};
export declare const MsgSetReceiverResponse: {
    encode(_: MsgSetReceiverResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetReceiverResponse;
    fromJSON(_: any): MsgSetReceiverResponse;
    toJSON(_: MsgSetReceiverResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetReceiverResponse>): MsgSetReceiverResponse;
};
export declare const MsgSetUnbondFee: {
    encode(message: MsgSetUnbondFee, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetUnbondFee;
    fromJSON(object: any): MsgSetUnbondFee;
    toJSON(message: MsgSetUnbondFee): unknown;
    fromPartial(object: DeepPartial<MsgSetUnbondFee>): MsgSetUnbondFee;
};
export declare const MsgSetUnbondFeeResponse: {
    encode(_: MsgSetUnbondFeeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetUnbondFeeResponse;
    fromJSON(_: any): MsgSetUnbondFeeResponse;
    toJSON(_: MsgSetUnbondFeeResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetUnbondFeeResponse>): MsgSetUnbondFeeResponse;
};
export declare const MsgLiquidityUnbond: {
    encode(message: MsgLiquidityUnbond, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgLiquidityUnbond;
    fromJSON(object: any): MsgLiquidityUnbond;
    toJSON(message: MsgLiquidityUnbond): unknown;
    fromPartial(object: DeepPartial<MsgLiquidityUnbond>): MsgLiquidityUnbond;
};
export declare const MsgLiquidityUnbondResponse: {
    encode(_: MsgLiquidityUnbondResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgLiquidityUnbondResponse;
    fromJSON(_: any): MsgLiquidityUnbondResponse;
    toJSON(_: MsgLiquidityUnbondResponse): unknown;
    fromPartial(_: DeepPartial<MsgLiquidityUnbondResponse>): MsgLiquidityUnbondResponse;
};
export declare const MsgSetUnbondCommission: {
    encode(message: MsgSetUnbondCommission, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetUnbondCommission;
    fromJSON(object: any): MsgSetUnbondCommission;
    toJSON(message: MsgSetUnbondCommission): unknown;
    fromPartial(object: DeepPartial<MsgSetUnbondCommission>): MsgSetUnbondCommission;
};
export declare const MsgSetUnbondCommissionResponse: {
    encode(_: MsgSetUnbondCommissionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSetUnbondCommissionResponse;
    fromJSON(_: any): MsgSetUnbondCommissionResponse;
    toJSON(_: MsgSetUnbondCommissionResponse): unknown;
    fromPartial(_: DeepPartial<MsgSetUnbondCommissionResponse>): MsgSetUnbondCommissionResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    AddNewPool(request: MsgAddNewPool): Promise<MsgAddNewPoolResponse>;
    RemovePool(request: MsgRemovePool): Promise<MsgRemovePoolResponse>;
    SetEraUnbondLimit(request: MsgSetEraUnbondLimit): Promise<MsgSetEraUnbondLimitResponse>;
    SetInitBond(request: MsgSetInitBond): Promise<MsgSetInitBondResponse>;
    SetChainBondingDuration(request: MsgSetChainBondingDuration): Promise<MsgSetChainBondingDurationResponse>;
    SetPoolDetail(request: MsgSetPoolDetail): Promise<MsgSetPoolDetailResponse>;
    SetLeastBond(request: MsgSetLeastBond): Promise<MsgSetLeastBondResponse>;
    ClearCurrentEraSnapShots(request: MsgClearCurrentEraSnapShots): Promise<MsgClearCurrentEraSnapShotsResponse>;
    SetCommission(request: MsgSetCommission): Promise<MsgSetCommissionResponse>;
    SetReceiver(request: MsgSetReceiver): Promise<MsgSetReceiverResponse>;
    SetUnbondFee(request: MsgSetUnbondFee): Promise<MsgSetUnbondFeeResponse>;
    LiquidityUnbond(request: MsgLiquidityUnbond): Promise<MsgLiquidityUnbondResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    SetUnbondCommission(request: MsgSetUnbondCommission): Promise<MsgSetUnbondCommissionResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    AddNewPool(request: MsgAddNewPool): Promise<MsgAddNewPoolResponse>;
    RemovePool(request: MsgRemovePool): Promise<MsgRemovePoolResponse>;
    SetEraUnbondLimit(request: MsgSetEraUnbondLimit): Promise<MsgSetEraUnbondLimitResponse>;
    SetInitBond(request: MsgSetInitBond): Promise<MsgSetInitBondResponse>;
    SetChainBondingDuration(request: MsgSetChainBondingDuration): Promise<MsgSetChainBondingDurationResponse>;
    SetPoolDetail(request: MsgSetPoolDetail): Promise<MsgSetPoolDetailResponse>;
    SetLeastBond(request: MsgSetLeastBond): Promise<MsgSetLeastBondResponse>;
    ClearCurrentEraSnapShots(request: MsgClearCurrentEraSnapShots): Promise<MsgClearCurrentEraSnapShotsResponse>;
    SetCommission(request: MsgSetCommission): Promise<MsgSetCommissionResponse>;
    SetReceiver(request: MsgSetReceiver): Promise<MsgSetReceiverResponse>;
    SetUnbondFee(request: MsgSetUnbondFee): Promise<MsgSetUnbondFeeResponse>;
    LiquidityUnbond(request: MsgLiquidityUnbond): Promise<MsgLiquidityUnbondResponse>;
    SetUnbondCommission(request: MsgSetUnbondCommission): Promise<MsgSetUnbondCommissionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
