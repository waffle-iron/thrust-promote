import React from 'react';
import classNames from 'classnames';


export default class Message extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            visible: true
        }
    }
    close() {
        this.setState((state, props) => {
            return {visible: false}
        })
    }
    render() {
        let _classes = classNames('message', {
            hide: !this.state.visible
        });
        return (
            <div className={_classes}>
                <span 
                    className="bfi flaticon-cancel close-icon"
                    onClick={this.close.bind(this)}
                    ></span>
                <div className="message-section">
                    <div className="message-header">
                        {this.props.header}
                    </div>
                    <div className="message-body">
                        {this.props.content}
                    </div>
                </div>
            </div>
        )
    }
}