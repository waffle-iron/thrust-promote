import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import SocialContentView from './SocialContentView';
import SocialSettingsView from './SocialSettingsView';

class SocialContainerView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            activeItem: 'settings'
        }
    }
    handleItemClick(e, { name }) {
        this.setState({ activeItem: name });
    }
    render() {
        return (
          <div>
            <SocialContentView
               socialName={this.props.socialName} 
               isActiveView={this.props.activeTab === 'post'}/>
            <SocialSettingsView
              socialName={this.props.socialName} 
              socialConnection={this.props.socialConnection}
              pages={this.props.pages}
              hidePageView={this.props.hidePageView}
              updateSocialState={this.props.updateSocialState}
              isActiveView={this.props.activeTab === 'settings'}/>
          </div>
        );
    }
}

export default SocialContainerView;