'use strict';

import React from 'react';

export default class BlogCard extends React.Component {
    render() {
        return (
            <div className="row">
                <div className="col-xs-12">
                    <div className="panel panel-success">
                        <div className="panel-heading">
                            {this.props.Title} ({this.props.ReleaseDate})
                        </div>
                        <div className="panel-body">
                            <span dangerouslySetInnerHTML={{__html: this.props.Body}} />
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
