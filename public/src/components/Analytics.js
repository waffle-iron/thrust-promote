import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../actions/auth';

function mapStateToProps(state) {
    return {
        isRegistering: state.auth.isRegistering,
        registerStatusText: state.auth.registerStatusText,
    };
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators(actionCreators, dispatch);
}

@connect(mapStateToProps, mapDispatchToProps)
class Analytics extends React.Component { // eslint-disable-line react/prefer-stateless-function
    render() {
        return (
            <div>
                <div className="top-section">
                </div>
                <div className="bottom-section">
                    <div className="overlay">
                        <div className="section-nav">
                            <div className="section-nav__brand">
                                Analytics
                            </div>
                        </div>
                        <div className="card">
                            Coming Soon
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default Analytics;
