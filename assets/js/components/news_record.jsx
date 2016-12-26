'use strict';

import React from 'react';

export default class NewsRecord extends React.Component {
    render() {
        var url = "/detail/" + this.props.Id;
        return (
            <li><a href={url}>{this.props.Title}</a></li>
        );
    }
}
