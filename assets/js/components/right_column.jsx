'use strict';

import React from 'react';
import NewsRecord from './news_record.jsx';
import $ from 'jquery';

export default class RightColumn extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            blog_list: {}
        };
    }

    componentDidMount() {
        this.requestBlogList();
    }

    requestBlogList() {
        $.ajax({
            type: "GET",
            url: "/?ajax=on",
        }).done((data) => {

            if (data.Data) {
                this.setState({blog_list: data.Data.Blogs})
            } else {
                this.setState({blog_list: {}});
            }
        }).fail((err) => {
            console.log(err);
        });
    }

    render() {
        var list = [];

        for (var i in this.state.blog_list) {
            list.push(<NewsRecord key={i} {...this.state.blog_list[i]} />);
        }

        if (list.length > 0) {
            return (
                <div className="panel panel-default">
                    <div className="panel-heading">新着記事</div>
                    <ul className="nav nav-pills nav-stacked">
                        {list}
                    </ul>
                </div>
            );
        } else {
            return (<div />);
        }
    }
}
