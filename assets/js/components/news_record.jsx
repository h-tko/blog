'use strict';

import React from 'react';
import { Router, Route, Link } from 'react-router';

export default class NewsRecord extends React.Component {
    render() {
        var url = "/detail/" + this.props.ID + "?ajax=on";
        return (
            <li><Link to={url}>{this.props.Title}</Link></li>
        );
    }
}
