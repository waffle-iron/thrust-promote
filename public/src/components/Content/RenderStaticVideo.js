import React from 'react';
import Dropzone from 'react-dropzone';
import Button from '../../lib/Button';
import Input from '../../lib/Input';
import TextArea from '../../lib/TextArea';
import Dropdown from '../../lib/Dropdown';

const styles = {
  uploadBox: {
    border: "1px #c9c9c9 solid",
    borderRadius: "5px",
    display: "inline-block",
    height: "150px",
    width: "100%",
    textAlign: "center",
    padding: "70px",
    marginTop: "30px",
    color: "#c9c9c9",
    cursor: "pointer"
  },
  dropzone: {
    border: "none",
    height: "150px",
  },
};


class RenderStaticVideo extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            releaseTitle: "Untitled",
            imageFile: null,
            audioFile: null,
            trackTitle: "Untitled Track",
            description: "",
            imageUploading: null,
            audioUploading: null,
        }
    }
    saveTrack(e) {
        // send request to save data
        // then close modal
        e.preventDefault();

        const token  = this.getToken();
        const currentState = this.state;
        const trackTitle = this.trackTitle.value;
        const description = this.description.value;

        fetch('api/save_track', {
            method: 'post',
            headers: {
                'Authorization': token,
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                release_title: currentState.releaseTitle,
                track_title: trackTitle,
                description: description,
                image_path: currentState.imageFile,
                audio_path: currentState.audioFile,
                render_video: currentState.renderVideo
            })
        }).then(res => {
            if (res.status === 200) {
                console.log('successful save');
            }
            this.props.closeModel();
        });
    }
    onImageDrop(acceptedFiles, rejectedFiles) {
      acceptedFiles.forEach((file, idx) => {
        this.uploadImage(file);
        if (idx === 0) 
          this.setState({ imageUploading: acceptedFiles[0] });
      });
    }
    onAudioDrop(acceptedFiles, rejectedFiles) {
      acceptedFiles.forEach((file, idx) => {
        this.uploadAudio(file);
        if (idx === 0)
            this.setState({ audioUploading: acceptedFiles[0] });
      });
    }
    uploadImage(file) {
        const token = localStorage.getItem('token');
        var data = new FormData();
        data.append('type', 'file');
        data.append('file', file);
        fetch('api/upload_image', {
            method: 'post',
            headers: {
                'Authorization': token
            },
            body: data
        }).then(res => {
            console.log('successful image upload')
            return res.json()
        }).then(data => {
            this.setState({
                imageFile: data.image_path,
                imageUploading: false
            })
        });
    }
    uploadAudio(file) {
        const token = localStorage.getItem('token');
        var data = new FormData();
        data.append('type', 'file');
        data.append('file', file);
        fetch('api/upload_audio', {
            method: 'post',
            headers: {
                'Authorization': token
            },
            body: data
        }).then(res => {
            console.log('successful image upload')
            // will need the env/unstaged/file
            return res.json();
        }).then(data => {
            this.setState({
                audioFile: data.audio_path,
                audioUploading: false
            })
        });
    }
    render() {
        return (
            <div className={this.props.isActive ? null : 'hide'}>
                Static Video
                <br/>
                    <form>
                        <div>
                            <Input 
                                placeholder="Song Title"
                                ref={(trackTitle) => this.trackTitle = trackTitle}
                            />
                            <br/>
                            <TextArea 
                                className="description" 
                                placeholder="Song Description"
                                ref={(desc) => this.description = desc}
                                autoHeight
                            />
                        </div>
                        <div style={{minHeight: "360px"}}
                            >
                            {/*<Dropdown />*/}
                            <Dropzone 
                                ref="imageDropzone"
                                accept="image/*"
                                style={styles.dropzone}
                                onDrop={this.onImageDrop.bind(this)} >
                                <div 
                                    style={styles.uploadBox}
                                    className="upload-box upload-image floatLeft"
                                >
                                    <i 
                                        name="image">
                                    </i>
                                    <br/>
                                    Choose Image
                                    {this.state.imageUploading ? 
                                        <div>
                                           Uploading Image... 
                                           <img src={this.state.imageUploading.preview} />
                                        </div>
                                    : null}
                                </div>
                            </Dropzone>
                            {/*<Dropdown />*/}
                            <Dropzone 
                                ref="audioDropzone"
                                accept="audio/aiff,audio/mp3,audio/wav"
                                style={styles.dropzone}
                                onDrop={this.onAudioDrop.bind(this)}>
                                <div 
                                    style={styles.uploadBox}
                                    className="upload-box upload-sound"
                                >
                                    <i 
                                        name="sound">
                                    </i>
                                    <br/>
                                    Choose Audio
                                    {this.state.audioUploading ? 
                                        <div>
                                           Uploading audio... 
                                           <audio 
                                                src={this.state.audioUploading.preview} 
                                                controls="controls">
                                           </audio>
                                        </div>
                                    : null}
                                </div>
                            </Dropzone>
                        </div>
                        <div className="bottom-form-buttons">
                            <Button 
                                style={{marginLeft: "15px"}}
                                floatRight
                                onClick={(e) => this.saveTrack(e)}
                                positive> 
                                Save
                            </Button>
                            <Button 
                                style={{marginLeft: "15px"}}
                                onClick={(e) => this.props.closeModal(e)}
                                floatRight
                                >
                                Cancel
                            </Button>
                        </div>
                    </form>
            </div>
        )
    }
}

export default RenderStaticVideo;