'use strict';

import React from 'react';

export default class NewsRecord extends React.Component {
    render() {
        return (
            <li><a href="/detail/{this.props.Id}">{this.props.Title}</a></li>
        );
    }
}
