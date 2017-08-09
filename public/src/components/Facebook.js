import React from 'react';
import axios from 'axios';


class Facebook extends React.Component {
    componentDidMount() {
      window.fbAsyncInit = function(){
        FB.init(appId: '<%= ENV["FACEBOOK_APP_ID"] %>', cookie: true)
      }
    }
    signIn(e) {
        e.preventDefault();
        FB.login(function(response){
          window.location = '/auth/facebook/callback' if response.authResponse
        });
    }

    signOut(e) {
        FB.getLoginStatus(function(response){
          FB.logout() if response.authResponse
        }
        return true;
    }
  
    render() {
        return (
            <div id="fb-root">
            </div>
        );
    }
}