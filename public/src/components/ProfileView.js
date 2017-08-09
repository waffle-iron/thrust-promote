import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Dropzone from 'react-dropzone';
import * as actionCreators from '../actions/data';
import Input from '../lib/Input';
import TextArea from '../lib/TextArea';
import Button from '../lib/Button';
import Flash from '../lib/Flash';

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
class ProfileView extends React.Component { // eslint-disable-line react/prefer-stateless-function
    constructor(props) {
        super(props);
        this.state = {
            disabled: true,
        }
    }
    componentDidMount() {
        this.fetchData();
    }
    fetchData() {
        this.props.fetchUserAndCurrentArtist(this.props.token);
    }
    componentWillReceiveProps(nextProps) {
        // console.log("NEW PROPS", nextProps);
    }
    save(e) {
       // TODO save  
       console.log('[save] called');
       e.preventDefault();
       let newData = {
            name: this.props.data.data.name,
            description: this.props.data.data.description,
            profileImageUrl: this.props.data.data.profileImageUrl
        }
       this.props.saveArtistData(this.props.token, newData);
    }
    uploadImage(file) {
        const token = localStorage.getItem('token');
        var data = new FormData();
        data.append('type', 'file');
        data.append('file', file);
        fetch('api/artists/upload_profile_image', {
            method: 'post',
            headers: {
                'Authorization': token
            },
            body: data
        }).then(res => {
            console.log('successful image upload');
            return res.json()
        }).then(data => {
            this.props.updateProfileImage(data.image_path);
        });
    }
    onImageDrop(acceptedFiles, rejectedFiles) {
      acceptedFiles.forEach((file, idx) => {
        this.uploadImage(file);
        if (idx === 0) {
            this.setState((state, props) => {
                return { imageUploading: acceptedFiles[0] };
            });
        }
      });
    }
    changeValue(e, type) {
        const value = e.target.value;
        this.props.changeArtistData({[type]: value});
    }
    render() {
        let data = this.props.data;
        if (typeof data.data === 'undefined' || data.data === null) {
            data.data = {
                profileImageUrl: `https://api.adorable.io/avatars/250/1jlr39.png`,
                name: '',
                description: ''
            }
        }
        return (
            <div>
                <div className="top-section">
                    {
                        this.props.isFetching ?
                        <Flash message="Loading Artist Data" />
                        : null
                    }
                </div>
                <div className="bottom-section">
                    <div className="overlay">
                        <div className="section-nav">
                            <div className="section-nav__brand">
                                Artist Profile
                            </div>
                        </div>
                        <div className="card">
                            <div className="profile-image">
                                <Dropzone 
                                    ref="imageDropzone"
                                    className="dropzone"
                                    accept="image/*"
                                    onDrop={this.onImageDrop.bind(this)} >
                                    <div className="circle image-view">
                                        <img 
                                            className="profile-img" 
                                            src={data.data.profileImageUrl} />
                                    </div>
                                </Dropzone>
                            </div>
                            <form>
                                <div className="artist-name">
                                    <field>
                                        <label>Artist Name</label>
                                        <Input
                                          placeholder="Name"
                                          type="name"
                                          value={data.data.name}
                                          onChange={(e) => this.changeValue(e, 'name')}
                                        />
                                    </field>
                                </div>
                                <div className="artist-bio">
                                    <field>
                                        <label>Artist Bio</label>
                                        <TextArea
                                          placeholder="bio"
                                          type="bio"
                                          value={data.data.description}
                                          onChange={(e) => this.changeValue(e, 'description')}
                                        />
                                    </field>
                                </div>

                                <Button
                                  ref={(btn) => {this.saveButton = btn; }}
                                  style={{'marginTop': '30px'}}
                                  onClick={(e) => this.save(e)}
                                  floatRight
                                  positive>
                                  Save
                                </Button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

ProfileView.propTypes = {
    fetchProtectedData: React.PropTypes.func,
    loaded: React.PropTypes.bool,
    userName: React.PropTypes.string,
    data: React.PropTypes.any,
    token: React.PropTypes.string,
};
export default ProfileView;