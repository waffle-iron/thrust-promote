import React from 'react';
import Modal from '../../lib/Modal';

import NewVideoView from './NewVideoView';

export default class NewVideoModalView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            open: false
        }
    }
    handleOpen(e){
        this.setState((state, props) => {
            return {open: true}
        });

    }
    handleClose(e) {
        if (e)
            e.preventDefault();

        this.setState((state, props) => {
            return {open: false}
        });

        if (this.props.onClose) {
            this.props.onClose();
        }
    }
    render() {
        let open = this.props.open || this.state.open
        return (
          <Modal 
            isOpen={open}
            containerClassName={"modal"}
            className={"modal-body"}
            onClose={this.handleClose.bind(this)}>
                <div className="modal-header">New Video</div>
                <div className="modal-content">
                      <NewVideoView closeModal={(e) => this.handleClose(e)} />
                </div>
          </Modal>
        );
    }
}