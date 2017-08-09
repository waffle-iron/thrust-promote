import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

class FacebookPageView extends React.Component {
  constructor(props) {
    super(props);
  }
  choosePage(e) {
      var token = localStorage.getItem('token');
      fetch('api/facebook/choose_page', {
        method: 'post',
        headers: {
          'Authorization': token
        },
        body: {
          'id': this.props.page.id,
          'access_token': this.props.page.access_token,
          'name': this.props.page.name
        }
      }).then(res => {
          console.log(res);
          // clear page data so the list no longer appears
          // TODO should analytics be in this page?? 
          this.props.hidePageView();

      });
  }
  render() {
      return (
          <li as='a' onClick={this.choosePage.bind(this)}>
              {this.props.page.name}
          </li>
      );
  }
}

export default FacebookPageView;