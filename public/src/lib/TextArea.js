import React from 'react';


export default class TextArea extends React.Component {
    render() {
        return (
            <div>
                <textarea 
                    className="textarea" 
                    {...this.props} >
                </textarea>
            </div>
        )
    }
}
