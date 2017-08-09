import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Dropzone from 'react-dropzone';
import * as actionCreators from '../../actions/data';

import Flash from '../../lib/Flash';
import Button from '../../lib/Button';

import ContentContainerView from './ContentContainerView';
import NewSongModalView from './NewSongModalView';
import NewImageModalView from './NewImageModalView';
import NewVideoModalView from './NewVideoModalView';

function mapStateToProps(state) {
    return {
        data: state.data,
        token: state.auth.token,
        loaded: state.data.loaded,
        isFetching: state.data.isFetching,
    };
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators(actionCreators, dispatch);
}

@connect(mapStateToProps, mapDispatchToProps)
class ContentView extends React.Component { // eslint-disable-line react/prefer-stateless-function
    constructor(props) {
        super(props);
        this.state = {
            audioModalOpen: false,
            imageModalOpen: false,
            videoModalOpen: false,
        }
    }
    openPage(route) {
        this.props.redirectToRoute(route);
    }
    openAudioModal(){
        this.setState((state, props) => {
            return {audioModalOpen: true}
        })
    }
    closeAudioModal() {
        this.setState((state, props) => {
            return {audioModalOpen: false}
        });
    }
    openImageModal() {
        this.setState((state, props) => {
            return {imageModalOpen: true}
        })
    }
    closeImageModal() {
        this.setState((state, props) => {
            return {imageModalOpen: false}
        });
    }
    openVideoModal() {
        this.setState((state, props) => {
            return {videoModalOpen: true}
        })
    }
    closeVideoModal() {
        this.setState((state, props) => {
            return {videoModalOpen: false}
        });
    }
    render() {
        /* TODO: Add a green "NEW SONG" UI button"
        * when clicked show a model for a new song
        * allow the user to click Audio Only or Image and Audio (for videos)
        * for Image and Audio have the option selected for createding
        * videos
        * for Audio Only have a Release aname and a track name
        * allow them to uplaod the file
        */

        return (
            <div>
                <div className="top-section">
                    {
                        this.props.isFetching ?
                        <Flash message="Loading Content Data" />
                        : null
                    }
                </div>
                <div className="bottom-section">
                    <div className="overlay">
                        <div className="section-nav">
                            <div className="section-nav__brand">
                                Content
                            </div>
                            <div className="section-nav__left">
                                <ul>
                                    <li>
                                        <a onClick={(e) => this.openPage('/content/audio') }>Audio</a>
                                    </li>
                                    <li>
                                        <a onClick={(e) => this.openPage('/content/images') }>Images</a>
                                    </li>
                                    <li>
                                        <a onClick={(e) => this.openPage('/content/video') }>Video</a>
                                    </li>
                                </ul>
                            </div>
                            <div className="section-nav__right">
                                 <Button
                                  isActive={!this.props.params.tab || this.props.params.tab === 'audio'}
                                  onClick={(e) => this.openAudioModal(e)}
                                  floatRight
                                  positive>
                                  New Audio
                                </Button>

                                 <Button
                                  isActive={this.props.params.tab === 'images'}
                                  onClick={(e) => this.openImageModal(e)}
                                  floatRight
                                  positive>
                                  New Image
                                </Button>

                                 <Button
                                  isActive={this.props.params.tab === 'video'}
                                  onClick={(e) => this.openVideoModal(e)}
                                  floatRight
                                  positive>
                                  New Video
                                </Button>
                            </div>
                        </div>
                        <div className="card">
                            <ContentContainerView
                                activeTab={this.props.params.tab || 'audio'}
                            />
                        </div>
                        <NewSongModalView 
                            open={this.state.audioModalOpen}
                            onClose={this.closeAudioModal.bind(this)}
                        />
                        <NewImageModalView 
                            open={this.state.imageModalOpen}
                            onClose={this.closeImageModal.bind(this)}
                        />
                        <NewVideoModalView 
                            open={this.state.videoModalOpen}
                            onClose={this.closeVideoModal.bind(this)}
                        />
                    </div>
                </div>
            </div>
        );
    }
}

export default ContentView;