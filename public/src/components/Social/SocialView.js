import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import * as actionCreators from '../../actions/data';
import { Link } from 'react-router';
import SocialContainerView from './SocialContainerView';
import Flash from '../../lib/Flash';
import Dropdown from '../../lib/Dropdown';

function mapStateToProps(state) {
    return {
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
            activeItem: {
                name: "twitter",
                social: "twitter-icon" 
            },
            items: [
              {
                name: "facebook",
                social: "facebook-icon"
              },
              {
                name: "youtube",
                social: "youtube-icon"
              },
              {
                name: "twitter",
                social: "twitter-icon",
              },
            ],
            pages: null
        };
    }
    handleItemClick(e, { name }) {
        this.setState({
            activeItem: name
        })
    }
    componentDidMount(){
        // this.getSocialData();
        this.props.fetchUserAndAllSocialData(this.props.token);
    }
    updateSocialState(social, pageData) {
        // check if social is connected
        var conn = Object.assign({},
            this.state.connection,
            social
        );
        // assert connection is actually valid
        if(pageData) {
          this.setState({
            connection: conn,
            pages: pageData
          });

        } else {
          this.setState({
            connection: conn
          });
        }
        // this.getSocialData();
    }
    hidePageView() {
      this.setState({ pages: null });
    }

    openPage(route) {
        this.props.redirectToRoute(route);
    }
    updateActiveItem(selectedItem) {
        this.setState((state, props) => {
          return {activeItem: selectedItem}
        })
    }
    render() {
        let data = this.props.data;
        if (!data.data) {
          data.data = {
            event: {
                message: "",
                sendAt: ""
            },
            account: {
                name: "",
                connection: "",
                page_id: "",
                page_access_token: "",
                page_name: ""
            }
          }
        }
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
                                           defaultSelected={this.state.activeItem}
                                           items={this.state.items}
                                           onChange={this.updateActiveItem.bind(this)}
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
                          <SocialContainerView
                            socialName={this.state.activeItem.name} 
                            activeTab={this.props.params.tab || 'settings'}
                            socialConnection={this.state.activeItem.name}
                            pages={this.state.pages}
                            hidePageView={this.hidePageView.bind(this)}
                            updateSocialState={this.updateSocialState.bind(this)}/> 
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default SocialView;