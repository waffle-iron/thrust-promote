import React from 'react';
import CSSTransitionGroup from 'react-transition-group/CSSTransitionGroup';

const Flash = ({...props}) => {
    return (
        <CSSTransitionGroup
          transitionName="flash-message"
          transitionAppear={true}
          transitionAppearTimeout={500}
          transitionEnterTimeout={500}
          transitionLeaveTimeout={1000}
          >
            <div className="flash-message">
                {props.message}
            </div>
        </CSSTransitionGroup>
    );
}

export default Flash;