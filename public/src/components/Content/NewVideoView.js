import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Dropzone from 'react-dropzone';
import * as actionCreators from '../../actions/data';
import classNames from 'classnames';

import Button from '../../lib/Button';
import Input from '../../lib/Input';
import TextArea from '../../lib/TextArea';

import UploadVideoFile from './UploadVideoFile';
import RenderStaticVideo from './RenderStaticVideo';

const styles = {
  chooseBox: {
    border: "1px #c9c9c9 solid",
    borderRadius: "5px",
    float: "left",
    display: "inline-block",
    height: "150px",
    width: "48%",
    textAlign: "center",
    padding: "70px",
    marginTop: "30px",
    marginLeft: "10px",
    color: "#c9c9c9",
    cursor: "pointer"
  },
}

const ChooseVideoView = (props) => {
    let _classes = classNames('choose-video-view', {
        'hide': !props.isActive,
    });
    return (
        <div className={_classes}>
            <div 
                className="video-file-upload-option"
                style={styles.chooseBox}
                onClick={(e) => props.onChoose('upload')}>
                Upload Video File
            </div>
            <div className="render-static-video-option"
                style={styles.chooseBox}
                onClick={(e) => props.onChoose('render')}>
                Create Static Video
            </div>
        </div>
    )
}

class NewVideoView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            view: 'default'
        }
    }
    chooseView(view) {
        this.setState((state, props) => {
            return {view: view}
        })
    }
    render() {
        // initial view is to choose either upload or create
        // then when one is chosen the correct component replaces
        // the choose view
        // which is the best way to implement?
        // since the same data will be passed to the other
        // does it make sense to have them children of teh Choose view
        // or does it mkae sense to have it this high a level
        return (
            <div>
                <div>
                    <ChooseVideoView 
                        isActive={!this.state.view || this.state.view == 'default'}
                        onChoose={this.chooseView.bind(this)}
                    />
                    <UploadVideoFile 
                        isActive={this.state.view == 'upload'} 
                        closeModal={this.props.closeModal}
                    />
                    <RenderStaticVideo 
                        isActive={this.state.view == 'render'} 
                        closeModal={this.props.closeModal}
                    />
                </div>
                {/*<div>
                    <Button 
                        style={{marginLeft: "15px"}}
                        onClick={(e) => this.props.closeModal(e)}
                        floatRight
                        >
                        Cancel
                    </Button>
                </div>*/}
            </div>
        );
    } 
}

export default NewVideoView;