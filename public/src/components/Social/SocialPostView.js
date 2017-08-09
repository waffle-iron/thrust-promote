import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../../actions/data';
import { Link } from 'react-router';
import SocialContentView from './SocialContentView';
import Flash from '../../lib/Flash';
import Dropdown from '../../lib/Dropdown';

function mapStateToProps(state) {
    return {
        isRegistering: state.auth.isRegistering,
        registerStatusText: state.auth.registerStatusText,
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
  tabs: {
    paddingRight: 16,
    paddingLeft: 16,
  },
};

@connect(mapStateToProps, mapDispatchToProps)
class SocialView extends React.Component { // eslint-disable-line react/prefer-stateless-function
    constructor(props) {
        super(props);
        this.state = {
            activeItem: "twitter",
            connection: {
              "facebook": null,
              "twitter": null,
              "youtube": null,
              // "soundcloud": null
            },
            pages: null
        };
    }
    handleItemClick(e, { name }) {
        this.setState({
            activeItem: name
        })
    }
    getSocialData(){
      const token = localStorage.getItem('token');
      fetch('/api/social/all', {
        headers: {
          'Authorization': token
        }
      }).then(res => {
        return res.json();
      }).then(data => {

         var conn = data.reduce((result, item) => {
            result[item.provider] = true;
            return result;
          }, this.state.connection);
         this.setState({
            connection: conn
         })
      })
    }
    componentDidMount(){
        this.getSocialData();
    }
    updateSocialState(social, pageData) {
        // check if social is connected
        var conn = Object.assign({},
            this.state.connection,
            social
        );
        // assert connection is actually valid
        if(pageData){
          this.setState({
            connection: conn,
            pages: pageData
          });

        } else {
          this.setState({
            connection: conn
          });
        }
        this.getSocialData();
    }
    hidePageView() {
      this.setState({ pages: null });
    }
    openPage(route) {
        this.props.redirectToRoute(route);
    }
    render() {
        return (
            <div>
                <div className="top-section">
                    {
                        this.props.isFetching ?
                        <Flash message="Loading Social Data" />
                        : null
                    }
                </div>
                <div className="bottom-section">
                    <div className="overlay">
                        <div className="section-nav">
                            <div className="section-nav__brand">
                                Social
                            </div>
                            <div className="section-nav__left">
                                <ul>
                                    <li>
                                        <Dropdown
                                           defaultSelected={this.props.data.activeItem}
                                           items={this.props.data.items}
                                          />
                                    </li>
                                    <li>
                                        <a onClick={(e) => this.openPage('/social/post') }>Post</a>
                                    </li>
                                    <li>
                                        <a onClick={(e) => this.openPage('/social/settings') }>Settings</a>
                                    </li>
                                </ul>
                            </div>
                            <div className="section-nav__right">
                            </div>
                        </div>
                        <div className="card">
                          <SocialContentView
                             event={this.props.data.data.event}
                             socialName={this.state.activeItem} />
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default SocialView;