import React from 'react';
import classNames from 'classnames';


export default class Button extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        let positive = this.props.positive ? true : false;
        let floatRight = this.props.floatRight ? true : false;
        let isActive = this.props.isActive
        if (isActive !== true && isActive !== false) {
            // if not explicitly set to false or true, set to true
            isActive = true;
        }
        let social = this.props.social;
        if (social)
            positive = false;
        let _classes = classNames('button', social, {
            'positive': positive,
            'floatRight': floatRight,
            'hide': !isActive
        });
        return (
            <button 
                className={_classes} 
                    {...this.props} >
            </button>
        )
    }
}