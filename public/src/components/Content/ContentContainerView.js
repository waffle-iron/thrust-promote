import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import NewSongView from './NewSongView';
import UploadedSongsView from './UploadedSongsView';

class ContentContainerView extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
          <div>
            <UploadedSongsView 
                isActiveView={this.props.activeTab === 'audio'}/>
          </div>
        );
    }
}

export default ContentContainerView;