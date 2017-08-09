import React from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import ReactTable from 'react-table';
import 'react-table/react-table.css'

const styles = {
  toolbar: {
    paddingRight: 30,
    paddingLeft: 30,
  }
};

class UploadedSongsView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            dataList: [
                {
                    title: "Ain't Strong",
                    release: "Mars Moses EP",
                    uploadedAt: "10/21/16",
                },
                {
                    title: "Glue",
                    release: "Mars Moses EP",
                    uploadedAt: "10/1/16",
                }
            ]
        }
    }
    render() {
        const data = this.state.dataList;
        const columns = [
            {
                Header: 'Name',
                accessor: 'name'
            },
            {
                Header: 'Email',
                accessor: 'email'
            },
            {
                Header: 'Links',
                accessor: 'links'
            },
            {
                Header: 'Source',
                accessor: 'source'
            },
            {
                Header: 'Tags',
                accessor: 'tags'
            },
            {
                Header: 'Notes',
                accessor: 'notes'
            },
        ]
        return (
            <div>
                <div>
                    <ReactTable 
                        data={data}
                        columns={columns}
                        showPagination={false} />
                </div>
            </div>
        );
    }
}

export default UploadedSongsView;
