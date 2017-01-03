require('material-design-lite/dist/material.deep_orange-amber.min.css');
require('material-design-lite/dist/material.min.js');
require('../css/app.css');
import hljs from 'highlight.js/lib/index.js';
import $ from 'jquery';
import 'jquery-ui/ui/core.js';
import 'jquery-ui/ui/widgets/datepicker.js';
import 'jquery-ui/themes/base/core.css';
import 'jquery-ui/themes/base/theme.css';
import 'jquery-ui/themes/base/datepicker.css';
require('highlight.js/styles/solarized-dark.css');

$(function() {
    hljs.initHighlightingOnLoad();

    $(".datepicker").datepicker({
        dateFormat: "yy/mm/dd"
    });

    $("button.clear").on('click', () => {
        $("input, textarea").val("");
    });

    $("button.regist-blog").on('click', () => {
        sendWriteData();
    });
});

const sendWriteData = () => {
    if (confirm("本当に登録してもよろしいですか？")) {
        var params = $('form[name=form-blog]').serialize();

        $.post("/regist_blog/", params, (result) => {
            if (result === "success") {
                alert("登録しました");

                $("button.clear").torriger("click");
            }
        });
    }
};
