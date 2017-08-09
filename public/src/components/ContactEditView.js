import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import Dropzone from 'react-dropzone';
import * as actionCreators from '../actions/data';

import ReactTable from 'react-table';
import 'react-table/react-table.css'

import Input from '../lib/Input';
import TextArea from '../lib/TextArea';
import Button from '../lib/Button';

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

const TextCell = ({rowIndex, data, col, ...props}) => {
    return (
        <Cell {...props}>
            {data[rowIndex][col]}
        </Cell>
    )
}

@connect(mapStateToProps, mapDispatchToProps)
class ContactEditView extends React.Component { // eslint-disable-line react/prefer-stateless-function
    constructor(props) {
        super(props);
        this.state = {}
    }
    changeValue(e, type) {
        const value = e.target.value;
        this.props.changeContactData({[type]: value});
    }
    componentDidMount() {
        this.props.fetchContactData(this.props.token, this.props.params.id)
    }
    save(e) {
        // save entity
        // push history to newly created contact
        e.preventDefault();
        let newData = {
            name: this.props.data.data.name,
            email: this.props.data.data.email,
            links: this.props.data.data.links,
            source: this.props.data.data.source,
            notes: this.props.data.data.notes,
        }
        this.props.saveContactData(this.props.token, this.props.params.id, newData);
    }
    
    openPage(route) {
        this.props.redirectToRoute(route);
    }
    render() {
        let data = this.props.data;
        if (typeof data.data === 'undefined' || data.data === null) {
            data.data = {
                name: '',
                email: '',
                links: '',
                source: '',
                notes: '',
            }
        }
        return (
            <div>
                <div className="top-section">
                </div>
                <div className="bottom-section">
                    <div className="overlay">
                        <div className="section-nav">

                            <div className="section-nav__brand">
                                Edit Contact
                            </div>
                            <div className="section-nav__left">
                                <ul>
                                    <li>
                                        <a onClick={(e) => this.openPage('/contacts') }>Overview</a>
                                    </li>
                                    <li>
                                        <a onClick={(e) => this.openPage('/contacts/all') }>All</a>
                                    </li>
                                </ul>
                            </div>
                            <div className="section-nav__right">
                            </div>
                        </div>
                        <div className="card">
                            <form>
                                <div className="contact-name">
                                    <field>
                                        <label>Contact Name</label>
                                        <Input
                                          placeholder="Name"
                                          type="name"
                                          value={data.data.name}
                                          onChange={(e) => this.changeValue(e, 'name')}
                                        />
                                    </field>
                                </div>
                                <div className="contact-email">
                                    <field>
                                        <label>Contact Email</label>
                                        <Input
                                          placeholder="Email"
                                          type="email"
                                          value={data.data.email}
                                          onChange={(e) => this.changeValue(e, 'email')}
                                        />
                                    </field>
                                </div>

                                <div className="contact-links">
                                    <field>
                                        <label>Contact Links</label>
                                        <Input
                                          placeholder="Links"
                                          type="links"
                                          value={data.data.links}
                                          onChange={(e) => this.changeValue(e, 'links')}
                                        />
                                    </field>
                                </div>
                                <div className="contact-source">
                                    <field>
                                        <label>Contact Source</label>
                                        <Input
                                          placeholder="Source"
                                          type="source"
                                          value={data.data.source}
                                          onChange={(e) => this.changeValue(e, 'source')}
                                        />
                                    </field>
                                </div>
                                <div className="contact-tags">
                                    <field>
                                        <label>Contact Tags</label>
                                        <Input
                                          placeholder="Tags"
                                          type="tags"
                                          value={data.data.tags}
                                          onChange={(e) => this.changeValue(e, 'tags')}
                                        />
                                    </field>
                                </div>
                                <div className="contact-notes">
                                    <field>
                                        <label>Contact Notes</label>
                                        <TextArea
                                          placeholder="notes"
                                          type="notes"
                                          value={data.data.notes}
                                          onChange={(e) => this.changeValue(e, 'notes')}
                                        />
                                    </field>
                                </div>

                                <Button
                                  ref={(btn) => {this.saveButton = btn; }}
                                  style={{'marginTop': '30px'}}
                                  onClick={(e) => this.save(e)}
                                  floatRight
                                  positive>
                                  Save
                                </Button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

ContactEditView.propTypes = {
    fetchProtectedData: React.PropTypes.func,
    loaded: React.PropTypes.bool,
    userName: React.PropTypes.string,
    data: React.PropTypes.any,
    token: React.PropTypes.string,
};
export default ContactEditView;
