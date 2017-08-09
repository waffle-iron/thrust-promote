import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Dropzone from 'react-dropzone';
import * as actionCreators from '../actions/auth';
import BigCalendar from 'react-big-calendar';
import moment from 'moment';

import Flash from '../lib/Flash';
import 'react-big-calendar/lib/css/react-big-calendar.css';

BigCalendar.setLocalizer(
  BigCalendar.momentLocalizer(moment)
);


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

// http://stackoverflow.com/questions/34587067/change-color-of-react-big-calendar-events
@connect(mapStateToProps, mapDispatchToProps)
class EventsView extends React.Component { // eslint-disable-line react/prefer-stateless-function
    constructor(props) {
        super(props);
        let date = new Date();
        let year = date.getFullYear();
        let month = date.getMonth()
        this.state = {
            events: [],
            currentMonth: new Date(year, month, 1)
        }
    }

    componentDidMount() {
        this.fetchData();
        this.getEvents().then(e => {
            this.setState((state, props) => {
                return {events: e}
            })
        });
    }
    fetchData() {
        const token = this.props.token;
        // this.props.fetchProtectedData(token);
    }
    async getEvents() {
        const token = localStorage.getItem('token');
        var response = await fetch('api/events/generate', {
            headers: {
                'Authorization': token
            }
        });
        var data = await response.json();
        return data;
    }
    render() {
        return (
            <div>
                <div className="top-section">
                    {
                        this.props.isFetching ?
                        <Flash message="Loading Contact Data" />
                        : null
                    }
                </div>
                <div className="bottom-section">
                    <div className="overlay">
                        <div className="section-nav">
                            <div className="section-nav__brand">
                                Events
                            </div>
                            {
                                this.props.registerStatusText &&
                                    <div className="alert alert-info">
                                        {this.props.registerStatusText}
                                    </div>
                            }
                        </div>
                        <div className="card">

                            <BigCalendar
                                {...this.props}
                                events={this.state.events}
                                defaultDate={this.state.currentMonth}
                              />
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

EventsView.propTypes = {
    fetchProtectedData: React.PropTypes.func,
    loaded: React.PropTypes.bool,
    userName: React.PropTypes.string,
    data: React.PropTypes.any,
    token: React.PropTypes.string,
};
export default EventsView;