/**
 * Wasp API
 * REST API for the Wasp node
 *
 * OpenAPI spec version: 0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { BurnRecord } from '../models/BurnRecord';
import { RequestJSON } from '../models/RequestJSON';
import { UnresolvedVMErrorJSON } from '../models/UnresolvedVMErrorJSON';
import { HttpFile } from '../http/http';

export class ReceiptResponse {
    'blockIndex': number;
    'errorMessage'?: string;
    /**
    * The gas budget (uint64 as string)
    */
    'gasBudget': string;
    'gasBurnLog': Array<BurnRecord>;
    /**
    * The burned gas (uint64 as string)
    */
    'gasBurned': string;
    /**
    * The charged gas fee (uint64 as string)
    */
    'gasFeeCharged': string;
    'rawError'?: UnresolvedVMErrorJSON;
    'request': RequestJSON;
    'requestIndex': number;
    /**
    * Storage deposit charged (uint64 as string)
    */
    'storageDepositCharged': string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "blockIndex",
            "baseName": "blockIndex",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "errorMessage",
            "baseName": "errorMessage",
            "type": "string",
            "format": "string"
        },
        {
            "name": "gasBudget",
            "baseName": "gasBudget",
            "type": "string",
            "format": "string"
        },
        {
            "name": "gasBurnLog",
            "baseName": "gasBurnLog",
            "type": "Array<BurnRecord>",
            "format": ""
        },
        {
            "name": "gasBurned",
            "baseName": "gasBurned",
            "type": "string",
            "format": "string"
        },
        {
            "name": "gasFeeCharged",
            "baseName": "gasFeeCharged",
            "type": "string",
            "format": "string"
        },
        {
            "name": "rawError",
            "baseName": "rawError",
            "type": "UnresolvedVMErrorJSON",
            "format": ""
        },
        {
            "name": "request",
            "baseName": "request",
            "type": "RequestJSON",
            "format": ""
        },
        {
            "name": "requestIndex",
            "baseName": "requestIndex",
            "type": "number",
            "format": "int32"
        },
        {
            "name": "storageDepositCharged",
            "baseName": "storageDepositCharged",
            "type": "string",
            "format": "string"
        }    ];

    static getAttributeTypeMap() {
        return ReceiptResponse.attributeTypeMap;
    }

    public constructor() {
    }
}
