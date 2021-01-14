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
    EntTypeproductEdges,
    EntTypeproductEdgesFromJSON,
    EntTypeproductEdgesFromJSONTyped,
    EntTypeproductEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTypeproduct
 */
export interface EntTypeproduct {
    /**
     * Typeproductname holds the value of the "Typeproductname" field.
     * @type {string}
     * @memberof EntTypeproduct
     */
    typeproductname?: string;
    /**
     * 
     * @type {EntTypeproductEdges}
     * @memberof EntTypeproduct
     */
    edges?: EntTypeproductEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntTypeproduct
     */
    id?: number;
}

export function EntTypeproductFromJSON(json: any): EntTypeproduct {
    return EntTypeproductFromJSONTyped(json, false);
}

export function EntTypeproductFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTypeproduct {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'typeproductname': !exists(json, 'Typeproductname') ? undefined : json['Typeproductname'],
        'edges': !exists(json, 'edges') ? undefined : EntTypeproductEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntTypeproductToJSON(value?: EntTypeproduct | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Typeproductname': value.typeproductname,
        'edges': EntTypeproductEdgesToJSON(value.edges),
        'id': value.id,
    };
}

