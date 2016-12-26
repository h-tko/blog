'use strict';

import React from 'react';
import BlogCard from './blog_card.jsx';

export default class BlogList extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            blog_list: window.DATA.Blogs
        };

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
