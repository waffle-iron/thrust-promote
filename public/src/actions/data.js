import { browserHistory } from 'react-router';
import { 
    FETCH_PROTECTED_DATA_REQUEST, 
    RECEIVE_PROTECTED_DATA,
    FETCH_ARTIST_DATA_REQUEST, 
    RECEIVE_ARTIST_DATA,
    ARTIST_DATA_FAILURE,
    ARTIST_DATA_SUCCESS,
    IMAGE_UPLOAD_SUCCESS,
    SAVE_ARTIST_DATA_REQUEST,
    CONTACT_DATA_NEW,
    FETCH_CONTACT_DATA_REQUEST,
    SAVE_CONTACT_DATA_REQUEST,
    RECEIVE_CONTACT_DATA,
    CONTACT_DATA_FAILURE,
    CONTACT_DATA_SUCCESS,
    SOCIAL_DATA_NEW,
    FETCH_SOCIAL_DATA_REQUEST,
    SAVE_SOCIAL_DATA_REQUEST,
    RECEIVE_SOCIAL_DATA,
    SOCIAL_DATA_FAILURE,
    SOCIAL_DATA_SUCCESS,
} from '../constants/index';
import { parseJSON } from '../utils/misc';
import { 
    data_about_user, 
    get_artist_data,
    save_artist_data,
    get_contact_data,
    save_contact_data,
    get_contacts_list,
    get_contacts_overview,
    get_social_data,
} from '../utils/http_functions';
import { redirectToRoute, logoutAndRedirect } from './auth';

import _ from 'lodash';

export function receiveProtectedData(data) {
    return {
        type: RECEIVE_PROTECTED_DATA,
        payload: {
            data,
        },
    };
}

export function fetchProtectedDataRequest() {
    return {
        type: FETCH_PROTECTED_DATA_REQUEST,
    };
}

export function fetchProtectedData(token) {
    return (dispatch) => {
        dispatch(fetchProtectedDataRequest());
        data_about_user(token)
            .then(parseJSON)
            .then(response => {
                dispatch(receiveProtectedData(response.result));
            })
            .catch(error => {
                if (error.status === 401) {
                    dispatch(logoutAndRedirect(error));
                }
            });
    };
}

export function fetchUserData(token) {
    return (dispatch) => {
        dispatch(fetchProtectedDataRequest());
        return data_about_user(token).then(parseJSON)
                .then(response => {
                    dispatch(receiveProtectedData(response.result));
                })
                .catch(error => {
                    console.error(error);
                });
    };
}

export function receiveArtistData(data) {
    let profile_image_url = data.profile_image_url 
        ? data.profile_image_url 
        : `https://api.adorable.io/avatars/250/${data.id}.png`

    let newData = Object.assign({}, _.omit(data, 'profile_image_url'), {
        profileImageUrl: profile_image_url
    });
    return {
        type: RECEIVE_ARTIST_DATA,
        payload: {
            data: newData,
        }
    };
}

export function fetchArtistDataRequest() {
    return {
        type: FETCH_ARTIST_DATA_REQUEST,
    };
}

