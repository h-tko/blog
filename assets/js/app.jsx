require('bootswatch/united/bootstrap.min.css');

import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, IndexRoute, hashHistory } from 'react-router';
import Layout from './components/layout.jsx';
import BlogList from './components/blog_list.jsx';
import Writer from './components/writer.jsx';

(() => {
    ReactDOM.render(
        <Router history={hashHistory}>
            <Route path="/" component={Layout}>
                <IndexRoute component={BlogList} />
                <Route path="/write_blog" component={Writer} />
            </Route>
        </Router>,
        document.getElementById('main')
    );
})();
