import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import FacebookPageView from './FacebookPageView';

class FacebookPages extends React.Component {
  constructor(props) {
    super(props);
  }
  render() {
    var pages = this.props.pages.map(page => {
      return (
          <FacebookPageView 
                hidePageView={this.props.hidePageView}
                page={page}/>
        );
    });
    return (
        <div>
          <div> Choose a facebook page to link</div>
          <ul>
            {pages}
          </ul>      
        </div>
    );
  }

}

export default FacebookPages;