export function saveArtistDataRequest() {
    return {
        type: SAVE_ARTIST_DATA_REQUEST,
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

export function saveArtistDataSuccess(data) {

    let profile_image_url = data.profile_image_url 
        ? data.profile_image_url 
        : `https://api.adorable.io/avatars/250/${data.id}.png`

    let newData = Object.assign({}, _.omit(data, 'profile_image_url'), {
        profileImageUrl: profile_image_url
    });
    return {
        type: ARTIST_DATA_SUCCESS,
        payload: {
            data: newData
        },
    };
}

export function saveArtistDataFailure(res) {
    return {
        type: ARTIST_DATA_FAILURE,
        payload: {
            status: res.status,
            statusText: res.error,
        },
    };
}

export function updateProfileImage(imageUrl) {
    return {
        type: IMAGE_UPLOAD_SUCCESS,
        payload: {
            data: {
                profileImageUrl: imageUrl
            }
        }
    }
}

export function changeArtistData(data) {
    return {
        type: ARTIST_DATA_SUCCESS,
        payload: {
            data
        },
    };
}

export function fetchArtistData(token, artist_id) {
    return (dispatch) => {
        dispatch(fetchArtistDataRequest());
        get_artist_data(token, artist_id)
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

export function fetchUserAndCurrentArtist(token) {
    return (dispatch, getState) => {
        return dispatch(fetchUserData(token)).then(() => {
            const artist_id = getState().data.data.artist;
            return dispatch(fetchArtistData(token, artist_id));
        })
    }
}

export function saveArtistData(token, artist_data) {
    return (dispatch, getState) => {
        dispatch(saveArtistDataRequest());
        dispatch(fetchUserAndCurrentArtist(token)).then(() => {
            const artist_id = getState().data.data.artist;
            artist_data.profile_image_url = artist_data.profileImageUrl;
            return save_artist_data(token, artist_id, artist_data)
                .then(parseJSON)
                .then(response => {
                    try {
                        dispatch(saveArtistDataSuccess(response));
                    } catch (e) {
                        alert(e);
                        dispatch(saveArtistDataFailure({
                            response: {
                                status: 403,
                                statusText: 'Invalid token',
                            },
                        }));
                    }
                })
                .catch(error => {
                    dispatch(saveArtistDataFailure(error));
                })

        })
    }
}

// Contact Actions

export function saveContactDataRequest() {
    return {
        type: SAVE_CONTACT_DATA_REQUEST,
    };
}

export function saveContactDataSuccess(data) {
    return {
        type: CONTACT_DATA_SUCCESS,
        payload: {
            data
        },
    };
}

export function saveContactDataFailure(res) {
    return {
        type: CONTACT_DATA_FAILURE,
        payload: {
            status: res.status,
            statusText: res.error,
        },
    };
}

export function receiveContactData(data) {
    return {
        type: RECEIVE_CONTACT_DATA,
        payload: {
            data,
        }
    };
}

export function fetchContactDataRequest() {
    return {
        type: FETCH_CONTACT_DATA_REQUEST,
    };
}

export function fetchContactDataFailure(error) {
    return {
        type: CONTACT_DATA_FAILURE,
        payload: {
            status: error.status,
            statusText: error.statusText,
        },
    };
}

export function createEmptyContact() {
    return {
        type: CONTACT_DATA_NEW,
        payload: {
            data: {
                name: '',
                email: '',
                links: '',
                source: '',
                notes: '',
            }
        },
    };
}

export function changeContactData(data) {
    return {
        type: CONTACT_DATA_SUCCESS,
        payload: {
            data
        },
    };
}

export function fetchContactData(token, contact_id) {
    return (dispatch) => {
        dispatch(fetchContactDataRequest());
        get_contact_data(token, contact_id)
            .then(parseJSON)
            .then(response => {
                dispatch(receiveContactData(response));
            })
            .catch(error => {
                if (error) {
                    dispatch(fetchContactDataFailure(error))
                }
            })
    }
}

export function fetchAllContactData(token, artist_id) {
    return (dispatch) => {
        dispatch(fetchContactDataRequest());
        get_contacts_list(token, artist_id)
            .then(parseJSON)
            .then(response => {
                dispatch(receiveContactData(response));
            })
            .catch(error => {
                if (error) {
                    dispatch(fetchContactDataFailure(error))
                }
            })
    }
}

export function fetchOverviewContactData(token, artist_id) {
    return (dispatch) => {
        dispatch(fetchContactDataRequest());
        get_contacts_overview(token, artist_id)
            .then(parseJSON)
            .then(response => {
                dispatch(receiveContactData(response));
            })
            .catch(error => {
                if (error) {
                    dispatch(fetchContactDataFailure(error))
                }
            })
    }
}

export function saveNewContactWithData(token, contactData) {
    return (dispatch, getState) => {
        dispatch(saveContactDataRequest());
        return dispatch(fetchUserData(token)).then(() => {
            const artist_id = getState().data.data.artist;
            contactData.artist_id = artist_id;
            save_contact_data(token, null, contactData)
                .then(parseJSON)
                .then(response => {
                    try {
                        dispatch(saveContactDataSuccess(response));
                        console.log("[/api/contacts/new]response", response);
                        // redirectToRoute(`/contacts/${response.id}`);
                        browserHistory.push('/contacts/'+response.id);

                    } catch (e) {
                        alert(e);
                        dispatch(saveContactDataFailure({
                            response: {
                                status: 403,
                                statusText: 'Invalid token',
                            },
                        }));
                    }
                })
                .catch(error => {
                    dispatch(saveContactDataFailure(error));
                })
        })
    }
}

export function saveContactData(token, contactId, contactData) {
    return (dispatch, getState) => {
        dispatch(saveContactDataRequest());
        return dispatch(fetchUserData(token)).then(() => {
            const artist_id = getState().data.data.artist;
            contactData.artist_id = artist_id;
            save_contact_data(token, contactId, contactData)
                .then(parseJSON)
                .then(response => {
                    try {
                        dispatch(saveContactDataSuccess(response));
                    } catch (e) {
                        alert(e);
                        dispatch(saveContactDataFailure({
                            response: {
                                status: 403,
                                statusText: 'Invalid token',
                            },
                        }));
                    }
                })
                .catch(error => {
                    dispatch(saveContactDataFailure(error));
                })
        })
    }
}

export function fetchUserAndAllContactData(token) {
    return (dispatch, getState) => {
        return dispatch(fetchUserData(token)).then(() => {
            const artist_id = getState().data.data.artist;
            return dispatch(fetchAllContactData(token, artist_id));
        })
    }
}

export function fetchUserAndOverviewContactData(token) {
    return (dispatch, getState) => {
        return dispatch(fetchUserData(token)).then(() => {
            const artist_id = getState().data.data.artist;
            return dispatch(fetchOverviewContactData(token, artist_id));
        })
    }
}

export function receiveSocialData(data) {
    return {
        type: RECEIVE_SOCIAL_DATA,
        payload: {
            data,
        }
    };
}

export function fetchSocialDataRequest() {
    return {
        type: FETCH_SOCIAL_DATA_REQUEST,
    };
}

export function fetchSocialDataFailure(error) {
    return {
        type: SOCIAL_DATA_FAILURE,
        payload: {
            status: error.status,
            statusText: error.statusText,
        },
    };
}


export function createEmptySocialData() {
    return {
        type: SOCIAL_DATA_NEW,
        payload: {
            data: {
                activeItem: {
                    name: "twitter",
                    social: "twitter-icon" 
                },
                items: [
                  {
                    name: "facebook",
                    social: "facebook-icon"
                  },
                  {
                    name: "youtube",
                    social: "youtube-icon"
                  },
                  {
                    name: "twitter",
                    social: "twitter-icon",
                  },
                ],
                data: {
                    event: {
                        message: "",
                        sendAt: ""
                    },
                    account: {
                        name: "",
                        connection: "",
                        page_id: "",
                        page_access_token: "",
                        page_name: ""
                    }
                }
            }
        },
    };
}

export function changeSocialData(activeItem) {
    let items = [
        {
            name: "facebook",
            social: "facebook-icon"
        },
        {
            name: "youtube",
            social: "youtube-icon"
        },
        {
            name: "twitter",
            social: "twitter-icon",
        },
    ]
    return {
        type: SOCIAL_DATA_SUCCESS,
        payload: {
            data: {
                activeItem,
                items
            }
        },
    };
}

export function fetchAllSocialData(token, artist_id) {
    return (dispatch) => {
        dispatch(fetchSocialDataRequest());
        get_social_data(token, artist_id)
            .then(parseJSON)
            .then(response => {
                dispatch(receiveSocialData(response));
            })
            .catch(error => {
                if (error) {
                    dispatch(fetchSocialDataFailure(error));
                }
            })
    }
}

export function fetchUserAndAllSocialData(token) {
    return (dispatch, getState) => {
        return dispatch(fetchUserData(token)).then(() => {
            const artist_id = getState().data.data.artist;
            return dispatch(fetchAllSocialData(token, artist_id));
        })
    }
}