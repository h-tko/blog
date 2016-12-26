require('bootswatch/united/bootstrap.min.css');

import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Routes, Route, IndexRoute, browserHistory } from 'react-router';
import Layout from './components/layout.jsx';
import BlogList from './components/blog_list.jsx';
import BlogDetail from './components/blog_detail.jsx';
import Writer from './components/writer.jsx';

(() => {
    ReactDOM.render(
        <Router history={browserHistory}>
            <Route path="/" component={Layout}>
                <IndexRoute component={BlogList} />
                <Route path="/detail/:id" component={BlogDetail} />
                <Route path="/write_blog" component={Writer} />
            </Route>
        </Router>
        ,document.getElementById('main')
    );
})();
