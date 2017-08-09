import React from 'react';
import classNames from 'classnames';


export default class Dropdown extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            visible: false,
            selected: null
        }
    }
    show() {
        this.setState((state, props) => {
            return {visible: true};
        })
    }
    hide() {
        this.setState((state, props) => {
            return {visible: false};
        })
    }
    selectItem(item) {
        this.setState((state, props) => {
            return {selected: item};
        })
        if (typeof this.props.onChange === 'function') 
            this.props.onChange(item);
        this.hide();
    }
    renderItems() {
        return this.props.items.map((item) => {

            let selectedItem = this.state.selected && item.name === this.state.selected.name ? true : false
            let _classes = classNames({
                selected: selectedItem
            })
            return (
                <div 
                    className={_classes}
                    onClick={this.selectItem.bind(this, item)}>
                    {item.social ? <i className={"smfi " + item.social}></i> : null}
                    <span className="item-padding">
                        {item.name}
                    </span>
                </div>
            );
        })
    }
    firstSelected() {
        return this.props.items[0] || {name: "No Items"};
    }
    toggleVisible() {
        if (this.state.visible === false) {
            this.show();
        } else {
            this.hide();
        }
    }
    render() {
        let selected = this.state.selected || this.props.defaultSelected || this.firstSelected();
        let _classes = classNames("dropdown-list", {
            hide: !this.state.visible
        })
        return (
            <div className="dropdown-container">
                <div className="dropdown-view">
                    {selected.social ? <i className={"smfi " + selected.social}></i> : null}
                    <span>&nbsp;{selected.name}</span>
                    <i className="bfi down-arrow" onClick={(e) => this.toggleVisible()}></i>
                </div>
                <div className={_classes}>
                    <div style={{width: "100px"}}>
                        {this.renderItems()}
                    </div>
                </div>
            </div>
        )
    }
}