require('material-design-lite/dist/material.deep_orange-amber.min.css');
require('material-design-lite/dist/material.min.js');
require('../css/app.css');
import hljs from 'highlight.js/lib/index.js';
import $ from 'jquery';
import MaterialDateTimePicker from 'material-datetime-picker/dist/js/datepicker.js';
require('material-datetime-picker/dist/css/datepicker.css')
require('highlight.js/styles/solarized-dark.css');

$(function() {
    hljs.initHighlightingOnLoad();

    const picker = new MaterialDateTimePicker();
    const pickerElem = document.querySelectorAll(".timepicker")

    pickerElem.on('click', () => picker.open());
});
