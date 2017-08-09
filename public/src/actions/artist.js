import { 
    FETCH_ARTIST_DATA_REQUEST, 
    RECEIVE_ARTIST_DATA,
    ARTIST_DATA_FAILURE } from '../constants/index';
import { parseJSON } from '../utils/misc';
import { get_artist_data, data_about_user } from '../utils/http_functions';

export function receiveArtistData(data) {
    return {
        type: RECEIVE_ARTIST_DATA,
        payload: {
            data,
        },
    };
}

export function fetchArtistDataRequest() {
    return {
        type: FETCH_ARTIST_DATA_REQUEST,
    };
}

export function fetchArtistDataFailure(error) {
    return {
        type: ARTIST_DATA_FAILURE,
        payload: {
            status: error.status,
            statusText: error.statusText,
        },
    };
}

export function fetchArtistData(token) {
    return (dispatch) => {
        dispatch(fetchArtistDataRequest());
        get_artist_data(token)
            .then(parseJSON)
            .then(response => {
                dispatch(receiveArtistData(response));
            })
            .catch(error => {
                if (error) {
                    dispatch(fetchArtistDataFailure(error))
                }
            })

    }
}