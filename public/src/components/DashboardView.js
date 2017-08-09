import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Dropzone from 'react-dropzone';
import * as actionCreators from '../actions/auth';

function mapStateToProps(state) {
    return {
        isRegistering: state.auth.isRegistering,
        registerStatusText: state.auth.registerStatusText,
        data: state.data,
        token: state.auth.token,
        loaded: state.data.loaded,
        isFetching: state.data.isFetching,
    };
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators(actionCreators, dispatch);
}

const styles = {
  paper: {
    width: "100%",
    height: "100%",
    padding: 20,
  },
  title: {
    textAlign: "center"
  },
};


@connect(mapStateToProps, mapDispatchToProps)
class DashboardView extends React.Component { // eslint-disable-line react/prefer-stateless-function
    constructor(props) {
        super(props);
        this.state = {
        }
    }

    componentDidMount() {
        this.fetchData();
    }
    fetchData() {
        const token = this.props.token;
        // this.props.fetchProtectedData(token);
    }
    render() {
        return (
            <div>
                <div className="top-section">
                </div>
                <div className="bottom-section">
                    <div className="overlay">
                        <div className="section-nav">
                            <div className="section-nav__brand">
                                Dashboard
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

DashboardView.propTypes = {
    fetchProtectedData: React.PropTypes.func,
    loaded: React.PropTypes.bool,
    userName: React.PropTypes.string,
    data: React.PropTypes.any,
    token: React.PropTypes.string,
};
export default DashboardView;