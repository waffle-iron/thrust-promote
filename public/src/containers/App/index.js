import React from 'react';
/* application components */
import { HeaderView } from '../../components/Header';
import { Footer } from '../../components/Footer';

/* global styles for app */
import './styles/app.scss';


class App extends React.Component { // eslint-disable-line react/prefer-stateless-function
    static propTypes = {
        children: React.PropTypes.node,
    };

    render() {
        return (
            <div>
                <div className="app-container">
                    <HeaderView location={this.props.location}/>
                    <div className="main">
                        {this.props.children}
                    </div>
                    <Footer />
                </div>
            </div>
        );
    }
}

export { App };
