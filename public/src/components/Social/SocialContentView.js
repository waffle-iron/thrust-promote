import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import Button from '../../lib/Button';
import TextArea from '../../lib/TextArea';

var styles = {
    float: 'left'
}

class SocialContentView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            twitterCount: 140
        }
    }
    toTitleCase(str) {
        return str.split(' ').map(i =>
            i[0].toUpperCase() + i.substring(1).toLowerCase()
        ).join(' ');
    }
    handleKeyPress(e) {
        var newLength = 140 - this.content.ref.value.length;
        styles.color = newLength >= 0 ? 'green' : 'red';
        this.setState({ 
            twitterCount: newLength
        });
    }
    render() {
        const options  = [
            { key: 'now', text: 'Send Now', value: 'now' },
            { key: 'later', text: 'Send Later', value: 'later' },
        ];
        return (
            <div className={this.props.isActiveView ? null : 'hide' }>
                <div>
                    <form>
                        <TextArea 
                            placeholder="Message Content"
                            onKeyPress={this.handleKeyPress.bind(this)}
                            ref={(content) => this.content = content}
                        />
                        <br/>
                        <Button
                            social={this.props.socialName}
                            style={{float: 'right'}}
                            target="_blank">
                            {"Send to " + this.props.socialName}
                        </Button>

                        {
                            this.props.socialName === 'twitter' ? 
                                <span style={{float: 'left'}}>
                                    {this.state.twitterCount}
                                </span>
                            :
                                null
                        }
                    </form>
                </div>
            </div>
        );
    }
}

export default SocialContentView;