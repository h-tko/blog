'use strict';

import React from 'react';
import $ from 'jquery';

export default class BlogDetail extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            willmount: false,
            blog: {}
        };
    }

    componentWillMount() {
        this.setState({willmount: true});
        this.requestDetail();
    }

    componentDidMount() {
        if (!this.state.willmount) {
            this.requestDetail();
        }

        this.setState({willmount: false});
    }

    requestDetail() {
        $.get(location.href + "?ajax=on", (result) => {
            this.setState({blog: result.blog});
        });
    }

    render() {

        var title = "";
        var release_date = "";

        if (Object.keys(this.state.blog).length > 0) {
            title = this.state.blog.Title;
            release_date = "(" + this.state.blog.ReleaseDate + ")";
        }

        return (
            <div className="row">
                <div className="col-xs-12">
                    <div className="panel panel-success">
                        <div className="panel-heading">
                            {title} {release_date}
                        </div>
                        <div className="panel-body">
                            <span dangerouslySetInnerHTML={{__html: this.state.blog.Body}} />
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
