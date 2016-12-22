'use strict';

import React from 'react';
import Writer from './writer.jsx';
import RightColumn from './right_column.jsx';
import BlogList from './blog_list.jsx';

export default class Layout extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            pageKey: window.DATA.PageKey
        };
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
                                <li className="active"><a href="#">LINK</a></li>
                                <li><a href="#">LINK</a></li>
                            </ul>
                        </div>
                    </div>
                </nav>

                <main>
                    <div className="container">
                        <div className="row">
                            <div className="col-xs-8">
                                {(() => {
                                    if (this.state.pageKey != "TOP") {
                                        return <Writer />;
                                    } else {
                                        return <BlogList />;
                                    }
                                })()}
                            </div>
                            <div className="col-xs-4">
                                <RightColumn />
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        );
    }
}
