import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

class SocialAnalyticsView extends React.Component {
    render() {
        return (
            <div style={{display: this.props.isActiveView === true ? '' : 'none'}} basic>
                <div>
                    <div class="message"
                        header={'Analytics Page ' + this.props.socialName}
                        content='Coming soon.'
                      >
                    </div>
                </div>
            </div>
        );
    }
}

export default SocialAnalyticsView;