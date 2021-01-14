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
/**
 * 
 * @export
 * @interface ControllersFixcomtype
 */
export interface ControllersFixcomtype {
    /**
     * 
     * @type {string}
     * @memberof ControllersFixcomtype
     */
    fixcomtypename?: string;
}

export function ControllersFixcomtypeFromJSON(json: any): ControllersFixcomtype {
    return ControllersFixcomtypeFromJSONTyped(json, false);
}

export function ControllersFixcomtypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersFixcomtype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fixcomtypename': !exists(json, 'fixcomtypename') ? undefined : json['fixcomtypename'],
    };
}

export function ControllersFixcomtypeToJSON(value?: ControllersFixcomtype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fixcomtypename': value.fixcomtypename,
    };
}

