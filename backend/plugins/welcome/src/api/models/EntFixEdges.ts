/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntAdminrepair,
    EntAdminrepairFromJSON,
    EntAdminrepairFromJSONTyped,
    EntAdminrepairToJSON,
    EntBrand,
    EntBrandFromJSON,
    EntBrandFromJSONTyped,
    EntBrandToJSON,
    EntCustomer,
    EntCustomerFromJSON,
    EntCustomerFromJSONTyped,
    EntCustomerToJSON,
    EntFixcomtype,
    EntFixcomtypeFromJSON,
    EntFixcomtypeFromJSONTyped,
    EntFixcomtypeToJSON,
    EntPersonal,
    EntPersonalFromJSON,
    EntPersonalFromJSONTyped,
    EntPersonalToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntFixEdges
 */
export interface EntFixEdges {
    /**
     * 
     * @type {EntBrand}
     * @memberof EntFixEdges
     */
    brand?: EntBrand;
    /**
     * 
     * @type {EntCustomer}
     * @memberof EntFixEdges
     */
    customer?: EntCustomer;
    /**
     * Fix holds the value of the fix edge.
     * @type {Array<EntAdminrepair>}
     * @memberof EntFixEdges
     */
    fix?: Array<EntAdminrepair>;
    /**
     * 
     * @type {EntFixcomtype}
     * @memberof EntFixEdges
     */
    fixcomtype?: EntFixcomtype;
    /**
     * 
     * @type {EntPersonal}
     * @memberof EntFixEdges
     */
    personal?: EntPersonal;
}

export function EntFixEdgesFromJSON(json: any): EntFixEdges {
    return EntFixEdgesFromJSONTyped(json, false);
}

export function EntFixEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntFixEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'brand': !exists(json, 'brand') ? undefined : EntBrandFromJSON(json['brand']),
        'customer': !exists(json, 'customer') ? undefined : EntCustomerFromJSON(json['customer']),
        'fix': !exists(json, 'fix') ? undefined : ((json['fix'] as Array<any>).map(EntAdminrepairFromJSON)),
        'fixcomtype': !exists(json, 'fixcomtype') ? undefined : EntFixcomtypeFromJSON(json['fixcomtype']),
        'personal': !exists(json, 'personal') ? undefined : EntPersonalFromJSON(json['personal']),
    };
}

export function EntFixEdgesToJSON(value?: EntFixEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'brand': EntBrandToJSON(value.brand),
        'customer': EntCustomerToJSON(value.customer),
        'fix': value.fix === undefined ? undefined : ((value.fix as Array<any>).map(EntAdminrepairToJSON)),
        'fixcomtype': EntFixcomtypeToJSON(value.fixcomtype),
        'personal': EntPersonalToJSON(value.personal),
    };
}


