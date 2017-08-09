import React from 'react';
import Dropzone from 'react-dropzone';
import Button from '../../lib/Button';
import Input from '../../lib/Input';
import TextArea from '../../lib/TextArea';

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
    width: "100%",
  },
};

export default class UploadVideoFile extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            videoUploading: null,
            videoFile:  ''
        }
    }

    onVideoDrop(acceptedFiles, rejectedFiles) {
      acceptedFiles.forEach((file, idx) => {
        this.uploadImage(file);
        if (idx === 0) 
          this.setState({ videoUploading: acceptedFiles[0] });
      });
    }
    uploadVideo(file) {
        const token = localStorage.getItem('token');
        var data = new FormData();
        data.append('type', 'file');
        data.append('file', file);
        fetch('api/upload_video', {
            method: 'post',
            headers: {
                'Authorization': token
            },
            body: data
        }).then(res => {
            console.log('successful video upload')
            // will need the env/unstaged/file
            return res.json();
        }).then(data => {
            this.setState({
                videoFile: data.video_path,
                videoUploading: false
            })
        });
    }
    render() {
        return (
            <div className={this.props.isActive ? null : 'hide'}>
                Video Dropzone
                <br />
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
                    <div style={{minHeight: "180px"}}>
                        <Dropzone 
                            ref="videoDropzone"
                            accept="video/*"
                            style={styles.dropzone}
                            onDrop={this.onVideoDrop.bind(this)} >
                            <div 
                                style={styles.uploadBox}
                                className="upload-box upload-video floatLeft"
                            >
                                <i 
                                    name="video">
                                </i>
                                <br/>
                                Choose Video
                                {this.state.videoUploading ? 
                                    <div>
                                       Uploading Video... 
                                       <img src={this.state.videoUploading.preview} />
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