require('material-design-lite/dist/material.lime-deep_purple.min.css');
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

    $("button.good").on('click', () => {
        incrementGood();
    });

    getRightColumn();
});

const sendWriteData = () => {
    if (confirm("本当に登録してもよろしいですか？")) {
        var params = $('form[name=form-blog]').serialize();

        $.post("/regist_blog/", params, (result) => {
            if (result === "success") {
                alert("登録しました");

                $("button.clear").trriger("click");
            }
        });
    }
};

const incrementGood = () => {
    var id = $("#blog_id").val();

    $.get("/increment_good/", {
        blog_id: id
    }, (result) => {
        $(".good").prop("disabled", true);
        $(".fab-count").text(result);
    });
};

const getRightColumn = () => {
    if ($("#right-column").length < 1) {
        return;
    }

    $.get("/popular_list/", (result) => {
        $("#right-column").empty();

        console.log(result.PopularList);
        $.each(result.PopularList, (i, val) => {
            $("#right-column").append($('<li class="mdl-list__item"></li>').append($('<a href="/detail/' + val.ID + '/" class="mdl-list__item-secondary-action"></a>').text(val.Title + " いいね:" + val.BlogCount.GoodCount)));
        });
    });
}
