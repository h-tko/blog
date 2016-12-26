'use strict';

import React from 'react';
import { Link } from 'react-router';
import Writer from './writer.jsx';
import BlogDetail from './blog_detail.jsx';
import RightColumn from './right_column.jsx';
import BlogList from './blog_list.jsx';

export default class Layout extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div>
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
                                <li className="active"><Link to="/">ブログ一覧</Link></li>
                                <li><Link to="/write_blog">ブログを書く</Link></li>
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
