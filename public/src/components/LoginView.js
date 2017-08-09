/* eslint camelcase: 0, no-underscore-dangle: 0 */

import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../actions/auth';
import { validateEmail } from '../utils/misc';

import Input from '../lib/Input';
import Button from '../lib/Button';

function mapStateToProps(state) {
    return {
        isAuthenticating: state.auth.isAuthenticating,
        statusText: state.auth.statusText,
    };
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators(actionCreators, dispatch);
}


const style = {
    marginTop: 50,
    paddingBottom: 50,
    paddingTop: 25,
    margin: "0 auto",
    width: '75%',
    display: 'inline-block',
};

@connect(mapStateToProps, mapDispatchToProps)
export default class LoginView extends React.Component {

    constructor(props) {
        super(props);
        const redirectRoute = '/login';
        this.state = {
            email: '',
            password: '',
            email_error_text: null,
            password_error_text: null,
            redirectTo: redirectRoute,
            disabled: true,
        };
    }

    isDisabled() {
        let email_is_valid = false;
        let password_is_valid = false;

        if (this.state.email === '') {
            this.setState({
                email_error_text: null,
            });
        } else if (validateEmail(this.state.email)) {
            email_is_valid = true;
            this.setState({
                email_error_text: null,
            });

        } else {
            this.setState({
                email_error_text: 'Sorry, this is not a valid email',
            });
        }

        if (this.state.password === '' || !this.state.password) {
            this.setState({
                password_error_text: null,
            });
        } else if (this.state.password.length >= 6) {
            password_is_valid = true;
            this.setState({
                password_error_text: null,
            });
        } else {
            this.setState({
                password_error_text: 'Your password must be at least 6 characters',
            });

        }

        if (email_is_valid && password_is_valid) {
            this.setState({
                disabled: false,
            });
        }

    }

    changeValue(e, type) {
        const value = e.target.value;
        const next_state = {};
        next_state[type] = value;
        this.setState(next_state, () => {
            this.isDisabled();
        });
    }

    _handleKeyPress(e) {
        if (e.key === 'Enter') {
            if (!this.state.disabled) {
                this.login(e);
            }
        }
    }

    login(e) {
        e.preventDefault();
        this.props.loginUser(this.state.email, this.state.password, this.state.redirectTo);
    }

    render() {
        return (
            <div>
                <div className="top-section">
                </div>
                <div className="bottom-section" onKeyPress={(e) => this._handleKeyPress(e)}>
                    <div className="overlay">
                        <div className="section-nav">
                            <h2>Login</h2>
                            {
                                this.props.statusText &&
                                    <div className="alert alert-info">
                                        {this.props.statusText}
                                    </div>
                            }
                        </div>
                        <div className="card">
                            <form role="form">
                                <div className="login-form">
                                    <field>
                                        <label>Email</label>
                                        <Input
                                          placeholder="Email"
                                          type="email"
                                          onChange={(e) => this.changeValue(e, 'email')}
                                        />
                                    </field>
                                    <field>
                                        <label>Password</label>
                                        <Input
                                          placeholder="Password"
                                          type="password"
                                          onChange={(e) => this.changeValue(e, 'password')}
                                        />
                                    </field>

                                    <Button
                                      disabled={this.state.disabled}
                                      style={{'marginTop': '30px'}}
                                      onClick={(e) => this.login(e)}
                                      floatRight
                                     positive>
                                     Submit
                                    </Button>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        );

    }
}

LoginView.propTypes = {
    loginUser: React.PropTypes.func,
    statusText: React.PropTypes.string,
};
