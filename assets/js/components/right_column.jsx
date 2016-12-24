'use strict';

import React from 'react';
import NewsRecord from './news_record.jsx';

export default class RightColumn extends React.Component {
    render() {
        var list = [];

        for (var i in window.DATA.Blogs) {
            list.push(<NewsRecord key={window.DATA.Blogs[i].Id} {...window.DATA.Blogs[i]} />);
        }

        return (
            <div className="panel panel-default">
                <div className="panel-heading">新着記事</div>
                <ul className="nav nav-pills nav-stacked">
                    {list}
                </ul>
            </div>
        );
    }
}
