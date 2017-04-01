import React from 'react';  
import Router from 'react-router';  
import { DefaultRoute, Link, Route, RouteHandler } from 'react-router';

import LoginHandler from './components/login.js';
import NewHandler from './components/new.js';

let App = React.createClass({  
  render() {
    return (
    	<div className="container">	
		  <div className="container nav-container">
		    <nav className="navbar">
		      <div className="container">
		        <ul className="nav-list">
		          <li>
		            <Link to="app">Home</Link>
		          </li>
		          <li>
		            <Link to="login">Login</Link>
		          </li>
		          <li>
		            <Link to="new">New</Link>
		          </li>
		        </ul>
		      </div>
		    </nav>
		  </div>
		  <div className="container main-container">
			  <RouteHandler/>
		  </div>
		</div>
    );
  }
});

let routes = (  
  <Route name="app" path="/" handler={App}>
    <Route name="login" path="/login" handler={LoginHandler}/>
    <Route name="new" path="/new" handler={NewHandler}/>
  </Route>
);

Router.run(routes, (Handler) => {  
  React.render(<Handler/>, document.body);
});