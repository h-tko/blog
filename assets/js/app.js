require('material-design-lite/dist/material.deep_orange-amber.min.css');
require('material-design-lite/dist/material.min.js');
require('../css/app.css');
import hljs from 'highlight.js/lib/index.js';
import $ from 'jquery';
require('highlight.js/styles/solarized-dark.css');

$(function() {
    hljs.initHighlightingOnLoad();
});
