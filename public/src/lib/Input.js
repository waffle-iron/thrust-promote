import React from 'react';


export default class Input extends React.Component {
    render() {
        return (
            <div>
                <input className="input" {...this.props} />
            </div>
        )
    }
}
