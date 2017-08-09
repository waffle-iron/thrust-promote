import { 
    RECEIVE_PROTECTED_DATA, 
    FETCH_PROTECTED_DATA_REQUEST,
    FETCH_ARTIST_DATA_REQUEST,
    RECEIVE_ARTIST_DATA,
    ARTIST_DATA_SUCCESS,
    ARTIST_DATA_FAILURE,
    SAVE_ARTIST_DATA_REQUEST,
    IMAGE_UPLOAD_SUCCESS,
    CONTACT_DATA_NEW,
    FETCH_CONTACT_DATA_REQUEST,
    SAVE_CONTACT_DATA_REQUEST,
    RECEIVE_CONTACT_DATA,
    CONTACT_DATA_SUCCESS,
    SOCIAL_DATA_NEW,
    FETCH_SOCIAL_DATA_REQUEST,
    SAVE_SOCIAL_DATA_REQUEST,
    RECEIVE_SOCIAL_DATA,
    SOCIAL_DATA_FAILURE,
    SOCIAL_DATA_SUCCESS,
} from '../constants';
import { createReducer } from '../utils/misc';

const initialState = {
    data: null,
    isFetching: false,
    loaded: false,
};

export default createReducer(initialState, {
    [RECEIVE_PROTECTED_DATA]: (state, payload) =>
        Object.assign({}, state, {
            data: payload.data,
            isFetching: false,
            loaded: true,
        }),
    [FETCH_PROTECTED_DATA_REQUEST]: (state) =>
        Object.assign({}, state, {
            isFetching: true,
        }),

    /*
        Artist Reducers
    */
    [ARTIST_DATA_SUCCESS]: (state, payload) => {
        return Object.assign({}, state, {
            data: {
                ...state.data,
                ...payload.data,
            }
        })
    },
    [IMAGE_UPLOAD_SUCCESS]: (state, payload) => {
        return Object.assign({}, state, {
            data: {
                ...state.data,
                ...payload.data,
            }
        })
    },
    [RECEIVE_ARTIST_DATA]: (state, payload) =>
        Object.assign({}, state, {
            data: payload.data,
            isFetching: false,
            loaded: true,
        }),
    [FETCH_ARTIST_DATA_REQUEST]: (state) =>
        Object.assign({}, state, {
            isFetching: true,
        }),

    /*
        Contact Reducers
    */

    [CONTACT_DATA_SUCCESS]: (state, payload) => {
        return Object.assign({}, state, {
            data: {
                ...state.data,
                ...payload.data,
            }
        })
    },
    [RECEIVE_CONTACT_DATA]: (state, payload) =>
        Object.assign({}, state, {
            data: payload.data,
            isFetching: false,
            loaded: true,
        }),
    [FETCH_CONTACT_DATA_REQUEST]: (state) =>
        Object.assign({}, state, {
            isFetching: true,
        }),
    /*
        Social Reducers
    */
    [SOCIAL_DATA_NEW]: (state, payload) => {
        return Object.assign({}, state, {
            activeItem: payload.activeItem,
            items: payload.items,
            data: {
                ...state.data,
                ...payload.data,
            }
        })
    },
    [SOCIAL_DATA_SUCCESS]: (state, payload) => {
        return Object.assign({}, state, {
            data: {
                ...state.data,
                ...payload.data,
            }
        });
    },
    [RECEIVE_SOCIAL_DATA]: (state, payload) =>
        Object.assign({}, state, {
            data: payload.data,
            isFetching: false,
            loaded: true,
        }),
    [FETCH_SOCIAL_DATA_REQUEST]: (state) =>
        Object.assign({}, state, {
            isFetching: true,
        }),
});
