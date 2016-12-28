'use strict';

import React from 'react';
import { Link } from 'react-router';
import RightColumn from './right_column.jsx';

export default class Layout extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            blog_list: null
        };
    }

    setBlogList(blog_list) {
        this.setState({blog_list: blog_list});
    }

    render() {
        return (
            <div>
                <base href="/" />
                <nav className="navbar navbar-default">
                    <div className="container">
                        <div className="navbar-header">
                            <a href="#" className="navbar-brand">blogs</a>
                            <button className="navbar-toggle" type="button" data-toggle="collapse" data-target="#navbar-main">
                                <span className="icon-bar"></span>
                            </button>
                        </div>
                        <div className="navbar-collapse collapse" id="navbar-main">
                            <ul className="nav navbar-nav">
                                <li className="active"><Link to="/">TOP</Link></li>
                            </ul>
                        </div>
                    </div>
                </nav>

                <main>
                    <div className="container">
                        <div className="row">
                            <div className="col-xs-8">
                                {this.props.children}
                            </div>
                            <div className="col-xs-4">
                                <RightColumn key="right-column" />
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        );
    }
}
