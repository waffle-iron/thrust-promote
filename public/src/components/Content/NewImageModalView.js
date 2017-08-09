import React from 'react';
import Modal from '../../lib/Modal';

import NewImageView from './NewImageView';

export default class NewImageModalView extends React.Component {
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
                <div className="modal-header">New Audio</div>
                <div className="modal-content">
                      <NewImageView closeModal={(e) => this.handleClose(e)} />
                </div>
          </Modal>
        );
    }
}