'use strict';

import React from 'react';
import $ from 'jquery';
import BlogCard from './blog_card.jsx';

export default class BlogList extends React.Component {
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
            list.push(<BlogCard key={i} {...this.state.blog_list[i]} />);
        }

        return (
            <div>
                {list}
            </div>
        );
    }
}
