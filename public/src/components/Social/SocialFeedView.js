import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

class SocialFeedView extends React.Component {
    render() {
        return (
            <div style={{display: this.props.isActiveView === true ? '' : 'none'}} basic>
                <div>
                    <div className="message"
                        header={'Feed page ' + this.props.socialName}
                        content='Coming soon.'
                      >
                    </div>
                </div>
            </div>
        );
    }
}

export default SocialFeedView;